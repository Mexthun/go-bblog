package service

import (
	"html/template"
	"me-bblog/common"
	"me-bblog/config"
	"me-bblog/dao"
	"me-bblog/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, errOfGetPost := dao.GetPostById(pid)
	if errOfGetPost != nil {
		return nil, errOfGetPost
	}
	username, errOfGetUserName := dao.GetUserNameById(post.UserId)
	if errOfGetUserName != nil {
		return nil, errOfGetUserName
	}
	categoryName, errOfGetCategoryName := dao.GetCategoryNameById(post.CategoryId)
	if errOfGetCategoryName != nil {
		return nil, errOfGetCategoryName
	}
	var postRes = &models.PostRes{}
	postRes.Article = models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		username,
		post.ViewCount,
		post.Type,
		common.DateDay(post.CreateAt),
		common.DateDay(post.UpdateAt),
	}

	postRes.Viewer = config.Cfg.Viewer
	postRes.SystemConfig = config.Cfg.System
	return postRes, nil

}
func SavePost(post *models.Post) error {
	err := dao.SavePost(post)
	if err != nil {
		return err
	}
	return nil
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
func DeletePostByPostId(pid int) error {
	err := dao.DeletePostByPostId(pid)
	return err
}

func GetPost(pid int) (post *models.Post, err error) {
	post, err = dao.GetPost(pid)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResps
}
func Writing() (*models.WritingRes, error) {
	var wr = &models.WritingRes{}
	wr.CdnURL = config.Cfg.System.CdnURL
	wr.Title = config.Cfg.Viewer.Title
	category, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	wr.Categorys = category
	return wr, nil
}
