&compiler.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Modules: []*compiler.Module{
    &compiler.Module{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Name: "main",
      End: &"\n",
      Functions: []*compiler.Function{
        &compiler.Function{
          Pos: Position{Filename: "", Offset: 12, Line: 2, Column: 1},
          Start: "\n",
          Type: "proc",
          Name: "add",
          Parameters: []compiler.FuncParameter{
            compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 22, Line: 3, Column: 10},
              Name: "a",
              Type: compiler.FuncParameterType{
                Pos: Position{Filename: "", Offset: 23, Line: 3, Column: 11},
                Name: "int",
              },
            },
            compiler.FuncParameter{
              Pos: Position{Filename: "", Offset: 30, Line: 3, Column: 18},
              Name: "b",
              Type: compiler.FuncParameterType{
                Pos: Position{Filename: "", Offset: 31, Line: 3, Column: 19},
                Name: "int",
              },
            },
          },
          ReturnType: "int",
          EndReturn: "\n",
          Body: []*compiler.Block{
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 47, Line: 4, Column: 3},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 51, Line: 4, Column: 7},
                    Reference: &"b",
                  },
                  Operator: &"+",
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 55, Line: 4, Column: 11},
                      Int: &4,
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
          End: "\n",
        },
        &compiler.Function{
          Pos: Position{Filename: "", Offset: 58, Line: 6, Column: 1},
          Start: "\n",
          Type: "proc",
          Name: "main",
          EndReturn: "\n",
          Body: []*compiler.Block{
            &compiler.Block{
              Exp: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 76, Line: 8, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 79, Line: 8, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 80, Line: 8, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 83, Line: 8, Column: 10},
                        Int: &3,
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
                  Pos: Position{Filename: "", Offset: 88, Line: 9, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 91, Line: 9, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 92, Line: 9, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 95, Line: 9, Column: 10},
                        Int: &3,
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
                  Pos: Position{Filename: "", Offset: 100, Line: 10, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 103, Line: 10, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 104, Line: 10, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 107, Line: 10, Column: 10},
                        Int: &3,
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
                  Pos: Position{Filename: "", Offset: 112, Line: 11, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 115, Line: 11, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 116, Line: 11, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 119, Line: 11, Column: 10},
                        Int: &3,
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
                  Pos: Position{Filename: "", Offset: 124, Line: 12, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 127, Line: 12, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 128, Line: 12, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 131, Line: 12, Column: 10},
                        Int: &3,
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
                  Pos: Position{Filename: "", Offset: 136, Line: 13, Column: 3},
                  Reference: &"add",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 139, Line: 13, Column: 6},
                    Params: []*compiler.Literal{
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 140, Line: 13, Column: 7},
                        Int: &1,
                      },
                      &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 143, Line: 13, Column: 10},
                        Int: &3,
                      },
                    },
                  },
                },
              },
              End: &"\n",
            },
          },
          End: "\n",
        },
      },
    },
  },
}
