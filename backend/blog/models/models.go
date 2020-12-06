package models

import (
         "time"
         "sync"
       )

type Attitude struct {
    Id int `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Name string `json:"name"  gorm:"<-:create gorm:"<-:update"`
    Code string `json:"code" gorm:"<-:create gorm:"<-:update"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update"`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update"`
}

type Post struct {
    Id int `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Title       string `json:"title" gorm:"<-:create gorm:"<-:update`
    Body        string `json:"body" gorm:"<-:create gorm:"<-:update`
    TimePublished time.Time `json:"time_published" gorm:"<-:create gorm:"<-:update`
    TimeLastEdited time.Time `json:"time_last_edited" gorm:"<-:create gorm:"<-:update`
    TimeLastCommented time.Time `json:"time_last_commented" gorm:"<-:create gorm:"<-:update`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update`
    AttitudeId int `json:"attitude_id" gorm:"foreignKey:AttitudeRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Attitude Attitude `json:",omitempty" gorm:"foreignKey:AttitudeId"`
}

type Comment struct {
    Text     string     `json:"text" gorm:"<-:create gorm:"<-:update`
    PostId int  `json:"parent_id" gorm:"foreignKey:AttitudeRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Parent *Comment `json:",omitempty" gorm:"foreignKey:ParentId"` 
    ParentId   int        `json:"parent_id" gorm:"foreignKey:CommentRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Id       int        `json:"id" gorm:"<-:create gorm:"<-:update gorm:"primaryKey`
    Score    int        `json:"score" gorm:"<-:create gorm:"<-:update`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update"`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update"`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update"`
    Hide bool `json:"hide" gorm:"<-:create gorm:"<-:update"`
    AttitudeId int `json:"attitude_id" gorm:"foreignKey:AttitudeRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Attitude Attitude `json:",omitempty" gorm:"foreignKey:AttitudeId"`
    mux      sync.Mutex
}
