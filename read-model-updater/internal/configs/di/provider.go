package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/pqxd"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
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
	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		panic(err)
	}
	return &awsConfig
}

func provideNewRelicApp() *newrelic.Application {
	app, err := newrelic.NewApplication(
		nrlambda.ConfigOption(),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_CONFIG_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_CONFIG_LICENSE")),
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

func provideRDBGORM() *rdb.DB {
	articleDB, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN_ARTICLE"))
	if err != nil {
		panic(err)
	}
	articleDialector := postgres.New(postgres.Config{Conn: articleDB})

	// default connection
	gormDB, err := gorm.Open(articleDialector)
	if err != nil {
		panic(err)
	}

	tagDB, err := sql.Open("nrpgx", os.Getenv("COCKROACHDB_DSN_TAG"))
	if err != nil {
		panic(err)
	}
	tagDialector := postgres.New(postgres.Config{Conn: tagDB})
	gormDB.Use(
		dbresolver.
			Register(
				dbresolver.Config{
					Sources: []gorm.Dialector{articleDialector}, TraceResolverMode: true,
				}, rdb.ArticleDBName,
			).
			Register(
				dbresolver.Config{
					Sources: []gorm.Dialector{tagDialector}, TraceResolverMode: true,
				}, rdb.TagDBName,
			),
	)
	return &rdb.DB{DB: gormDB}
}

func provideDynamoDBGORM(awsConfig *aws.Config) *dynamo.DB {
	db := sql.OpenDB(pqxd.NewConnector(*awsConfig))
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	dynamoDialector := dynmgrm.New(dynmgrm.WithConnection(db))

	// default connection
	gormDB, err := gorm.Open(dynamoDialector)
	if err != nil {
		panic(err)
	}
	return &dynamo.DB{DB: gormDB}
}

func provideBlogPublisher() *githubactions.BlogPublisher {
	endpoint := os.Getenv("BLOG_PUBLISH_ENDPOINT")
	token := os.Getenv("GITHUB_TOKEN")
	return githubactions.NewBlogPublisher(endpoint, token, http.DefaultClient)
}

func provideSynUsecaseSet(
	rdbGorm *rdb.DB,
	dynamodbGorm *dynamo.DB,
	bloggingEventQueryService query.BloggingEventService,
	articleCommandService command.Article,
	tagCommandService command.Tag,
	blogAPIPublisher externalapi.BlogPublisher,
) *usecase.Sync {
	return usecase.NewSync(
		rdbGorm,
		dynamodbGorm,
		bloggingEventQueryService,
		articleCommandService,
		tagCommandService,
		blogAPIPublisher,
	)
}
