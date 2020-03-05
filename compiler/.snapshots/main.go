&compiler.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*compiler.Module{
    &compiler.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: "wasi_unstable",
      End: &"\n",
      ExternFuncs: []*compiler.ExternFunc{
        &compiler.ExternFunc{
          Pos: Position{Filename: "", Offset: 21, Line: 2, Column: 1},
          Start: "\n",
          Name: "fd_write",
          Parameters: []*compiler.FuncParameter{
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 43, Line: 3, Column: 22},
              Name: "a",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 44, Line: 3, Column: 23},
                Name: "i32",
              },
            },
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 51, Line: 3, Column: 30},
              Name: "b",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 52, Line: 3, Column: 31},
                Name: "i32",
              },
            },
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 59, Line: 3, Column: 38},
              Name: "c",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 60, Line: 3, Column: 39},
                Name: "i32",
              },
            },
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 67, Line: 3, Column: 46},
              Name: "d",
              Type: &compiler.Type{
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
    &compiler.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: "main",
      End: &"\n",
      Imports: []*compiler.ImportStatement{
        &compiler.ImportStatement{
          Pos: Position{Filename: "", Offset: 12, Line: 2, Column: 1},
          Start: &"\n",
          Url: "github.com/yum/wasi_unstable",
          End: &"\n",
        },
      },
      Funs: []*compiler.FunDecl{
        &compiler.FunDecl{
          Pos: Position{Filename: "", Offset: 51, Line: 4, Column: 1},
          Start: "\n",
          Name: "add",
          Parameters: []*compiler.FuncParameter{
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 58, Line: 5, Column: 7},
              Name: "a",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 59, Line: 5, Column: 8},
                Name: "num",
              },
            },
            &compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 66, Line: 5, Column: 15},
              Name: "b",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 67, Line: 5, Column: 16},
                Name: "num",
              },
            },
          },
          ReturnTypes: []string{
            "num",
            "\n",
          },
          Body: []*compiler.Block{
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 82, Line: 6, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 86, Line: 6, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
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
        &compiler.FunDecl{
          Pos: Position{Filename: "", Offset: 92, Line: 7, Column: 1},
          Start: "\n",
          Name: "_start",
          ReturnTypes: []string{
            "\n",
          },
          Body: []*compiler.Block{
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 107, Line: 9, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 110, Line: 9, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 111, Line: 9, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 114, Line: 9, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 119, Line: 10, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 122, Line: 10, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 123, Line: 10, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 126, Line: 10, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 131, Line: 11, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 134, Line: 11, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 135, Line: 11, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 138, Line: 11, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 143, Line: 12, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 146, Line: 12, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 147, Line: 12, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 150, Line: 12, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 155, Line: 13, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 158, Line: 13, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 159, Line: 13, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 162, Line: 13, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 167, Line: 14, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 170, Line: 14, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 171, Line: 14, Column: 7},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 174, Line: 14, Column: 10},
                        Num: &2,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 179, Line: 15, Column: 3},
                  Reference: &"fd_write",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 187, Line: 15, Column: 11},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 188, Line: 15, Column: 12},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 191, Line: 15, Column: 15},
                        Str: &"hello world",
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 206, Line: 15, Column: 30},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 209, Line: 15, Column: 33},
                        Num: &20,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 215, Line: 16, Column: 3},
                  Reference: &"fd_write",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 223, Line: 16, Column: 11},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 224, Line: 16, Column: 12},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 227, Line: 16, Column: 15},
                        Str: &"hello world",
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 242, Line: 16, Column: 30},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 245, Line: 16, Column: 33},
                        Num: &20,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 251, Line: 17, Column: 3},
                  Reference: &"fd_write",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 259, Line: 17, Column: 11},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 260, Line: 17, Column: 12},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 263, Line: 17, Column: 15},
                        Str: &"hello world",
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 278, Line: 17, Column: 30},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 281, Line: 17, Column: 33},
                        Num: &20,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 287, Line: 18, Column: 3},
                  Reference: &"fd_write",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 295, Line: 18, Column: 11},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 296, Line: 18, Column: 12},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 299, Line: 18, Column: 15},
                        Str: &"sqdasd",
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 309, Line: 18, Column: 25},
                        Num: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 312, Line: 18, Column: 28},
                        Num: &20,
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
