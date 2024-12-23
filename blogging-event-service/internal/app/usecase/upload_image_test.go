package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/mock/app/usecase/storage"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/pkg"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestUploadImage_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *dto.UploadImageInDto
	}
	type want struct {
		out *dto.UploadImageOutDto
		err error
	}
	type testCase struct {
		args              args
		want              want
		setupMockUploader func(u *storage.MockUploader)
	}
	errUnhappyPath := errors.New("unhappy_path")

	tests := map[string]testCase{
		"happy_path": {
			args: func() args {
				in := dto.NewUploadImageInDto("example.png", []byte{}, "image/png")
				return args{
					ctx: context.Background(),
					in:  &in,
				}
			}(),
			want: func() want {
				out := dto.NewUploadImageOutDto(*pkg.MustParseURL("http://example.com/example.png"))
				return want{
					out: &out,
				}
			}(),
			setupMockUploader: func(u *storage.MockUploader) {
				u.EXPECT().
					Upload(gomock.Any(), "example.png", []byte{}, "image/png").
					Return(pkg.MustParseURL("http://example.com/example.png"), nil).
					Times(1)
			},
		},
		"unhappy_path": {
			args: func() args {
				in := dto.NewUploadImageInDto("example.png", []byte{}, "image/png")
				return args{
					ctx: context.Background(),
					in:  &in,
				}
			}(),
			want: want{
				err: errUnhappyPath,
			},
			setupMockUploader: func(u *storage.MockUploader) {
				u.EXPECT().
					Upload(gomock.Any(), "example.png", []byte{}, "image/png").
					Return(nil, errUnhappyPath).
					Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uploader := storage.NewMockUploader(ctrl)
			tt.setupMockUploader(uploader)

			u := NewUploadImage(uploader)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.want.err)
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
