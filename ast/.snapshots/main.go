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
      Classes: []*ast.Class{
        &ast.Class{
          Pos: Position{Filename: "", Offset: 51, Line: 4, Column: 1},
          Start: "\n",
          Name: "String",
          Next: &"\n",
          CompilerFuncs: []*ast.CompilerFunc{
            &ast.CompilerFunc{
              Pos: Position{Filename: "", Offset: 65, Line: 6, Column: 1},
              Start: "\n",
              Name: "ptr",
              ReturnTypes: []string{
                "i32",
                "\n",
              },
              End: &"\n",
            },
            &ast.CompilerFunc{
              Pos: Position{Filename: "", Offset: 102, Line: 9, Column: 1},
              Start: "\n",
              Name: "length",
              ReturnTypes: []string{
                "i32",
                "\n",
              },
              End: &"\n",
            },
          },
          End: &"\n",
        },
      },
      Funs: []*ast.FunDecl{
        &ast.FunDecl{
          Pos: Position{Filename: "", Offset: 143, Line: 13, Column: 1},
          Start: "\n",
          Name: "add",
          Parameters: []*ast.FuncParameter{
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 150, Line: 14, Column: 7},
              Name: "a",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 151, Line: 14, Column: 8},
                Name: "num",
              },
            },
            &ast.FuncParameter{
              Pos: Position{Filename: "", Offset: 158, Line: 14, Column: 15},
              Name: "b",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 159, Line: 14, Column: 16},
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
                  Pos: Position{Filename: "", Offset: 174, Line: 15, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 178, Line: 15, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &ast.Expression{
                    Left: &ast.Literal{
                      Pos: Position{Filename: "", Offset: 182, Line: 15, Column: 11},
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
          Pos: Position{Filename: "", Offset: 184, Line: 16, Column: 1},
          Start: "\n",
          Name: "_start",
          ReturnTypes: []string{
            "\n",
          },
          Body: []*ast.Block{
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 199, Line: 18, Column: 3},
                  Num: &1,
                },
                Operator: &"+",
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 203, Line: 18, Column: 7},
                    Num: &1,
                  },
                },
              },
              End: &"\n",
            },
            &ast.Block{
              Exp: &ast.Expression{
                Left: &ast.Literal{
                  Pos: Position{Filename: "", Offset: 207, Line: 19, Column: 3},
                  Reference: &"add",
                },
                Right: &ast.Expression{
                  Left: &ast.Literal{
                    Pos: Position{Filename: "", Offset: 210, Line: 19, Column: 6},
                    Params: []*ast.Literal{
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 211, Line: 19, Column: 7},
                        Num: &1,
                      },
                      &ast.Literal{
                        Pos: Position{Filename: "", Offset: 214, Line: 19, Column: 10},
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
