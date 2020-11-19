package defaults

import "testing"

func TestUint8(t *testing.T) {
	type args struct {
		first  uint8
		others []uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []uint8{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []uint8{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []uint8{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []uint8{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		first  uint16
		others []uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []uint16{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []uint16{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []uint16{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []uint16{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		first  uint32
		others []uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []uint32{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []uint32{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []uint32{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []uint32{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		first  uint64
		others []uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []uint64{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []uint64{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []uint64{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []uint64{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		first  uint
		others []uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []uint{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []uint{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []uint{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []uint{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}
