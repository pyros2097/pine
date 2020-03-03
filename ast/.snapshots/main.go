&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*ast.Module{
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: "wasi_unstable",
      End: &"\n",
      ExternFuncs: []*ast.ExternFunc{
        &ast.ExternFunc{
          Pos: Position{Filename: "", Offset: 21, Line: 2, Column: 1},
          Start: "\n",
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
          ReturnType: "i32",
          End: &"\n",
        },
      },
    },
    &ast.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: "main",
      End: &"\n",
      Imports: []*ast.ImportStatement{
        &ast.ImportStatement{
          Pos: Position{Filename: "", Offset: 12, Line: 2, Column: 1},
          Start: &"\n",
          Url: "github.com/yum/wasi_unstable",
          End: &"\n",
        },
      },
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 51, Line: 4, Column: 1},
          Start: "\n",
          Name: "add",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 58, Line: 5, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 59, Line: 5, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 66, Line: 5, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 67, Line: 5, Column: 16},
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
                  Pos: Position{Filename: "", Offset: 82, Line: 6, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 86, Line: 6, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 90, Line: 6, Column: 11},
                      Num: &4,
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
        },
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 92, Line: 7, Column: 1},
          Start: "\n",
          Name: "_start",
          ReturnTypes: []string{
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 107, Line: 9, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 110, Line: 9, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 111, Line: 9, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 114, Line: 9, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 119, Line: 10, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 122, Line: 10, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 123, Line: 10, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 126, Line: 10, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 131, Line: 11, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 134, Line: 11, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 135, Line: 11, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 138, Line: 11, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 143, Line: 12, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 146, Line: 12, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 147, Line: 12, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 150, Line: 12, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 155, Line: 13, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 158, Line: 13, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 159, Line: 13, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 162, Line: 13, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 167, Line: 14, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 170, Line: 14, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 171, Line: 14, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 174, Line: 14, Column: 10},
                        Num: &2,
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
                  Pos: Position{Filename: "", Offset: 179, Line: 15, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 182, Line: 15, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 183, Line: 15, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 186, Line: 15, Column: 10},
                        Num: &2,
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
  },
}
