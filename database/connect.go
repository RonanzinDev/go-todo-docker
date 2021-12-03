package database

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/ronanzindev/go-todo-docker/config"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)
	configData := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open("postgres", configData)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed connect database")
	}
}
