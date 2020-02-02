package defaults

import "testing"

func TestUint8(t *testing.T) {
	type args struct {
		value        uint8
		defaultValue uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 100},
			want: 100,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: 123, defaultValue: 100},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		value        uint16
		defaultValue uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 1000},
			want: 1000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: 12345, defaultValue: 100},
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		value        uint32
		defaultValue uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 100000},
			want: 100000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: 100000, defaultValue: 100},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		value        uint64
		defaultValue uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 0x7000000000000000},
			want: 0x7000000000000000,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: 1234567890, defaultValue: 0x6000000000000000},
			want: 1234567890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		value        uint
		defaultValue uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "When 'value' is 0",
			args: args{value: 0, defaultValue: 12345678},
			want: 12345678,
		},
		{
			name: "When 'value' is not 0",
			args: args{value: 12345678, defaultValue: 12345678},
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}
