package code_gen

import (
	"os"
	"strings"
	"testing"
	"yum/ast"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

var testCases = []string{
	"main.yum",
}

func TestCode(t *testing.T) {
	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
	// snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory(".snapshots"))
	for _, fileName := range testCases {
		t.Run(fileName, func(t *testing.T) {
			ast, err := ast.ParseFile("../examples/" + fileName)
			require.NoError(t, err)

			file, err := os.Create("./.snapshots/" + strings.Replace(fileName, ".yum", ".wasm", 1))
			if err != nil {
				panic(err)
			}
			defer file.Close()
			GenerateCode(file, ast)

			// buffer := bytes.NewBuffer(nil)
			// GenerateCode(buffer, ast)
			// snapshotter.SnapshotTName(t, strings.Replace(fileName, ".yum", ".wasm", 1), buffer.Bytes())
		})
	}
}
