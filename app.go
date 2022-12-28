package main

import (
	"context"
	"fmt"

	"my-redis-cli/internal/define"
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
func (a *App) ConnectionList() H {

	connList, err := service.ConnectionList()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "Error: " + err.Error(),
		}
	}

	return M{
		"code": 200,
		"msg":  "success",
		"data": connList,
	}
}

// ConnectionCreate 新建连接
func (a *App) ConnectionCreate(conn *define.Connection) H {
	err := service.ConnectionCreate(conn)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "Error: " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "新建成功",
	}
}

// ConnectionEdit 编辑连接
func (a *App) ConnectionEdit(conn *define.Connection) H {
	err := service.ConnectionEdit(conn)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "Error: " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "修改成功",
	}
}

// ConnectionDelete 删除连接
func (a *App) ConnectionDelete(identity string) H {
	err := service.ConnectionDelete(identity)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "Error: " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"msg":  "删除成功",
	}
}
