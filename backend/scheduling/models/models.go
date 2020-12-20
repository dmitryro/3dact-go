package models

import (
         "time"
         //"3dact.com/user/models"
       )

type Message struct {
    Id int `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Email    string `json:"email"`
    Username string `json:"username"`
    Message  string `json:"message"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update"`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update"`
}}


