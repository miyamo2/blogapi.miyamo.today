package provider

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/wire"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/godynamo"
	"gorm.io/gorm"
)

func GormDialector(awsConfig *aws.Config) *gorm.Dialector {
	godynamo.RegisterAWSConfig(*awsConfig)

	db, err := sql.Open("godynamo", "")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err) // because they are critical errors
	}

	gormDialector := dynmgrm.New(dynmgrm.WithConnection(db))
	return &gormDialector
}

var GormSet = wire.NewSet(GormDialector)
