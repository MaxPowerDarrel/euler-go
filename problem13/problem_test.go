package problem13

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Base case",
			want: "5537376230",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
