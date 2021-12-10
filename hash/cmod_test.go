package hash

import "testing"

func Test_Mod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: struct {
				a int
				b int
			}{a: 10, b: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mod() = %v, want %v", got, tt.want)
			}
		})
	}
}
