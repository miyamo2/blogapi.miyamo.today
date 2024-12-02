package usecase

import (
	"context"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	blogapictx "github.com/miyamo2/blogapi.miyamo.today/core/context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	mgrpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/infra/grpc/tag"
	"go.uber.org/mock/gomock"
)

func TestTags_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagsInDto
	}
	type want struct {
		out dto.TagsOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) grpc.TagServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestTags := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/next_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
					dto.TagsOutDtoWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithLast(1), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
					dto.TagsOutDtoWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithLast(1), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
		"happy_path/execute": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.TagsInDto{},
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
				),
			},
		},
		"unhappy_path/execute_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.TagsInDto{},
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewTags(aSvcClt)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_executeNextPaging(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagsInDto
	}
	type want struct {
		out dto.TagsOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) grpc.TagServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestTags := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/next_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
					dto.TagsOutDtoWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewTags(aSvcClt)
			got, err := u.executeNextPaging(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("executeNextPaging() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("executeNextPaging() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("executeNextPaging() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("executeNextPaging() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}

func TestTags_executePrevPaging(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagsInDto
	}
	type want struct {
		out dto.TagsOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) grpc.TagServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestTags := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithLast(2), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
					dto.TagsOutDtoWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDto {
					in, _ := dto.NewTagsInDto(dto.TagsInWithLast(2), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewTags(aSvcClt)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTags_execute(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type want struct {
		out dto.TagsOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) grpc.TagServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestTags := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/all_articles": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "Tag1",
								Name: "Tag1",
								Articles: []*grpc.Article{
									{
										Id:           "Article1",
										Title:        "Article1",
										ThumbnailUrl: "example.test",
										CreatedAt:    "2020-01-01T00:00:00Z",
										UpdatedAt:    "2020-01-01T00:00:00Z",
									},
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z"),
							}),
					},
				),
			},
		},
		"unhappy_path/all_articles_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				aSvcClt := mgrpc.NewMockTagServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{}, errTestTags).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.TagsOutDto{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewTags(aSvcClt)
			got, err := u.execute(tt.args.ctx)
			if tt.wantErr {
				if err == nil {
					t.Errorf("execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("execute() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}
