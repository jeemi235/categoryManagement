package middlewares

import (
	"categorymanagement/models"
	"errors"
	"reflect"
)

func InsertCategoryBodyValidation(data models.Categories) error {
	name := reflect.TypeOf(data.Name).Kind()
	if name != reflect.String {
		err := errors.New("Invalid value passed in body")
		return err
	}

	icon := reflect.TypeOf(data.Name).Kind()
	if icon != reflect.String {
		err := errors.New("Invalid value passed in body")
		return err
	}

	types := [4]string{"object", "material", "texture", "mep"}
	for i := 0; i < len(types); i++ {
		if types[i] == data.Type {
			return nil
		}
	}
	err := errors.New("Invalid value passed in body")
	return err
}

func UpdateCategoryValidation(newType string, newName string) error {
	if newName == "" {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	if newType == "" {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	ntype := reflect.TypeOf(newType).Kind()
	if ntype != reflect.String {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	name := reflect.TypeOf(newName).Kind()
	if name != reflect.String {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	return nil
}

func DeleteCategoryValidation(newName string) error {
	if newName == "" {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	name := reflect.TypeOf(newName).Kind()
	if name != reflect.String {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	return nil
}

func GetCategoryValidation(name string) error {
	if name == "" {
		err := errors.New("Invalid value passed in URL")
		return err
	}

	checkname := reflect.TypeOf(name).Kind()
	if checkname != reflect.String {
		err := errors.New("Invalid value passed in URL")
		return err
	}
	
	return nil
}
