package main

import (
	"testing"
)

func TestConnectFail(t *testing.T) {
	_, err := getClient(nil)
	if err == nil {
		t.Error("getClient(): got no error, want error")
	}
}
