package blog

type User struct {
    Id int `json:"id"`
    Username string `json: "username"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    IsStaff bool `json:"is_staff"`
    IsActive bool `json:"is_active"`
    IsSuperuser bool `json:"is_superuser"`
    LastLogin string `json: "last_login"`
    DateJoined string `json: "date_joined"`
    Bio string `json: "bio"`    
    Category string `json: "category"`
    Password string `json: "password"`
}

type Topic struct {

}

type Attitude struct {
}

type Group struct {
    Id int `json:"id"`
    Name string `json:"name"`
}


type Address struct {
    Id int `json:"id"`
    Address1 string `json:"address1"`
    Address2 string `json:"address2"`
    City string `json:"city"`
    State string `json:"state"`
    PostalCode string `json:"postal_code"`
    Country string `json: "country"`   
}

type Item struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Category string `json:"category"`
    Price float32 `json:"price"`
}

type State struct {
    Id int `json:"id"`
    Name string `json: "name"`
    Code string `json: "code`
}

