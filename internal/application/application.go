package application

import (
    "bufio"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/GuLiKK/CalcService/pkg/calculation"
)

type Config struct {
    Addr string
}

func ConfigFromEnv() *Config {
    c := new(Config)
    c.Addr = os.Getenv("PORT")
    if c.Addr == "" {
        c.Addr = "8080"
    }
    return c
}

type Application struct {
    config *Config
}

func New() *Application {
    return &Application{config: ConfigFromEnv()}
}

func (a *Application) Run() error {
    reader := bufio.NewReader(os.Stdin)
    for {
        text, err := reader.ReadString('\n')
        if err != nil {
            continue
        }
        text = strings.TrimSpace(text)
        if text == "exit" {
            return nil
        }
        result, err := calculation.Calc(text)
        if err != nil {
            log.Println("error:", err)
        } else {
            log.Printf("%s = %g\n", text, result)
        }
    }
}

type Request struct {
    Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
        return
    }
    result, err := calculation.Calc(req.Expression)
    if err != nil {
        if errors.Is(err, calculation.ErrInvalidExpression) ||
           errors.Is(err, calculation.ErrDivisionByZero) {
            http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"result":"%g"}`, result)
}

func (a *Application) RunServer() error {
    http.HandleFunc("/api/v1/calculate", CalcHandler)
    log.Printf("Starting server on :%s\n", a.config.Addr)
    return http.ListenAndServe(":"+a.config.Addr, nil)
}
