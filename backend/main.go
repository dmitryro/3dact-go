package main


import (
    "fmt"
    "3dact.com/blog"
    "3dact.com/blog/models"
    "3dact.com/blog/api"
)

func main() {
    // Get a greeting message and print it.
    api.Register()
    p := models.Post{ // b == Student{"Bob", 0}
      Id:1,
      Title:"Very cool Title",
      Body: "Very cool Body",
    }
    s := fmt.Sprintf("%d %s %s",p.Id, p.Title, p.Body)
    fmt.Println("LET US SEE "+s)
    message := blog.Hello("Gladys")
    fmt.Println(message)
}
