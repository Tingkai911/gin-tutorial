package main

type Book struct {
	// uppercase field names -> exported fields, fields that can be read outside the module
	// `json:"id"` to declare the json key as lower case
	ID       string `json:"id" gorm:"id"`
	Title    string `json:"title" gorm:"title"`
	Author   string `json:"author" gorm:"author"`
	Quantity int    `json:"quantity" gorm:"quanity"`
}

// overwrites table name
func (Book) TableName() string {
	return "book"
}
