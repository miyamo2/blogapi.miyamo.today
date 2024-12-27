package entity

import (
	"database/sql"
	"database/sql/driver"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/goccy/go-json"
)

var ErrNotCompatibleWithArticles = errors.New("not compatible with articles")

type Tag struct {
	ID        string               `gorm:"primaryKey; <-:false"`
	Name      string               `gorm:"<-:false"`
	Articles  Articles             `gorm:"<-:false"`
	CreatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
	UpdatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
}

type Article struct {
	ID        string               `gorm:"<-:false" json:"id"`
	Title     string               `gorm:"<-:false" json:"title"`
	Thumbnail string               `gorm:"<-:false" json:"thumbnail"`
	CreatedAt synchro.Time[tz.UTC] `gorm:"<-:false" json:"created_at"`
	UpdatedAt synchro.Time[tz.UTC] `gorm:"<-:false" json:"updated_at"`
}

var (
	_ sql.Scanner   = (*Articles)(nil)
	_ driver.Valuer = (*Articles)(nil)
)

type Articles []Article

func (a *Articles) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, a)
	}
	return ErrNotCompatibleWithArticles
}

func (a Articles) Value() (driver.Value, error) {
	return json.Marshal(a)
}
