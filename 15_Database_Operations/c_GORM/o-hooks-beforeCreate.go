package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*

Types of Hooks in GORM
1. Create Hooks
	These hooks are triggered during the creation of a record.
	Hook Name		Description
	-------------------------
	BeforeCreate	Executed before a record is created in the database.
	AfterCreate		Executed after a record is created in the database.

2. Update Hooks
	These hooks are triggered during the updating of a record.
	-----------------------
	Hook Name		Description
	-----------------------
	BeforeUpdate	Executed before a record is updated in the database.
	AfterUpdate		Executed after a record is updated in the database.

3. Save Hooks
	These hooks are triggered during both creation and updating of a record.
	-----------------------
	Hook Name	Description
	-----------------------
	BeforeSave	Executed before a record is created or updated in the database.
	AfterSave	Executed after a record is created or updated in the database.

4. Delete Hooks
	These hooks are triggered during the deletion of a record.
	-----------------------
	Hook Name		Description
	-----------------------
	BeforeDelete	Executed before a record is deleted from the database.
	AfterDelete	Executed after a record is deleted from the database.

5. Query Hooks
	These hooks are triggered during the querying of records.

	-----------------------
	Hook Name	Description
	-----------------------
	AfterFind	Executed after a record is queried from the database.



*/

type User struct {
	ID   uint
	Name string
	Age  int
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before creating user:", u.Name)
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&User{Name: "Eve", Age: 28})
}
