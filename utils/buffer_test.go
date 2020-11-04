package utils

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadInteger(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
		length uint8
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"TestReadInteger",
			args{bytes.NewBuffer([]byte{0x4e, 0x00, 0x00,}), 3},
			78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadInteger(tt.args.buffer, tt.args.length); got != tt.want {
				t.Errorf("ReadInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadLengthCodedBinary(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
	}
	tests := []struct {
		name          string
		args          args
		wantByteSlice []byte
		wantLength    uint64
	}{
		{
			"TestReadLengthCodedBinary",
			args{bytes.NewBuffer(
				[]byte{
					0x14,
					0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
					0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
					0xad, 0x23, 0x40, 0x43,
				},
			)},
			[]byte{
				0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
				0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
				0xad, 0x23, 0x40, 0x43,
			},
			0x14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotByteSlice, gotLength := ReadLengthCodedBinary(tt.args.buffer)
			if !reflect.DeepEqual(gotByteSlice, tt.wantByteSlice) {
				t.Errorf("ReadLengthCodedBinary() gotByteSlice = %v, want %v", gotByteSlice, tt.wantByteSlice)
			}
			if gotLength != tt.wantLength {
				t.Errorf("ReadLengthCodedBinary() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func TestReadNullTerminatedString(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{bytes.NewBuffer([]byte{0x72, 0x6f, 0x6f, 0x74, 0x00, })},
			"root",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadNullTerminatedString(tt.args.buffer); got != tt.want {
				t.Errorf("ReadNullTerminatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteInteger(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
		length uint8
		number uint64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestWriteInteger",
			args{bytes.NewBuffer([]byte{}), 3, 78},
			[]byte{0x4e, 0x00, 0x00,},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteInteger(tt.args.buffer, tt.args.length, tt.args.number)
			got := tt.args.buffer.Bytes()
			want := tt.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("TestWriteInteger() = %v, want %v", got, want)
			}
		})
	}
}

func TestWriteLengthCodedBinary(t *testing.T) {
	type args struct {
		buffer    *bytes.Buffer
		byteSlice []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestWriteLengthCodedBinary",
			args{
				bytes.NewBuffer([]byte{}),
				[]byte{
					0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
					0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
					0xad, 0x23, 0x40, 0x43,
				},
			},
			[]byte{
				0x14,
				0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
				0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
				0xad, 0x23, 0x40, 0x43,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteLengthCodedBinary(tt.args.buffer, tt.args.byteSlice)
			got := tt.args.buffer.Bytes()
			want := tt.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("TestWriteNullTerminatedString() = %v, want %v", got, want)
			}
		})
	}
}

func TestWriteNullTerminatedString(t *testing.T) {
	type args struct {
		buffer *bytes.Buffer
		str    string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestWriteNullTerminatedString",
			args{bytes.NewBuffer([]byte{}), "root"},
			[]byte{0x72, 0x6f, 0x6f, 0x74, 0x00, },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteNullTerminatedString(tt.args.buffer, tt.args.str)
			got := tt.args.buffer.Bytes()
			want := tt.want
			if !reflect.DeepEqual(got, want) {
				t.Errorf("TestWriteNullTerminatedString() = %v, want %v", got, want)
			}
		})
	}
}

func TestWriteRepeat(t *testing.T) {
	type args struct {
		buffer    *bytes.Buffer
		byteSlice []byte
		count     int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestWriteRepeat",
			args{
				bytes.NewBuffer([]byte{}),
				[]byte{0},
				8,
			},
			[]byte{0,0,0,0,0,0,0,0,},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
