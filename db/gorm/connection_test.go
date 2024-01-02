package gorm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/db/gorm/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
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
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = &dialector
			},
		},
		"unhappy_path/gorm_connection_is_not_initialized": {
			want:    want{err: ErrDialectorNotInitialized},
			wantErr: true,
			beforeFunc: func() {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			_, err := Get()
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
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = nil
			},
		},
		"happy_path/gorm_connection_is_already_initialized": {
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
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = &dialector
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
				if internal.Dialector != args {
					t.Errorf("InitializeDialector() is not updating the dialector.")
					return
				}
			} else if internal.Dialector == args {
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
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = &dialector
			},
		},
		"happy_path/gorm_connection_is_not_initialized": {
			beforeFunc: func() {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Dialector = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			InvalidateDialector()
			if internal.Dialector != nil {
				t.Errorf("InvalidateDialector() is not invalidate the dialector.")
				return
			}
		})
	}
}
