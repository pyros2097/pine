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
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 19, Line: 3, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 20, Line: 3, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 27, Line: 3, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 28, Line: 3, Column: 16},
                Name: "num",
              },
            },
          },
          ReturnTypes: []string{
            "num",
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 43, Line: 4, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 47, Line: 4, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 51, Line: 4, Column: 11},
                      Num: &4,
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
          End: &"\n",
        },
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 54, Line: 6, Column: 1},
          Name: "another",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 64, Line: 6, Column: 11},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 65, Line: 6, Column: 12},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 72, Line: 6, Column: 19},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 73, Line: 6, Column: 20},
                Name: "num",
              },
            },
          },
          ReturnTypes: []string{
            "num",
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 88, Line: 7, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 91, Line: 7, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 92, Line: 7, Column: 7},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 95, Line: 7, Column: 10},
                        Reference: &"b",
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
          End: &"\n",
        },
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 99, Line: 9, Column: 1},
          Name: "main",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 106, Line: 9, Column: 8},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 107, Line: 9, Column: 9},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 114, Line: 9, Column: 16},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 115, Line: 9, Column: 17},
                Name: "num",
              },
            },
          },
          ReturnTypes: []string{
            "num",
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 130, Line: 10, Column: 3},
                  Reference: &"another",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 137, Line: 10, Column: 10},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 138, Line: 10, Column: 11},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 141, Line: 10, Column: 14},
                        Reference: &"b",
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
          End: &"\n",
        },
      },
    },
  },
}
