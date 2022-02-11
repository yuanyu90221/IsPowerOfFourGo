package power

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "4",
			args: args{n: 4},
			want: true,
		},
		{
			name: "0",
			args: args{n: 0},
			want: false,
		},
		{
			name: "1024",
			args: args{n: 1024},
			want: true,
		},
		{
			name: "192",
			args: args{n: 192},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
