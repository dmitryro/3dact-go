package main


import (
    "context"
    "time"
    "log"
    "os"
    "syscall"
    "os/signal"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/go-redis/redis"
    blog_api "3dact.com/blog/api"
    user_api "3dact.com/user/api"
)

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func Register() {
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

    // Create Server and Route Handlers
    r := mux.NewRouter()

    blog_api.Register(r)
    user_api.Register(r)

    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    // Start Server
    go func() {
        log.Println("Starting Server")
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    // Graceful Shutdown
    waitForShutdown(srv)
}

func main() {
    // Get a greeting message and print it.
     Register()
}

func waitForShutdown(srv *http.Server) {
    interruptChan := make(chan os.Signal, 1)
    signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // Block until we receive our signal.
    <-interruptChan

    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    srv.Shutdown(ctx)

    log.Println("Shutting down")
    os.Exit(0)
}

