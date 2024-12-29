package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	mcommand "blogapi.miyamo.today/blogging-event-service/internal/mock/app/usecase/command"
	mdb "blogapi.miyamo.today/blogging-event-service/internal/mock/core/db"
	"blogapi.miyamo.today/core/db"
	"context"
	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestCreateArticle_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *dto.CreateArticleInDto
	}
	type want struct {
		out *dto.CreateArticleOutDto
		err error
	}
	type testCase struct {
		args                args
		want                want
		setupCommandService func(cs *mcommand.MockBloggingEventService, in model.CreateArticleEvent, stmt *mdb.MockStatement)
	}
	errUnhappyPath := errors.New("unhappy_path")

	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				in: func() *dto.CreateArticleInDto {
					v := dto.NewCreateArticleInDto("happy_path", "## happy_path", "thumbnail", []string{"tag1", "tag2"})
					return &v
				}(),
			},
			setupCommandService: func(cs *mcommand.MockBloggingEventService, in model.CreateArticleEvent, stmt *mdb.MockStatement) {
				cs.EXPECT().CreateArticle(gomock.Any(), in, gomock.Any()).DoAndReturn(
					func(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
						stmt.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, _ ...db.ExecuteOption) error {
							v := model.NewBloggingEventKey("1", "1")
							out.Set(&v)
							return nil
						}).Times(1)
						return stmt
					}).Times(1)
			},
			want: func() want {
				o := dto.NewCreateArticleOutDto("1", "1")
				return want{out: &o, err: nil}
			}(),
		},
		"unhappy_path": {
			args: args{
				ctx: context.Background(),
				in: func() *dto.CreateArticleInDto {
					v := dto.NewCreateArticleInDto("unhappy_path", "## unhappy_path", "thumbnail", []string{"tag1", "tag2"})
					return &v
				}(),
			},
			setupCommandService: func(cs *mcommand.MockBloggingEventService, in model.CreateArticleEvent, stmt *mdb.MockStatement) {
				cs.EXPECT().CreateArticle(gomock.Any(), in, gomock.Any()).DoAndReturn(
					func(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
						stmt.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, _ ...db.ExecuteOption) error {
							return errUnhappyPath
						}).Times(1)
						return stmt
					}).Times(1)
			},
			want: want{out: nil, err: errUnhappyPath},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cs := mcommand.NewMockBloggingEventService(ctrl)
			stmt := mdb.NewMockStatement(ctrl)
			tt.setupCommandService(cs, model.NewCreateArticleEvent(tt.args.in.Title(), tt.args.in.Body(), tt.args.in.ThumbnailUrl(), tt.args.in.TagNames()), stmt)

			u := NewCreateArticle(cs)
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
