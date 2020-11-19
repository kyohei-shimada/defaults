package defaults

import "testing"

func TestString(t *testing.T) {
	type args struct {
		first  string
		others []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "When arguments have a empty string, return empty string",
			args: args{first: "", others: []string{}},
			want: "",
		},
		{
			name: "When arguments have a non-empty string, return first value",
			args: args{first: "Hello World!", others: []string{}},
			want: "Hello World!",
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are emPty string, return empty string",
			args: args{first: "", others: []string{"", "", ""}},
			want: "",
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: "", others: []string{"", "foo", "bar", ""}},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
