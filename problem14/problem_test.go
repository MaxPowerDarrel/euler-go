package problem14

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "Base case",
			want:  837799,
			want1: 525,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Solve()
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Solve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
