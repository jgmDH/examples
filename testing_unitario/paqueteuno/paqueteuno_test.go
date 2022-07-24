package paqueteuno_test

import (
	"testing"

	"github.com/bootcamp-go/testing/paqueteuno"
	"github.com/stretchr/testify/assert"
)

func TestPaqueteUno(t *testing.T) {
	want := "Soy el paquete uno"
	got := paqueteuno.PaqueteUno()

	if want != got {
		t.Errorf("Fallo, se esperaba %s, pero se obtuvo %s", want, got)
	}
}

func TestPaqueteUnoTestify(t *testing.T) {
	want := "Soy el paquete uno"
	got := paqueteuno.PaqueteUno()

	assert.Equal(t, want, got)
}
