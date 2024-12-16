package dynamo

import (
	"testing"
)

func TestBloggingEventCommandService_CreateArticle(t *testing.T) {
	// do not work with sqlmock

	//type args struct {
	//	ctx context.Context
	//	in  model.CreateArticleEvent
	//	out *db.SingleStatementResult[*model.BloggingEventKey]
	//}
	//type test struct {
	//	args        args
	//	execOpt     func() []db.ExecuteOption
	//	want        error
	//	expectedOut *db.SingleStatementResult[*model.BloggingEventKey]
	//}
	//
	//articleID := ulid.MustParse("01JF0RDJYN8NJ57RN65G7FNGHS")
	//eventID := ulid.MustParse("01JF0REBGD4QKPFGN1SX2STY4M")
	//
	//tests := map[string]test{
	//	"happy_path": {
	//		args: args{
	//			ctx: context.Background(),
	//			in:  model.NewCreateArticleEvent("abc", "hello world", "https://example.com/example.png", []string{"tag1", "tag2"}),
	//			out: db.NewSingleStatementResult[*model.BloggingEventKey](),
	//		},
	//		execOpt: func() []db.ExecuteOption {
	//			sqlDB, mock, err := sqlmock.New()
	//			if err != nil {
	//				panic(err)
	//			}
	//			mock.ExpectExec(`INSERT INTO "blogging_event" VALUE {'event_id' : ?, 'article_id' : ?, 'title' : ?, 'content' : ?, 'thumbnail' : ?, 'tags' : ?}`).
	//				WithArgs("01JF0RDJYN8NJ57RN65G7FNGHS", "01JF0REBGD4QKPFGN1SX2STY4M", "abc", "hello world", "https://example.com/example.png", []string{"tag1", "tag2"}).
	//				WillReturnResult(driver.ResultNoRows)
	//			dialector := dynmgrm.New(dynmgrm.WithConnection(sqlDB))
	//			gormDB, err := gorm.Open(dialector, &gorm.Config{
	//				PrepareStmt:            false,
	//				SkipDefaultTransaction: true,
	//			})
	//			if err != nil {
	//				panic(err)
	//			}
	//			return []db.ExecuteOption{gwrapper.WithTransaction(gormDB)}
	//		},
	//		want: nil,
	//		expectedOut: func() *db.SingleStatementResult[*model.BloggingEventKey] {
	//			out := db.NewSingleStatementResult[*model.BloggingEventKey]()
	//			key := model.NewBloggingEventKey(fmt.Sprintf("%s", eventID), fmt.Sprintf("%s", articleID))
	//			out.Set(&key)
	//			return out
	//		}(),
	//	},
	//}
	//for name, tt := range tests {
	//	t.Run(name, func(t *testing.T) {
	//		t.Setenv("BLOGGING_EVENTS_TABLE_NAME", "blogging_event")
	//		i := 0
	//		var mockULIDGenerator pkg.ULIDGenerator = func() ulid.ULID {
	//			if i == 0 {
	//				i++
	//				return eventID
	//			}
	//			return articleID
	//		}
	//		sut := NewBloggingEventCommandService(&mockULIDGenerator)
	//		err := sut.CreateArticle(tt.args.ctx, tt.args.in, tt.args.out).Execute(tt.args.ctx, tt.execOpt()...)
	//		if !errors.Is(err, tt.want) {
	//			t.Errorf("BloggingEventCommandService.CreateArticle() error = %v, want %v", err, tt.want)
	//		}
	//		if !reflect.DeepEqual(tt.expectedOut.StrictGet(), tt.args.out.StrictGet()) {
	//			t.Errorf("want %+v but got %+v", tt.expectedOut.StrictGet(), tt.args.out.StrictGet())
	//			return
	//		}
	//	})
	//}
}
