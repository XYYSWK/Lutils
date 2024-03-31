package token

import (
	"errors"
	"time"
)

var ErrSecretLen = errors.New("密钥长度不正确")

type MakerToken interface {
	// CreateToken 生成 MakerToken
	CreateToken(content []byte, expireDate time.Duration) (string, *Payload, error)
	// VerifyToken 解析 MakerToken
	VerifyToken(token string) (*Payload, error)
}
