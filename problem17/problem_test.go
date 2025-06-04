package problem17

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Base case",
			args: args{
				input: 5,
			},
			want: 19,
		},
		{
			name: "Test case",
			args: args{
				input: 1000,
			},
			want: 21124,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.input); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
