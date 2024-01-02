package gorm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
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
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = &dialector

				singletonDB.mu.Lock()
				defer singletonDB.mu.Unlock()
				db, err := gorm.Open(dialector)
				if err != nil {
					panic(err)
				}
				singletonDB.db = db
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
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = &dialector

				singletonDB.mu.Lock()
				defer singletonDB.mu.Unlock()
				if err != nil {
					panic(err)
				}
				singletonDB.db = nil
			},
		},
		"unhappy_path/gorm_dialector_is_not_initialized": {
			want:    want{err: ErrDialectorNotInitialized},
			wantErr: true,
			beforeFunc: func() {
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = nil
				singletonDB.mu.Lock()
				defer singletonDB.mu.Unlock()
				singletonDB.db = nil
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

				singletonDB.mu.Lock()
				defer singletonDB.mu.Unlock()
				db, err := gorm.Open(dialector)
				if err != nil {
					panic(err)
				}
				singletonDB.db = db
			},
		},
		"happy_path/gorm_connection_is_not_initialized": {
			beforeFunc: func() {
				singletonDB.mu.Lock()
				defer singletonDB.mu.Unlock()
				singletonDB.db = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			Invalidate()
			if singletonDB.db != nil {
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
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = nil
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
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = &dialector
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
				if singletonDialector.dialector != args {
					t.Errorf("InitializeDialector() is not updating the dialector.")
					return
				}
			} else if singletonDialector.dialector == args {
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
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = &dialector
			},
		},
		"happy_path/gorm_dialector_is_not_initialized": {
			beforeFunc: func() {
				singletonDialector.mu.Lock()
				defer singletonDialector.mu.Unlock()
				singletonDialector.dialector = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.beforeFunc()
			}
			InvalidateDialector()
			if singletonDialector.dialector != nil {
				t.Errorf("InvalidateDialector() is not invalidate the dialector.")
				return
			}
		})
	}
}
