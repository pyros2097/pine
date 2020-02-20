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
          Name: "sub",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 23, Line: 3, Column: 10},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 24, Line: 3, Column: 11},
                Name: "float",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 33, Line: 3, Column: 20},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 34, Line: 3, Column: 21},
                Name: "float",
              },
            },
          },
          ReturnType: &ast.ReturnType{
            Pos: Position{Filename: "", Offset: 43, Line: 3, Column: 30},
            Name: "float",
          },
          Body: []*ast.Statement{
            &ast.Statement{
              Pos: Position{Filename: "", Offset: 54, Line: 4, Column: 3},
              ReturnStatement: &ast.ReturnStatement{
                Pos: Position{Filename: "", Offset: 54, Line: 4, Column: 3},
                ReturnLiteral: &ast.ReturnLiteral{
                  Pos: Position{Filename: "", Offset: 61, Line: 4, Column: 10},
                  Literal: a,
                },
              },
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 70, Line: 8, Column: 1},
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 70, Line: 8, Column: 1},
          Name: "main",
          Parameters: []*ast.MethodParameter{
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 80, Line: 8, Column: 11},
              ID: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 81, Line: 8, Column: 12},
                Name: "int",
              },
            },
            &ast.MethodParameter{
              Pos: Position{Filename: "", Offset: 88, Line: 8, Column: 19},
              ID: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 89, Line: 8, Column: 20},
                Name: "int",
              },
            },
          },
          ReturnType: &ast.ReturnType{
            Pos: Position{Filename: "", Offset: 96, Line: 8, Column: 27},
            Name: "int",
          },
          Body: []*ast.Statement{
            &ast.Statement{
              Pos: Position{Filename: "", Offset: 105, Line: 9, Column: 3},
              ReturnStatement: &ast.ReturnStatement{
                Pos: Position{Filename: "", Offset: 105, Line: 9, Column: 3},
                ReturnLiteral: &ast.ReturnLiteral{
                  Pos: Position{Filename: "", Offset: 112, Line: 9, Column: 10},
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
