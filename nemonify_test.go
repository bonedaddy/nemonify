package nemonify

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	data := []byte("hello world")
	encoded := Encode(data)
	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if string(decoded) != string(data) {
		t.Fatal("Bad decoded data")
	}
}

func Test_Mnemonic(t *testing.T) {
	data := []byte("hello world")
	encoded := Encode(data)
	phrase, err := ToMnemonic(encoded)
	if err != nil {
		t.Fatal(err)
	}
	encodedFromPhrase, err := FromMnemonic(phrase)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := Decode(encodedFromPhrase)
	if err != nil {
		t.Fatal(err)
	}
	if string(decoded) != string(data) {
		t.Fatal("Bad decoded data")
	}
}

func Test_GenerateDecodeMnemonic(t *testing.T) {
	var (
		data              = []byte("hello world")
		fileName          = "input.file"
		savePathGenerated = "output.generated.file"
		savePathDecoded   = "output.decoded.file"
	)
	// setup the test data
	if err := ioutil.WriteFile(
		fileName,
		data,
		os.FileMode(0644),
	); err != nil {
		t.Fatal(err)
	}
	// defer removal of all test data
	defer os.RemoveAll(fileName)
	defer os.RemoveAll(savePathGenerated)
	defer os.RemoveAll(savePathDecoded)
	// the mnemonic phrase of our data should end up
	// being saved in the path of savePathGenerated
	if err := generateMnemonic(fileName, savePathGenerated); err != nil {
		t.Fatal(err)
	}
	// the decoded mnemonic phrase (our original data) should be read from
	// savePathGenerated and stored in savePathDecoded
	if err := decodeMnemonic(savePathGenerated, savePathDecoded); err != nil {
		t.Fatal(err)
	}
}
