package dao

import (
    "fmt"
    "gorm.io/gorm"
    "3dact.com/user/models"
    "3dact.com/config/dbconnect"
    "gorm.io/driver/postgres"
)

func CreateUser(user models.User) models.User{
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Omit("Id", "CreatedAt").Create(&user) // pass pointer of data to Crea
     return user
}


func ReadAllUsers() []models.User{
     var users = []models.User{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.Find(&users)
     return users
}

func ReadUserById(id int) models.User{
     var user = models.User{}
     conn := dbconnect.Connect()
     db, err := gorm.Open(postgres.Open(conn.Dsn), &gorm.Config{})
     if err != nil {
         msg := fmt.Sprintf("Error connecting to database - %s", err)
         fmt.Println(msg)
     }
     db.First(&user, id)
     return user
}



func initialize() {

}
