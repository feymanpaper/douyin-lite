package favorite

import (
	"douyin-lite/repository"
	"douyin-lite/service/favorite_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	UserDao = repository.NewUserDaoInstance()
)

// FavoriteList 获取列表方法
func FavoriteList(c *gin.Context) {
	//user_id获取
	getUserId, _ := c.Get("user_id")
	var userIdHost uint
	if v, ok := getUserId.(uint); ok {
		userIdHost = v
	}
	userIdStr := c.Query("user_id") //自己id或别人id
	userId, _ := strconv.ParseUint(userIdStr, 10, 10)
	userIdNew := uint(userId)
	if userIdNew == 0 {
		userIdNew = userIdHost
	}

	//函数调用及响应
	videoList, err := favorite_service.Favorite_List(userIdNew)
	videoListNew := make([]FavoriteVideo, 0)
	for _, m := range videoList {
		var author = FavoriteAuthor{}
		// var getUser = repository.User{}
		user, err := UserDao.QueryUserById(m.AuthorId)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 403,
				StatusMsg:  "找不到作者！",
			})
			c.Abort()
			return
		}
		//isfollowing
		isfollowing, err := repository.NewFollowingDaoInstance().QueryisFollow(userIdHost, m.AuthorId)
		//isfavorite
		isfavorite := favorite_service.Check_Favorite(userIdHost, m.ID)
		//作者信息
		author.Id = user.ID
		author.Name = user.Name
		author.FollowCount = user.FollowingCount
		author.FollowerCount = user.FollowerCount
		author.IsFollow = isfollowing
		//组装
		var video = FavoriteVideo{}
		video.Id = m.ID //类型转换
		video.Author = author
		video.PlayUrl = m.PlayUrl
		video.CoverUrl = m.CoverUrl
		video.FavoriteCount = m.FavoriteCount
		video.CommentCount = m.CommentCount
		video.IsFavorite = isfavorite
		video.Title = m.Title

		videoListNew = append(videoListNew, video)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "查找失败！",
			},
			VideoList: nil,
		})
	} else {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "找到列表！",
			},
			VideoList: videoListNew,
		})
	}
}
