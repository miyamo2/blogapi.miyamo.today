package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/mock/infra/s3"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/pkg"
	"go.uber.org/mock/gomock"
	"io"
	"net/url"
	"reflect"
	"testing"
)

func TestUploader_Upload(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		data        []byte
		contentType string
	}
	type want struct {
		uri *url.URL
		err error
	}
	type testCase struct {
		args              args
		want              want
		setupMockS3Client func(client *s3.MockClient)
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx:         context.Background(),
				name:        "example.png",
				data:        []byte("abcd"),
				contentType: "image/png",
			},
			want: want{
				uri: pkg.MustParseURL("https://example.com/example.png"),
				err: nil,
			},
			setupMockS3Client: func(client *s3.MockClient) {
				client.EXPECT().
					PutObject(gomock.Any(), NewPutObjectInputMatcher(&awss3.PutObjectInput{
						Bucket:      aws.String("example"),
						Key:         aws.String("example.png"),
						Body:        bytes.NewBuffer([]byte("abcd")),
						ContentType: aws.String("image/png"),
					}), gomock.Any()).
					Return(nil, nil).
					Times(1)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Setenv("S3_BUCKET", "example")
			t.Setenv("CDN_HOST", "https://example.com")

			gomock.NewController(t)
			client := s3.NewMockClient(gomock.NewController(t))
			tt.setupMockS3Client(client)

			u := NewUploader(client)
			uri, err := u.Upload(tt.args.ctx, tt.args.name, tt.args.data, tt.args.contentType)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Upload() = %v, want %v", err, tt.want)
			}
			if !reflect.DeepEqual(uri, tt.want.uri) {
				t.Errorf("Upload() = %v, want %v", uri, tt.want)
			}
		})
	}
}

type PutObjectInputMatcher struct {
	gomock.Matcher
	expect *awss3.PutObjectInput
}

func NewPutObjectInputMatcher(expect *awss3.PutObjectInput) gomock.Matcher {
	return &PutObjectInputMatcher{
		expect: expect,
	}
}

func (m *PutObjectInputMatcher) Matches(x interface{}) bool {
	var cmpOpts = []cmp.Option{
		cmp.AllowUnexported(awss3.PutObjectInput{}),
		cmpopts.IgnoreFields(awss3.PutObjectInput{}, "Body"),
	}
	switch x := x.(type) {
	case *awss3.PutObjectInput:
		diff := cmp.Diff(x, m.expect, cmpOpts...)
		if diff != "" {
			return false
		}
		expectBody, err := io.ReadAll(m.expect.Body)
		if err != nil {
			return false
		}
		xBody, err := io.ReadAll(x.Body)
		if err != nil {
			return false
		}
		return reflect.DeepEqual(xBody, expectBody)
	}
	return false
}

func (m *PutObjectInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
