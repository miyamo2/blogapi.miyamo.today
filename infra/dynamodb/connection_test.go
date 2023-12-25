package dynamodb

import (
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi-core/infra/dynamodb/client"
	"github.com/miyamo2/blogapi-core/infra/dynamodb/internal"
	mclient "github.com/miyamo2/blogapi-core/internal/mock/infra/dynamodb/client"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGet(t *testing.T) {
	type want struct {
		err error
	}

	type testCase struct {
		want       want
		wantErr    bool
		client     func(ctrl *gomock.Controller) client.Client
		beforeFunc func(client client.Client)
	}

	tests := map[string]testCase{
		"happy_path": {
			client: func(ctrl *gomock.Controller) client.Client {
				return mclient.NewMockClient(ctrl)
			},
			beforeFunc: func(client client.Client) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = client
			},
		},
		"unhappy_path/client_is_not_initialized": {
			want:    want{err: ErrClientNotInitialized},
			wantErr: true,
			client: func(ctrl *gomock.Controller) client.Client {
				return nil
			},
			beforeFunc: func(client client.Client) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = client
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			tt.beforeFunc(tt.client(mockCtrl))
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

func TestInitialize(t *testing.T) {
	type testCase struct {
		client        func(ctrl *gomock.Controller) client.Client
		beforeFunc    func(client client.Client)
		wantOverwrite bool
	}

	tests := map[string]testCase{
		"happy_path": {
			client: func(ctrl *gomock.Controller) client.Client {
				return mclient.NewMockClient(ctrl)
			},
			wantOverwrite: true,
			beforeFunc: func(client client.Client) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = nil
			},
		},
		"happy_path/connection_is_already_initialized": {
			client: func(ctrl *gomock.Controller) client.Client {
				return mclient.NewMockClient(ctrl)
			},
			beforeFunc: func(client client.Client) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = &mclient.MockClient{}
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			clt := tt.client(mockCtrl)
			tt.beforeFunc(clt)
			Initialize(clt)
			if tt.wantOverwrite {
				if internal.Client != clt {
					t.Errorf("Initialize() is not updating the client.")
					return
				}
			} else if internal.Client == clt {
				t.Errorf("Initialize() is updating the client.")
				return
			}
		})
	}
}

func TestInvalidate(t *testing.T) {
	type testCase struct {
		beforeFunc func(ctrl *gomock.Controller)
	}

	tests := map[string]testCase{
		"happy_path": {
			beforeFunc: func(ctrl *gomock.Controller) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = mclient.NewMockClient(ctrl)
			},
		},
		"happy_path/gorm_connection_is_not_initialized": {
			beforeFunc: func(ctrl *gomock.Controller) {
				internal.Lock.Lock()
				defer internal.Lock.Unlock()
				internal.Client = nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			tt.beforeFunc(mockCtrl)
			Invalidate()
			if internal.Client != nil {
				t.Errorf("Invalidate() is not invalidate the dialector.")
				return
			}
		})
	}
}
