package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate: Setting default values")
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

// AfterCreate hook
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate: Logging user creation")
	log.Printf("User created: %v\n", u)
	return
}

// BeforeUpdate hook
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate: Updating timestamp")
	u.UpdatedAt = time.Now()
	return
}

// AfterFind hook
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind: Logging queried user")
	log.Printf("User queried: %v\n", u)
	return
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the schema
	db.AutoMigrate(&User{})

	// Create a user
	user := User{Name: "John", Email: "john@example.com"}
	db.Create(&user)

	// Query the user
	var queriedUser User
	db.First(&queriedUser, user.ID)

	// Update the user
	db.Model(&queriedUser).Update("Name", "John Doe")

	// Delete the user
	db.Delete(&queriedUser)
}