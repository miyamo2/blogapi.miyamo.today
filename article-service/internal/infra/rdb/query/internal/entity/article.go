package entity

import (
	"database/sql"
	"database/sql/driver"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/goccy/go-json"
)

var ErrNotCompatibleWithTags = errors.New("not compatible with tags")

type Article struct {
	ID        string               `gorm:"primaryKey; <-:false"`
	Title     string               `gorm:"<-:false"`
	Body      string               `gorm:"<-:false"`
	Thumbnail string               `gorm:"<-:false"`
	Tags      Tags                 `gorm:"<-:false"`
	CreatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
	UpdatedAt synchro.Time[tz.UTC] `gorm:"<-:false"`
}

type Tag struct {
	ID   string `gorm:"<-:false" json:"id"`
	Name string `gorm:"<-:false" json:"name"`
}

var (
	_ sql.Scanner   = (*Tags)(nil)
	_ driver.Valuer = (*Tags)(nil)
)

type Tags []Tag

func (t *Tags) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, t)
	}
	return ErrNotCompatibleWithTags
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}
