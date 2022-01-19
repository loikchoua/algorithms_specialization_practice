package merge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Same length",
			args: args{
				a: []int{1, 2},
				b: []int{3, 4},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Different length",
			args: args{
				a: []int{5},
				b: []int{3, 4},
			},
			want: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.a, tt.args.b)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Ok more than 1 element",
			args: args{
				input: []int{3, 7, 6},
			},
			want: []int{3, 6, 7},
		},
		{
			name: "Ok 1 element",
			args: args{
				input: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sort(tt.args.input)
			require.Equal(t, tt.want, got)
		})
	}
}
