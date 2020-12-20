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
    "github.com/gorilla/mux"
    "3dact.com/user/dao"
    "3dact.com/user/models"
)


// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshaller struct {
    Value interface{}
}

type UserRequest struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json: "email"`
    Username string `json: "username"`
    Password string `json: "password"`
    Bio string `json:"bio"`
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


func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := dao.ReadAllUsers()
    respondWithJson(w, http.StatusOK, users)
}


func createUserHandler(w http.ResponseWriter, r *http.Request) {
    // Double check it's a post request being made
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method")
        return
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var ur  UserRequest
    json.Unmarshal(reqBody, &ur)
   
    currentTime := time.Now()
    dateString := currentTime.Format("Mon, 02 Jan 2006 15:04:05 MST")
    time, err := time.Parse(time.RFC1123, dateString);

    if err!=nil {
        fmt.Println("Error while parsing date :", err);
    }

    var user = models.User{
                            Username: ur.Username,
                            Password: ur.Password,
                            FirstName: ur.FirstName,
                            LastName: ur.LastName,
                            Email: ur.Email,
                            CreatedAt: time,
                            UpdatedAt: time,
              }

    dao.CreateUser(user)
    respondWithJson(w, http.StatusOK, user)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var user = dao.ReadUserById(id)
    respondWithJson(w, http.StatusOK, user)
}


func Register(r *mux.Router) {
    // Create Redis Client
    r.HandleFunc("/users/{id}", getUserByIdHandler).Methods("GET")
    r.HandleFunc("/users", getAllUsersHandler).Methods("GET")
    r.HandleFunc("/users", createUserHandler).Methods("POST")
    // Graceful Shutdown
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
