// Package service
// @author    : MuXiang123
// @time      : 2022/6/11 17:09
package service

import (
	"SimpleDY/dao"
	"SimpleDY/global"
	"SimpleDY/status"
	"errors"
	"gorm.io/gorm"
)

type FollowingService struct {
}

// InsertFollowing 插入关注
func InsertFollowing(following *dao.Follow) error {
	err := global.Db.Model(&dao.Follow{}).Create(&following).Error
	if err != nil {
		return err
	}
	return nil
}

// AddFollowCount 增加关注数
func AddFollowCount(hostId uint64) error {
	err := global.Db.Model(&dao.User{}).
		Where("id = ?", hostId).
		Update("follow_count", gorm.Expr("follow_count + ?", 1)).
		Error
	if err != nil {
		return err
	}
	return nil
}

// AddFollowerCount 增加粉丝数
func AddFollowerCount(guestId uint64) error {
	err := global.Db.Model(&dao.User{}).
		Where("id = ?", guestId).
		Update("follower_count", gorm.Expr("follower_count + ?", 1)).
		Error
	if err != nil {
		return nil
	}
	return err
}

// DecreaseFollowerCount 减少粉丝数
func DecreaseFollowerCount(guest uint64) error {
	err := global.Db.Model(&dao.User{}).Where("id = ?", guest).
		Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		return nil
	}
	return err
}

// DecreaseFollowCount 减少关注数
func DecreaseFollowCount(hostId uint64) error {
	err := global.Db.Model(&dao.User{}).Where("id = ?", hostId).
		Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		return nil
	}
	return err
}

// DeleteFollowing 删除关注表记录
func DeleteFollowing(following *dao.Follow) error {
	err := global.Db.Model(&dao.Follow{}).
		Where("host_id=? AND guest_id=?", following.HostId, following.GuestId).
		Delete(&following).Error
	if err != nil {
		return err
	}
	return nil
}

// IsFollowing 判断host是否关注guest
func IsFollowing(hostId uint64, guestId uint64) bool {
	var relationExist = &dao.Follow{}
	//判断关注是否存在
	err := global.Db.Model(&dao.Follow{}).
		Where("host_id=? AND guest_id=?", hostId, guestId).
		First(&relationExist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//关注不存在
		return false
	}
	//关注存在
	return true

}

// FollowAction 关注和取消关注操作
func (followingservice FollowingService) FollowAction(hostId uint64, guestId uint64, actionType uint) (int64, error) {
	newFollowing := dao.Follow{
		HostId:  hostId,
		GuestId: guestId,
	}
	if actionType == 1 {
		//判断关注是否存在，避免重复关注
		if IsFollowing(hostId, guestId) == true {
			return status.AttentionExistsError, errors.New("关注已存在")
		} else {
			//关注不存在,创建关注
			errTransaction := global.Db.Transaction(func(db *gorm.DB) error {
				//添加关注
				err1 := InsertFollowing(&newFollowing)
				if err1 != nil {
					return err1
				}
				//更新关注数
				err2 := AddFollowCount(hostId)
				if err2 != nil {
					return err2
				}
				//更新粉丝数
				err3 := AddFollowerCount(guestId)
				if err3 != nil {
					return err3
				}
				return nil
			})
			if errTransaction != nil {
				return status.UnknownError, errTransaction
			}
		}
	}
	// 取消关注操作
	if actionType == 2 {
		//判断关注是否存在
		if IsFollowing(hostId, guestId) == true {
			//关注存在,取消关注，先删除关注表的记录，再减少关注数和粉丝数
			errTransaction := global.Db.Transaction(func(db *gorm.DB) error {
				err1 := DeleteFollowing(&newFollowing)
				if err1 != nil {
					return err1
				}
				//减少host_id的关注数
				err2 := DecreaseFollowCount(hostId)
				if err2 != nil {
					return err2
				}
				//减少guest_id的粉丝数
				err3 := DecreaseFollowerCount(guestId)
				if err3 != nil {
					return err3
				}
				return nil
			})
			if errTransaction != nil {
				return status.UnknownError, errTransaction
			}
		} else {
			//关注不存在
			return status.AttentionNullError, nil
		}
	}
	return status.Success, nil
}

// GetFollowingList GetFollowing 获取关注列表
func (followingservice FollowingService) GetFollowingList(hostId uint64) ([]dao.User, error) {
	//用户列表
	var userList []dao.User
	//粉丝集合
	var guestList []uint64
	//子查询，先查关注表，再查用户表，获取所有粉丝
	errGuestList := global.Db.Model(&dao.Follow{}).Select("guest_id").Where("host_id = ?", hostId).Scan(&guestList).Error
	if errGuestList != nil {
		return userList, nil
	}
	errUserList := global.Db.Model(&dao.User{}).Where("id IN ?", guestList).Scan(&userList).Error
	if errUserList != nil {
		return userList, errUserList
	}
	return userList, nil
}

// GetFollowerList 获取粉丝列表
func (followingservice FollowingService) GetFollowerList(guest uint64) ([]dao.User, error) {
	//用户列表
	var userList []dao.User
	//粉丝集合
	var hostList []uint64
	//子查询，先查关注表，再查用户表，获取所有关注
	errHostList := global.Db.Model(&dao.Follow{}).Select("host_id").Where("guest_id = ?", guest).Scan(&hostList).Error
	if errHostList != nil {
		return userList, nil
	}
	errUserList := global.Db.Model(&dao.User{}).Where("id IN ?", hostList).Scan(&userList).Error
	if errUserList != nil {
		return userList, errUserList
	}
	return userList, nil
}
