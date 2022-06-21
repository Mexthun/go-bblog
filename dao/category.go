package dao

import (
	"me-bblog/models"
	"strconv"
)

func GetAllCategory() ([]models.Category, error) {

	rows, err := db.Raw("select * from blog_category").Rows()
	if err != nil {
		return nil, err
	}
	var categorys []models.Category

	for rows.Next() {
		db.ScanRows(rows, &categorys)
	}
	return categorys, nil
}

func GetCategoryNameById(id int) (string, error) {
	rows, err := db.Raw("select name from blog_category where cid=?", id).Rows()
	if err != nil {
		return "nil", err
	}
	var name string
	for rows.Next() {
		db.ScanRows(rows, &name)
	}
	return name, nil

}
func GetCategoryIdByCategoryName(cname string) (int, error) {
	rows, err := db.Raw("select cid from blog_category where name=?", cname).Rows()
	if err != nil {
		return 0, err
	}
	var cid string
	for rows.Next() {
		db.ScanRows(rows, &cid)
	}

	cidInt, errOfGetCid := strconv.Atoi(cid)
	if errOfGetCid != nil {
		return 0, errOfGetCid
	}
	return cidInt, nil

}
