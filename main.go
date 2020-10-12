package main

import (
	"essential/dao"
	"essential/models"
	"essential/router"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()                 // 获取当前工作目录
	viper.SetConfigName("application")       //设置读取配置文件名字
	viper.SetConfigType("yml")               // 设置读取配置文件类型
	viper.AddConfigPath(workDir + "/config") // 设置读取配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func main() {
	InitConfig()
	if err := dao.InitDB(); err != nil {
		panic(err)
	}

	defer dao.DB.Close()
	r := router.SetUpRouter()
	dao.DB.AutoMigrate(&models.User{})

	//if err := r.Run(); err != nil {
	//	fmt.Println("gin run error:", err)
	//	return
	//}

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}
