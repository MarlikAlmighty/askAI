package main

import (
	"github.com/MarlikAlmighty/kickHisAss/internal/app"
	"testing"
)

func Test_Main(t *testing.T) {
	if err := app.Run(); err == nil {
		t.Errorf("Want err != nil, got %v == nil\n", err)
	}
}
