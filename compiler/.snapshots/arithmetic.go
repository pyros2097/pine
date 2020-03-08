&compiler.Module{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Name: "main",
  Enums: []*compiler.Enum{
    &compiler.Enum{
      Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
      Name: "Direction",
      Start: "III",
      Value: []*compiler.EnumValue{
        &compiler.EnumValue{
          Pos: Position{Filename: "", Offset: 30, Line: 4, Column: 3},
          Name: "Left",
        },
        &compiler.EnumValue{
          Pos: Position{Filename: "", Offset: 39, Line: 5, Column: 3},
          Name: "Right",
        },
      },
      End: "III",
    },
    &compiler.Enum{
      Pos: Position{Filename: "", Offset: 48, Line: 7, Column: 1},
      Name: "Position",
      Start: "III",
      Value: []*compiler.EnumValue{
        &compiler.EnumValue{
          Pos: Position{Filename: "", Offset: 64, Line: 8, Column: 3},
          Name: "Top",
        },
        &compiler.EnumValue{
          Pos: Position{Filename: "", Offset: 72, Line: 9, Column: 3},
          Name: "Bottom",
        },
      },
      End: "III",
    },
  },
  Types: []*compiler.Type{
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 82, Line: 11, Column: 1},
      Name: "User",
      Start: "III",
      Fields: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 94, Line: 12, Column: 3},
          Name: "name",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 109, Line: 13, Column: 3},
          Name: "age",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 123, Line: 14, Column: 3},
          Name: "were",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 138, Line: 15, Column: 3},
          Name: "params",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 155, Line: 16, Column: 3},
          Name: "address",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 173, Line: 17, Column: 3},
          Name: "createdAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 193, Line: 18, Column: 3},
          Name: "updatedAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 213, Line: 19, Column: 3},
          Name: "deletedAt",
          Value: "string",
        },
      },
      End: "III",
    },
  },
  Functions: []*compiler.Function{
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 296, Line: 23, Column: 1},
      Type: "proc",
      Name: "addFloatRef",
      Parameters: []compiler.FuncParameter{
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 313, Line: 23, Column: 18},
          Name: "a",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 314, Line: 23, Column: 19},
            Name: "float",
          },
        },
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 323, Line: 23, Column: 28},
          Name: "b",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 324, Line: 23, Column: 29},
            Name: "float",
          },
        },
      },
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 332, Line: 23, Column: 37},
        Name: "float",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 344, Line: 24, Column: 3},
              Reference: &"a",
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 348, Line: 24, Column: 7},
                Reference: &"b",
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 351, Line: 26, Column: 1},
      Type: "proc",
      Name: "addIntRef",
      Parameters: []compiler.FuncParameter{
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 366, Line: 26, Column: 16},
          Name: "a",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 367, Line: 26, Column: 17},
            Name: "int",
          },
        },
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 374, Line: 26, Column: 24},
          Name: "b",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 375, Line: 26, Column: 25},
            Name: "int",
          },
        },
      },
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 381, Line: 26, Column: 31},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 391, Line: 27, Column: 3},
              Reference: &"a",
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 395, Line: 27, Column: 7},
                Reference: &"b",
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 398, Line: 29, Column: 1},
      Type: "proc",
      Name: "addInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 411, Line: 29, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 421, Line: 30, Column: 3},
              Int: &5,
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 425, Line: 30, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 428, Line: 32, Column: 1},
      Type: "proc",
      Name: "subInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 441, Line: 32, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 451, Line: 33, Column: 3},
              Int: &5,
            },
            Operator: &"-",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 455, Line: 33, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 458, Line: 35, Column: 1},
      Type: "proc",
      Name: "mulInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 471, Line: 35, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 481, Line: 36, Column: 3},
              Int: &5,
            },
            Operator: &"*",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 485, Line: 36, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 488, Line: 38, Column: 1},
      Type: "proc",
      Name: "divInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 501, Line: 38, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 511, Line: 39, Column: 3},
              Int: &5,
            },
            Operator: &"/",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 515, Line: 39, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 518, Line: 41, Column: 1},
      Type: "test",
      Name: "all",
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 533, Line: 42, Column: 3},
              Reference: &"addIntRef",
            },
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 542, Line: 42, Column: 12},
                Params: []*compiler.Literal{
                  &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 543, Line: 42, Column: 13},
                    Int: &1,
                  },
                  &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 546, Line: 42, Column: 16},
                    Int: &3,
                  },
                },
              },
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 551, Line: 43, Column: 3},
                  Reference: &"addInt",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 557, Line: 43, Column: 9},
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 562, Line: 44, Column: 3},
                      Reference: &"subInt",
                    },
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 568, Line: 44, Column: 9},
                      },
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 573, Line: 45, Column: 3},
                          Reference: &"mulInt",
                        },
                        Right: &compiler.Expression{
                          Left: &compiler.Literal{
                            Pos: Position{Filename: "", Offset: 579, Line: 45, Column: 9},
                          },
                          Right: &compiler.Expression{
                            Left: &compiler.Literal{
                              Pos: Position{Filename: "", Offset: 584, Line: 46, Column: 3},
                              Reference: &"divInt",
                            },
                            Right: &compiler.Expression{
                              Left: &compiler.Literal{
                                Pos: Position{Filename: "", Offset: 590, Line: 46, Column: 9},
                              },
                            },
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
          },
        },
      },
      End: "III",
    },
  },
}
