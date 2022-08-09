package api

import (
	"SimpleDY/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var FavoriteService service.FavoriteService

// Favorite 点赞或取消赞操作
func Favorite(c *gin.Context) {
	//获取当前用户id
	strUserId := c.Query("user_id")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	strVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	strActionType := c.Query("action_type")
	actionType, _ := strconv.ParseInt(strActionType, 10, 64)

	FavoriteService.FavoriteAction(videoId, userId, int32(actionType))

}
func FavoriteList(c *gin.Context) {

}
