package entity

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

type Tag struct {
	ID        string               `gorm:"primaryKey; <-:false"`
	Name      string               `gorm:"<-:false"`
	CreatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
	UpdatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
}

type TagArticle struct {
	Tag
	ArticleID        *string               `gorm:"<-:false"`
	ArticleTitle     *string               `gorm:"<-:false"`
	ArticleThumbnail *string               `gorm:"<-:false"`
	ArticleCreatedAt *synchro.Time[tz.UTC] `gorm:"<-:false"`
	ArticleUpdatedAt *synchro.Time[tz.UTC] `gorm:"<-:false"`
}
