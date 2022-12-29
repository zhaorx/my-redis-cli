package helper

import (
	"encoding/json"
	"errors"
	"os"

	"my-redis-cli/internal/define"
)

func GetConnection(identity string) (*define.Connection, error) {
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	bytes, err := os.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}

	// 存在配置文件 则追加conn
	err = json.Unmarshal(bytes, conf)
	if err != nil {
		return nil, err
	}

	for _, v := range conf.Connections {
		if v.Identity == identity {
			return v, nil
		}
	}

	return nil, errors.New("连接数据不存在")
}
