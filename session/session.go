package session

import (
	"Taichi/log"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Session struct {
	ExpTime    int64  `json:"exp_time"`
	PreLoad    any    `json:"pre_load"`
	SessionKey string `json:"session_key"`
}

var L = log.NewLog(nil, nil)

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
	create, err := os.Create(s.SessionKey)
	if err != nil {
		return "", err
	}
	_, err = create.WriteString(string(marshal))
	if err != nil {
		return "", err
	}
	create.Close()
	return s.SessionKey, nil
}

func VerifySession(session string, preload any) error {
	content, err := os.ReadFile(session)
	if err != nil {
		L.INFO(err.Error())
		return err
	}
	var s Session
	err = json.Unmarshal(content, &s)
	if err != nil {
		return err
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
