package request

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims structure
type CustomClaims struct {
	UUID        string
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
