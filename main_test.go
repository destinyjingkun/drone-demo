package main

import(
	"testing"
	"os"
)

func TestMain(m *testing.M){
	os.Exit(m.Run())
}

func TestHandler(t *testing.T){
	t.Error("error here")
}