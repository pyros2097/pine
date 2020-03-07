package compiler

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wasmerio/go-ext-wasm/wasmer"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"

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
			snapshotter.SnapshotTName(t, strings.Replace(filename, ".yum", ".go", 1), resultString)

			file, err := os.Create("./.snapshots/" + strings.Replace(filename, ".yum", ".wasm", 1))
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
			if filename == "arithmetic.yum" {
				for _, v := range []struct {
					name   string
					left   interface{}
					right  interface{}
					result string
				}{
					{"addInt", nil, nil, "8"},
					{"subInt", nil, nil, "2"},
					{"mulInt", nil, nil, "15"},
					{"divInt", nil, nil, "1"},
					{"addIntRef", 1, 3, "4"},
				} {
					require.NotNil(t, instance.Exports[v.name])
					var result wasmer.Value
					var err error
					if v.left != nil && v.right != nil {
						result, err = instance.Exports[v.name](v.left, v.right)
					} else {
						result, err = instance.Exports[v.name]()
					}
					if err != nil {
						require.NoError(t, err)
					}
					require.NotNil(t, result)
					assert.Equal(t, v.result, result.String())
				}
			}
		})
	}
}
