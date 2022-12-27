package service

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
	"my-redis-cli/internal/define"
)

// ConnectionList 连接列表
func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	conf := new(define.Config)
	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, errors.New("json parse to config error：" + err.Error())
	}

	return conf.Connections, nil
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

	_uuid, _ := uuid.NewUUID()
	conn.Identity = _uuid.String()

	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)

	// 不存在配置文件 创建配置文件
	if errors.Is(err, os.ErrNotExist) {
		conf.Connections = []*define.Connection{conn}
		bytes, _ = json.Marshal(conf)
		os.MkdirAll(nowPath, 0666)
		os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)

		return nil
	}

	// 存在配置文件 则追加conn
	json.Unmarshal(bytes, conf)
	conf.Connections = append(conf.Connections, conn)
	bytes, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)

	return nil
}

// ConnectionEdit 编辑链接
func ConnectionEdit(conn *define.Connection) error {
	if conn.Identity == "" {
		return errors.New("连接唯一标识不能为空")
	}
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
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, conf)
	for i, v := range conf.Connections {
		if v.Identity == conn.Identity {
			conf.Connections[i] = conn
		}
	}

	bytes, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)

	return nil
}

// ConnectionDelete 编辑链接
func ConnectionDelete(identity string) error {
	if identity == "" {
		return errors.New("连接唯一标识不能为空")
	}

	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return err
	}

	json.Unmarshal(bytes, conf)
	for i, v := range conf.Connections {
		if v.Identity == identity {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
			break
		}
	}

	bytes, _ = json.Marshal(conf)
	os.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, bytes, 0666)

	return nil
}
