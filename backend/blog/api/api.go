package api

import (
    "encoding/json"
    "time"
    "io/ioutil"
    "fmt"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "3dact.com/blog/dao"
    "3dact.com/blog/models"
    "3dact.com/utils"
)

var att map[string]models.Attitude

type PostRequest struct {
    Title       string `json:"title"`
    Body        string `json:"body"`
    AttitudeId    int `json:"attitude_id"`    
}


type AttitudeRequest struct {
    Name        string `json:"name"`
    Code        string `json:"code"`
}

type CommentRequest struct {
    Text string `json:"text"`
    AttitudeId    int `json:"attitude_id"`
    ParentId    int `json:"parent_id"`
    PostId    int `json:"post_id"`
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome! Please hit the `/qod` API to get the quote of the day."))
}

func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
    posts := dao.ReadAllPosts()
    utils.RespondWithJson(w, http.StatusOK, posts)
}

func getAllAttitudesHandler(w http.ResponseWriter, r *http.Request) {
    posts := dao.ReadAllAttitudes()
    utils.RespondWithJson(w, http.StatusOK, posts)
}

func getAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
    comments := dao.ReadAllComments()
    utils.RespondWithJson(w, http.StatusOK, comments)
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
    utils.RespondWithJson(w, http.StatusOK, att)
}

func createCommentHandler(w http.ResponseWriter, r *http.Request) {
    // Double check it's a post request being made
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "invalid_http_method")
        return
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var cr  CommentRequest
    json.Unmarshal(reqBody, &cr)
  
    currentTime := time.Now()
    dateString := currentTime.Format("Mon, 02 Jan 2006 15:04:05 MST")
    time, err := time.Parse(time.RFC1123, dateString);

    if err!=nil {
        fmt.Println("Error while parsing date :", err);
    }

    var comment = models.Comment{
                            Text: cr.Text,
                            ParentId: cr.ParentId,
                            PostId: cr.PostId,
                            AttitudeId: cr.AttitudeId,
                            CreatedAt: time,
                            UpdatedAt: time,
              }

    dao.CreateComment(comment)
    utils.RespondWithJson(w, http.StatusOK, comment)
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
                            Title: pr.Title,
                            Body: pr.Body,
                            Attitude: attitude,
                            TimePublished: time,
                            TimeLastEdited: time,
                            CreatedAt: time,
                            UpdatedAt: time,
                  }

    dao.CreatePost(p) 
    utils.RespondWithJson(w, http.StatusOK, p)
}

func getCommentByIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var comment = dao.ReadCommentById(id)
    utils.RespondWithJson(w, http.StatusOK, comment)
}


func getAllCommentsByPostIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var comments = dao.ReadCommentsByPostId(id)
    utils.RespondWithJson(w, http.StatusOK, comments)
}

func getAllCommentsByParentIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var comments = dao.ReadCommentsByParentId(id)
    utils.RespondWithJson(w, http.StatusOK, comments)
}

func getPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var post = dao.ReadPostById(id)
    utils.RespondWithJson(w, http.StatusOK, post)
}

func getAttitudeByIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var attitude = dao.ReadAttitudeById(id)
    utils.RespondWithJson(w, http.StatusOK, attitude)
}

func getAttitudeByCodeHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    code := string(params["code"])
    var attitude = dao.ReadAttitudeByCode(code)
    utils.RespondWithJson(w, http.StatusOK, attitude)
}



func Register(r *mux.Router) {
    // Create Redis Client

    init_attitudes() 

    r.HandleFunc("/", indexHandler)
    r.HandleFunc("/posts/{id}", getPostByIdHandler).Methods("GET")
    r.HandleFunc("/posts", getAllPostsHandler).Methods("GET")
    r.HandleFunc("/posts", createPostHandler).Methods("POST")

    r.HandleFunc("/comments", getAllCommentsHandler).Methods("GET")
    r.HandleFunc("/comments_post/{id}", getAllCommentsByPostIdHandler).Methods("GET")
    r.HandleFunc("/comments_parent/{id}", getAllCommentsByParentIdHandler).Methods("GET")
    r.HandleFunc("/comment/{id}", getCommentByIdHandler).Methods("GET")    
    r.HandleFunc("/comments", createCommentHandler).Methods("POST") 
    
    r.HandleFunc("/attitudes/{id}", getAttitudeByIdHandler).Methods("GET")
    r.HandleFunc("/attitudes", getAllAttitudesHandler).Methods("GET")
    r.HandleFunc("/attitudes", createAttitudeHandler).Methods("POST")
    r.HandleFunc("/attitudesbc/{code}", getAttitudeByCodeHandler).Methods("GET")
}
