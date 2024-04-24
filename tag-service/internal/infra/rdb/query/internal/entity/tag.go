package entity

type Tag struct {
	ID        string `gorm:"primaryKey; <-:false"`
	Name      string `gorm:"<-:false"`
	CreatedAt string `gorm:"<-:false"`
	UpdatedAt string `gorm:"<-:false"`
}

type TagArticle struct {
	Tag
	ArticleID        *string `gorm:"<-:false"`
	ArticleTitle     *string `gorm:"<-:false"`
	ArticleThumbnail *string `gorm:"<-:false"`
	ArticleCreatedAt *string `gorm:"<-:false"`
	ArticleUpdatedAt *string `gorm:"<-:false"`
}
