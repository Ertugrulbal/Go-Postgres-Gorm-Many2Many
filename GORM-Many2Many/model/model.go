package model

import (
	"gorm.io/gorm"
	"fmt"
)

// var db *gorm.DB

type Roles struct {
	//gorm.Model
	ID        int        `json:"id" ;gorm:"primaryKey"`
	Name      string     `json:"name" ;gorm:"size:61;not null"`
	Processes []*Process `gorm:"many2many:process_roles"`
}

type Process struct {
	//gorm.Model
	ID   int      `json:"id" ;gorm:"primaryKey"`
	Name string   `json:"name" gorm:"size:50;not null"`
	Role []*Roles `json:"roles" gorm:"many2many:process_roles"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Roles{})
	db.AutoMigrate(&Process{})
	return db
}

func Seed(db *gorm.DB) {
	create := &Process{
		Name: "Create",
	}
	read := &Process{
		Name: "Read",
	}
	update := &Process{
		Name: "Update",
	}
	delete := &Process{
		Name: "Delete",
	}
	db.Save(create)
	db.Save(read)
	db.Save(update)
	db.Save(delete)

	Admin := &Roles{
		Name: "Admin",
		Processes: []*Process{
			create, read, update, delete,
		},
	}
	Editor := &Roles{
		Name: "Editor",
		Processes: []*Process{
			read, update,
		},
	}
	Viewer := &Roles{
		Name: "Viewer",
		Processes: []*Process{
			read,
		},
	}

	db.Save(Admin)
	db.Save(Editor)
	db.Save(Viewer)
	fmt.Printf("Roles  created:\n%v\n", Admin)
	fmt.Printf("Roles  created:\n%v\n", Editor)
	fmt.Printf("Roles  created:\n%v\n", Viewer)
	fmt.Printf("Processes created:\n%v\n", []*Process{create, read, update, delete})

}

func ListRoles(db *gorm.DB) {
	var roles []Roles
	db.Preload("Processes").Find(&roles)

	for _, role := range roles {
		fmt.Println("Role data: \n", role)
		for _, role := range role.Processes {

			fmt.Printf("Role-Process data: %v\n", role)
		}
	}
}

func ListProcess(db *gorm.DB) {
	var processes []Process
	db.Preload("Roles").Find(&processes)

	for _, process := range processes {
		fmt.Println("Process data: \n", process)
		for _, process := range process.Role {

			fmt.Printf("Role-Process data: %v\n", process)
		}
	}
}

func ClearEverything(db *gorm.DB) {
	err1 := db.Delete(&Roles{}).Error
	err2 := db.Delete(&Process{}).Error
	fmt.Printf("Deleting the records:\n%v\n", err1)
	fmt.Printf("Deleting the records:\n%v\n", err2)
}

