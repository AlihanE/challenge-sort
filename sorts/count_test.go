package sorts

import "testing"

func TestCountingSort(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "lol",
			args: args{
				words: []string{"b", "a", "c"},
			},
		},
		{
			name: "lol2",
			args: args{
				words: []string{"ab", "aa", "ac"},
			},
		},
		{
			name: "lol3",
			args: args{
				words: []string{"a", "aa", "b", "ac"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountingSort(tt.args.words)
		})
	}
}
