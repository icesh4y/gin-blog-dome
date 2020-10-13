package controller

import (
	"essential/dao"
	models "essential/models"
	"essential/response"
	"essential/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

type IPostController interface {
	RestController
}
type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "数据验证错误")
		log.Panicln(err.Error())
		return
	}
	user, _ := c.Get("user")

	post := models.Post{
		UserId:     user.(models.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(c, nil, "文章创建成功")
}

func (p PostController) Delete(c *gin.Context) {
	postId := c.Params.ByName("id")

	var post models.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID
	if userID != post.UserId {
		response.Fail(c, nil, "该文章不属于当前用用户")
		return
	}

	p.DB.Delete(&post)
	response.Success(c, nil, "成功")
}

func (p PostController) Show(c *gin.Context) {
	postId := c.Params.ByName("id")

	var post models.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}
	response.Success(c, gin.H{"post": post}, "成功")
}

func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "数据验证错误")
		log.Panicln(err.Error())
		return
	}

	postId := c.Params.ByName("id")

	var post models.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID
	if userID != post.UserId {
		response.Fail(c, nil, "该文章不属于当前用于")
		return
	}

	if err := p.DB.Model(&post).Update(&requestPost).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, gin.H{"requestPost": requestPost}, "更新成功")

}

func NewPostsRouterController() IPostController {
	db := dao.DB
	db.AutoMigrate(&models.Post{})
	return PostController{DB: db}
}
