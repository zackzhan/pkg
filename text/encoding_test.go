package text

import "testing"

func TestUtf72Utf8(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "中文",
			args: args{
				name: "中文",
			},
			want: "中文",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf72Utf8(tt.args.name); got != tt.want {
				t.Errorf("Utf72Utf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
