package api

import (
    "context"
    "encoding/json"
    "regexp"
    "time"
    "unicode"
    "unicode/utf8"
    "io/ioutil"
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "syscall"
    "github.com/callicoder/go-docker-compose/model"
    "github.com/go-redis/redis"
    "github.com/gorilla/mux"
    "3dact.com/blog/dao"
    "3dact.com/blog/models"
)

var att map[string]models.Attitude

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshaller struct {
    Value interface{}
}

type PostRequest struct {
    Title       string `json:"title"`
    Body        string `json:"body"`
    AttitudeId    int `json:"attitude_id"`    
}


type AttitudeRequest struct {
    Name        string `json:"name"`
    Code        string `json:"code"`
}

func init_attitudes() {
    att = make(map[string]models.Attitude)
    att["love"] = models.Attitude{
        Id:1, Name:"love", Code: "love",
    }
    att["meh"] = models.Attitude{
        Id:2, Name:"meh", Code: "meh",
    }
    att["hate"] = models.Attitude{
       Id:3, Name:"hate", Code: "hate",
    } 
}


func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
    marshalled, err := json.Marshal(c.Value)

    converted := keyMatchRegex.ReplaceAllFunc(
        marshalled,
        func(match []byte) []byte {
            // Empty keys are valid JSON, only lowercase if we do not have an
            // empty key.
            if len(match) > 2 {
                // Decode first rune after the double quotes
                r, width := utf8.DecodeRune(match[1:])
                r = unicode.ToLower(r)
                utf8.EncodeRune(match[1:width+1], r)
            }
            return match
        },
    )

    return converted, err
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome! Please hit the `/qod` API to get the quote of the day."))
}

func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
    posts := dao.ReadAllPosts()
    respondWithJson(w, http.StatusOK, posts)
}

func getAllAttitudesHandler(w http.ResponseWriter, r *http.Request) {
    posts := dao.ReadAllAttitudes()
    respondWithJson(w, http.StatusOK, posts)
}

func createAttitudeHandler(w http.ResponseWriter, r *http.Request) {
    // Double check it's a post request being made
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method")
        return
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var ar  AttitudeRequest
    json.Unmarshal(reqBody, &ar)
   
    currentTime := time.Now()
    dateString := currentTime.Format("Mon, 02 Jan 2006 15:04:05 MST")
    time, err := time.Parse(time.RFC1123, dateString);

    if err!=nil {
        fmt.Println("Error while parsing date :", err);
    }

    var att = models.Attitude{
                            Name: ar.Name,
                            Code: ar.Code,
                            CreatedAt: time,
                            UpdatedAt: time,
              }

    dao.CreateAttitude(att)
    respondWithJson(w, http.StatusOK, att)
}


func createPostHandler(w http.ResponseWriter, r *http.Request) {

    // Double check it's a post request being made
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method")
        return
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var pr PostRequest
    json.Unmarshal(reqBody, &pr)
    var attitude = dao.ReadAttitudeById(pr.AttitudeId)
    msg := fmt.Sprintf("Attitude ID found - %d", pr.AttitudeId)  
    fmt.Println(msg)

    //attitude := att[pr.Attitude]

    
    currentTime := time.Now()
    dateString := currentTime.Format("Mon, 02 Jan 2006 15:04:05 MST")
    time, err := time.Parse(time.RFC1123, dateString);

    if err!=nil {
        fmt.Println("Error while parsing date :", err);
    }

    attitude.CreatedAt = time
    attitude.UpdatedAt = time

    var p = models.Post{
                            Id:1,
                            Title: pr.Title,
                            Body: pr.Body,
                            Attitude: attitude,
                            TimePublished: time,
                            TimeLastEdited: time,
                            CreatedAt: time,
                            UpdatedAt: time,
                  }

    dao.CreatePost(p) 
    respondWithJson(w, http.StatusOK, p)
}

func getPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var post = dao.ReadPostById(id)
    respondWithJson(w, http.StatusOK, post)
}

func getAttitudeByIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var attitude = dao.ReadAttitudeById(id)
    respondWithJson(w, http.StatusOK, attitude)
}

func getAttitudeByCodeHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    code := string(params["code"])
    var attitude = dao.ReadAttitudeByCode(code)
    respondWithJson(w, http.StatusOK, attitude)
}



func Register() {
    // Create Redis Client

    init_attitudes() 
    client := redis.NewClient(&redis.Options{
        Addr:     getEnv("REDIS_URL", "localhost:6379"),
        Password: getEnv("REDIS_PASSWORD", ""),
        DB:       0,
    })

    _, err := client.Ping().Result()
    if err != nil {
        log.Fatal(err)
    }

    // Create Server and Route Handlers
    r := mux.NewRouter()

    r.HandleFunc("/", indexHandler)
    r.HandleFunc("/posts/{id}", getPostByIdHandler).Methods("GET")
    r.HandleFunc("/posts", getAllPostsHandler).Methods("GET")
    r.HandleFunc("/posts", createPostHandler).Methods("POST")
     
    r.HandleFunc("/attitudes/{id}", getAttitudeByIdHandler).Methods("GET")
    r.HandleFunc("/attitudes", getAllAttitudesHandler).Methods("GET")
    r.HandleFunc("/attitudes", createAttitudeHandler).Methods("POST")
    r.HandleFunc("/attitudesbc/{code}", getAttitudeByCodeHandler).Methods("GET")
    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    // Start Server
    go func() {
        log.Println("Starting Server")
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    // Graceful Shutdown
    waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
    interruptChan := make(chan os.Signal, 1)
    signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // Block until we receive our signal.
    <-interruptChan

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    srv.Shutdown(ctx)

    log.Println("Shutting down")
    os.Exit(0)
}

func getQuoteFromAPI() (*model.QuoteResponse, error) {
    API_URL := "http://quotes.rest/qod.json"
    resp, err := http.Get(API_URL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    log.Println("Quote API Returned: ", resp.StatusCode, http.StatusText(resp.StatusCode))

    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        quoteResp := &model.QuoteResponse{}
        json.NewDecoder(resp.Body).Decode(quoteResp)
        return quoteResp, nil
    } else {
        return nil, errors.New("Could not get quote from API")
    }

}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// method to print error output for http respon
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// method to print output for http respon
// parameter  [w (Http.RestponWriter), http.statuscode, payload/data/msg]
// payload is data credential which will be trans to other part
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.MarshalIndent(conventionalMarshaller{payload}, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
