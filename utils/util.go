package utils

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

func SendCode(toUserEmail, code string)error{
	e := email.NewEmail()
	e.From = "Get <ch_sheng5@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	// 返回 EOF 时，关闭SSL重试
	// password开启STMP的授权码
	log.Println(EmailPassword)
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "ch_sheng5@163.com", EmailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
// 生成随机code
func GetRand()string{
	rand.Seed(time.Now().UnixNano())
	s:=""
	for i:=0;i<6;i++{
		s +=strconv.Itoa(rand.Intn(10))
	}
	return s
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("email-code")

// GenerateToken
// 生成 token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}