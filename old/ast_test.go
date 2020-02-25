package compiler

// import (
// 	"io/ioutil"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/alecthomas/assert"
// 	"github.com/alecthomas/repr"
// 	"github.com/pyros2097/cupaloy"
// 	"github.com/stretchr/testify/require"
// )

// func parseFile(t *testing.T, file string) *Ast {
// 	data, err := ioutil.ReadFile("./examples/" + file)
// 	require.NoError(t, err)
// 	ast := &Ast{}
// 	err = parser.ParseBytes(data, ast)
// 	require.NoError(t, err)
// 	return ast
// }

// var testCases = []string{
// 	"basic.pony",
// }

// func TestAst(t *testing.T) {
// 	assert.NoError(t, os.Setenv("UPDATE_SNAPSHOTS", "true"))
// 	snapshotter := cupaloy.New(cupaloy.SnapshotSubdirectory("examples"))
// 	for _, fileName := range testCases {
// 		t.Run(fileName, func(t *testing.T) {
// 			result := parseFile(t, fileName)
// 			resultString := repr.String(result, repr.Indent("  "))
// 			snapshotter.SnapshotTName(t, strings.Replace(fileName, ".pony", ".go", 1), resultString)
// 		})
// 	}
// }
