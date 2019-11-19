package model

import (
	"github.com/jinzhu/gorm"

	"github.com/joshqu1985/fireman/pkg/utime"
	"github.com/joshqu1985/protos/pkg/model"
)

type Follower struct {
	model.Base
	Uid string `json:"uid"`
	Fid string `json:"fid"`
}

func (f *Follower) TableName() string {
	return "followers"
}

func (f *Follower) BeforeCreate(scope *gorm.Scope) error {
	now := utime.Millisec()
	scope.SetColumn("CreatedTime", now)
	scope.SetColumn("UpdatedTime", now)
	return nil
}

type Following struct {
	model.Base
	Uid string `json:"uid"`
	Fid string `json:"fid"`
}

func (f *Following) TableName() string {
	return "followings"
}

func (f *Following) BeforeCreate(scope *gorm.Scope) error {
	now := utime.Millisec()
	scope.SetColumn("CreatedTime", now)
	scope.SetColumn("UpdatedTime", now)
	return nil
}

type Counter struct {
	model.Base
	FollowerCount  int `json:"follower_count"`
	FollowingCount int `json:"following_count"`
}

func (c *Counter) TableName() string {
	return "counters"
}

func (c *Counter) BeforeCreate(scope *gorm.Scope) error {
	now := utime.Millisec()
	scope.SetColumn("CreatedTime", now)
	scope.SetColumn("UpdatedTime", now)
	return nil
}
