package di

import (
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
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
	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")))
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
			Register(dbresolver.Config{
				Sources: []gorm.Dialector{articleDialector}, TraceResolverMode: true,
			}, rdb.ArticleDBName).
			Register(dbresolver.Config{
				Sources: []gorm.Dialector{tagDialector}, TraceResolverMode: true,
			}, rdb.TagDBName),
	)
	return &rdb.DB{DB: gormDB}
}

func provideDynamoDBGORM(awsConfig *aws.Config) *dynamo.DB {
	db := sql.OpenDB(pqxd.NewConnector(*awsConfig))
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
