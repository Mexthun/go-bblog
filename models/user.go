package models

import "time"

type Users struct {
	Uid      int
	UserName string
	Passwd   string
	Avatar   string
	CreateAt time.Time
	Update   time.Time
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
