&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"main\n\n",
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 62, Line: 5, Column: 1},
          Name: "add",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 68, Line: 5, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 69, Line: 5, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 76, Line: 5, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 77, Line: 5, Column: 16},
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
                  Pos: Position{Filename: "", Offset: 92, Line: 6, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 96, Line: 6, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 100, Line: 6, Column: 11},
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
          Pos: Position{Filename: "", Offset: 103, Line: 8, Column: 1},
          Name: "another",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 113, Line: 8, Column: 11},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 114, Line: 8, Column: 12},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 121, Line: 8, Column: 19},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 122, Line: 8, Column: 20},
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
                  Pos: Position{Filename: "", Offset: 137, Line: 9, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 140, Line: 9, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 141, Line: 9, Column: 7},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 144, Line: 9, Column: 10},
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
          Pos: Position{Filename: "", Offset: 263, Line: 14, Column: 1},
          Name: "main",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 270, Line: 14, Column: 8},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 271, Line: 14, Column: 9},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 278, Line: 14, Column: 16},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 279, Line: 14, Column: 17},
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
                  Pos: Position{Filename: "", Offset: 294, Line: 15, Column: 3},
                  Num: &1,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 298, Line: 15, Column: 7},
                    Num: &2,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 302, Line: 16, Column: 3},
                  Num: &3,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 306, Line: 16, Column: 7},
                    Num: &4,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 331, Line: 18, Column: 3},
                  Reference: &"another",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 338, Line: 18, Column: 10},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 339, Line: 18, Column: 11},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 342, Line: 18, Column: 14},
                        Reference: &"b",
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 347, Line: 19, Column: 3},
                  Reference: &"another",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 354, Line: 19, Column: 10},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 355, Line: 19, Column: 11},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 358, Line: 19, Column: 14},
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
