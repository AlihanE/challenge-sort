package sorts

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "lo",
			args: args{
				arr: []string{"b", "a", "c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "lol",
			args: args{
				arr: []string{"b", "a", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "lol2",
			args: args{
				arr: []string{"ab", "aa", "ac"},
			},
			want: []string{"aa", "ab", "ac"},
		},
		{
			name: "lol3",
			args: args{
				arr: []string{"a", "aa", "b", "ac"},
			},
			want: []string{"a", "aa", "ac", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
