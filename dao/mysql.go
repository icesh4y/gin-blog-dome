package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var DB *gorm.DB

func InitDB() (err error) {
	driver := viper.GetString("datasource.dirverName")
	host := viper.GetString("datasource.host")
	name := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.databases")
	charset := viper.GetString("datasource.charset")
	local := viper.GetString("local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", name, password, host, port, database, charset,url.QueryEscape(local))
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		fmt.Println("gorm open databases error", err)
		return
	}

	DB = db

	return  nil
}
