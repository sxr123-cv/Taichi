package sdk

import (
	"Taichi/config"
	"Taichi/log"
	redis2 "Taichi/redis"
	"github.com/redis/go-redis/v9"
)

var Log = log.NewLog(nil, nil)

var Redis *redis.Client = nil

const PubKey = "-----BEGIN RSA PUBLIC KEY-----\nMIGJAoGBAKFPBI3KEF6+Avi9Zi3BCVBh6PnDqIiHE74rTAg78LFyUHFDCYdxVxxG\nYfWsVZia9d2rLVfYg2hYCBm7wDD62d9Y9puiZq9Bnr8O+eRJmrTTORLEXiSgVimw\nCMl6l3JjN/y8yxERTz8LbZ6JAEkrGLlVmpacCxsj+UHdTKJWWFzvAgMBAAE=\n-----END RSA PUBLIC KEY-----"
const PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQChTwSNyhBevgL4vWYtwQlQYej5w6iIhxO+K0wIO/CxclBxQwmH\ncVccRmH1rFWYmvXdqy1X2INoWAgZu8Aw+tnfWPabomavQZ6/DvnkSZq00zkSxF4k\noFYpsAjJepdyYzf8vMsREU8/C22eiQBJKxi5VZqWnAsbI/lB3UyiVlhc7wIDAQAB\nAoGAZXLtuULo0r0L32Y2mfX6ppw9Sr+8Acl6KDLQyajw8pijcOgpWQ52K6k0OK8e\n0jA7CyN6C/J2iqw6w/xpniRV2vtTpqT6AZiW0PdJzmIlwu+8AWrnnXQgfqcSChQ+\ny7xPl1BynRSr3w/jSVownTnHQ2xngie0OCxb5BVa+Xke+8ECQQDOa99RPevOghQE\nvu8YuW8f/SI+K9AuKpUeuXFfpAWFzRjgG5enJ/QnHOKo8JCRMOo0gOk0HTluj6EK\nhvG7DBshAkEAyA1Sm40H4Qtt14/odv7RTv6kVxxPPiQJCh9htKy04ykkeAxwC/Hk\nSxl6pOBwrmYb79wUQMqMQooKc+krGAAGDwJBAMUcYwocG/F9awpBHOW6JPAh8zH0\n+n0rMmw9XpKaeJ+VAtz13DHFSDKVNadm2FWcpPhv5MBb67y0sG3yACDB5aECQQCA\nBN6PQnH/VsQFBUebFrg1GAls4WXoe0D5pFlvOHJNSB/ZXwQ48KPKV1S/vAz/3cIU\n/b5MNBIL5rCHunfkVOeLAkA+sTHuJAy2fSzGizR9VK2zl7IPYTcK+7FxCfokBKIz\n+1t3tiJ/4LKijHe9bl2t5msOn9bZLATWR/W5nPBri1Ev\n-----END RSA PRIVATE KEY-----"

func InitRedis(redisConfig config.Redis) {
	Redis = redis2.GetRedis(redisConfig)
}
