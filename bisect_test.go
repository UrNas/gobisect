package gobisect

import (
	"reflect"
	"testing"
)

func Test_InsertInt(t *testing.T) {
	type args struct {
		slice []int
		idx   int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "InsertInt_1",
			args: args{
				slice: []int{1, 3},
				idx:   1,
				value: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "InsertInt_2",
			args: args{
				slice: []int{2, 4},
				idx:   1,
				value: 3,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "InsertInt_3",
			args: args{
				slice: []int{7, 8, 9},
				idx:   3,
				value: 10,
			},
			want: []int{7, 8, 9, 10},
		},
		{
			name: "InsertInt_4",
			args: args{
				slice: []int{0, 0, 0, 0},
				idx:   2,
				value: 1,
			},
			want: []int{0, 0, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertInt(tt.args.slice, tt.args.idx, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertString(t *testing.T) {
	type args struct {
		s     []string
		idx   int
		value string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "insertString_1",
			args: args{
				s:     []string{"firstname", "lastname"},
				idx:   1,
				value: "midlename",
			},
			want: []string{"firstname", "midlename", "lastname"},
		},
		{
			name: "insertString_2",
			args: args{
				s:     []string{"cat", "dog"},
				idx:   0,
				value: "go",
			},
			want: []string{"go", "cat", "dog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertString(tt.args.s, tt.args.idx, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBisecRightInt(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "BisectRightInt_1",
			args: args{
				slice: []int{2, 3, 5},
				value: 1,
			},
			want: 0,
		},
		{
			name: "BisectRightInt_2",
			args: args{
				slice: []int{2, 3, 5},
				value: 4,
			},
			want: 2,
		},
		{
			name: "BisectRightInt_3",
			args: args{
				slice: []int{2, 3, 5},
				value: 5,
			},
			want: 3,
		},
		{
			name: "BisectRightInt_4",
			args: args{
				slice: []int{2, 3, 5},
				value: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BisecRightInt(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("BisecRightInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBisecLeftInt(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "BisectLeftInt_1",
			args: args{
				slice: []int{1, 2},
				value: 0,
			},
			want: 0,
		},
		{
			name: "BisectLeftInt_1",
			args: args{
				slice: []int{1, 2, 4},
				value: 2,
			},
			want: 1,
		},
		{
			name: "BisectLeftInt_1",
			args: args{
				slice: []int{1, 1, 1},
				value: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BisecLeftInt(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("BisecLeftInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsortRightInt(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "InsortRightInt_1",
			args: args{
				slice: []int{1, 3},
				value: 2,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "InsortRightInt_2",
			args: args{
				slice: []int{1, 3},
				value: 4,
			},
			want: []int{1, 3, 4},
		},
		{
			name: "InsortRightInt_3",
			args: args{
				slice: []int{1, 3, 4, 5},
				value: 4,
			},
			want: []int{1, 3, 4, 4, 5},
		},
		{
			name: "InsortRightInt_4",
			args: args{
				slice: []int{0, 0, 0},
				value: 0,
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsortRightInt(tt.args.slice, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsortRightInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsortLeftInt(t *testing.T) {
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "InsortLeftInt_1",
			args: args{
				slice: []int{2, 3},
				value: 1,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "InsortLeftInt_2",
			args: args{
				slice: []int{2, 3},
				value: 4,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "InsortLeftInt_3",
			args: args{
				slice: []int{2, 4},
				value: 3,
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsortLeftInt(tt.args.slice, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsortLeftInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
