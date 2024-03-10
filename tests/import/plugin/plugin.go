//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/khulnasoft-lab/go-plugin/tests/import/proto/bar"
	"github.com/khulnasoft-lab/go-plugin/tests/import/proto/foo"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	foo.RegisterFoo(TestPlugin{})
}

type TestPlugin struct{}

var _ foo.Foo = (*TestPlugin)(nil)

func (p TestPlugin) Hello(_ context.Context, request *foo.Request) (*bar.Reply, error) {
	return &bar.Reply{
		A: request.GetA() + ", bar",
	}, nil
}
