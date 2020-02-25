package compiler

import (
	"strings"
)

func GenerateCode(ast *Ast) (string, error) {
	sb := strings.Builder{}
	for _, entry := range ast.Entries {
		// if entry.Package != nil {
		// 	sb.WriteString("// package ")
		// 	sb.WriteString(*entry.Package)
		// 	sb.WriteString("\n")
		// }
		for _, importStmt := range entry.Imports {
			if importStmt != nil {
				sb.WriteString("const ")
				sb.WriteString(importStmt.Name)
				sb.WriteString(" = ")
				sb.WriteString("require('" + importStmt.Name + "');")
				sb.WriteString("\n")
			}
		}
		// for _, st := range entry.Structs {
		// 	if st != nil {
		// 		sb.WriteString("function  "st.Name+"")
		// 		sb.WriteString(importStmt.Name)
		// 	}
		// }
	}
	return sb.String(), nil
}
