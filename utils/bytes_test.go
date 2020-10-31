package utils

import (
	"bytes"
	"testing"
)

func TestByteSliceToInt(t *testing.T) {
	type TestCase struct {
		input []byte
		except int
	}
	testCase := []TestCase{
		{
			input: []byte{76,0,},
			except: 76,
		},
		{
			input: []byte{76,},
			except: 76,
		},
		{
			input: []byte{76,0,0,0,},
			except: 76,
		},
		{
			input: []byte{76,0,0,0,0,0,0,0,},
			except: 76,
		},
		{
			input: []byte{76,0,0,0,0,0,0,0,},
			except: 76,
		},

		{
			input: []byte{1,0,0,0,0,0,0,0,},
			except: 1,
		},
		{
			input: []byte{1,1,0,0,0,0,0,0,},
			except: 1 + 256,
		},
		{
			input: []byte{1,1,1,0,0,0,0,0,},
			except: 1 + 256 + 65536,
		},
	}

	for _, v := range testCase {
		actual, err := ByteSliceToInt(v.input)
		except := int64(v.except)
		if actual != except || err != nil {
			t.Errorf("input: [%+v], actual: [%+v], except: [%+v], error: [%+v] \n", v.input, actual, except, err)
		}
	}
}

func TestIntToByteSlice(t *testing.T) {
	type TestCase struct {
		input int
		except []byte
	}
	testCase := []TestCase{
		{
			input: 76,
			except: []byte{76,0,0,0,0,0,0,0,},
		},
		{
			input: 76,
			except: []byte{76,0,0,0,0,0,0,0,},
		},
		{
			input: 76,
			except: []byte{76,0,0,0,0,0,0,0,},
		},
		{
			input: 76,
			except: []byte{76,0,0,0,0,0,0,0,},
		},
		{
			input: 76,
			except: []byte{76,0,0,0,0,0,0,0,},
		},

		{
			input: 1,
			except: []byte{1,0,0,0,0,0,0,0,},
		},
		{
			input: 1 + 256,
			except: []byte{1,1,0,0,0,0,0,0,},
		},
		{
			input: 1 + 256 + 65536,
			except: []byte{1,1,1,0,0,0,0,0,},
		},
	}

	for _, v := range testCase {
		actual, err := IntToByteSlice(int64(v.input))
		except := v.except
		if !bytes.Equal(actual, except) {
			t.Errorf("input: [%+v], actual: [%+v], except: [%+v], error: [%+v] \n", v.input, actual, except, err)
		}
	}
}
