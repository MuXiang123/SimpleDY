// Package service
// @author    : MuXiang123
// @time      : 2022/7/7 21:11
package service

import (
	"SimpleDY/middleware/redis"
	"log"
	"strconv"
	"time"
)

type FavoriteService struct {
	VideoService
	UserService
}

// IsFavorite 根据userId和videoId查询点赞状态，先走redis，再走mysql
func (favorite *FavoriteService) IsFavorite(videoId int64, userId int64) (bool, error) {

}

// FavoriteAction 点赞或者取消赞操作
func (favorite *FavoriteService) FavoriteAction(videoId int64, userId int64, actionType int32) error {
	//转成字符串
	strUserId := strconv.FormatInt(userId, 10)
	strVideoId := strconv.FormatInt(videoId, 10)
	//点赞操作
	if actionType == 1 {
		//先走缓存，查看是否有对应的videoid和userid
		if n, err := redis.RdbLikeUserId.Exists(redis.Ctx, strUserId).Result(); n > 0 {
			if err != nil {
				log.Printf("function: FavoriteAction RdbLikeUserId query key fail: %v", err)
				return err
			}
			//查询到有这个userid，则将userid和videoId先添加到redis 再加入mysql
			if _, err := redis.RdbLikeVideoId.SAdd(redis.Ctx, strUserId, strVideoId).Result(); err != nil {
				log.Printf("function: FavoriteAction RdbLikeVideoId query key fail: %v", err)
				return err
			}
		} else {
			//缓存不存在userid 先将userid添加到redis，防止缓存击穿
			//设置value == -1，防止数据库出现脏读
			if _, err := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, -1).Result(); err != nil {
				log.Printf("function: FavoriteAction RdbLikeUserId add value fail: %v", err)
				redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
				return err
			}
			//设置过期时间
			if _, err := redis.RdbLikeUserId.Expire(redis.Ctx, strUserId,
				time.Duration(60*60*24*30)*time.Second).Result(); err != nil {
				log.Printf("function: FavoriteAction RdbLikeUserId set expire fail: %v", err)
				return err
			}
			//查userid的全部点赞视频
			dao
			videoIdList, err1 := dao.GetLikeVideoIdList(userId)
			if err1 != nil {
				return err1
			}

			//将当前userId点赞的全部视频都加入到redis，保证数据库和缓存一致性
			for _, vId := range videoIdList {
				if _, err := redis.RdbLikeVideoId.SAdd(redis.Ctx, strUserId, vId).Result(); err != nil {
					log.Printf("function: FavoriteAction RdbLikeVideoId add fail: %v", err)
					return err
				}
			}

			if _, err2 := redis.RdbLikeUserId.SRem(redis.Ctx, strUserId, videoId).Result(); err2 != nil {
				log.Printf("方法:FavouriteAction RedisLikeUserId del value失败：%v", err2)
				return err2
			}

		}
	} else if actionType == 2 {

	}
}
