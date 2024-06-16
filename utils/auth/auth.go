package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"kube-resource-manager/internal/db/models"
	"net/http"
	"strings"
	"time"
)

// HashPassword  takes a plain text password and returns its bcrypt hash.
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// ComparePasswords  compares a plain text password with a bcrypt hashed password.
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

var jwtSecret = []byte("kube-resource-manager")

func GenerateToken(user *models.User) (string, error) {
	claims := CreateClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func CreateClaims(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"UserId":   user.ID,
		"Username": user.Username,
		"Exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为24小时后
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供Token"})
			c.Abort()
			return
		}

		// 验证JWT令牌
		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// 检查签名方法是否正确以及使用正确的秘钥
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil // 这里jwtSecret是从之前定义的全局变量获取的
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效"})
			c.Abort()
			return
		}

		// 检查Token是否解析成功且有效
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将claims附加到Context中，以便后续处理可以访问
			c.Set("claims", claims)
			c.Next() // 如果验证通过，则继续处理请求
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效或已过期"})
			c.Abort()
		}
	}
}
