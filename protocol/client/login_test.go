package client

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLoginBuild(t *testing.T) {
	input := new(Login)
	input.ClientCapabilities = 0xa685
	input.ExtendedClientCapabilities = 0x007f
	input.MAXPacket = 1073741824
	input.Charset = 33
	input.Username = "root"
	input.PasswordLength = uint8(len([]byte{
		0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
		0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
		0xad, 0x23, 0x40, 0x43,
	}))
	input.Password = []byte{
		0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
		0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
		0xad, 0x23, 0x40, 0x43,
	}
	input.ClientAuthPlugin = "mysql_native_password"
	input.ConnectionAttributes = []byte{
		0x03, 0x5f, 0x6f, 0x73, 0x05, 0x57, 0x69,
		0x6e, 0x36, 0x34, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
		0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
		0x08, 0x6c, 0x69, 0x62, 0x6d, 0x79, 0x73, 0x71,
		0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x32,
		0x35, 0x36, 0x37, 0x32, 0x07, 0x5f, 0x74, 0x68,
		0x72, 0x65, 0x61, 0x64, 0x05, 0x31, 0x36, 0x37,
		0x34, 0x30, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74,
		0x66, 0x6f, 0x72, 0x6d, 0x05, 0x41, 0x4d, 0x44,
		0x36, 0x34, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65,
		0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
		0x6f, 0x6e, 0x07, 0x31, 0x30, 0x2e, 0x31, 0x2e,
		0x32, 0x34,
	}

	actual := input.Build()
	except := []byte{
		0x85, 0xa6, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x40,
		0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x8d, 0xfb,
		0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b, 0xe8, 0x1d,
		0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7, 0xad, 0x23,
		0x40, 0x43, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f,
		0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
		0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x00,
		0x61, 0x03, 0x5f, 0x6f, 0x73, 0x05, 0x57, 0x69,
		0x6e, 0x36, 0x34, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
		0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
		0x08, 0x6c, 0x69, 0x62, 0x6d, 0x79, 0x73, 0x71,
		0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x32,
		0x35, 0x36, 0x37, 0x32, 0x07, 0x5f, 0x74, 0x68,
		0x72, 0x65, 0x61, 0x64, 0x05, 0x31, 0x36, 0x37,
		0x34, 0x30, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74,
		0x66, 0x6f, 0x72, 0x6d, 0x05, 0x41, 0x4d, 0x44,
		0x36, 0x34, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65,
		0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
		0x6f, 0x6e, 0x07, 0x31, 0x30, 0x2e, 0x31, 0x2e,
		0x32, 0x34,
	}

	if !bytes.Equal(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}

func TestLoginResolve(t *testing.T) {
	input := []byte{
		0x85, 0xa6, 0x7f, 0x00, 0x00, 0x00, 0x00, 0x40,
		0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x8d, 0xfb,
		0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b, 0xe8, 0x1d,
		0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7, 0xad, 0x23,
		0x40, 0x43, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f,
		0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
		0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x00,
		0x61, 0x03, 0x5f, 0x6f, 0x73, 0x05, 0x57, 0x69,
		0x6e, 0x36, 0x34, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
		0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
		0x08, 0x6c, 0x69, 0x62, 0x6d, 0x79, 0x73, 0x71,
		0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x32,
		0x35, 0x36, 0x37, 0x32, 0x07, 0x5f, 0x74, 0x68,
		0x72, 0x65, 0x61, 0x64, 0x05, 0x31, 0x36, 0x37,
		0x34, 0x30, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74,
		0x66, 0x6f, 0x72, 0x6d, 0x05, 0x41, 0x4d, 0x44,
		0x36, 0x34, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65,
		0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
		0x6f, 0x6e, 0x07, 0x31, 0x30, 0x2e, 0x31, 0x2e,
		0x32, 0x34,
	}

	actual := new(Login)
	actual.Resolve(input)

	except := new(Login)
	except.ClientCapabilities = 0xa685
	except.ExtendedClientCapabilities = 0x007f
	except.MAXPacket = 1073741824
	except.Charset = 33
	except.Username = "root"
	except.PasswordLength = uint8(len([]byte{
		0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
		0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
		0xad, 0x23, 0x40, 0x43,
	}))
	except.Password = []byte{
		0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
		0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
		0xad, 0x23, 0x40, 0x43,
	}
	except.ClientAuthPlugin = "mysql_native_password"
	except.ConnectionAttributes = []byte{
		0x03, 0x5f, 0x6f, 0x73, 0x05, 0x57, 0x69,
		0x6e, 0x36, 0x34, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
		0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
		0x08, 0x6c, 0x69, 0x62, 0x6d, 0x79, 0x73, 0x71,
		0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x32,
		0x35, 0x36, 0x37, 0x32, 0x07, 0x5f, 0x74, 0x68,
		0x72, 0x65, 0x61, 0x64, 0x05, 0x31, 0x36, 0x37,
		0x34, 0x30, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74,
		0x66, 0x6f, 0x72, 0x6d, 0x05, 0x41, 0x4d, 0x44,
		0x36, 0x34, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65,
		0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
		0x6f, 0x6e, 0x07, 0x31, 0x30, 0x2e, 0x31, 0x2e,
		0x32, 0x34,
	}

	if !reflect.DeepEqual(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}

func TestMysqlNativePassword(t *testing.T) {
	type args struct {
		scramble []byte
		password string
	}
	tests := []struct {
		name string
		args args
		want []byte
	} {
		{
			"root cw1997",
			args{
				[]byte{
					0x6c, 0x72, 0x3b, 0x26, 0x15, 0x3f, 0x4b, 0x2a,
					0x50, 0x40, 0x67, 0x18, 0x6c, 0x48, 0x2a, 0x4f,
					0x72, 0x6e, 0x2c, 0x11,
				},
				"cw1997",
			},
			[]byte{
				0x8d, 0xfb, 0x92, 0x8e, 0xa8, 0x79, 0x59, 0x6b,
				0xe8, 0x1d, 0xfe, 0xe4, 0x65, 0xa4, 0xc2, 0xe7,
				0xad, 0x23, 0x40, 0x43,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MysqlNativePassword(tt.args.scramble, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlNativePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}