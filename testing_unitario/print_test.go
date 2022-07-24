package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintMessage(t *testing.T) {
	bootcamper := "Miguel"
	want := "Bootcamper " + bootcamper
	got := PrintMessage(bootcamper)

	if want != got {
		t.Errorf("Test fallido, se esperaba: %s\n, pero se obtuvo: %s\n", want, got)
	}
}

func TestPrintMessageTestify(t *testing.T) {
	bootcamper := "Miguel"
	want := "Bootcamper " + bootcamper
	got := PrintMessage(bootcamper)

	assert.Equal(t, want, got)
}
