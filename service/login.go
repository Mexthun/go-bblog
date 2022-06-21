package service

import (
	"errors"
	"me-bblog/dao"
	"me-bblog/models"
	"me-bblog/utils"
)

func Login(username, passwd string) (*models.LoginRes, error) {

	user, err := dao.GetUserByUsernameAndPasswd(username, passwd)
	if err != nil {
		return nil, errors.New("账户名或密码错误")
	}

	var loginRes = &models.LoginRes{}
	token, err := utils.Award(&user.Uid)
	if err != nil {
		return nil, errors.New("生成token失败")
	}
	loginRes.Token = token
	loginRes.UserInfo.Uid = user.Uid
	loginRes.UserInfo.Avatar = user.Avatar
	loginRes.UserInfo.UserName = user.UserName

	return loginRes, nil

}
