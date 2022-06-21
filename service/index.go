package service

import (
	"fmt"
	"html/template"
	"me-bblog/common"
	"me-bblog/config"
	"me-bblog/dao"
	"me-bblog/models"
)

func GetAllIndexInfo(slug string, page int, pageSize int) (*models.HomeResponse, error) {
	categorys, errOfGetCategorys := dao.GetAllCategory()

	if errOfGetCategorys != nil {
		return nil, errOfGetCategorys
	}
	fmt.Println(categorys)
	var posts []*models.Post
	var err error
	if slug == "" {
		posts, err = dao.GetPagePost(page, pageSize)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(posts)
	fmt.Println(slug)

	total, errCountPosts := dao.GetCountAllPostsWithSlug(slug)
	if errCountPosts != nil {
		return nil, errCountPosts
	}
	fmt.Println(total)

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName, _ := dao.GetCategoryNameById(post.CategoryId)

		userName, _ := dao.GetUserNameById(post.UserId)
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
	return hr, nil

}
