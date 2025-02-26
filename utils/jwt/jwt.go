package jwt

import (
	"BriefChat/global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

var JwtSecret []byte

type Claims struct {
	Account string `json:"account"`
	//Password string `json:"password"`
	jwt.MapClaims
}

const accessTokenExpiration = 7 * 24 * time.Hour

func GenerateToken(account, password string) (string, error) {
	currentTime := time.Now()
	expiredTime := currentTime.Add(accessTokenExpiration)

	claims := Claims{
		account,
		//password,
		jwt.MapClaims{
			"exp": expiredTime.UnixMilli(),
			//"role": role,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用 HS256 签名算法生成 token
	token, err := tokenClaims.SignedString(JwtSecret)
	// 返回签名后的 token 字符串
	jwtToken := "Bearer " + token
	return jwtToken, err
}

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header required"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			// 验证 JWT 使用的签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// 将 JWT 中的用户名存储到上下文中
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("account", claims.Account) // 将用户名存储到上下文中
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}

}

func ParseToken(token string) (*Claims, error) {
	token = strings.Replace(token, global.JWT_TOKEN_PREFIX, "", -1)
	log.Printf("Token: %s\n", token)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
