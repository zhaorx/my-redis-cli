package main

import (
    "context"
    "fmt"

    "my-redis-cli/internal/service"
)

// App struct
type App struct {
    ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
    return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
    return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ConnectionList 连接列表
func (a *App) ConnectionList() interface{} {

    connList, err := service.ConnectionList()
    if err != nil {
        return map[string]interface{}{
            "code": -1,
            "msg":  "Error: " + err.Error(),
        }
    }

    return map[string]interface{}{
        "code": 200,
        "msg":  "success",
        "data": connList,
    }
}
