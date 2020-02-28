&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"wasi_unstable",
      Next: &"\n\n",
      ExternFuncs: []*ast.ExternFunc{
        &ast.ExternFunc{
          Pos: Position{Filename: "", Offset: 22, Line: 3, Column: 1},
          Name: "fd_write",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 43, Line: 3, Column: 22},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 44, Line: 3, Column: 23},
                Name: "i32",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 51, Line: 3, Column: 30},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 52, Line: 3, Column: 31},
                Name: "i32",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 59, Line: 3, Column: 38},
              Name: "c",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 60, Line: 3, Column: 39},
                Name: "i32",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 67, Line: 3, Column: 46},
              Name: "d",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 68, Line: 3, Column: 47},
                Name: "i32",
              },
            },
          },
          ReturnTypes: []string{
            "i32",
            "\n",
          },
          End: &"\n",
        },
      },
    },
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: &"main",
      Next: &"\n\n",
      Imports: []*ast.ImportStatement{
        &ast.ImportStatement{
          Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
          Url: "github.com/yum/wasi_unstable",
          End: &"\n",
        },
      },
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 100, Line: 6, Column: 1},
          Name: "add",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 106, Line: 6, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 107, Line: 6, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 114, Line: 6, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 115, Line: 6, Column: 16},
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
                  Pos: Position{Filename: "", Offset: 130, Line: 7, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 134, Line: 7, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 138, Line: 7, Column: 11},
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
          Pos: Position{Filename: "", Offset: 141, Line: 9, Column: 1},
          Name: "another",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 151, Line: 9, Column: 11},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 152, Line: 9, Column: 12},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 159, Line: 9, Column: 19},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 160, Line: 9, Column: 20},
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
                  Pos: Position{Filename: "", Offset: 175, Line: 10, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 178, Line: 10, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 179, Line: 10, Column: 7},
                        Reference: &"a",
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 182, Line: 10, Column: 10},
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
          Pos: Position{Filename: "", Offset: 301, Line: 15, Column: 1},
          Name: "main",
          ReturnTypes: []string{
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 313, Line: 16, Column: 3},
                  Num: &1,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 317, Line: 16, Column: 7},
                    Num: &2,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 321, Line: 17, Column: 3},
                  Num: &3,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 325, Line: 17, Column: 7},
                    Num: &4,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 329, Line: 18, Column: 3},
                  Reference: &"another",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 336, Line: 18, Column: 10},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 337, Line: 18, Column: 11},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 340, Line: 18, Column: 14},
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
