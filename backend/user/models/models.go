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
    IsSuperuser bool `json:"is_superuser"`
    LastLogin time.Time `json: "last_login"`
    DateJoined time.Time `json: "date_joined"`
    LastVisit time.Time `json: "last_visit"`
    Bio string `json: "bio"`
    Category string `json: "category"`
    Password string `json: "password"`
}

