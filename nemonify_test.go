package main

import "testing"

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
