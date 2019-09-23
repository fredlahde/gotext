package gotext

import "testing"

func TestLoadModel(t *testing.T) {
	err := loadModel("/home/lothar/research/fasttext/wiki.en.bin")
	if err != nil {
		t.Fatal(err)
	}
}
