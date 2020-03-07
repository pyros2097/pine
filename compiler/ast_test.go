package compiler

import (
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

func TestAst(t *testing.T) {
	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory(".snapshots"))
	for _, filename := range listFiles(t) {
		t.Run(filename, func(t *testing.T) {
			result, err := ParseFile("./tests/" + filename)
			require.NoError(t, err)
			resultString := repr.String(result, repr.Indent("  "))
			snapshotter.SnapshotTName(t, strings.Replace(filename, ".yum", ".go", 1), resultString)
		})
	}
}
