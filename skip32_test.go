package skip32_test

import (
	"encoding/hex"
	"testing"

	"github.com/jmhobbs/skip32"
)

func TestSkip32Default(t *testing.T) {
	var (
		key       [10]uint8 = [10]uint8{0x00, 0x99, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}
		input     uint32    = 0x33221100
		encrypted uint32    = 0x819d5f1f
	)

	actual := skip32.Encrypt(key, input)
	if actual != encrypted {
		t.Errorf("error encrypting\nexpected: %v\n   actual: %v", encrypted, actual)
	}

	iso := skip32.Decrypt(key, actual)
	if iso != input {
		t.Errorf("error decrypting\nexpected: %v\n   actual: %v", input, iso)
	}
}

// https://github.com/alestic/Crypt-Skip32/blob/master/t/01basic.t
func TestSkip32PerlSuite(t *testing.T) {
	rawKey, _ := hex.DecodeString("DE2624BD4FFC4BF09DAB")
	key := [10]byte{}
	copy(key[:], rawKey)

	tests := []struct{
		Input uint32
		Expected uint32
	} {
		{0,   78612854},
		{3, 3719912389},
		{21, 1463300585},
		{147, 1277082297},
		{1029, 2878029910},
		{7203, 4086218104},
		{50421, 2588160464},
		{352947, 2703568194},
		{2470629, 2600508864},
		{17294403, 4119915301},
		{121060821, 4266122367},
		{847425747, 2671425558},
		{4294967295,  949651845},
	}

	for _, test := range tests {
		actual := skip32.Encrypt(key, test.Input)
		if actual != test.Expected {
			t.Errorf("error encrypting\nexpected: %v\n   actual: %v", test.Expected, actual)
		}
		
		decrypted := skip32.Decrypt(key, actual)
		if decrypted != test.Input {
			t.Errorf("error decrypting\nexpected: %v\n   actual: %v", test.Input, decrypted)
		}		
	}
}
