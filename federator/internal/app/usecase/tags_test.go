package usecase

import (
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"blogapi.miyamo.today/federator/internal/utils"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	mgrpc "blogapi.miyamo.today/federator/internal/mock/infra/grpc/tag"
	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestTags_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagsInDTO
	}
	type want struct {
		out dto.TagsOutDTO
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
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
					dto.TagsOutDTOWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithLast(1), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
					dto.TagsOutDTOWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithLast(1), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
		"happy_path/execute": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.TagsInDTO{},
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
				),
			},
		},
		"unhappy_path/execute_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.TagsInDTO{},
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewTags(articleServiceClient)
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
		in  dto.TagsInDTO
	}
	type want struct {
		out dto.TagsOutDTO
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
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
					dto.TagsOutDTOWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetNextTagResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithFirst(1), dto.TagsInWithAfter("Tag0"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewTags(articleServiceClient)
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
		in  dto.TagsInDTO
	}
	type want struct {
		out dto.TagsOutDTO
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
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithLast(2), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
					dto.TagsOutDTOWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetPrevTagResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.TagsInDTO {
					in, _ := dto.NewTagsInDTO(dto.TagsInWithLast(2), dto.TagsInWithBefore("Tag2"))
					return in
				}(),
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewTags(articleServiceClient)
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
		out dto.TagsOutDTO
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
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
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
										ThumbnailUrl: "example.com/example.png",
										CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
										UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									},
								},
							},
						},
					}, nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.NewTagsOutDTO(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									utils.MustURLParse("example.com/example.png"),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					},
				),
			},
		},
		"unhappy_path/all_articles_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.TagServiceClient {
				articleServiceClient := mgrpc.NewMockTagServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllTags(gomock.Any(), gomock.Any()).
					Return(&grpc.GetAllTagsResponse{}, errTestTags).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.TagsOutDTO{},
				err: errTestTags,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewTags(articleServiceClient)
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
