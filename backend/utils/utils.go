package utils

import (
    "context"
    "encoding/json"
    "errors"
    "github.com/callicoder/go-docker-compose/model"
    "net/http"
    "log"
    "os"
    "os/signal"
    "regexp"
    "syscall"
    "time"
    "unicode"
    "unicode/utf8"
)

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshaller struct {
    Value interface{}
}

func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// method to print error output for http respon
func RespondWithError(w http.ResponseWriter, code int, msg string) {
    RespondWithJson(w, code, map[string]string{"error": msg})
}

// method to print output for http respon
// parameter  [w (Http.RestponWriter), http.statuscode, payload/data/msg]
// payload is data credential which will be trans to other part
func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.MarshalIndent(conventionalMarshaller{payload}, "", "  ")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
    marshalled, err := json.Marshal(c.Value)

    converted := keyMatchRegex.ReplaceAllFunc(
        marshalled,
        func(match []byte) []byte {
            // Empty keys are valid JSON, only lowercase if we do not have an
            // empty key.
            if len(match) > 2 {
                // Decode first rune after the double quotes
                r, width := utf8.DecodeRune(match[1:])
                r = unicode.ToLower(r)
                utf8.EncodeRune(match[1:width+1], r)
            }
            return match
        },
    )

    return converted, err
}

func WaitForShutdown(srv *http.Server) {
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

func GetQuoteFromAPI() (*model.QuoteResponse, error) {
    API_URL := "http://quotes.rest/qod.json"
    resp, err := http.Get(API_URL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    log.Println("Quote API Returned: ", resp.StatusCode, http.StatusText(resp.StatusCode))

    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        quoteResp := &model.QuoteResponse{}
        json.NewDecoder(resp.Body).Decode(quoteResp)
        return quoteResp, nil
    } else {
        return nil, errors.New("Could not get quote from API")
    }

}

