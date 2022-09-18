# 模拟验证码功能

[![build workflow](https://github.com/go-redis/redis/actions/workflows/build.yml/badge.svg)](https://github.com/go-redis/redis/actions/workflows/build.yml/badge.svg) [![PkgGoDev](https://camo.githubusercontent.com/4917695de7771a4295a6fdfd7105b904cf1ebb9b3056b277736a49f036ad8d3b/68747470733a2f2f706b672e676f2e6465762f62616467652f6769746875622e636f6d2f676f2d72656469732f72656469732f7638)](https://pkg.go.dev/github.com/go-redis/redis/v8?tab=doc) [![Documentation](https://camo.githubusercontent.com/7692019ac4eff10a035bdfa643ca2e90eb68f34120ac94264a0b5bf4f05edddf/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f72656469732d646f63756d656e746174696f6e2d696e666f726d6174696f6e616c)](https://redis.uptrace.dev/)https://godoc.org/github.com/jordan-wright/email)

## 前提

> golang 1.19
>
> mysql 8.0.24 
>
> redis-v8 : go get github.com/go-redis/redis/v8
>
> gin:  go get github.com/gin-gonic/gin 
>
> gorm : gorm.io/gorm

## 要求

1. 实现Email六位验证码；
2. 时效2分钟；

## 实现

### 结构

```bash
|  go.mod
│  go.sum
│  main.go
│
├─config
│      config.ini # 配置文件，读取数据库、路由、email授权码等
│
├─docs	# swagger 文件
│      docs.go
│      swagger.json
│      swagger.yaml
│
├─models	
│      init.go 	# 数据库初始化，mysql,redisDB
│      user.go	# 用户信息
│
├─routers
│      router.go # 路由，
│
├─services
│      ping.go # 测试路由接口
│      user.go # 用户路由，发送验证码以及用户注册
│
├─test	# 测试
└─utils
        setting.go 	# 配置文件读取
        util.go 	# 公共 


```

### 数据库

+ 用户

> gorm.io/gorm

```go
type User struct {
	gorm.Model
	Identity  string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
	Name      string `gorm:"column:name;type:varchar(100);" json:"name"`        // 用户名
	Password  string `gorm:"column:password;type:varchar(32);" json:"-"`        // 密码
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`       // 手机号
	Mail      string `gorm:"column:mail;type:varchar(100);" json:"mail"`        // 邮箱
}
```

### 路由

> github.com/gin-gonic/gin 
>
>  swaggerfiles "github.com/swaggo/files"
>
>  ginSwagger "github.com/swaggo/gin-swagger"

```go
r.POST("/send-code", services.SendCode) // 发送验证码
r.POST("/register", services.Register)	// 用户注册
```

### 验证🐎

> github.com/dgrijalva/jwt-go
>
> github.com/jordan-wright/email 
>
> github.com/satori/go.uuid



```go
func SendCode(){
    e := email.NewEmail()
	e.From = "Get <发送邮件邮箱账户>"
	e.To = []string{"接收邮件邮箱账户"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码：<b>123456</b>")
	// 返回 EOF 时，关闭SSL重试
	// password开启STMP的授权码
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "发送邮件邮箱账户", "发送邮件邮箱账户STMP授权码", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
```

### 常用邮箱服务器地址与端口

+ 163邮箱

| 收件服务器 | POP        | pop.163.com  | 995           | 110             |
| ---------- | ---------- | ------------ | ------------- | --------------- |
| 收件服务器 | IMAP       | imap.163.com | 993           | 143             |
| 类型       | 服务器名称 | 服务器地址   | SSL协议端口号 | 非SSL协议端口号 |
| 发件服务器 | SMTP       | smtp.163.com | 465/994       | 25              |

+ qq邮箱

| 类型       | 服务器名称 | 服务器地址  | 非SSL协议端口号 | SSL协议端口号 |
| ---------- | ---------- | ----------- | --------------- | ------------- |
| 发件服务器 | SMTP       | smtp.qq.com | 25              | 465/587       |
| 收件服务器 | POP        | pop.qq.com  | 110             | 995           |
| 收件服务器 | IMAP       | imap.qq.com | 143             | 993           |

+ gmail

```bash
gmail(google.com) ：

POP3服务器地址:pop.gmail.com（SSL启用 端口：995） 

SMTP服务器地址:smtp.gmail.com（SSL启用 端口：587） 
```

