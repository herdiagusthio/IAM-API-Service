package config

import (
	"fmt"
	"strconv"

	"iam-api-service/repository/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabaseConnection(config *AppConfig) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": config.DbUsername,
		"DB_Password": config.DbPassword,
		"DB_Port":     strconv.Itoa(config.DbPort),
		"DB_Host":     config.DbAddress,
		"DB_Name":     config.DbName,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	fmt.Println(connectionString)

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}
