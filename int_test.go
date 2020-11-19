package defaults

import "testing"

func TestInt8(t *testing.T) {
	type args struct {
		first  int8
		others []int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []int8{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []int8{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []int8{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []int8{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		first  int16
		others []int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []int16{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []int16{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []int16{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []int16{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		first  int32
		others []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []int32{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []int32{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []int32{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []int32{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		first  int64
		others []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []int64{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []int64{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []int64{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []int64{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		first  int
		others []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "When arguments have a zero, return zero",
			args: args{first: 0, others: []int{}},
			want: 0,
		},
		{
			name: "When arguments have a non-zero, return first value",
			args: args{first: 123, others: []int{}},
			want: 123,
		},
		{
			name: "When arguments have 2 values or more and all of 'values' are 0, return 0",
			args: args{first: 0, others: []int{0, 0, 0, 0}},
			want: 0,
		},
		{
			name: "When arguments have 2 values or more and 'values' include non-zero, return first non-zero value",
			args: args{first: 0, others: []int{0, 3, 1, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
