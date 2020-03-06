package compiler

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
	"github.com/wasmerio/go-ext-wasm/wasmer"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
)

func loadFile(t *testing.T, fileName string) wasmer.Instance {
	ast, err := ParseFile("./tests/" + fileName)
	if err != nil {
		require.NoError(t, err)
	}
	data, err := NewEmitter(ast).EmitAll()
	if err != nil {
		require.NoError(t, err)
	}
	instance, err := wasm.NewInstance(data.Bytes())
	if err != nil {
		require.NoError(t, err)
	}
	return instance
}

func TestArithmetic(t *testing.T) {
	instance := loadFile(t, "arithmetic.yum")
	defer instance.Close()

	result, err := instance.Exports["addIntRef"](1, 2)
	if err != nil {
		require.NoError(t, err)
	}
	assert.Equal(t, "3", result.String())
	results := []string{"8", "2", "15", "1"}
	for i, fn := range []string{"addInt", "subInt", "mulInt", "divInt"} {
		result, err = instance.Exports[fn]()
		if err != nil {
			require.NoError(t, err)
		}
		assert.Equal(t, results[i], result.String())
	}
}

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
				require.NoError(t, err)
			}
			defer file.Close()
			data, err := NewEmitter(ast).EmitAll()
			if err != nil {
				require.NoError(t, err)
			}
			io.Copy(file, bytes.NewBuffer(data.Bytes()))
			instance, err := wasm.NewInstance(data.Bytes())
			if err != nil {
				require.NoError(t, err)
			}
			defer instance.Close()

			mainFunc := instance.Exports["main"]
			result, err := mainFunc()
			if err != nil {
				require.NoError(t, err)
			}
			assert.Equal(t, "void", result.String())
		})
	}
}

func TestEmitStore(t *testing.T) {
}

func TestEmitString(t *testing.T) {
}
