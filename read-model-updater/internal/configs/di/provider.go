package di

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/miyamo2/dynmgrm"
	"github.com/miyamo2/pqxd"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx5"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func provideAWSConfig() *aws.Config {
	ctx := context.Background()
	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		panic(err)
	}
	nraws.AppendMiddlewares(&awsConfig.APIOptions, nil)
	return &awsConfig
}

func provideNewRelicApp() *newrelic.Application {
	app, err := newrelic.NewApplication(
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

func provideStreamARN() StreamARN {
	v := os.Getenv("STREAM_ARN")
	return &v
}

func provideDynamoDBStreamClient(awsConfig *aws.Config) *dynamodbstreams.Client {
	return dynamodbstreams.NewFromConfig(*awsConfig)
}

func provideSynUsecaseSet(
	rdbGorm *rdb.DB,
	dynamodbGorm *dynamo.DB,
	bloggingEventQueryService query.BloggingEventService,
	articleCommandService command.ArticleService,
	tagCommandService command.TagService,
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
