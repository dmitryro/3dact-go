package dbconnect

import (
  "fmt"
  "os"
)


type connect struct {
    Username string
    Password string
    DbHost string
    DbName string
    TimeZone string
    SslMode string
    Dsn string
}

func Connect() *connect{
    var c connect
    c.Username = os.Getenv("POSTGRES_USER")
    c.Password = os.Getenv("POSTGRES_PASSWORD")
    c.DbName = os.Getenv("POSTGRES_DB")
    c.DbHost = os.Getenv("POSTGRES_HOST")
    c.SslMode = os.Getenv("POSTGRES_SSL_MODE")
    c.TimeZone = os.Getenv("POSTGRES_TIME_ZONE")
    c.Dsn = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s sslmode=%s TimeZone=%s",
                        c.DbHost, c.Username, c.DbName, c.Password, c.SslMode, c.TimeZone)
    return &c
}

