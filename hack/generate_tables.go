package main

/*
This is a code generator for the `table_test.go` file.
It generates test tables for Skip32 using the original C
binary.  You will need to download and compile the C
source (`gcc skip32.c -o main`) and place the executable
in this directory.

Then run `go run . | gofmt >| ../table_test.go` to generate
the test file.  The random tests will change with each run.

This will take quite a while to run with large numbers
of tests, so be patient.
*/

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"text/template"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

var outputTemplate string = `package skip32_test

/* GENERATED FILE - DO NOT EDIT */
/* See hack/generate_tables.go */

import (
	"encoding/hex"
	"testing"

	"github.com/jmhobbs/skip32"
)

var testKey [10]byte = [10]byte{}

func init () {
	key, err := hex.DecodeString("00998877665544332211")
	if err != nil {
		panic(err)
	}
	copy(testKey[:], key)
}

var sequentialTestValues = []struct{
	Input uint32
	Expected uint32
} { {{range .Sequential }}
	{ 0x{{.Input}}, 0x{{.Output}} },
{{- end}}
}

var randomTestValues = []struct{
	Key string
	Input uint32
	Expected uint32
} { {{range .Random }}
	{ "{{.Key}}", 0x{{.Input}}, 0x{{.Output}} },
{{- end}}
}

func TestSequentialTable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large table test")
	}
	for _, test := range sequentialTestValues {
		actual := skip32.Encrypt(testKey, test.Input)
		if actual != test.Expected {
			t.Errorf("error encrypting\nexpected: %v\n   actual: %v", test.Expected, actual)
		}
		
		decrypted := skip32.Decrypt(testKey, actual)
		if decrypted != test.Input {
			t.Errorf("error decrypting\nexpected: %v\n   actual: %v", test.Input, decrypted)
		}		
	}
}

func TestRandomTable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large table test")
	}
	var key [10]byte = [10]byte{}
	for _, test := range randomTestValues {		
		decodeKey, err := hex.DecodeString(test.Key)
		if err != nil {
			panic(err)
		}
		copy(key[:], decodeKey)

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
`

type testValue struct {
	Key string
	Input string
	Output string
}

func main() {
	var (
		sequentialKey *string = flag.String("key", "00998877665544332211", "Sequential test key")
		sequentialCount *int = flag.Int("sequential", 10000, "How many sequential tests to generate")
		randomCount *int = flag.Int("random", 1000, "How many random tests to generate")
	)
	flag.Parse()

	tmpl, err := template.New("table_tests").Parse(outputTemplate)
	if err != nil {
		panic(err)
	}

	sequentialTestValues := []testValue{}
	randomTestValues := []testValue{}

	fmt.Fprint(os.Stderr, "Generating sequential: 0 ------------------------------ 100%\n")
	fmt.Fprint(os.Stderr, "                         ")
	for i := 0; i < *sequentialCount; i++ {
		if i % (*sequentialCount/30) == 0 {
			fmt.Fprint(os.Stderr, "*")
		}
		input := fmt.Sprintf("%08x", i)
		output := getCanonicalValue(*sequentialKey, input)
		sequentialTestValues = append(sequentialTestValues, testValue{Input: input, Output: output})
	}
	fmt.Fprintln(os.Stderr, "")


	fmt.Fprint(os.Stderr, "    Generating random: 0 ------------------------------ 100%\n")
	fmt.Fprint(os.Stderr, "                         ")
	for j := 0; j < *randomCount; j++ {
		if j % (*randomCount/10) == 0 {
			fmt.Fprint(os.Stderr, "***")
		}
		var i uint32 = 0
		for {
			i = rand.Uint32()
			if i + uint32(*sequentialCount) > math.MaxUint32 {
				continue
			} else {
				i = i + uint32(*sequentialCount)
				break
			}
		}

		keyBytes := make([]byte, 10)
		_, err := rand.Read(keyBytes)
		if err != nil {
			panic(err)
		}
		key := hex.EncodeToString(keyBytes)

		input := fmt.Sprintf("%08x", i)
		output := getCanonicalValue(key, input)
		randomTestValues = append(randomTestValues, testValue{
			Key: key,
			Input: input,
			Output: output,
		})
	}
	fmt.Fprintln(os.Stderr, "")

	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"Sequential": sequentialTestValues,
		"Random": randomTestValues,
	})

	if err != nil {
		panic(err)
	}
}

func getCanonicalValue(key, input string) string {
	cmd := exec.Command("./main", "e", key, input)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(stdoutStderr[:len(stdoutStderr)-1])
}
