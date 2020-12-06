package main


import (
    "fmt"
    "lovehate.io/blog"
    "lovehate.io/blog/models"
    blogapi "lovehate.io/blog/api"
)

func main() {
    // Get a greeting message and print it.
    blogapi.Register()
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
