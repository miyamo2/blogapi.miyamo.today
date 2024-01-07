package entity

type Article struct {
	ID        string `gorm:"primaryKey; <-:false"`
	Title     string `gorm:"<-:false"`
	Body      string `gorm:"<-:false"`
	Thumbnail string `gorm:"<-:false"`
	CreatedAt string `gorm:"<-:false"`
	UpdatedAt string `gorm:"<-:false"`
}

type ArticleTag struct {
	Article
	TagID   *string `gorm:"<-:false"`
	TagName *string `gorm:"<-:false"`
}
