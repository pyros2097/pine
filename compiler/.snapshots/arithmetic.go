&compiler.Module{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Name: "main",
  Types: []*compiler.Type{
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
      Name: "Direction",
      Struct: &compiler.Struct{
        Pos: Position{Filename: "", Offset: 37, Line: 4, Column: 3},
      },
    },
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 37, Line: 4, Column: 3},
      Enum: &compiler.Enum{
        Pos: Position{Filename: "", Offset: 37, Line: 4, Column: 3},
        Value: []string{
          "Left",
          "Right",
        },
      },
    },
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 53, Line: 7, Column: 1},
      Name: "Position",
      Struct: &compiler.Struct{
        Pos: Position{Filename: "", Offset: 76, Line: 8, Column: 3},
      },
    },
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 76, Line: 8, Column: 3},
      Enum: &compiler.Enum{
        Pos: Position{Filename: "", Offset: 76, Line: 8, Column: 3},
        Value: []string{
          "Top",
          "Bottom",
          "Left",
          "Right",
        },
      },
    },
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 107, Line: 11, Column: 1},
      Name: "User",
      Struct: &compiler.Struct{
        Pos: Position{Filename: "", Offset: 128, Line: 12, Column: 3},
        Fields: []*compiler.TypeField{
          &compiler.TypeField{
            Pos: Position{Filename: "", Offset: 128, Line: 12, Column: 3},
            ValueField: &compiler.ValueField{
              Pos: Position{Filename: "", Offset: 128, Line: 12, Column: 3},
              Name: "name",
              Type: "string",
            },
          },
          &compiler.TypeField{
            Pos: Position{Filename: "", Offset: 145, Line: 13, Column: 3},
            FuncField: &compiler.FuncField{
              Pos: Position{Filename: "", Offset: 145, Line: 13, Column: 3},
              Name: "setName",
              Parameters: []*compiler.TypeField{
                &compiler.TypeField{
                  Pos: Position{Filename: "", Offset: 154, Line: 13, Column: 12},
                  ValueField: &compiler.ValueField{
                    Pos: Position{Filename: "", Offset: 154, Line: 13, Column: 12},
                    Name: "v",
                    Type: "string",
                  },
                },
              },
            },
          },
        },
      },
    },
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 164, Line: 14, Column: 1},
      Enum: &compiler.Enum{
        Pos: Position{Filename: "", Offset: 164, Line: 14, Column: 1},
      },
    },
  },
  Functions: []*compiler.Function{
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 167, Line: 16, Column: 1},
      Name: "createUser",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 187, Line: 16, Column: 21},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 187, Line: 16, Column: 21},
            Name: "a",
            Type: "string",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 197, Line: 16, Column: 31},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 197, Line: 16, Column: 31},
            Name: "b",
            Type: "string",
          },
        },
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 214, Line: 17, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 214, Line: 17, Column: 3},
            Name: "a",
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 219, Line: 17, Column: 8},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 219, Line: 17, Column: 8},
                  Map: []*compiler.MapLiteral{
                    &compiler.MapLiteral{
                      Pos: Position{Filename: "", Offset: 225, Line: 18, Column: 5},
                      Key: &"name",
                      Value: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 231, Line: 18, Column: 11},
                        Str: &"",
                      },
                    },
                    &compiler.MapLiteral{
                      Pos: Position{Filename: "", Offset: 239, Line: 19, Column: 5},
                      Key: &"setName",
                      Value: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 250, Line: 19, Column: 16},
                        Str: &"",
                      },
                    },
                  },
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 259, Line: 21, Column: 3},
                    Reference: &"return",
                  },
                },
              },
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 269, Line: 24, Column: 1},
      Name: "User",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 283, Line: 24, Column: 15},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 283, Line: 24, Column: 15},
            Name: "name",
            Type: "string",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 296, Line: 24, Column: 28},
          FuncField: &compiler.FuncField{
            Pos: Position{Filename: "", Offset: 296, Line: 24, Column: 28},
            Name: "callback",
            Parameters: []*compiler.TypeField{
              &compiler.TypeField{
                Pos: Position{Filename: "", Offset: 306, Line: 24, Column: 38},
                FuncField: &compiler.FuncField{
                  Pos: Position{Filename: "", Offset: 306, Line: 24, Column: 38},
                  Name: "age",
                  Parameters: []*compiler.TypeField{
                    &compiler.TypeField{
                      Pos: Position{Filename: "", Offset: 311, Line: 24, Column: 43},
                      ValueField: &compiler.ValueField{
                        Pos: Position{Filename: "", Offset: 311, Line: 24, Column: 43},
                        Name: "g",
                        Type: "string",
                      },
                    },
                  },
                },
              },
            },
          },
        },
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 330, Line: 25, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 330, Line: 25, Column: 3},
            Name: "fullName",
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 342, Line: 25, Column: 15},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 342, Line: 25, Column: 15},
                  Str: &"Mr. ${name}",
                },
              },
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 479, Line: 35, Column: 1},
      Name: "fullName",
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 499, Line: 35, Column: 21},
        Name: "string",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 585, Line: 38, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 585, Line: 38, Column: 3},
            Name: "map",
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 592, Line: 38, Column: 10},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 592, Line: 38, Column: 10},
                  Map: []*compiler.MapLiteral{
                    &compiler.MapLiteral{
                      Pos: Position{Filename: "", Offset: 598, Line: 39, Column: 5},
                      Key: &"123",
                      Value: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 605, Line: 39, Column: 12},
                        Str: &"123",
                      },
                    },
                    &compiler.MapLiteral{
                      Pos: Position{Filename: "", Offset: 616, Line: 40, Column: 5},
                      Key: &"444",
                      Value: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 623, Line: 40, Column: 12},
                        Map: []*compiler.MapLiteral{
                          &compiler.MapLiteral{
                            Pos: Position{Filename: "", Offset: 631, Line: 41, Column: 7},
                            Key: &"444",
                            Value: &compiler.Literal{
                              Pos: Position{Filename: "", Offset: 638, Line: 41, Column: 14},
                              Str: &"444",
                            },
                          },
                        },
                      },
                    },
                    &compiler.MapLiteral{
                      Pos: Position{Filename: "", Offset: 655, Line: 43, Column: 5},
                      Key: &"test",
                      Value: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 663, Line: 43, Column: 13},
                        Reference: &"u.name",
                      },
                    },
                  },
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 677, Line: 45, Column: 3},
                    Reference: &"return",
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 684, Line: 45, Column: 10},
                      Reference: &"u.name",
                    },
                    Operator: &"+",
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 693, Line: 45, Column: 19},
                        Str: &" ",
                      },
                      Operator: &"+",
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 699, Line: 45, Column: 25},
                          Reference: &"u.age",
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
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 708, Line: 48, Column: 1},
      Name: "newState",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 726, Line: 48, Column: 19},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 726, Line: 48, Column: 19},
            Name: "value",
            Type: "string",
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 854, Line: 57, Column: 1},
      Name: "render",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 973, Line: 62, Column: 1},
      Name: "Counter",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 990, Line: 62, Column: 18},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 990, Line: 62, Column: 18},
            Name: "a",
            Type: "int",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 997, Line: 62, Column: 25},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 997, Line: 62, Column: 25},
            Name: "b",
            Type: "int",
          },
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 1004, Line: 62, Column: 32},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1016, Line: 63, Column: 3},
          ReturnStatement: &compiler.ReturnStatement{
            Pos: Position{Filename: "", Offset: 1016, Line: 63, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1023, Line: 63, Column: 10},
              Xml: &compiler.XmlLiteral{
                Pos: Position{Filename: "", Offset: 1023, Line: 63, Column: 10},
                Name: "View",
                EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                Children: []*compiler.XmlLiteral{
                  &compiler.XmlLiteral{
                    Pos: Position{Filename: "", Offset: 1036, Line: 64, Column: 7},
                    Name: "View",
                    EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                    Children: []*compiler.XmlLiteral{
                      &compiler.XmlLiteral{
                        Pos: Position{Filename: "", Offset: 1051, Line: 65, Column: 9},
                        Name: "Text",
                        EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                        Value: "Hello",
                        Close: "Text",
                      },
                    },
                    Close: "View",
                  },
                  &compiler.XmlLiteral{
                    Pos: Position{Filename: "", Offset: 1108, Line: 69, Column: 7},
                    Name: "View",
                    EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                    Children: []*compiler.XmlLiteral{
                      &compiler.XmlLiteral{
                        Pos: Position{Filename: "", Offset: 1123, Line: 70, Column: 9},
                        Name: "Text",
                        EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                        Value: "Hello",
                        Close: "Text",
                      },
                    },
                    Close: "View",
                  },
                },
                Close: "View",
              },
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 1189, Line: 77, Column: 1},
      Name: "adder",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 1204, Line: 77, Column: 16},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 1204, Line: 77, Column: 16},
            Name: "a",
            Type: "int",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 1211, Line: 77, Column: 23},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 1211, Line: 77, Column: 23},
            Name: "b",
            Type: "int",
          },
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 1218, Line: 77, Column: 30},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1273, Line: 81, Column: 3},
          ReturnStatement: &compiler.ReturnStatement{
            Pos: Position{Filename: "", Offset: 1273, Line: 81, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1280, Line: 81, Column: 10},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1280, Line: 81, Column: 10},
                  Reference: &"a",
                },
                Operator: &"+",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1284, Line: 81, Column: 14},
                    Reference: &"b",
                  },
                },
              },
            },
          },
        },
      },
    },
  },
  Tests: []*compiler.Test{
    &compiler.Test{
      Pos: Position{Filename: "", Offset: 1289, Line: 84, Column: 1},
      Name: "for equal or not equal",
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1323, Line: 85, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1323, Line: 85, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1331, Line: 85, Column: 11},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1331, Line: 85, Column: 11},
                  Int: &1,
                },
                Operator: &"==",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1336, Line: 85, Column: 16},
                    Int: &1,
                  },
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1340, Line: 86, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1340, Line: 86, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1348, Line: 86, Column: 11},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1348, Line: 86, Column: 11},
                  Int: &2,
                },
                Operator: &"==",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1353, Line: 86, Column: 16},
                    Int: &2,
                  },
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1357, Line: 87, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1357, Line: 87, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1365, Line: 87, Column: 11},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1365, Line: 87, Column: 11},
                  Int: &4,
                },
                Operator: &"==",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1370, Line: 87, Column: 16},
                    Int: &4,
                  },
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1374, Line: 88, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1374, Line: 88, Column: 3},
            AssignmentLiteral: &compiler.AssignmentLiteral{
              Pos: Position{Filename: "", Offset: 1382, Line: 88, Column: 11},
              Expression: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1382, Line: 88, Column: 11},
                  Int: &3,
                },
                Operator: &"!=",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1387, Line: 88, Column: 16},
                    Int: &13,
                  },
                },
              },
            },
          },
        },
      },
    },
  },
}
