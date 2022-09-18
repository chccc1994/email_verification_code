package services

import (
	"email_code/models"
	"email_code/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email ==""{
		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"Parameter error",
		})
		return
	}

	code :=utils.GetRand()
	models.InitRedisDB().Set(c,email,code,time.Second*120)

	err:=utils.SendCode(email,code)
	if err!=nil{
		log.Println("Send code error",err.Error())
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"Send code error"+err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"Sending succeeded",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param mail formData string true "mail"
// @Param code formData string true "code"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func Register(c *gin.Context){
	// 用户名，密码，邮箱，手机
	mail := c.PostForm("mail")
	userCode := c.PostForm("code")
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	if mail =="" ||userCode==""||
	name == "" ||password ==""{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"Parameter error",
		})
		return 
	}
	sysRDBCode,err :=models.InitRedisDB().Get(c,mail).Result()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"Get Code Error:"+ err.Error(),
		})
		return
	}
	if sysRDBCode !=userCode{
		c.JSON(http.StatusOK,gin.H{
			"code":-1,
			"msg":"Verification code mismatch",
		})
		return
	}

	var cnt int64

	err  =models.DB.Where("mail=?",mail).Model(new(models.User)).Count(&cnt).Error
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get User Error:" + err.Error(),
		})
		return
	}
	if cnt>0{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该邮箱已被注册",
		})
		return
	}

	userIdentity := utils.GetUUID()

	data := &models.User{
		Identity:  userIdentity,
		Name:      name,
		Password:  utils.GetMd5(password),
		Phone:     phone,
		Mail:      mail,
	}

	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Crete User Error:" + err.Error(),
		})
		return
	}

	token,err:=utils.GenerateToken(userIdentity,name)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Generate Token Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}