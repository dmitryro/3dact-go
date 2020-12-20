package api

import (
    "encoding/json"
    "time"
    "io/ioutil"
    "fmt"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "3dact.com/user/dao"
    "3dact.com/user/models"
    "3dact.com/utils"
)


type UserRequest struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json: "email"`
    Username string `json: "username"`
    Password string `json: "password"`
    Bio string `json:"bio"`
}


func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := dao.ReadAllUsers()
    utils.RespondWithJson(w, http.StatusOK, users)
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
    utils.RespondWithJson(w, http.StatusOK, user)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r) // mux library to get all parameters
    id, _ := strconv.Atoi(params["id"])
    var user = dao.ReadUserById(id)
    utils.RespondWithJson(w, http.StatusOK, user)
}


func Register(r *mux.Router) {
    // Create Redis Client
    r.HandleFunc("/users/{id}", getUserByIdHandler).Methods("GET")
    r.HandleFunc("/users", getAllUsersHandler).Methods("GET")
    r.HandleFunc("/users", createUserHandler).Methods("POST")
    // Graceful Shutdown
}


