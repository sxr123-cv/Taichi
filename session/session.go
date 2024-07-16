package session

import (
	"Taichi/sdk"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Session struct {
	ExpTime    int64  `json:"exp_time"`
	PreLoad    any    `json:"pre_load"`
	SessionKey string `json:"session_key"`
}

func SaveSession(PreLoad any, sec int64) (string, error) {
	sha := sha256.New()
	s := Session{
		ExpTime:    time.Now().Unix() + sec,
		PreLoad:    PreLoad,
		SessionKey: "",
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	sha.Write(marshal)
	bytes := sha.Sum(nil)
	s.SessionKey = fmt.Sprintf("%x", bytes)
	ctx := context.WithValue(context.Background(), "q1111", "2222") //协程 通道
	err = WriteDataToRedis(s.SessionKey, string(marshal), ctx)
	if err != nil {
		return "", err
	}
	return s.SessionKey, nil
}

func VerifySession(session string, preload any) error {
	content, err := ReadDataFormRedis(session)
	if err != nil {
		sdk.Log.INFO(err.Error())
		return err
	}
	var s Session
	err = json.Unmarshal(content, &s)
	if err != nil {
		return err
	}
	if time.Now().Unix() >= s.ExpTime {
		os.Remove(session)
		return errors.New("会话过期")
	}
	marshal, err := json.Marshal(s.PreLoad)
	if err != nil {
		return err
	}
	json.Unmarshal(marshal, preload)

	return nil

}

type Preload struct {
	Role   string `json:"role"`
	UserId int64  `json:"user_id"`
}

func ReadDataFormRedis(session string) (bytes []byte, err error) {
	ctx := context.Background()
	result, err := sdk.Redis.Get(ctx, session).Result()
	if err != nil {
		return nil, err
	}
	return []byte(result), err
}

func WriteDataToRedis(session string, jsonData string, ctx context.Context) error {
	sdk.Log.INFO(ctx.Value("q1111"))
	return sdk.Redis.Set(ctx, session, jsonData, 6*time.Hour).Err()

}
