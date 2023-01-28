package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"my-redis-cli/internal/define"
	"my-redis-cli/internal/helper"
)

func KeyList(req define.KeyListRequest) ([]string, error) {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})

	count := 100
	if req.Keyword != "" {
		count = 2000
	}
	res, _, err := rdb.Scan(context.Background(), 0, "*"+req.Keyword, int64(count)).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetKeyValue(req define.KeyValueRequest) (*define.KeyValueReply, error) {
	conn, err := helper.GetConnection(req.ConnIdentity)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     conn.Addr + ":" + conn.Port,
		Username: conn.Username,
		Password: conn.Password,
		DB:       req.Db,
	})

	_type, err := rdb.Type(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}

	v, err := rdb.Get(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}
	ttl, err := rdb.TTL(context.Background(), req.Key).Result()
	if err != nil {
		return nil, err
	}

	return &define.KeyValueReply{
		Value: v,
		Type:  _type,
		TTL:   ttl,
	}, nil
}
