package entity

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

type Article struct {
	ID        string               `gorm:"primaryKey; <-:false"`
	Title     string               `gorm:"<-:false"`
	Body      string               `gorm:"<-:false"`
	Thumbnail string               `gorm:"<-:false"`
	CreatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
	UpdatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
}

type ArticleTag struct {
	Article
	TagID   *string `gorm:"<-:false"`
	TagName *string `gorm:"<-:false"`
}
