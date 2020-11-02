package utils

import (
	"reflect"
	"testing"
)

func TestSHA1(t *testing.T) {
	type args struct {
		byteSlice []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"TestSHA1 root",
			args{[]byte("root")},
			[]byte{220, 118, 233, 240, 192, 0, 110, 143, 145, 158, 12, 81, 92, 102, 219, 186, 57, 130, 247, 133},
		},
		{
			"TestSHA1 cw1997",
			args{[]byte("cw1997")},
			[]byte{76, 27, 238, 135, 113, 255, 224, 183, 157, 191, 131, 105, 173, 153, 234, 199, 182, 247, 14, 186},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.byteSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}