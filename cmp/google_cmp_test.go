package cmp

import "testing"

func TestGoogle_cmp(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "比较通过",
			args: args{
				a:1,
				b:1,
			},
			want: true,
		},
		{
			name: "比较不通过通过",
			args: args{
				a:1,
				b:2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Google_cmp(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Google_cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGoogle_cmp(b *testing.B) {
	x := map[string]interface{}{
		"a":1,
		"b":[]int{1,2,3},
	}
	y:= map[string]interface{}{
		"a":1,
		"b":[]int{1,2,3},
		"c":1.25,
	}
	for i := 0; i < b.N; i++ {
		Google_cmp(x, y)
	}
}