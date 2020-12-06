package models

import (
         "time"
         "3dact.com/user/models"
       )

type Room struct {
    Id int `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Name string `json:"name"  gorm:"<-:create gorm:"<-:update"`
    CreatedBy models.User  `json:",omitempty" gorm:"foreignKey:CreatedById"`
    CreatedById int  `json:"created_by_id" gorm:"foreignKey:CreatedbyRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update"`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update"`
    Level int `json:"level" gorm:"<-:create gorm:"<-:update"`
}

type Message struct {
    Id int `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Email    string `json:"email"`
    Username string `json:"username"`
    Message  string `json:"message"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update"`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update"`
}}


