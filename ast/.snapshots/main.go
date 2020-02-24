&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"main\n\n",
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
          Name: "add",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 19, Line: 3, Column: 7},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 20, Line: 3, Column: 8},
                Name: "int",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 27, Line: 3, Column: 15},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 28, Line: 3, Column: 16},
                Name: "int",
              },
            },
          },
          ReturnTypes: []string{
            "int",
          },
          Body: []*ast.Block{
            &ast.Block{
              End: &"\n",
            },
            &ast.Block{
              Statements: []*ast.Expression{
                &ast.Expression{
                  Left: a,
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: b,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              End: &"\n",
            },
          },
        },
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 50, Line: 6, Column: 1},
          Name: "main",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 57, Line: 6, Column: 8},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 58, Line: 6, Column: 9},
                Name: "int",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 65, Line: 6, Column: 16},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 66, Line: 6, Column: 17},
                Name: "int",
              },
            },
          },
          ReturnTypes: []string{
            "int",
          },
          Body: []*ast.Block{
            &ast.Block{
              End: &"\n",
            },
            &ast.Block{
              Statements: []*ast.Expression{
                &ast.Expression{
                  Left: a,
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: b,
                  },
                },
              },
              End: &"\n",
            },
          },
        },
      },
    },
  },
}
