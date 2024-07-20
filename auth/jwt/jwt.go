package auth

import (
	"Taichi/sdk"
	"Taichi/utils"
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type SignType struct {
	SignType string `json:"sign_type"`
	ExpTime  int64  `json:"exp_time"`
}

type PreloadInterFace interface {
	GetAuthId() string
}

type SignValue string

func GetJwt(preLoad PreloadInterFace, sec int64) (string, error) {
	var signType = SignType{
		SignType: "RSA1024",
		ExpTime:  time.Now().Unix() + sec,
	}
	marshal, err := json.Marshal(signType)
	if err != nil {
		return "", err
	}
	header := base64.URLEncoding.EncodeToString(marshal)
	jsonData, err := json.Marshal(preLoad)
	if err != nil {
		return "", err
	}
	data := base64.URLEncoding.EncodeToString(jsonData)
	key := ""
	if v, err := sdk.Redis.Get(context.Background(), preLoad.GetAuthId()).Result(); err == nil && v != "" {
		key = v
	} else {
		key = fmt.Sprintf("%s.%d", preLoad.GetAuthId(), time.Now().Unix())
	}
	bytes := utils.Hash(crypto.SHA256, []byte(fmt.Sprintf("%s.%s.%s", header, data, key)))
	sign, err := utils.RsaSign([]byte(sdk.PrivateKey), crypto.SHA256, bytes)
	if err != nil {
		return "", err
	}
	signBase64 := base64.URLEncoding.EncodeToString(sign)
	sdk.Redis.Set(context.Background(), preLoad.GetAuthId(), key, 72*time.Hour)

	return fmt.Sprintf("%s.%s.%s", header, data, signBase64), nil

}
func VerifyJWT(jwt string, preload PreloadInterFace) error {
	res := strings.Split(jwt, ".")

	if len(res) != 3 {
		return errors.New("验证错误")
	}
	decodeString, err := base64.URLEncoding.DecodeString(res[2])
	if err != nil {
		return err
	}
	header, err := base64.URLEncoding.DecodeString(res[0])
	if err != nil {
		return err
	}
	data, err := base64.URLEncoding.DecodeString(res[1])
	if err != nil {
		return err
	}
	json.Unmarshal(data, preload)
	if v, err := sdk.Redis.Get(context.Background(), preload.GetAuthId()).Result(); err == nil && v != "" {
		bytes := utils.Hash(crypto.SHA256, []byte(fmt.Sprintf("%s.%s.%s", res[0], res[1], v)))
		var signType SignType
		err = json.Unmarshal(header, &signType)
		if err != nil {
			return err
		}
		if time.Now().Unix() >= signType.ExpTime {
			return errors.New("验证错误")
		}
		fmt.Println(fmt.Sprintf("%x", bytes))
		if jwtIs, err := sdk.Redis.Get(context.Background(), fmt.Sprintf("%x", bytes)).Result(); err == nil && jwtIs != "" {
			return errors.New("验证错误")
		}
		err = utils.RsaVerifySign([]byte(sdk.PubKey), crypto.SHA256, bytes, decodeString)
		if err != nil {
			return errors.New("验证错误")
		}
		return nil
	} else {
		return errors.New("验证错误")
	}

}
func Logout(jwt string, preload PreloadInterFace) error {
	res := strings.Split(jwt, ".")
	if len(res) != 3 {
		return errors.New("验证错误")
	}
	header, err := base64.URLEncoding.DecodeString(res[0])
	if err != nil {
		return err
	}
	data, err := base64.URLEncoding.DecodeString(res[1])
	if err != nil {
		return err
	}
	json.Unmarshal(data, preload)
	if v, err := sdk.Redis.Get(context.Background(), preload.GetAuthId()).Result(); err == nil && v != "" {
		bytes := utils.Hash(crypto.SHA256, []byte(fmt.Sprintf("%s.%s.%s", res[0], res[1], v)))
		var signType SignType
		err = json.Unmarshal(header, &signType)
		if err != nil {
			return err
		}
		if time.Now().Unix() >= signType.ExpTime {
			return nil
		}
		fmt.Println(fmt.Sprintf("%x", bytes))
		sdk.Redis.Set(context.Background(), fmt.Sprintf("%x", bytes), "1", time.Duration(signType.ExpTime-time.Now().Unix())*time.Second)
		return nil
	} else {
		return errors.New("验证错误")
	}

}
