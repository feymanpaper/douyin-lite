package main

import (
	"douyin-lite/handler"
	"douyin-lite/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.GET("/douyin/relation/follow/list/query/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//tokenStr := c.Param("token")
		followListResp, err := handler.QueryFollowListHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, followListResp)
			return
		}
		c.JSON(http.StatusOK, followListResp)
	})

	r.GET("/douyin/relation/follower/list/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//tokenStr := c.Param("token")
		followListResp, err := handler.QueryFollowerListHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, followListResp)
			return
		}
		c.JSON(http.StatusOK, followListResp)
	})

	r.POST("/douyin/user/register/", func(c *gin.Context) {
		userName := c.Query("username")
		userPassword := c.Query("password")
		registerResp, err := handler.RegisterUserHandler(userName, userPassword)
		if err != nil {
			c.JSON(http.StatusOK, registerResp)
			return
		}
		c.JSON(http.StatusOK, registerResp)
	})

	r.POST("/douyin/user/login/", func(c *gin.Context) {
		userName := c.Query("username")
		userPassword := c.Query("password")
		loginResp, err := handler.LoginUserHandler(userName, userPassword)
		if err != nil {
			c.JSON(http.StatusOK, loginResp)
			return
		}
		c.JSON(http.StatusOK, loginResp)
	})

	r.GET("/douyin/user/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//userToken := c.Query("token")
		userInfoResp, err := handler.QueryUserInfoHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, userInfoResp)
			return
		}
		c.JSON(http.StatusOK, userInfoResp)
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}
