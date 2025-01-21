package api

import (
	"ginBlog/models"
	"ginBlog/pkg/e"
	"ginBlog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAuth(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}

	a := models.Auth{Username: userName, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]any)
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(userName, password)
		if isExist {
			token, err := util.GenerateToken(userName, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
