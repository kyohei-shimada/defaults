package defaults

import "testing"

func TestInt8(t *testing.T) {
	type args struct {
		value        int8
		defaultValue int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 100},
			want: 100,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: -123, defaultValue: 100},
			want: -123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		value        int16
		defaultValue int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 1000},
			want: 1000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: -12345, defaultValue: 100},
			want: -12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		value        int32
		defaultValue int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 100000},
			want: 100000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: -100000, defaultValue: 100},
			want: -100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		value        int64
		defaultValue int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 0x7000000000000000},
			want: 0x7000000000000000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: -1234567890, defaultValue: 0x6000000000000000},
			want: -1234567890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		value        int
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 12345678},
			want: 12345678,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: -12345678, defaultValue: 12345678},
			want: -12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
