package compiler

// import (
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/alecthomas/assert"
// 	"github.com/pyros2097/cupaloy"
// 	"github.com/stretchr/testify/require"
// )

// func TestCode(t *testing.T) {
// 	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
// 	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("examples"))
// 	for _, fileName := range testCases {
// 		t.Run(fileName, func(t *testing.T) {
// 			result := parseFile(t, fileName)
// 			generatedCode, err := GenerateCode(result)
// 			require.NoError(t, err)
// 			snapshotter.SnapshotTName(t, strings.Replace(fileName, ".pony", ".js", 1), generatedCode)
// 		})
// 	}
// }
