&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Entries: []*ast.Entry{
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Package: &"main",
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 14, Line: 3, Column: 1},
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 14, Line: 3, Column: 1},
          Name: "main",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 24, Line: 3, Column: 11},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 25, Line: 3, Column: 12},
                Name: "int",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 32, Line: 3, Column: 19},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 33, Line: 3, Column: 20},
                Name: "int",
              },
            },
          },
          ReturnType: &ast.ReturnType{
            Pos: Position{Filename: "", Offset: 40, Line: 3, Column: 27},
            Name: "int",
          },
          Body: []*ast.Statement{
            &ast.Statement{
              Pos: Position{Filename: "", Offset: 49, Line: 4, Column: 3},
              ReturnStatement: &ast.ReturnStatement{
                Pos: Position{Filename: "", Offset: 49, Line: 4, Column: 3},
                ReturnLiteral: &ast.ReturnLiteral{
                  Pos: Position{Filename: "", Offset: 56, Line: 4, Column: 10},
                  Literal: a,
                },
              },
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
  },
}
