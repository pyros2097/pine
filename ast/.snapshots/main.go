&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"main\n\n",
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
          Name: "main",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 20, Line: 3, Column: 8},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 21, Line: 3, Column: 9},
                Name: "int",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 28, Line: 3, Column: 16},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 29, Line: 3, Column: 17},
                Name: "int",
              },
            },
          },
          ReturnTypes: []string{
            "int",
          },
          Body: &ast.Block{
            Statements: []*ast.Expression{
              &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 44, Line: 4, Column: 3},
                  Int: &2,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 48, Line: 4, Column: 7},
                    Int: &6,
                  },
                  Operator: &"-",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 52, Line: 4, Column: 11},
                      Int: &2,
                    },
                    Operator: &"*",
                    Right: &ast.Expression{
                      Left: &ast.Literal{
                        Pos: Position{Filename: "", Offset: 56, Line: 4, Column: 15},
                        Int: &10,
                      },
                      Operator: &"/",
                      Right: &ast.Expression{
                        Left: &ast.Literal{
                          Pos: Position{Filename: "", Offset: 61, Line: 4, Column: 20},
                          Int: &7,
                        },
                      },
                    },
                  },
                },
              },
            },
            End: &"\n",
          },
        },
      },
    },
  },
}
