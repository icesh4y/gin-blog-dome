package controller

import (
	"essential/dto"
	"essential/jwt"
	"essential/models"
	"essential/response"
	"essential/util"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	phone := c.PostForm("phone")
	if len(phone) != 11 {
		c.JSON(http.StatusForbidden, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}

	if models.SelectUserPhone(phone) {
		response.Response(c, http.StatusForbidden, 403, nil, "手机号已存在")
		//c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "手机号已存在"})
		return
	}

	if len(pwd) < 6 {
		response.Response(c, http.StatusForbidden, 403, nil, "密码长度必须大于6位")
		//c.JSON(http.StatusFailedDependency, gin.H{"code": 422, "msg": "密码长度必须大于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandString(10)
		//c.JSON(http.StatusOK, gin.H{"code": 200, "msg": name})
	}
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密错误")
		//c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密错误"})
		return
	}
	models.CreateNewUser(name, phone, string(hasePassword))

	response.Success(c, nil, "注册成功")
	//c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
	log.Println(name, pwd, phone)
}

func Login(c *gin.Context) {
	pwd := c.PostForm("pwd")
	phone := c.PostForm("phone")
	if len(phone) != 11 {
		response.Response(c, http.StatusInternalServerError, 422, nil, "手机号必须是11位")
		//c.JSON(http.StatusForbidden, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}
	if len(pwd) < 6 {
		response.Response(c, http.StatusInternalServerError, 422, nil, "密码长度必须大于6位")
		//c.JSON(http.StatusFailedDependency, gin.H{"code": 422, "msg": "密码长度必须大于6位"})
		return
	}
	user := models.SelectUser(phone)
	if user.ID == 0 {
		response.Response(c, http.StatusInternalServerError, 422, nil, "用户不存在")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 402, "msg": "用户不存在"})
		return
	}
	fmt.Println(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd))
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 422, nil, "密码错误")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 402, "msg": "密码错误"})
		return
	}
	token, err := jwt.ReleaseToken(*user)
	if err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "异常")
		//c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "异常"})
		return
	}
	response.Response(c, http.StatusOK, 200, gin.H{"data": gin.H{"user": user}, "token": token}, "登录成功")

	//c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}, "token": token})
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	//log.Println(c.Get("user"))
	response.Response(c, http.StatusOK, 200, gin.H{"code": 200, "data": gin.H{"user": dto.UserToDto(user.(models.User))}}, "")
	//c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.UserToDto(user.(models.User))}})
}
