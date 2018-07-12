package gofunctiontest

import (
	"testing"
)

func TestMp3FromBase64(t *testing.T) {
	jsonFile := "./track.json"
	outdir := "./tmp"
	if err := Mp3FromBase64(jsonFile, outdir); err != nil {
		t.Errorf("error:\n%s", err.Error())
	}
}
