&compiler.Module{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Name: "main",
  End: "\n",
  ImportSection: compiler.ImportSection{
    Pos: Position{Filename: "", Offset: 12, Line: 2, Column: 1},
    Start: &"\n",
    End: &"\n",
  },
  AliasSection: compiler.AliasSection{
    Pos: Position{Filename: "", Offset: 14, Line: 4, Column: 1},
    Start: &"\n",
    Aliases: []*compiler.Alias{
      &compiler.Alias{
        Pos: Position{Filename: "", Offset: 15, Line: 5, Column: 1},
        Name: "Int",
        Value: &"i32",
        End: &"\n",
      },
      &compiler.Alias{
        Pos: Position{Filename: "", Offset: 28, Line: 6, Column: 1},
        Name: "Float",
        Value: &"f32",
        End: &"\n",
      },
    },
    End: &"\n",
  },
  EnumSection: compiler.EnumSection{
    Pos: Position{Filename: "", Offset: 44, Line: 8, Column: 1},
    Start: &"\n",
    Enums: []*compiler.Enum{
      &compiler.Enum{
        Pos: Position{Filename: "", Offset: 45, Line: 9, Column: 1},
        Name: "Direction",
        Value: []*compiler.EnumValue{
          &compiler.EnumValue{
            Pos: Position{Filename: "", Offset: 59, Line: 9, Column: 15},
            Start: "\n",
            Name: "Left",
          },
          &compiler.EnumValue{
            Pos: Position{Filename: "", Offset: 68, Line: 10, Column: 9},
            Start: "\n",
            Name: "Right",
          },
        },
        End: &"\n",
      },
    },
    End: &"\n",
  },
  TypeSection: compiler.TypeSection{
    Pos: Position{Filename: "", Offset: 80, Line: 13, Column: 1},
    Start: &"\n",
    Types: []*compiler.Type{
      &compiler.Type{
        Pos: Position{Filename: "", Offset: 81, Line: 14, Column: 1},
        Name: "Person",
        Fields: []compiler.TypeField{
          compiler.TypeField{
            Pos: Position{Filename: "", Offset: 92, Line: 14, Column: 12},
            Start: "\n",
            Name: "age",
            Value: "string",
          },
          compiler.TypeField{
            Pos: Position{Filename: "", Offset: 106, Line: 15, Column: 14},
            Start: "\n",
            Name: "name",
            Value: "string",
          },
        },
      },
    },
    End: &"\n",
  },
  FunctionSection: compiler.FunctionSection{
    Pos: Position{Filename: "", Offset: 192, Line: 21, Column: 1},
    Start: &"\n",
    Functions: []*compiler.Function{
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 193, Line: 22, Column: 1},
        Type: "proc",
        Name: "addFloatRef",
        Parameters: []compiler.FuncParameter{
          compiler.FuncParameter{
            Pos: Position{Filename: "", Offset: 210, Line: 22, Column: 18},
            Name: "a",
            Type: compiler.FuncParameterType{
              Pos: Position{Filename: "", Offset: 211, Line: 22, Column: 19},
              Name: "float",
            },
          },
          compiler.FuncParameter{
            Pos: Position{Filename: "", Offset: 220, Line: 22, Column: 28},
            Name: "b",
            Type: compiler.FuncParameterType{
              Pos: Position{Filename: "", Offset: 221, Line: 22, Column: 29},
              Name: "float",
            },
          },
        },
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 229, Line: 22, Column: 37},
          Name: "float",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 241, Line: 23, Column: 3},
                Reference: &"a",
              },
              Operator: &"+",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 245, Line: 23, Column: 7},
                  Reference: &"b",
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 248, Line: 25, Column: 1},
        Type: "proc",
        Name: "addIntRef",
        Parameters: []compiler.FuncParameter{
          compiler.FuncParameter{
            Pos: Position{Filename: "", Offset: 263, Line: 25, Column: 16},
            Name: "a",
            Type: compiler.FuncParameterType{
              Pos: Position{Filename: "", Offset: 264, Line: 25, Column: 17},
              Name: "int",
            },
          },
          compiler.FuncParameter{
            Pos: Position{Filename: "", Offset: 271, Line: 25, Column: 24},
            Name: "b",
            Type: compiler.FuncParameterType{
              Pos: Position{Filename: "", Offset: 272, Line: 25, Column: 25},
              Name: "int",
            },
          },
        },
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 278, Line: 25, Column: 31},
          Name: "int",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 288, Line: 26, Column: 3},
                Reference: &"a",
              },
              Operator: &"+",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 292, Line: 26, Column: 7},
                  Reference: &"b",
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 295, Line: 28, Column: 1},
        Type: "proc",
        Name: "addInt",
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 308, Line: 28, Column: 14},
          Name: "int",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 318, Line: 29, Column: 3},
                Int: &5,
              },
              Operator: &"+",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 322, Line: 29, Column: 7},
                  Int: &3,
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 325, Line: 31, Column: 1},
        Type: "proc",
        Name: "subInt",
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 338, Line: 31, Column: 14},
          Name: "int",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 348, Line: 32, Column: 3},
                Int: &5,
              },
              Operator: &"-",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 352, Line: 32, Column: 7},
                  Int: &3,
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 355, Line: 34, Column: 1},
        Type: "proc",
        Name: "mulInt",
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 368, Line: 34, Column: 14},
          Name: "int",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 378, Line: 35, Column: 3},
                Int: &5,
              },
              Operator: &"*",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 382, Line: 35, Column: 7},
                  Int: &3,
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
      &compiler.Function{
        Pos: Position{Filename: "", Offset: 385, Line: 37, Column: 1},
        Type: "proc",
        Name: "divInt",
        ReturnType: &compiler.FuncParameterType{
          Pos: Position{Filename: "", Offset: 398, Line: 37, Column: 14},
          Name: "int",
        },
        EndReturn: "\n",
        Body: []*compiler.Block{
          &compiler.Block{
            Exp: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 408, Line: 38, Column: 3},
                Int: &5,
              },
              Operator: &"/",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 412, Line: 38, Column: 7},
                  Int: &3,
                },
              },
            },
            End: &"\n",
          },
        },
        End: "\n",
      },
    },
    End: &"\n",
  },
}
