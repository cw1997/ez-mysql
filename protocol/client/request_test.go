package client

import (
	"bytes"
	"reflect"
	"testing"
)

func TestRequestBuild(t *testing.T) {
	input := new(Request)
	input.Command = 3
	input.Statement = "select version()"

	actual := input.Build()
	except := []byte{
		0x03, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x20,
		0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x28,
		0x29,
	}

	if !bytes.Equal(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}

func TestRequestResolve(t *testing.T) {
	input := []byte{
		0x03, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x20,
		0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x28,
		0x29,
	}

	actual := new(Request)
	actual.Resolve(input)

	except := new(Request)
	except.Command = 3
	except.Statement = "select version()"

	if !reflect.DeepEqual(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}
