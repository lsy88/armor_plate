package jwt

import (
	"armor_plate/core"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	EmployeeID   uint64 `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	//AuthorityId  uint   //角色类别
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24 * 365

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(core.Conf.Jwt.SigningKey),
	}
}

//生成token
func (j *JWT) GenToken(employeeName string) (aToken, rToken string, err error) {
	//创建一个我们自己的声明
	claims := CustomClaims{
		EmployeeName: employeeName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "armor_plate",                              // 签发人
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
		Issuer:    "armor_plate",                           // 签发人
	}).SignedString(j.SigningKey)
	return
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			return nil, TokenMalformed
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			//token已经到期
			return nil, TokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			//token还未激活
			return nil, TokenNotValidYet
		} else {
			return nil, TokenInvalid
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

//刷新token
func (j *JWT) RefreshToken(token string) (aToken, rToken string, err error) {
	//从旧的token中解析出claims数据
	var claims CustomClaims
	_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	return j.GenToken(claims.EmployeeName)
}
