package db

import (
	"reflect"
	"testing"
)

func TestPagination_Cursor(t *testing.T) {
	type testCase struct {
		sut  *Pagination
		want string
	}

	tests := map[string]testCase{
		"happy_path": {
			sut: &Pagination{
				cursor: "happy_path",
			},
			want: "happy_path",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := tt.sut
			if got := p.Cursor(); got != tt.want {
				t.Errorf("Cursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_Limit(t *testing.T) {
	type testCase struct {
		sut  *Pagination
		want int
	}

	tests := map[string]testCase{
		"happy_path": {
			sut: &Pagination{
				limit: 10,
			},
			want: 10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := tt.sut
			if got := p.Limit(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_NextPaging(t *testing.T) {
	type testCase struct {
		sut  *Pagination
		want bool
	}

	tests := map[string]testCase{
		"happy_path": {
			sut: &Pagination{
				isNextPaging: true,
			},
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := tt.sut
			if got := p.IsNextPaging(); got != tt.want {
				t.Errorf("IsNextPaging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_PreviousPaging(t *testing.T) {
	type testCase struct {
		sut  *Pagination
		want bool
	}

	tests := map[string]testCase{
		"happy_path": {
			sut: &Pagination{
				isPreviousPaging: true,
			},
			want: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p := tt.sut
			if got := p.IsPreviousPaging(); got != tt.want {
				t.Errorf("IsPreviousPaging() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithNextPaging(t *testing.T) {
	type args struct {
		limit  int
		cursor string
	}
	type testCase struct {
		args args
		sut  *Pagination
		want Pagination
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				limit:  10,
				cursor: "happy_path",
			},
			sut: &Pagination{},
			want: Pagination{
				cursor:           "happy_path",
				limit:            10,
				isPreviousPaging: false,
				isNextPaging:     true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			WithNextPaging(tt.args.limit, &tt.args.cursor)(tt.sut)
			if !reflect.DeepEqual(*tt.sut, tt.want) {
				t.Errorf("sut expected after executing WithNextPaging() = %v, actual %v", tt.want, *tt.sut)
			}
		})
	}
}

func TestWithPreviousPaging(t *testing.T) {
	type args struct {
		limit  int
		cursor string
	}
	type testCase struct {
		args args
		sut  *Pagination
		want Pagination
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				limit:  10,
				cursor: "happy_path",
			},
			sut: &Pagination{},
			want: Pagination{
				cursor:           "happy_path",
				limit:            10,
				isPreviousPaging: true,
				isNextPaging:     false,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			WithPreviousPaging(tt.args.limit, &tt.args.cursor)(tt.sut)
			if !reflect.DeepEqual(*tt.sut, tt.want) {
				t.Errorf("sut expected after executing WithPreviousPaging() = %v, actual %v", tt.want, *tt.sut)
			}
		})
	}
}
