package controller

import (
	"essential/dao"
	"essential/models"
	"essential/response"
	"essential/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategorycontroller interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategorycontroller {
	dao.DB.AutoMigrate(&models.Category{})

	return CategoryController{DB: dao.DB}
}

func (c2 CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	category := models.Category{Name: requestCategory.Name}
	c2.DB.Create(&category)

	response.Success(c, gin.H{"requestCategory": requestCategory}, "")
}

func (c2 CategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if err := c2.DB.Delete(models.Category{}, id).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}

	response.Success(c, nil, "")
}

func (c2 CategoryController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var Category models.Category
	if c2.DB.First(&Category, id).RecordNotFound() {
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"Category": Category}, "")
}

func (c2 CategoryController) Update(c *gin.Context) {
	// 绑定body 中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	// 绑定 path 中的参数
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var updateCategory models.Category
	if c2.DB.First(&updateCategory, id).RecordNotFound() {
		response.Fail(c, nil, "分类不存在")
		return
	}

	c2.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(c, gin.H{"updateCategory": updateCategory}, "")
}
