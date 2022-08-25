package main

import (
	"fmt"
	"log"

	"github.com/ertugrulbal/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	dsn := "host=18.185.93.196 user=postgres password=postgres dbname=testErtugrul port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	if err != nil {
		log.Fatal("Could not connect database")
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	app := &app.App{}
	app.Initialize()
	//model.Seed(db)
	// model.ListProcess(db)
	// model.ListRoles(db)
	// model.ClearEverything(db)
	app.Run(":3000")
}
