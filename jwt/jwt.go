package jwt

import (
	"Taichi/sdk"
	"Taichi/utils"
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

type Preload any

type SignValue string

func GetJwt(preLoad Preload, sec int64) (string, error) {
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
	sign, err := utils.RsaSign([]byte(sdk.PrivateKey), crypto.SHA256, []byte(fmt.Sprintf("%s.%s", header, data)))
	if err != nil {
		return "", err
	}
	signBase64 := base64.URLEncoding.EncodeToString(sign)
	return fmt.Sprintf("%s.%s.%s", header, data, signBase64), nil

}
func VerifyJWT(jwt string, preload any) error {
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
	err = utils.RsaVerifySign([]byte(sdk.PubKey), crypto.SHA256, []byte(fmt.Sprintf("%s.%s", res[0], res[1])), decodeString)
	if err != nil {
		return errors.New("验证错误")
	}
	var signType SignType
	err = json.Unmarshal(header, &signType)
	if err != nil {
		return err
	}
	if time.Now().Unix() >= signType.ExpTime {
		return errors.New("验证错误")
	}
	json.Unmarshal(data, preload)
	return nil

}
