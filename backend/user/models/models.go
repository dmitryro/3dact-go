package models

import (
         "time"
       )

type User struct {
    Id int `json:"id"`
    Username string `json: "username"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    IsStaff bool `json:"is_staff"`
    IsAdmin bool `json: "is_admin"`
    IsActive bool `json:"is_active"`
    IsCleared bool `json:"is_cleared"`
    IsSuspended bool `json: "is_suspended"`
    IsSuperuser bool `json:"is_superuser"`
    InChat bool `json:"in_chat"`
    LastLogin time.Time `json: "last_login"`
    DateJoined time.Time `json: "date_joined"`
    Bio string `json: "bio"`
    Category string `json: "category"`
    Password string `json: "password"`
    ActivationKey string `json: "activation_key"`
    PasswordRecoveryKey string `json: "password_recovery_key"`
    ProfileImage string `json: "profile_image"`
    AvatarImage string `json: "avatar_image"`
    AvatarThumbnail string `json: "avatar_thumbnail"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update`
}

type Relationship struct {
    Id int `json:"id"`
    Code string `json:"code"`
    Name string `json:"name"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update`
}

type Peer struct {
    Id int `json:"id"`
    CreatedAt    time.Time `json:"created_at" gorm:"<-:create gorm:"<-:update`
    UpdatedAt    time.Time `json:"updated_at" gorm:"<-:create gorm:"<-:update`
    DeletedAt time.Time `json:"deleted_at" gorm:"<-:create gorm:"<-:update`
    RelationId int `json:"relation_id" gorm:"foreignKey:RelationRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Relation Relationship `json:",omitempty" gorm:"foreignKey:RelationId"`
    InitiatorId int `json:"initiator_id" gorm:"foreignKey:InitiatorRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Initiator User `json:",omitempty" gorm:"foreignKey:InitiatorId"`
    AcceptorId int `json:"acceptor_id" gorm:"foreignKey:AcceptorRefer" constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
    Acceptor User `json:",omitempty" gorm:"foreignKey:AcceptorId"`     
}
