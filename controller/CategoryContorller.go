package controller

import (
	"essential/models"
	"essential/repository"
	"essential/response"
	"essential/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategorycontroller interface {
	RestController
}

type CategoryController struct {
	repository repository.CategoryRepository
}

func NewCategoryController() ICategorycontroller {
	repository := repository.NewCategoryRepository()

	repository.DB.AutoMigrate(&models.Category{})

	return CategoryController{repository: repository}
}

func (c2 CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	category, err := c2.repository.Create(requestCategory.Name)
	if err !=nil{
		panic(err)
		return
	}
	response.Success(c, gin.H{"category": category}, "")
}

func (c2 CategoryController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if err := c2.repository.DeleteById(id);err!=nil{
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "")
}

func (c2 CategoryController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	category, err := c2.repository.SelectById(id)
	if err != nil{
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"category": category}, "")
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

	updateCategory, err := c2.repository.SelectById(id)
	if err !=nil{
		response.Fail(c, nil, "分类不存在")
		return
	}
	category, err := c2.repository.Update(*updateCategory, requestCategory.Name)
	if err != nil{
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"category": category}, "修改成功")
}
