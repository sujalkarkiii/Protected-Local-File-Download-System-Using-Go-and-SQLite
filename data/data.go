package data

import "gorm.io/gorm"

type Handlelingauthetication struct{
	gorm.Model
	RollNo int16
	Department string
}

type File struct {
	Filename string `json:"filename"`
}