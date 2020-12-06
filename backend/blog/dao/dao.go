package dao

import (
    "fmt"
    "gorm.io/gorm"
    "3dact.com/blog/models"
    "3dact.com/config/dbconnect"
    "gorm.io/driver/postgres"
)


func CreatePost(post models.Post) models.Post{
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Omit("Id", "CreatedAt").Create(&post) // pass pointer of data to Crea
     return post
}

func CreateAttitude(attitude models.Attitude) models.Attitude{
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Omit("Id", "CreatedAt").Create(&attitude) // pass pointer of data to Crea
     return attitude
}

func ReadAllAttitudes() []models.Attitude{
     var attitudes = []models.Attitude{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Find(&attitudes)
     return attitudes
}

func ReadAllPosts() []models.Post{
     posts := []models.Post{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Find(&posts)
     for i:=0; i < len(posts); i++ {
          var attitude = models.Attitude{}
          db.First(&attitude, posts[i].AttitudeId) 
          posts[i].Attitude = attitude
     }
     return posts
}

func ReadAttitudeByCode(code string) models.Attitude{
     var attitude = models.Attitude{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Where(&models.Attitude{Code: code}).First(&attitude)
     return attitude
}


func ReadAttitudeById(id int) models.Attitude{
     var attitude = models.Attitude{} 
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
    db.Where(&models.Attitude{Id: id}).First(&attitude)
    return attitude
}

func ReadPostById(id int) models.Post{
     var post = models.Post{}
     var attitude = models.Attitude{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.First(&post, id)
     db.First(&attitude, post.AttitudeId)
     post.Attitude = attitude
     return post
}

func updatePost(id int, post models.Post) {

}

func deletePost(id int) {

}

func initialize() {

}
