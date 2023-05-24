package bll

import (
	"categorymanagement/api/dal"
	"categorymanagement/models"
	"database/sql"
)

func Insertcategory(data models.Categories, db *sql.DB, id string) error {
	err := dal.AddCategory(data, db, id)
	return err
}

func UpdateCategory(db *sql.DB, newType string, newName string) error {
	err := dal.UpdateCategory(db, newType, newName)
	return err
}

func DeleteCategory(db *sql.DB, newName string) error {
	err := dal.DeleteCategory(db, newName)
	return err
}

func GetCategories(db *sql.DB, name string) ([]models.Types, error) {
	data, err := dal.GetCategories(db, name)
	return data, err
}
