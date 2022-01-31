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

func TestSortAndCountInversions(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name                   string
		args                   args
		wantSorted             []int
		wantNumberOfInversions int
	}{
		{
			name: "1 inversion",
			args: args{
				[]int{2, 1},
			},
			wantSorted:             []int{1, 2},
			wantNumberOfInversions: 1,
		},
		{
			name: "3 inversions",
			args: args{
				[]int{1, 3, 5, 2, 4, 6},
			},
			wantSorted:             []int{1, 2, 3, 4, 5, 6},
			wantNumberOfInversions: 3,
		},
		{
			name: "4 inversions",
			args: args{
				[]int{3, 2, 1, 5, 7, 6},
			},
			wantSorted:             []int{1, 2, 3, 5, 6, 7},
			wantNumberOfInversions: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSorted, gotNumberOfInversions := SortAndCountInversions(tt.args.a)
			require.Equal(t, tt.wantSorted, gotSorted)
			require.Equal(t, tt.wantNumberOfInversions, gotNumberOfInversions)
		})
	}
}
