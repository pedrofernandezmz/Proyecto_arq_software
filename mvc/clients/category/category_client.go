package category

import (
	"mvc/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB //ORM para la base de datos

func GetCategoryById(id int) model.Category {
	var category model.Category
	Db.Where("id = ?", id).First(&category) //traduccion y seteo en category
	log.Debug("Category: ", category)

	return category
}

func GetCategories() model.Categories {
	var categories model.Categories
	Db.Find(&categories) //busca y guarda todo en categories
	log.Debug("Categories: ", categories)

	return categories
}
