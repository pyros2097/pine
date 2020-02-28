&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"main\n\n",
      Imports: []*ast.ImportStatement{
        &ast.ImportStatement{
          Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
          Name: "github.com/yum/std\n",
        },
      },
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 90, Line: 6, Column: 1},
          Name: "add",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 96, Line: 6, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 97, Line: 6, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 104, Line: 6, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 105, Line: 6, Column: 16},
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
                  Pos: Position{Filename: "", Offset: 120, Line: 7, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 124, Line: 7, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 128, Line: 7, Column: 11},
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
          Pos: Position{Filename: "", Offset: 131, Line: 9, Column: 1},
          Name: "another",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 141, Line: 9, Column: 11},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 142, Line: 9, Column: 12},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 149, Line: 9, Column: 19},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 150, Line: 9, Column: 20},
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
                  Pos: Position{Filename: "", Offset: 165, Line: 10, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 168, Line: 10, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 169, Line: 10, Column: 7},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 172, Line: 10, Column: 10},
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
          Pos: Position{Filename: "", Offset: 291, Line: 15, Column: 1},
          Name: "main",
          ReturnTypes: []string{
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 303, Line: 16, Column: 3},
                  Num: &1,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 307, Line: 16, Column: 7},
                    Num: &2,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 311, Line: 17, Column: 3},
                  Num: &3,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 315, Line: 17, Column: 7},
                    Num: &4,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 319, Line: 18, Column: 3},
                  Reference: &"another",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 326, Line: 18, Column: 10},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 327, Line: 18, Column: 11},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 330, Line: 18, Column: 14},
                        Num: &5,
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
