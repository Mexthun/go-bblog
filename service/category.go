package service

import (
	"html/template"
	"me-bblog/common"
	"me-bblog/config"
	"me-bblog/dao"
	"me-bblog/models"
)

func GetPostsByCategoryIdOrCategoryName(cName string, cId, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, errOfeGetCategorys := dao.GetAllCategory()
	if errOfeGetCategorys != nil {
		return nil, errOfeGetCategorys
	}
	var posts []models.Post
	var err error
	if cName != "" {
		cId, err = dao.GetCategoryIdByCategoryName(cName)
		if err != nil {
			return nil, err
		}

	}

	posts, err = dao.GetPostPageByCategoryId(cId, page, pageSize)

	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName, errOfGetCategoryNameById := dao.GetCategoryNameById(post.CategoryId)
		if errOfGetCategoryNameById != nil {
			return nil, errOfGetCategoryNameById
		}
		userName, errOfGetGetUserNameById := dao.GetUserNameById(post.UserId)
		if errOfGetGetUserNameById != nil {
			return nil, errOfGetGetUserNameById
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			common.DateDay(post.CreateAt),
			common.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//11  10 2  10 1 9 1  21 3
	//  (11-1)/10 + 1 = 2
	total := dao.CountGetAllPostByCategoryId(cId)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	categoryName, errOfGetCategoryNameById := dao.GetCategoryNameById(cId)
	if errOfGetCategoryNameById != nil {
		return nil, errOfGetCategoryNameById
	}
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
