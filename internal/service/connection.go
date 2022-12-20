package service

import (
	"encoding/json"
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

// ConnectionCreate 创建链接
func ConnectionCreate(conn *define.Connection) error {
	if conn.Addr == "" {
		return errors.New("连接地址不能为空")
	}
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Name = "6379"
	}

	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)

	// 不存在配置文件 创建配置文件
	if errors.Is(err, os.ErrNotExist) {
		conf.Connections = []*define.Connection{conn}
		bytes, _ = json.Marshal(conf)
		os.MkdirAll(nowPath, 0666)
		os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)
	}
	// 存在配置文件 则追加conn
	json.Unmarshal(bytes, conf)
	conf.Connections = append(conf.Connections, conn)
	bytes, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)
	return nil
}
