package server

import (
	"bytes"
	"reflect"
	"testing"
)

func TestResponseOKBuild(t *testing.T) {
	input := new(ResponseOK)
	input.ResponseCode = 0x00
	input.AffectedRows = 0
	input.LastInsertID = 0
	input.ServerStatus = 0x0002
	input.Warnings = 0

	actual := input.Build()
	except := []byte{
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
	}

	if !bytes.Equal(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}

func TestResponseOKResolve(t *testing.T) {
	input := []byte{
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
	}

	actual := new(ResponseOK)
	actual.Resolve(input)

	except := new(ResponseOK)
	except.ResponseCode = 0x00
	except.AffectedRows = 0
	except.LastInsertID = 0
	except.ServerStatus = 0x0002
	except.Warnings = 0

	if !reflect.DeepEqual(actual, except) {
		t.Errorf("input: \n [%+v] \n actual \n [%+v] \n except \n [%+v] \n \n", input, actual, except)
	}
}
