package models

import (
    "github.com/jinzhu/gorm"
)

type Todo struct {
    gorm.Model
    Text    string  `gorm:"size:200" json:"text"`
    Checked bool
}