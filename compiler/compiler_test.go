package compiler

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/repr"
	"github.com/pyros2097/cupaloy"
)

func listFiles(t *testing.T) []string {
	var files []string
	root := "./tests"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	require.NoError(t, err)
	return files
}

func TestSnapshotAll(t *testing.T) {
	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory(".snapshots"))
	for _, filename := range listFiles(t) {
		t.Run(filename, func(t *testing.T) {
			ast, err := ParseFile("./tests/" + filename)
			require.NoError(t, err)
			resultString := repr.String(ast, repr.Indent("  "))
			snapshotter.SnapshotTName(t, strings.Replace(filename, ".pi", ".go", 1), resultString)

			jsFile, err := os.Create("./.snapshots/" + strings.Replace(filename, ".pi", ".js", 1))
			if err != nil {
				require.NoError(t, err)
			}
			defer jsFile.Close()
			data, err := NewEmitter(ast).EmitJS()
			if err != nil {
				require.NoError(t, err)
			}
			io.Copy(jsFile, bytes.NewBuffer(data.Bytes()))

			// file, err := os.Create("./.snapshots/" + strings.Replace(filename, ".pi", ".wasm", 1))
			// if err != nil {
			// 	require.NoError(t, err)
			// }
			// defer file.Close()
			// TODO: later
			// data, err := NewEmitter(ast).EmitWASM()
			// if err != nil {
			// 	require.NoError(t, err)
			// }
			// io.Copy(file, bytes.NewBuffer(data.Bytes()))
			// instance, err := wasm.NewInstance(data.Bytes())
			// if err != nil {
			// 	require.NoError(t, err)
			// }
			// defer instance.Close()
			// for _, fun := range ast.Functions {
			// 	if strings.HasPrefix(fun.Name, "test_") {
			// 		require.NotNil(t, instance.Exports[fun.Name])
			// 		result, err := instance.Exports[fun.Name]()
			// 		require.NoError(t, err)
			// 		require.NotNil(t, result)
			// 		assert.Equal(t, "void", result.String())
			// 	}
			// }
		})
	}
}
