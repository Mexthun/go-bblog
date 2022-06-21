package dao

import (
	"me-bblog/models"
)

func GetUserByUsernameAndPasswd(username, passwd string) (*models.Users, error) {
	var user = &models.Users{}
	err := db.Table("blog_user").Where("user_name=? and passwd =?", username, passwd).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserNameById(id int) (string, error) {
	rows, err := db.Raw("select user_name from blog_user where uid =?", id).Rows()
	if err != nil {
		return "", err
	}

	var name string
	for rows.Next() {

		db.ScanRows(rows, &name)
	}

	return name, err
}
