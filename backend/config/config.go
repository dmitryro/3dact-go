package config

import (
	"log" // to print log
    "os"
    "github.com/go-redis/redis"
)

// Represent database server n credentials
type Config struct {
	Server   string
	Database string
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}


func Start() {

    // Create Redis Client

    client := redis.NewClient(&redis.Options{
        Addr:     getEnv("REDIS_URL", "localhost:6379"),
        Password: getEnv("REDIS_PASSWORD", ""),
        DB:       0,
    })

    _, err := client.Ping().Result()
    if err != nil {
        log.Fatal(err)
    }

}



