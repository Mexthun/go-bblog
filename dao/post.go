package dao

import (
	"database/sql"
	"log"
	"me-bblog/models"
	"strconv"
)

func GetPagePost(page, pageSize int) ([]*models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := db.Raw("select * from blog_post limit ?,?", page, pageSize).Rows()
	if err != nil {
		return nil, err
	}
	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		db.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return posts, nil
}

func GetAllPosts() ([]*models.Post, error) {
	rows, err := db.Raw("select * from blog_post").Rows()
	if err != nil {
		return nil, err
	}
	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		db.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return posts, nil
}

func SavePost(post *models.Post) (err error) {

	_, err = db.Raw("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	).Rows()

	if err != nil {
		return err
	}
	post.Pid, err = getLastPid()
	if err != nil {
		return err
	}
	return nil
}

func DeletePostByPostId(pid int) error {
	var post = &models.Post{
		Pid: pid,
	}
	err := db.Table("blog_post").Where("pid", pid).Delete(post).Error
	return err
}

func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := db.Raw("select * from blog_post where title like ?", "%"+condition+"%").Rows()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPost(pid int) (*models.Post, error) {
	rows, err := db.Raw("select * from blog_post where pid = ?", pid).Rows()
	var post = &models.Post{}

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		db.ScanRows(rows, post)
	}
	return post, nil
}
func getLastPid() (int, error) {
	row, err := db.Raw("select pid from blog_post order by pid desc limit 1").Rows()
	if err != nil {
		return 0, err
	}
	var pidStr string
	for row.Next() {
		db.ScanRows(row, &pidStr)
	}
	pid, errOfStrToInt := strconv.Atoi(pidStr)
	if errOfStrToInt != nil {
		return 0, errOfStrToInt
	}
	return pid, nil

}
func UpdatePost(post *models.Post) {

	_, err := db.Raw("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	).Rows()
	if err != nil {
		log.Println(err)
	}
}
func GetCountAllPostsWithSlug(slug string) (int, error) {
	var err error
	var rows *sql.Rows
	if slug == "" {
		rows, err = db.Raw("select count(*) from blog_post").Rows()
	} else {
		rows, err = db.Raw("select count(*) from blog_post where slug=?", slug).Rows()
	}

	if err != nil {
		return 0, err
	}
	var totalStr string
	for rows.Next() {
		db.ScanRows(rows, &totalStr)
	}
	total, err := strconv.Atoi(totalStr)
	if err != nil {
		return 0, err
	}
	return total, nil

}

func GetPostById(pid int) (*models.Post, error) {
	rows, err := db.Raw("select * from blog_post where pid =?", pid).Rows()
	if err != nil {
		return nil, err
	}
	var post models.Post
	for rows.Next() {
		db.ScanRows(rows, &post)
	}
	return &post, nil
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := db.Raw("select count(1) from blog_post where category_id=?", cId)
	_ = rows.Scan(&count)

	return
}

func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := db.Raw("select * from blog_post where category_id = ? limit ?,?", cId, page, pageSize).Rows()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
