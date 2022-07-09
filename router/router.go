package router

import (
	"SimpleDY/api"
	"SimpleDY/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20 //短视频大小最大不超过 10M
	//主路由
	douyinGroup := r.Group("/douyin")
	{
		//用户操作路由
		userGroup := douyinGroup.Group("/user")
		{
			userGroup.POST("/register/", api.Register)
			userGroup.POST("/login/", api.Login)
			userGroup.GET("/", jwt.JwtMiddleWare(), api.GetInfo)
		}
		//发布视频
		publishGroup := douyinGroup.Group("/publish")
		{
			publishGroup.POST("/action/", jwt.JwtMiddleWare(), api.Publish)
			publishGroup.GET("/list/", jwt.JwtMiddleWare(), api.GetPublishListByAuthorId)
		}
		//feed
		r.GET("/feed/", api.Feed)

		//关注
		relationGroup := douyinGroup.Group("/relation")
		{
			relationGroup.POST("/action/", jwt.JwtMiddleWare(), api.RelationAction)
			relationGroup.GET("/follow/list/", jwt.JwtMiddleWare(), api.FollowList)
			relationGroup.GET("/follower/list/", jwt.JwtMiddleWare(), api.FollowerList)
		}
		//点赞
		favoriteGroup := douyinGroup.Group("favorite")
		{
			favoriteGroup.POST("/action/", jwt.JwtMiddleWare(), api.Favorite)
			favoriteGroup.GET("/list/", jwt.JwtMiddleWare(), api.FavoriteList)
		}

		// 评论
		douyinGroup.Group("/comment")
		{
			//commentGroup.POST("/action/", middleware.JwtMiddleware(), api.CommentAction)
			//commentGroup.GET("/list/", middleware.JwtMiddleware(), api.CommentList)
		}
	}

	return r
}
