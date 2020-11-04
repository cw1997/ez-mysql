package server

import (
	"reflect"
	"testing"
)

func TestResponseField_Build(t *testing.T) {
	tests := []struct {
		name   string
		fields ResponseField
		want   []byte
	}{
		{
			"TestResponseField_Build",
			ResponseField{
				"def",
				"test",
				"user",
				"user",
				"id",
				"id",
				63,
				10,
				3,
				0x4223,
				0,
			},
			[]byte{
				0x03, 0x64, 0x65, 0x66,
				0x04, 0x74, 0x65, 0x73, 0x74, 0x04, 0x75, 0x73,
				0x65, 0x72, 0x04, 0x75, 0x73, 0x65, 0x72, 0x02,
				0x69, 0x64, 0x02, 0x69, 0x64, 0x0c, 0x3f, 0x00,
				0x0a, 0x00, 0x00, 0x00, 0x03, 0x23, 0x42, 0x00,
				0x00, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet := &ResponseField{
				Catalog:       tt.fields.Catalog,
				Database:      tt.fields.Database,
				Table:         tt.fields.Table,
				OriginalTable: tt.fields.OriginalTable,
				Name:          tt.fields.Name,
				OriginalName:  tt.fields.OriginalName,
				CharsetNumber: tt.fields.CharsetNumber,
				Length:        tt.fields.Length,
				Type:          tt.fields.Type,
				Flags:         tt.fields.Flags,
				Decimals:      tt.fields.Decimals,
			}
			if got := packet.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = \n%v, \n want \n%v \n", got, tt.want)
			}
		})
	}
}

func TestResponseField_Resolve(t *testing.T) {
	type args struct {
		sliceByte []byte
	}
	tests := []struct {
		name   string
		fields ResponseField
		args   args
	}{
		{
			"TestResponseField_Resolve",
			ResponseField{
				"def",
				"test",
				"user",
				"user",
				"id",
				"id",
				63,
				10,
				3,
				0x4223,
				0,
			},
			args{
				[]byte{
					0x03, 0x64, 0x65, 0x66,
					0x04, 0x74, 0x65, 0x73, 0x74, 0x04, 0x75, 0x73,
					0x65, 0x72, 0x04, 0x75, 0x73, 0x65, 0x72, 0x02,
					0x69, 0x64, 0x02, 0x69, 0x64, 0x0c, 0x3f, 0x00,
					0x0a, 0x00, 0x00, 0x00, 0x03, 0x23, 0x42, 0x00,
					0x00, 0x00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new(ResponseField)
			got.Resolve(tt.args.sliceByte)
			if !reflect.DeepEqual(*got, tt.fields) {
				t.Errorf("Build() = %v, want %v", *got, tt.fields)
			}
		})
	}
}
