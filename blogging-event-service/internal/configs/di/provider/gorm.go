package provider

import (
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/wire"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/pqxd"
	"gorm.io/gorm"
)

func GormDialector(awsConfig *aws.Config) *gorm.Dialector {
	db := sql.OpenDB(pqxd.NewConnector(*awsConfig))
	if err := db.Ping(); err != nil {
		panic(err)
	}
	gormDialector := dynmgrm.New(dynmgrm.WithConnection(db))
	return &gormDialector
}

var GormSet = wire.NewSet(GormDialector)
