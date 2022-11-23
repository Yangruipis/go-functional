package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
)

func TestRange(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "00",
			args: args{
				start: 0,
				end:   4,
				step:  1,
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "01",
			args: args{
				start: 0,
				end:   4,
				step:  2,
			},
			want: []int{0, 2},
		},
		{
			name: "02",
			args: args{
				start: 0,
				end:   4,
				step:  3,
			},
			want: []int{0, 3},
		},
		{
			name: "03",
			args: args{
				start: 0,
				end:   4,
				step:  4,
			},
			want: []int{0},
		},
		{
			name: "04",
			args: args{
				start: 0,
				end:   4,
				step:  5,
			},
			want: []int{0},
		},
		{
			name: "05",
			args: args{
				start: 0,
				end:   0,
				step:  1,
			},
			want: []int{},
		},
		{
			name: "06",
			args: args{
				start: 0,
				end:   4,
				step:  0,
			},
			want: []int{},
		},
		{
			name: "07",
			args: args{
				start: 4,
				end:   0,
				step:  -1,
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Range(tt.args.start, tt.args.end, tt.args.step); !reflect.DeepEqual(fn.ToSlice(got), tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
