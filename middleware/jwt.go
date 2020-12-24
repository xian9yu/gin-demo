package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwt/models"
	"jwt/utils"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Invalid token")
	c                      = utils.NewCfg()
	cfg                    = c.InitConfig()
	SecretKey              = cfg.GetString("JWT.SecretKey") // 签名信息应该设置成动态从配置文件中获取
	expireTime             = cfg.GetInt64("JWT.ExpireAt")   // token有效期(时间戳/s)

)

// JWT基本数据结构
// 签名的signkey
type JWT struct {
	SecretKey []byte
}

// 定义载荷
type Claims struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}

// 初始化JWT实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(SecretKey),
	}
}

// token生成器
func (j *JWT) GenerateToken(c *gin.Context, user models.User) (string, error) {
	// 构造用户claims信息(负荷)
	claims := &Claims{
		Id:       user.ID,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + expireTime), // 签名过期时间(时间戳/s)
			Issuer:    user.UserName,                         // 签名颁发者
			IssuedAt:  time.Now().Unix(),                     //签名时间
		},
	}

	// 根据claims生成token对象
	return j.CreateToken(claims)

}

// 创建Token(基于用户的基本信息claims)
// 使用HS256算法进行token生成
// 使用用户基本信息claims以及签名key(signkey)生成token
func (j *JWT) CreateToken(u *Claims) (string, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#Token
	// 返回一个token的结构体指针
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	return token.SignedString(j.SecretKey)
}

//func (j *JWT) GenerateToken(u Claims) (string, error) {
//	claims := Claims{
//		Id:       u.Id,
//		UserName: u.UserName,
//
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: u.ExpiresAt, // 过期时间时间戳/秒  //在userLogin里定义
//			Issuer:    u.UserName,  //签名的发行者
//		},
//	}
//
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	token, err := tokenClaims.SignedString(SecretKey)
//	return token, err
//}

// token解析
// Couldn't handle this token:
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}

		}
	}

	// 将token中的claims信息解析出来和用户原始数据进行校验
	// 做以下类型断言，将token.Claims转换成具体用户自定义的Claims结构体
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, TokenInvalid

}

// 更新token TODO
//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//
//	token, _ := jwt.ParseWithClaims(
//		tokenString, &Claims{},
//		func(token *jwt.Token) (interface{}, error) {
//			return j.SecretKey, nil
//		})
//
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		//claims.StandardClaims.NotBefore = time.Now().Add(1 * time.Second).Unix()
//		claims.StandardClaims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
//		return j.CreateToken(claims)
//	}
//	return "", TokenInvalid
//}

// 退出登录
//func (j *JWT) LogoutToken(tokenString string) (string, error) {
//	// 拿到token基础数据
//	token, err := jwt.ParseWithClaims(
//		tokenString, &Claims{},
//		func(token *jwt.Token) (interface{}, error) {
//			return j.SecretKey, nil
//		},
//	)
//	claims, ok := token.Claims.(*Claims)
//	if !ok || !token.Valid {
//		return "", err
//	}
//
//	newClaims := Claims{
//		claims.Id,
//		claims.UserName,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Second * 1000).Unix(),
//			Issuer:    claims.UserName,
//			IssuedAt:  time.Now().Unix(),
//		},
//	}
//	tokens, err := j.CreateToken(newClaims)
//	if err != nil {
//		return "", err
//	}
//	return tokens, err
//
//}

//TODO 本来想实现刷新令牌的，但是不知道为什么调用这个方法后所有token都失效了，重新登录生成的token也是无效的
func (j *JWT) LogoutToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, _ := jwt.ParseWithClaims(
		tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.SecretKey, nil
		})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//claims.StandardClaims.NotBefore = time.Now().Add(1 * time.Second).Unix()
		claims.StandardClaims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
		return j.CreateToken(claims)
	}
	return "", TokenInvalid
}
