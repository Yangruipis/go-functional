package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
)

func TestRepeat(t *testing.T) {
	type args struct {
		num  int
		item string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "00",
			args: args{
				num:  4,
				item: "a",
			},
			want: []string{"a", "a", "a", "a"},
		},
		{
			name: "01",
			args: args{
				num:  0,
				item: "a",
			},
			want: []string{},
		},
		{
			name: "02",
			args: args{
				num:  -1,
				item: "a",
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Repeat(tt.args.item, tt.args.num); !reflect.DeepEqual(fn.ToSlice(got), tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
