package favorite

import (
	"douyin-lite/service/favorite_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FavoriteAuthor struct { //从user中获取,getUser函数
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"` //从following或follower中获取
}

type FavoriteVideo struct { //从video中获取
	Id            uint           `json:"id,omitempty"`
	Author        FavoriteAuthor `json:"author,omitempty"`
	PlayUrl       string         `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	FavoriteCount uint           `json:"favorite_count,omitempty"`
	CommentCount  uint           `json:"comment_count,omitempty"`
	IsFavorite    bool           `json:"is_favorite,omitempty"` //true
	Title         string         `json:"title,omitempty"`
}

type FavoriteListResponse struct {
	Response
	VideoList []FavoriteVideo `json:"video_list,omitempty"`
}

// 点赞视频
func FavoriteAction(c *gin.Context) {
	//参数绑定
	//user_id获取
	getUserId, _ := c.Get("user_id")
	var userId uint
	if v, ok := getUserId.(uint); ok {
		userId = v
	}
	//参数获取
	actionTypeStr := c.Query("action_type")
	actionType, _ := strconv.ParseUint(actionTypeStr, 10, 10)
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseUint(videoIdStr, 10, 10)

	//函数调用及响应
	err := favorite_service.Favorite_Action(userId, uint(videoId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "操作成功！",
		})
	}
}
