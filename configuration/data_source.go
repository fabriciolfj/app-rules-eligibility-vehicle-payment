package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/fabriciolfj/rules-elegibility/data"
	"github.com/magiconair/properties"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProviderDataBase() *gorm.DB {
	p := properties.MustLoadFile("database.properties", properties.UTF8)
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = p.MustGetString("db_host")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		p.MustGetString("db_user"),
		p.MustGetString("db_password"),
		host,
		p.MustGetString("db_port"),
		p.MustGetString("db_name"))

	log.Printf("dsn: %s", dsn)

	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil || DB == nil {
		log.Fatal("failed to connect database", err)
	}

	err = DB.AutoMigrate(&data.CustomerData{})

	if err != nil {
		log.Println("failed to migrate database", err)
	}

	return DB
}
