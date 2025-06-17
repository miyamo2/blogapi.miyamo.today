module blogapi.miyamo.today/blogging-event-service

go 1.24.0

require (
	blogapi.miyamo.today/core v0.24.0
	blogapi.miyamo.today/core/echo v0.7.0
	connectrpc.com/connect v1.18.1
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.3.0
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.12
	github.com/aws/aws-sdk-go-v2/service/s3 v1.79.0
	github.com/cockroachdb/errors v1.11.3
	github.com/goccy/go-json v0.10.5
	github.com/google/go-cmp v0.7.0
	github.com/google/wire v0.6.0
	github.com/joho/godotenv v1.5.1
	github.com/labstack/echo/v4 v4.13.4
	github.com/miyamo2/altnrslog v0.4.2
	github.com/miyamo2/dynmgrm v0.10.0
	github.com/miyamo2/godynamo v1.4.0
	github.com/miyamo2/sqldav v0.2.1
	github.com/newrelic/go-agent/v3 v3.38.0
	github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2 v1.2.3
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.1.3
	github.com/newrelic/go-agent/v3/integrations/nrpkgerrors v1.1.0
	github.com/oklog/ulid/v2 v2.1.0
	go.uber.org/mock v0.5.0
	golang.org/x/net v0.40.0
	google.golang.org/protobuf v1.36.6
	gorm.io/gorm v1.25.12
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.65 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.18.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.30 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.42.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.25.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.25.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.30.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.17 // indirect
	github.com/aws/smithy-go v1.22.3 // indirect
	github.com/btnguyen2k/consu/g18 v0.1.0 // indirect
	github.com/btnguyen2k/consu/reddo v0.1.9 // indirect
	github.com/cockroachdb/logtags v0.0.0-20241215232642-bb51bb14a506 // indirect
	github.com/cockroachdb/redact v1.1.6 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0 // indirect
	github.com/getsentry/sentry-go v0.31.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.2 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/httprc/v3 v3.0.0-beta1 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/jwx v1.2.30 // indirect
	github.com/lestrrat-go/jwx/v3 v3.0.0 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter v1.0.1 // indirect
	github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrwriter v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/grpc v1.71.0 // indirect
	gorm.io/plugin/dbresolver v1.5.3 // indirect
)
