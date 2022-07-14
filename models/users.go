package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        //uid, createdby, updatedby
	Name       string `json:"name,omitempty"     gorm:"type:varchar(40)"`
	Email      string `json:"email,omitempty"    gorm:"type:varchar(100);unique_index"`
	Password   string `json:"password,omitempty" gorm:"not null"`
	Admin      bool   `json:"admin,omitempty" gorm:"default:false"`
}
