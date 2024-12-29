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

func TestNewUpdateArticleBody(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *dto.UpdateArticleBodyInDto
	}
	type want struct {
		out *dto.UpdateArticleBodyOutDto
		err error
	}
	type testCase struct {
		args                args
		want                want
		setupCommandService func(cs *mcommand.MockBloggingEventService, in model.UpdateArticleBodyEvent, stmt *mdb.MockStatement)
	}
	errUnhappyPath := errors.New("unhappy_path")

	tests := map[string]testCase{
		"happy_path": {
			args: func() args {
				in := dto.NewUpdateArticleBodyInDto("article_id", "body")
				return args{
					ctx: context.Background(),
					in:  &in,
				}
			}(),
			want: func() want {
				out := dto.NewUpdateArticleBodyOutDto("event_id", "article_id")
				return want{
					out: &out,
				}
			}(),
			setupCommandService: func(cs *mcommand.MockBloggingEventService, in model.UpdateArticleBodyEvent, stmt *mdb.MockStatement) {
				cs.EXPECT().UpdateArticleBody(gomock.Any(), in, gomock.Any()).DoAndReturn(
					func(ctx context.Context, in model.UpdateArticleBodyEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
						stmt.EXPECT().Execute(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, _ ...db.ExecuteOption) error {
							v := model.NewBloggingEventKey("event_id", "article_id")
							out.Set(&v)
							return nil
						}).Times(1)
						return stmt
					}).Times(1)
			},
		},
		"unhappy_path": {
			args: func() args {
				in := dto.NewUpdateArticleBodyInDto("article_id", "body")
				return args{
					ctx: context.Background(),
					in:  &in,
				}
			}(),
			want: want{
				err: errUnhappyPath,
			},
			setupCommandService: func(cs *mcommand.MockBloggingEventService, in model.UpdateArticleBodyEvent, stmt *mdb.MockStatement) {
				cs.EXPECT().UpdateArticleBody(gomock.Any(), in, gomock.Any()).DoAndReturn(
					func(ctx context.Context, in model.UpdateArticleBodyEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
						stmt.EXPECT().Execute(gomock.Any(), gomock.Any()).Return(errUnhappyPath).Times(1)
						return stmt
					}).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			cs := mcommand.NewMockBloggingEventService(ctrl)
			stmt := mdb.NewMockStatement(ctrl)
			tt.setupCommandService(cs, model.NewUpdateArticleBodyEvent(tt.args.in.ID(), tt.args.in.Body()), stmt)

			u := NewUpdateArticleBody(cs)
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
