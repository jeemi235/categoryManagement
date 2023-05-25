package handlers

import (
	bll "categorymanagement/api/bll"
	e "categorymanagement/api/utils/errrors"
	"categorymanagement/cache"
	"categorymanagement/middlewares"
	"categorymanagement/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")

	db := r.Context().Value("database").(*sql.DB)

	var category models.Categories
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	data := category
	err = middlewares.InsertCategoryBodyValidation(data)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	err = bll.Insertcategory(data, db, id)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}
	w.Write([]byte("Data added successfully"))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	newType := r.URL.Query().Get("newType")
	newName := r.URL.Query().Get("newName")

	err := middlewares.UpdateCategoryValidation(newType, newName)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	db := r.Context().Value("database").(*sql.DB)

	err = bll.UpdateCategory(db, newType, newName)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data updated successfully"))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	newName := r.URL.Query().Get("newName")

	err := middlewares.DeleteCategoryValidation(newName)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	db := r.Context().Value("database").(*sql.DB)

	err = bll.DeleteCategory(db, newName)
	if err != nil {
		fmt.Println(err)
		e.ErrorGenerator(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data deleted successfully"))
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	id := name

	err := middlewares.GetCategoryValidation(name)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	db := r.Context().Value("database").(*sql.DB)

	data, err := bll.GetCategories(db, name)
	if err != nil {
		e.ErrorGenerator(w, err)
		return
	}

	//If we are calling api for the first time will add data in the cache
	cache.UpdateCache(id, data)
	
	middlewares.ResponseWithJsonPayload(w, data)
}
