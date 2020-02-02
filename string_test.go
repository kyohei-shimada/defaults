package defaults

import "testing"

func TestString(t *testing.T) {
	type args struct {
		value        string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "When value is empty string",
			args: args{value: "", defaultValue: "This is a default value."},
			want: "This is a default value.",
		},
		{
			name: "When value is not empty string",
			args: args{value: "Hello world!", defaultValue: "This is a default value."},
			want: "Hello world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
