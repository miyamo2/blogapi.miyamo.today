package converter

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"testing"
)

var cmpOpts = []cmp.Option{
	cmp.AllowUnexported(gqlscalar.URL{}),
	cmp.AllowUnexported(gqlscalar.UTC{}),
}

func TestConverter_ToArticle(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.ArticleOutDto
	}
	type want struct {
		out *model.ArticleNode
		ok  bool
	}
	type testCase struct {
		sut  func() *Converter
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/single_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/single_tag",
						"## happy_path/single_tag",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						}),
				),
			},
			want: want{
				&model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/single_tag",
					Content:      "## happy_path/single_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges: []*model.ArticleTagEdge{
							{
								Cursor: "Tag1",
								Node: &model.ArticleTagNode{
									ID:   "Tag1",
									Name: "Tag1",
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Tag1",
							EndCursor:   "Tag1",
						},
						TotalCount: 1,
					},
				},
				true,
			},
		},
		"happy_path/multi_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/multi_tag",
						"## happy_path/multi_tag",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
							dto.NewTag("Tag2", "Tag2"),
						})),
			},
			want: want{
				&model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/multi_tag",
					Content:      "## happy_path/multi_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges: []*model.ArticleTagEdge{
							{
								Cursor: "Tag1",
								Node: &model.ArticleTagNode{
									ID:   "Tag1",
									Name: "Tag1",
								},
							},
							{
								Cursor: "Tag2",
								Node: &model.ArticleTagNode{
									ID:   "Tag2",
									Name: "Tag2",
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Tag1",
							EndCursor:   "Tag2",
						},
						TotalCount: 2,
					},
				},
				true,
			},
		},
		"happy_path/no_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/no_tag",
						"## happy_path/no_tag",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{}),
				),
			},
			want: want{
				&model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/no_tag",
					Content:      "## happy_path/no_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges:    []*model.ArticleTagEdge{},
						PageInfo: &model.PageInfo{},
					},
				},
				true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, ok := c.ToArticle(tt.args.ctx, tt.args.from)
			if ok != tt.want.ok {
				t.Errorf("ToArticle() ok = %v, want %v", ok, tt.want.ok)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestConverter_articleNodeFromArticleTagDto(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.ArticleTag
	}
	type want struct {
		out *model.ArticleNode
		err error
	}
	type testCase struct {
		sut     func() *Converter
		args    args
		want    want
		wantErr bool
	}
	tests := map[string]testCase{
		"happy_path/single_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleTag(
					"Article1",
					"happy_path/single_tag",
					"## happy_path/single_tag",
					"example.test",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{
						dto.NewTag("Tag1", "Tag1"),
					}),
			},
			want: want{
				out: &model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/single_tag",
					Content:      "## happy_path/single_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges: []*model.ArticleTagEdge{
							{
								Cursor: "Tag1",
								Node: &model.ArticleTagNode{
									ID:   "Tag1",
									Name: "Tag1",
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Tag1",
							EndCursor:   "Tag1",
						},
						TotalCount: 1,
					},
				},
			},
		},
		"happy_path/multi_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleTag(
					"Article1",
					"happy_path/multi_tag",
					"## happy_path/multi_tag",
					"example.test",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{
						dto.NewTag("Tag1", "Tag1"),
						dto.NewTag("Tag2", "Tag2"),
					}),
			},
			want: want{
				out: &model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/multi_tag",
					Content:      "## happy_path/multi_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges: []*model.ArticleTagEdge{
							{
								Cursor: "Tag1",
								Node: &model.ArticleTagNode{
									ID:   "Tag1",
									Name: "Tag1",
								},
							},
							{
								Cursor: "Tag2",
								Node: &model.ArticleTagNode{
									ID:   "Tag2",
									Name: "Tag2",
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Tag1",
							EndCursor:   "Tag2",
						},
						TotalCount: 2,
					},
				},
			},
		},
		"happy_path/no_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticleTag(
					"Article1",
					"happy_path/no_tag",
					"## happy_path/no_tag",
					"example.test",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					[]dto.Tag{}),
			},
			want: want{
				out: &model.ArticleNode{
					ID:           "Article1",
					Title:        "happy_path/no_tag",
					Content:      "## happy_path/no_tag",
					ThumbnailURL: "example.test",
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					Tags: &model.ArticleTagConnection{
						Edges:    []*model.ArticleTagEdge{},
						PageInfo: &model.PageInfo{},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, err := c.articleNodeFromArticleTagDto(tt.args.ctx, tt.args.from)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("articleNodeFromArticleTagDto() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestConverter_ToArticles(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.ArticlesOutDto
	}
	type want struct {
		out *model.ArticleConnection
		ok  bool
	}
	type testCase struct {
		sut  func() *Converter
		args args
		want want
	}
	ptrue := func() *bool {
		v := true
		return &v
	}()
	pfalse := func() *bool {
		v := false
		return &v
	}()
	_ = pfalse
	tests := map[string]testCase{
		"happy_path/single_article/single_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/single_tag",
							"## happy_path/single_article/single_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							},
						),
					},
				),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/single_tag",
								Content:      "## happy_path/single_article/single_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article1",
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/single_article/multi_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/multi_tag",
							"## happy_path/single_article/multi_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					},
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/multi_tag",
								Content:      "## happy_path/single_article/multi_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article1",
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/multi_article/single_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/single_tag",
							"## happy_path/multi_article/single_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/single_tag",
							"## happy_path/multi_article/single_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/single_tag",
								Content:      "## happy_path/multi_article/single_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/single_tag",
								Content:      "## happy_path/multi_article/single_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article2",
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/multi_article/multi_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/multi_tag",
							"## happy_path/multi_article/multi_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/multi_tag",
							"## happy_path/multi_article/multi_tag",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					}),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/multi_tag",
								Content:      "## happy_path/multi_article/multi_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/multi_tag",
								Content:      "## happy_path/multi_article/multi_tag",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article2",
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/single_article/single_tag/has_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/single_tag/with_next_paging",
							"## happy_path/single_article/single_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							},
						),
					},
					dto.ArticlesOutDtoWithHasNext(true),
				),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/single_tag/with_next_paging",
								Content:      "## happy_path/single_article/single_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article1",
						HasNextPage: ptrue,
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/single_article/multi_tag/with_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/multi_tag/with_next_paging",
							"## happy_path/single_article/multi_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					},
					dto.ArticlesOutDtoWithHasNext(true),
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/multi_tag/with_next_paging",
								Content:      "## happy_path/single_article/multi_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article1",
						HasNextPage: ptrue,
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/multi_article/single_tag/with_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/single_tag/with_next_paging",
							"## happy_path/multi_article/single_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/single_tag/with_next_paging",
							"## happy_path/multi_article/single_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasNext(true),
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/single_tag/with_next_paging",
								Content:      "## happy_path/multi_article/single_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/single_tag/with_next_paging",
								Content:      "## happy_path/multi_article/single_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article2",
						HasNextPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/multi_article/multi_tag/with_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/multi_tag/with_next_paging",
							"## happy_path/multi_article/multi_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/multi_tag/with_next_paging",
							"## happy_path/multi_article/multi_tag/with_next_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					},
					dto.ArticlesOutDtoWithHasNext(true)),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/multi_tag/with_next_paging",
								Content:      "## happy_path/multi_article/multi_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/multi_tag/with_next_paging",
								Content:      "## happy_path/multi_article/multi_tag/with_next_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article2",
						HasNextPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/single_article/single_tag/has_prev_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/single_tag/with_prev_paging",
							"## happy_path/single_article/single_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							},
						),
					},
					dto.ArticlesOutDtoWithHasPrev(true),
				),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/single_tag/with_prev_paging",
								Content:      "## happy_path/single_article/single_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor:     "Article1",
						EndCursor:       "Article1",
						HasPreviousPage: ptrue,
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/single_article/multi_tag/with_prev_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/single_article/multi_tag/with_prev_paging",
							"## happy_path/single_article/multi_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					},
					dto.ArticlesOutDtoWithHasPrev(true),
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/single_article/multi_tag/with_prev_paging",
								Content:      "## happy_path/single_article/multi_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor:     "Article1",
						EndCursor:       "Article1",
						HasPreviousPage: ptrue,
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/multi_article/single_tag/with_prev_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/single_tag/with_prev_paging",
							"## happy_path/multi_article/single_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/single_tag/with_prev_paging",
							"## happy_path/multi_article/single_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasPrev(true),
				)},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/single_tag/with_prev_paging",
								Content:      "## happy_path/multi_article/single_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/single_tag/with_prev_paging",
								Content:      "## happy_path/multi_article/single_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor:     "Article1",
						EndCursor:       "Article2",
						HasPreviousPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/multi_article/multi_tag/with_prev_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/multi_article/multi_tag/with_prev_paging",
							"## happy_path/multi_article/multi_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
						dto.NewArticleTag(
							"Article2",
							"happy_path/multi_article/multi_tag/with_prev_paging",
							"## happy_path/multi_article/multi_tag/with_prev_paging",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
								dto.NewTag("Tag2", "Tag2"),
							}),
					},
					dto.ArticlesOutDtoWithHasPrev(true)),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "happy_path/multi_article/multi_tag/with_prev_paging",
								Content:      "## happy_path/multi_article/multi_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Article2",
							Node: &model.ArticleNode{
								ID:           "Article2",
								Title:        "happy_path/multi_article/multi_tag/with_prev_paging",
								Content:      "## happy_path/multi_article/multi_tag/with_prev_paging",
								ThumbnailURL: "example.test",
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
										{
											Cursor: "Tag2",
											Node: &model.ArticleTagNode{
												ID:   "Tag2",
												Name: "Tag2",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor:     "Article1",
						EndCursor:       "Article2",
						HasPreviousPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/no_article": {
			sut: NewConverter,
			args: args{
				ctx:  context.Background(),
				from: dto.NewArticlesOutDto([]dto.ArticleTag{}),
			},
			want: want{
				ok: true,
				out: &model.ArticleConnection{
					Edges:      []*model.ArticleEdge{},
					PageInfo:   &model.PageInfo{},
					TotalCount: 0,
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, ok := c.ToArticles(tt.args.ctx, tt.args.from)
			if ok != tt.want.ok {
				t.Errorf("ToArticles() ok = %v, want %v", ok, tt.want.ok)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestConverter_ToTag(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.TagOutDto
	}
	type want struct {
		out *model.TagNode
		err error
	}
	type testCase struct {
		sut     func() *Converter
		args    args
		want    want
		wantErr bool
	}
	tests := map[string]testCase{
		"happy_path/single_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								"example.test",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))})),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges: []*model.TagArticleEdge{
							{
								Cursor: "Article1",
								Node: &model.TagArticleNode{
									ID:           "Article1",
									Title:        "Article1",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Article1",
							EndCursor:   "Article1",
						},
						TotalCount: 1,
					},
				},
			},
		},
		"happy_path/multiple_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								"example.test",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							dto.NewArticle(
								"Article2",
								"Article2",
								"",
								"example.test",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))})),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges: []*model.TagArticleEdge{
							{
								Cursor: "Article1",
								Node: &model.TagArticleNode{
									ID:           "Article1",
									Title:        "Article1",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
							{
								Cursor: "Article2",
								Node: &model.TagArticleNode{
									ID:           "Article2",
									Title:        "Article2",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Article1",
							EndCursor:   "Article2",
						},
						TotalCount: 2,
					},
				},
			},
		},
		"happy_path/no_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{})),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges:      []*model.TagArticleEdge{},
						PageInfo:   &model.PageInfo{},
						TotalCount: 0,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, err := c.ToTag(tt.args.ctx, tt.args.from)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToTag() error = %v, want =  %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestConverter_tagNodeFromTagArticleDto(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.TagArticle
	}
	type want struct {
		out *model.TagNode
		err error
	}
	type testCase struct {
		sut     func() *Converter
		args    args
		want    want
		wantErr bool
	}

	tests := map[string]testCase{
		"happy_path/single_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagArticle(
					"Tag1",
					"Tag1",
					[]dto.Article{
						dto.NewArticle(
							"Article1",
							"Article1",
							"## Article1",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
					}),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges: []*model.TagArticleEdge{
							{
								Cursor: "Article1",
								Node: &model.TagArticleNode{
									ID:           "Article1",
									Title:        "Article1",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Article1",
							EndCursor:   "Article1",
						},
						TotalCount: 1,
					},
				},
			},
		},
		"happy_path/multiple_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagArticle(
					"Tag1",
					"Tag1",
					[]dto.Article{
						dto.NewArticle(
							"Article1",
							"Article1",
							"## Article1",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
						dto.NewArticle(
							"Article2",
							"Article2",
							"## Article2",
							"example.test",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
					}),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges: []*model.TagArticleEdge{
							{
								Cursor: "Article1",
								Node: &model.TagArticleNode{
									ID:           "Article1",
									Title:        "Article1",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
							{
								Cursor: "Article2",
								Node: &model.TagArticleNode{
									ID:           "Article2",
									Title:        "Article2",
									ThumbnailURL: "example.test",
									CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
									UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								},
							},
						},
						PageInfo: &model.PageInfo{
							StartCursor: "Article1",
							EndCursor:   "Article2",
						},
						TotalCount: 2,
					},
				},
			},
		},
		"happy_path/no_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagArticle(
					"Tag1",
					"Tag1",
					[]dto.Article{}),
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
					Articles: &model.TagArticleConnection{
						Edges:      []*model.TagArticleEdge{},
						PageInfo:   &model.PageInfo{},
						TotalCount: 0,
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, err := c.tagNodeFromTagArticleDto(tt.args.ctx, tt.args.from)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("tagNodeFromTagArticleDto() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func TestConverter_ToTags(t *testing.T) {
	type args struct {
		ctx  context.Context
		from dto.TagsOutDto
	}
	type want struct {
		out *model.TagConnection
		err error
	}
	type testCase struct {
		sut     func() *Converter
		args    args
		want    want
		wantErr bool
	}
	ptrue := func() *bool {
		v := true
		return &v
	}()
	tests := map[string]testCase{
		"happy_path/single_tag/single_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					}),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/single_tag/multiple_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
							}),
					}),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/multiple_tag/single_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
						dto.NewTagArticle(
							"Tag2",
							"Tag2",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
					}),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article1",
									},
									TotalCount: 1,
								},
							},
						},
						{
							Cursor: "Tag2",
							Node: &model.TagNode{
								ID:   "Tag2",
								Name: "Tag2",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag2",
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/multiple_tag/multiple_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
						dto.NewTagArticle(
							"Tag2",
							"Tag2",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
					}),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Tag2",
							Node: &model.TagNode{
								ID:   "Tag2",
								Name: "Tag2",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag2",
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/single_tag/single_article/has_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
					},
					dto.TagsOutDtoWithHasNext(true)),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
						HasNextPage: ptrue,
					},
					TotalCount: 1,
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/has_next_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
						dto.NewTagArticle(
							"Tag2",
							"Tag2",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
					},
					dto.TagsOutDtoWithHasNext(true)),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Tag2",
							Node: &model.TagNode{
								ID:   "Tag2",
								Name: "Tag2",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag2",
						HasNextPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/has_prev_paging": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
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
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
						dto.NewTagArticle(
							"Tag2",
							"Tag2",
							[]dto.Article{
								dto.NewArticle(
									"Article1",
									"Article1",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								dto.NewArticle(
									"Article2",
									"Article2",
									"",
									"example.test",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0))}),
					},
					dto.TagsOutDtoWithHasPrev(true)),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
						{
							Cursor: "Tag2",
							Node: &model.TagNode{
								ID:   "Tag2",
								Name: "Tag2",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
										{
											Cursor: "Article2",
											Node: &model.TagArticleNode{
												ID:           "Article2",
												Title:        "Article2",
												ThumbnailURL: "example.test",
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Article1",
										EndCursor:   "Article2",
									},
									TotalCount: 2,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor:     "Tag1",
						EndCursor:       "Tag2",
						HasPreviousPage: ptrue,
					},
					TotalCount: 2,
				},
			},
		},
		"happy_path/no_tag": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
					[]dto.TagArticle{}),
			},
			want: want{
				out: &model.TagConnection{
					Edges:      []*model.TagEdge{},
					PageInfo:   &model.PageInfo{},
					TotalCount: 0,
				},
			},
		},
		"happy_path/single_tag/no_article": {
			sut: NewConverter,
			args: args{
				ctx: context.Background(),
				from: dto.NewTagsOutDto(
					[]dto.TagArticle{
						dto.NewTagArticle(
							"Tag1",
							"Tag1",
							[]dto.Article{},
						),
					}),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges:      []*model.TagArticleEdge{},
									PageInfo:   &model.PageInfo{},
									TotalCount: 0,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
					},
					TotalCount: 1,
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := tt.sut()
			got, err := c.ToTags(tt.args.ctx, tt.args.from)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToTags() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
