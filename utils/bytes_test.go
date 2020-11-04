package utils

import (
	"reflect"
	"testing"
)


func TestByteSliceToInt(t *testing.T) {
	type args struct {
		byteSlice []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"TestByteSliceToInt",
			args{[]byte{76,},},
			76,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{76,0,},},
			76,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{76,0,0,0,},},
			76,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{76,0,0,0,0,0,0,0,},},
			76,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{76,0,0,0,0,0,0,0,},},
			76,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{1,0,0,0,0,0,0,0,},},
			1,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{1,1,0,0,0,0,0,0,},},
			1 + 256,
		},
		{
			"TestByteSliceToInt",
			args{[]byte{1,1,1,0,0,0,0,0,},},
			1 + 256 + 65536,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteSliceToInt(tt.args.byteSlice); got != tt.want {
				t.Errorf("ByteSliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToByteSlice(t *testing.T) {
	type args struct {
		num uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{

		{
			"TestIntToByteSlice",
			args{76},
			[]byte{76,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{76},
			[]byte{76,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{76},
			[]byte{76,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{76},
			[]byte{76,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{76},
			[]byte{76,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{1},
			[]byte{1,0,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{1 + 256},
			[]byte{1,1,0,0,0,0,0,0,},
		},
		{
			"TestIntToByteSlice",
			args{1 + 256 + 65536},
			[]byte{1,1,1,0,0,0,0,0,},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToByteSlice(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
