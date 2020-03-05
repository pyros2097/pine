package compiler

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func TestCode(t *testing.T) {
	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
	// snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory(".snapshots"))
	for _, fileName := range testCases {
		t.Run(fileName, func(t *testing.T) {
			ast, err := ParseFile("../examples/" + fileName)
			require.NoError(t, err)

			wasmFileName := strings.Replace(fileName, ".yum", ".wasm", 1)
			file, err := os.Create("./.snapshots/" + wasmFileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			data, err := NewEmitter(ast).EmitAll()
			if err != nil {
				panic(err)
			}
			io.Copy(file, bytes.NewBuffer(data.Bytes()))
			instance, err := wasm.NewInstance(data.Bytes())
			if err != nil {
				panic(err)
			}
			defer instance.Close()

			mainFunc := instance.Exports["main"]
			result, err := mainFunc()
			if err != nil {
				panic(err)
			}
			assert.Equal(t, "void", result.String())
		})
	}
}

func TestEmitStore(t *testing.T) {
}

func TestEmitString(t *testing.T) {
}
