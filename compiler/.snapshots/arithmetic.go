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
      },
      End: "III",
    },
  },
  Functions: []*compiler.Function{
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 190, Line: 17, Column: 1},
      Type: "proc",
      Name: "addFloatRef",
      Parameters: []compiler.FuncParameter{
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 207, Line: 17, Column: 18},
          Name: "a",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 208, Line: 17, Column: 19},
            Name: "float",
          },
        },
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 217, Line: 17, Column: 28},
          Name: "b",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 218, Line: 17, Column: 29},
            Name: "float",
          },
        },
      },
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 226, Line: 17, Column: 37},
        Name: "float",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 238, Line: 18, Column: 3},
              Reference: &"a",
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 242, Line: 18, Column: 7},
                Reference: &"b",
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 245, Line: 20, Column: 1},
      Type: "proc",
      Name: "addIntRef",
      Parameters: []compiler.FuncParameter{
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 260, Line: 20, Column: 16},
          Name: "a",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 261, Line: 20, Column: 17},
            Name: "int",
          },
        },
        compiler.FuncParameter{
          Pos: Position{Filename: "", Offset: 268, Line: 20, Column: 24},
          Name: "b",
          Type: compiler.FuncParameterType{
            Pos: Position{Filename: "", Offset: 269, Line: 20, Column: 25},
            Name: "int",
          },
        },
      },
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 275, Line: 20, Column: 31},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 285, Line: 21, Column: 3},
              Reference: &"a",
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 289, Line: 21, Column: 7},
                Reference: &"b",
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 292, Line: 23, Column: 1},
      Type: "proc",
      Name: "addInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 305, Line: 23, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 315, Line: 24, Column: 3},
              Int: &5,
            },
            Operator: &"+",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 319, Line: 24, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 322, Line: 26, Column: 1},
      Type: "proc",
      Name: "subInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 335, Line: 26, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 345, Line: 27, Column: 3},
              Int: &5,
            },
            Operator: &"-",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 349, Line: 27, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 352, Line: 29, Column: 1},
      Type: "proc",
      Name: "mulInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 365, Line: 29, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 375, Line: 30, Column: 3},
              Int: &5,
            },
            Operator: &"*",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 379, Line: 30, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 382, Line: 32, Column: 1},
      Type: "proc",
      Name: "divInt",
      ReturnType: &compiler.FuncParameterType{
        Pos: Position{Filename: "", Offset: 395, Line: 32, Column: 14},
        Name: "int",
      },
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 405, Line: 33, Column: 3},
              Int: &5,
            },
            Operator: &"/",
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 409, Line: 33, Column: 7},
                Int: &3,
              },
            },
          },
        },
      },
      End: "III",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 412, Line: 35, Column: 1},
      Type: "test",
      Name: "all",
      Start: "III",
      Body: []*compiler.Block{
        &compiler.Block{
          Exp: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 427, Line: 36, Column: 3},
              Reference: &"addIntRef",
            },
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 436, Line: 36, Column: 12},
                Params: []*compiler.Literal{
                  &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 437, Line: 36, Column: 13},
                    Int: &1,
                  },
                  &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 440, Line: 36, Column: 16},
                    Int: &3,
                  },
                },
              },
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 445, Line: 37, Column: 3},
                  Reference: &"addInt",
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 451, Line: 37, Column: 9},
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 456, Line: 38, Column: 3},
                      Reference: &"subInt",
                    },
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 462, Line: 38, Column: 9},
                      },
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 467, Line: 39, Column: 3},
                          Reference: &"mulInt",
                        },
                        Right: &compiler.Expression{
                          Left: &compiler.Literal{
                            Pos: Position{Filename: "", Offset: 473, Line: 39, Column: 9},
                          },
                          Right: &compiler.Expression{
                            Left: &compiler.Literal{
                              Pos: Position{Filename: "", Offset: 478, Line: 40, Column: 3},
                              Reference: &"divInt",
                            },
                            Right: &compiler.Expression{
                              Left: &compiler.Literal{
                                Pos: Position{Filename: "", Offset: 484, Line: 40, Column: 9},
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
