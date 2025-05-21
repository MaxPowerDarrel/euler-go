package problem8

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
				input: 4,
			},
			want: 5832,
		},
		{
			name: "Test case",
			args: args{
				input: 13,
			},
			want: 23514624000,
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
