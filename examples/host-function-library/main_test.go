package main

import (
	"testing"

	"github.com/khulnasoft-lab/go-plugin/tests"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	got := tests.TestStdout(t, main)
	want := "Hello, Sato. This is Yamada-san (age 20)."
	assert.Equal(t, want, got)
}
