package main

import (
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/godynamo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"net/http"
	"os"
	"time"
)

func provideAWSConfig() *aws.Config {
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}
	return &awsConfig
}

func provideNewRelicApp() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_CONFIG_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_CONFIG_LICENSE")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		panic(err)
	}
	err = app.WaitForConnection(5 * time.Second)
	if err != nil {
		panic(err)
	}
	return app
}

func provideGORMDB(awsConfig *aws.Config) *gorm.DB {
	godynamo.RegisterAWSConfig(*awsConfig)

	dynamoDialector := dynmgrm.New()

	// default connection
	gormDB, err := gorm.Open(dynamoDialector)
	if err != nil {
		panic(err)
	}

	articleDB, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN_ARTICLE"))
	if err != nil {
		panic(err)
	}
	articleDialector := postgres.New(postgres.Config{Conn: articleDB})

	tagDB, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN_TAG"))
	if err != nil {
		panic(err)
	}
	tagDialector := postgres.New(postgres.Config{Conn: tagDB})

	gormDB.Use(
		dbresolver.
			Register(dbresolver.Config{Sources: []gorm.Dialector{dynamoDialector}}, dynamo.DBName).
			Register(dbresolver.Config{Sources: []gorm.Dialector{articleDialector}}, rdb.ArticleDBName).
			Register(dbresolver.Config{Sources: []gorm.Dialector{tagDialector}}, rdb.TagDBName),
	)
	return gormDB
}

func provideBlogPublisher() *githubactions.BlogPublisher {
	endpoint := os.Getenv("BLOG_PUBLISH_ENDPOINT")
	token := os.Getenv("GITHUB_TOKEN")
	return githubactions.NewBlogPublisher(endpoint, token, http.DefaultClient)
}

func provideSynUsecaseSet(
	bloggingEventQueryService query.BloggingEventService,
	articleCommandService command.ArticleService,
	tagCommandService command.TagService,
	blogAPIPublisher externalapi.BlogPublisher,
) *usecase.Sync {
	return usecase.NewSync(
		gw.Manager(),
		bloggingEventQueryService,
		articleCommandService,
		tagCommandService,
		blogAPIPublisher,
	)
}
