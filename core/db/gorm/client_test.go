package gorm

import (
	"context"
	"github.com/miyamo2/api.miyamo.today/core/db/gorm/internal/conn"
	"github.com/miyamo2/api.miyamo.today/core/db/gorm/internal/dial"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGet(t *testing.T) {
	type want struct {
		err error
	}

	type testCase struct {
		want       want
		wantErr    bool
		beforeFunc func()
	}

	tests := map[string]testCase{
		"happy_path": {
			beforeFunc: func() {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = &dialector

				conn.Mu.Lock()
				defer conn.Mu.Unlock()
				db, err := gorm.Open(dialector)
				if err != nil {
					panic(err)
				}
				conn.Instance = db
			},
		},
		"happy_path/connection_is_not_initialized": {
			beforeFunc: func() {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = &dialector

				conn.Mu.Lock()
				defer conn.Mu.Unlock()
				if err != nil {
					panic(err)
				}
				conn.Instance = nil
			},
		},
		"unhappy_path/gorm_dialector_is_not_initialized": {
			want:    want{err: ErrDialectorNotInitialized},
			wantErr: true,
			beforeFunc: func() {
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = nil
				conn.Mu.Lock()
				defer conn.Mu.Unlock()
				conn.Instance = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			_, err := Get(context.Background())
			if tt.wantErr {
				if !errors.Is(tt.want.err, err) {
					t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("Get() returns error. error = %v", err)
				return
			}
		})
	}
}

func TestInvalidate(t *testing.T) {
	type testCase struct {
		beforeFunc func()
	}

	tests := map[string]testCase{
		"happy_path": {
			beforeFunc: func() {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})

				conn.Mu.Lock()
				defer conn.Mu.Unlock()
				db, err := gorm.Open(dialector)
				if err != nil {
					panic(err)
				}
				conn.Instance = db
			},
		},
		"happy_path/gorm_connection_is_not_initialized": {
			beforeFunc: func() {
				conn.Mu.Lock()
				defer conn.Mu.Unlock()
				conn.Instance = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			Invalidate()
			if conn.Instance != nil {
				t.Errorf("Invalidate() is not invalidate the dialector.")
				return
			}
		})
	}
}

func TestInitializeDialector(t *testing.T) {
	type testCase struct {
		args          func() *gorm.Dialector
		wantOverwrite bool
		beforeFunc    func()
	}

	tests := map[string]testCase{
		"happy_path": {
			args: func() *gorm.Dialector {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				return &dialector
			},
			wantOverwrite: true,
			beforeFunc: func() {
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = nil
			},
		},
		"happy_path/gorm_dialector_is_already_initialized": {
			args: func() *gorm.Dialector {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				return &dialector
			},
			beforeFunc: func() {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = &dialector
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			args := tt.args()
			InitializeDialector(args)
			if tt.wantOverwrite {
				if dial.Instance != args {
					t.Errorf("InitializeDialector() is not updating the dialector.")
					return
				}
			} else if dial.Instance == args {
				t.Errorf("InitializeDialector() is updating the dialector.")
				return
			}
		})
	}
}

func TestInvalidateDialector(t *testing.T) {
	type testCase struct {
		beforeFunc func()
	}

	tests := map[string]testCase{
		"happy_path": {
			beforeFunc: func() {
				sqlDB, _, err := sqlmock.New()
				if err != nil {
					panic(err)
				}
				dialector := postgres.New(postgres.Config{
					Conn: sqlDB,
				})
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = &dialector
			},
		},
		"happy_path/gorm_dialector_is_not_initialized": {
			beforeFunc: func() {
				dial.Mu.Lock()
				defer dial.Mu.Unlock()
				dial.Instance = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			Invalidate()
			InvalidateDialector()
			if dial.Instance != nil {
				t.Errorf("InvalidateDialector() is not invalidate the dialector.")
				return
			}
		})
	}
}
