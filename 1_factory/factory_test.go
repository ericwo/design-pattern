package factory

import (
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Eric")
	if s != "Hi, Eric" {
		t.Fatal("Type1 test failed")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Eric")
	if s != "Hello, Eric" {
		t.Fatal("Type2 test failed")
	}
}
