package model

import (
	"selefra-demo/internal/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	u := model.User{}
	u.CreateUser(&u)
}
