package compiler

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/repr"
	"github.com/pyros2097/cupaloy"
)

var testCases = []string{
	// "basic.pony",
	// "index.yum",
	"main.yum",
}

func TestAst(t *testing.T) {
	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory(".snapshots"))
	for _, fileName := range testCases {
		t.Run(fileName, func(t *testing.T) {
			result, err := ParseFile("../examples/" + fileName)
			require.NoError(t, err)
			resultString := repr.String(result, repr.Indent("  "))
			snapshotter.SnapshotTName(t, strings.Replace(fileName, ".yum", ".go", 1), resultString)
		})
	}
}
