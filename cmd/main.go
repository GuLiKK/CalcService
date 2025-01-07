package main

import (
    "github.com/GuLiKK/CalcService/internal/application"
)

func main() {
    app := application.New()
    // Выберите режим: CLI или Web-сервер
    // _ = app.Run() // CLI
    _ = app.RunServer() // Web
}
