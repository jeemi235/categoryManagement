package dal

import (
	"categorymanagement/models"
	"database/sql"
	"errors"
	"log"
	"time"
)

func AddCategory(data models.Categories, db *sql.DB, id string) error {
	sqlQuery := `INSERT INTO category(type,name,parent_id,icon,user_id,created_at) values($1,$2,$3,$4,$5,$6)`

	_, err := db.Exec(
		sqlQuery,
		data.Type,
		data.Name,
		data.ParentId,
		data.Icon,
		id,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(db *sql.DB, newType string, newName string) error {
	sqlQuery := `UPDATE category set type=$1, updated_at=$2 Where name=$3 `

	_, err := db.Exec(
		sqlQuery,
		newType,
		time.Now(),
		newName,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, newName string) error {
	sqlQuery := `delete from category where name=$1`

	rows, err := db.Exec(
		sqlQuery,
		newName,
	)
	if err != nil {
		return err
	}

	// This will count the number of rows present which we want to delete
	row, _ := rows.RowsAffected()

	//If rows are 0 will give error
	if row == 0 {
		err = errors.New("category does not exists")
		return err
	}

	return nil
}

func GetCategories(db *sql.DB, name string) ([]models.Types, error) {

	rows, err := db.Query(`select type, ifnull(parent_id,0), id, name, icon, user_id, created_at from category where name LIKE '%' || $1 || '%' OR type::string like '%' || $1 || '%'`,
		name,
	)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	categories := []models.Types{}

	for rows.Next() {
		var category models.Types
		if err := rows.Scan(&category.Type, &category.ParentCategory.ParentID, &category.ParentCategory.SubCategory.CategoryID, &category.ParentCategory.SubCategory.Name, &category.ParentCategory.SubCategory.Icon, &category.ParentCategory.SubCategory.UserID, &category.ParentCategory.SubCategory.CreatedAt); err != nil {
			log.Println(err)
			return []models.Types{}, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return []models.Types{}, err
	}
	return categories, err
}
