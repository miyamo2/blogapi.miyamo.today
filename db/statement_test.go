package db

import (
	"reflect"
	"testing"
)

func TestMultipleStatementResult_Get(t *testing.T) {
	type testCase struct {
		sut  func() *MultipleStatementResult[string]
		want interface{}
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				r.value = []string{"happy_path"}
				return r
			},
			want: []string{"happy_path"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleStatementResult_HasNext(t *testing.T) {
	type testCase struct {
		sut  func() *MultipleStatementResult[string]
		want bool
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				r.value = []string{"happy_path", "happy_path"}
				return r
			},
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.HasNext(); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleStatementResult_Next(t *testing.T) {
	type testCase struct {
		name string
		sut  func() *MultipleStatementResult[string]
		want string
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				r.value = []string{"1", "2"}
				return r
			},
			want: "1",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleStatementResult_Set(t *testing.T) {
	type testCase struct {
		sut    func() *MultipleStatementResult[string]
		args   interface{}
		expect []string
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				return r
			},
			args:   []string{"happy_path", "happy_path"},
			expect: []string{"happy_path", "happy_path"},
		},
		"happy_path/value_already_set": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				r.hasBeenSet = true
				r.value = []string{"happy_path/value_already_set", "happy_path/value_already_set"}
				return r
			},
			args:   []string{"overwriting"},
			expect: []string{"happy_path/value_already_set", "happy_path/value_already_set"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			sut.Set(tt.args)
			if got := sut.value; !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("Next() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestMultipleStatementResult_StrictGet(t *testing.T) {
	type testCase struct {
		name string
		sut  func() *MultipleStatementResult[string]
		want []string
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *MultipleStatementResult[string] {
				r := NewMultipleStatementResult[string]()
				r.value = []string{"happy_path"}
				return r
			},
			want: []string{"happy_path"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.StrictGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrictGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleStatementResult_Get(t *testing.T) {
	type testCase struct {
		name string
		sut  func() *SingleStatementResult[string]
		want interface{}
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *SingleStatementResult[string] {
				r := NewSingleStatementResult[string]()
				r.value = "happy_path"
				return r
			},
			want: "happy_path",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleStatementResult_Set(t *testing.T) {
	type testCase struct {
		sut    func() *SingleStatementResult[string]
		args   interface{}
		expect string
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *SingleStatementResult[string] {
				r := NewSingleStatementResult[string]()
				return r
			},
			args:   "happy_path",
			expect: "happy_path",
		},
		"happy_path/value_already_set": {
			sut: func() *SingleStatementResult[string] {
				r := NewSingleStatementResult[string]()
				r.hasBeenSet = true
				r.value = "happy_path/value_already_set"
				return r
			},
			args:   "overwriting",
			expect: "happy_path/value_already_set",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			sut.Set(tt.args)
			if got := sut.value; !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("Next() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestSingleStatementResult_StrictGet(t *testing.T) {
	type testCase struct {
		name string
		sut  func() *SingleStatementResult[string]
		want string
	}
	tests := map[string]testCase{
		"happy_path": {
			sut: func() *SingleStatementResult[string] {
				r := NewSingleStatementResult[string]()
				r.value = "happy_path"
				return r
			},
			want: "happy_path",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sut := tt.sut()
			if got := sut.StrictGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrictGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
