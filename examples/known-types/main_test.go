package main

import (
	"testing"

	"github.com/khulnasoft-lab/go-plugin/tests"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	got := tests.TestStdout(t, main)
	want := `I love Sushi
I love Tempura
Duration: 1h0m0s`
	assert.Equal(t, want, got)
}
