package models

import (
	"github.com/atanda0x/go-book-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&book{})
}

func (b *book) CreateBook() *book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []book {
	var Books []book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*book, *gorm.DB) {
	var getBook book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) book {
	var Book book
	db.Where("ID=?", ID).Delete(Book)
	return Book
}
