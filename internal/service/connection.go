package service

import (
    "errors"
    "os"

    "github.com/tidwall/gjson"
    "my-redis-cli/internal/define"
)

// ConnectionList 连接列表
func ConnectionList() ([]*define.Connection, error) {
    nowPath, _ := os.Getwd()
    bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
    if errors.Is(err, os.ErrNotExist) {
        return nil, err
    }

    config, ok := gjson.Parse(string(bytes)).Value().(define.Config)
    if !ok {
        return nil, errors.New("json parse to config error")
    }

    return config.Connections, nil
}
