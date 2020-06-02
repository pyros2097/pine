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
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 220, Line: 17, Column: 9},
                Map: []*compiler.MapLiteral{
                  &compiler.MapLiteral{
                    Pos: Position{Filename: "", Offset: 226, Line: 18, Column: 5},
                    Key: &"name",
                    Value: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 232, Line: 18, Column: 11},
                      Str: &"",
                    },
                  },
                  &compiler.MapLiteral{
                    Pos: Position{Filename: "", Offset: 240, Line: 19, Column: 5},
                    Key: &"setName",
                    Value: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 251, Line: 19, Column: 16},
                      Str: &"",
                    },
                  },
                },
              },
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 260, Line: 21, Column: 3},
                  Reference: &"return",
                },
              },
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 270, Line: 24, Column: 1},
      Name: "User",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 284, Line: 24, Column: 15},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 284, Line: 24, Column: 15},
            Name: "name",
            Type: "string",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 297, Line: 24, Column: 28},
          FuncField: &compiler.FuncField{
            Pos: Position{Filename: "", Offset: 297, Line: 24, Column: 28},
            Name: "callback",
            Parameters: []*compiler.TypeField{
              &compiler.TypeField{
                Pos: Position{Filename: "", Offset: 307, Line: 24, Column: 38},
                FuncField: &compiler.FuncField{
                  Pos: Position{Filename: "", Offset: 307, Line: 24, Column: 38},
                  Name: "age",
                  Parameters: []*compiler.TypeField{
                    &compiler.TypeField{
                      Pos: Position{Filename: "", Offset: 312, Line: 24, Column: 43},
                      ValueField: &compiler.ValueField{
                        Pos: Position{Filename: "", Offset: 312, Line: 24, Column: 43},
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
          Pos: Position{Filename: "", Offset: 331, Line: 25, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 331, Line: 25, Column: 3},
            Name: "fullName",
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 343, Line: 25, Column: 15},
                Str: &"Mr. ${name}",
              },
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 480, Line: 35, Column: 1},
      Name: "fullName",
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 500, Line: 35, Column: 21},
        Name: "string",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 565, Line: 37, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 565, Line: 37, Column: 3},
            Name: "name",
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 573, Line: 37, Column: 11},
                Reference: &"u.name",
              },
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 583, Line: 38, Column: 3},
                  Reference: &"map",
                },
                Operator: &":=",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 590, Line: 38, Column: 10},
                    Map: []*compiler.MapLiteral{
                      &compiler.MapLiteral{
                        Pos: Position{Filename: "", Offset: 596, Line: 39, Column: 5},
                        Key: &"123",
                        Value: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 603, Line: 39, Column: 12},
                          Str: &"123",
                        },
                      },
                      &compiler.MapLiteral{
                        Pos: Position{Filename: "", Offset: 614, Line: 40, Column: 5},
                        Key: &"444",
                        Value: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 621, Line: 40, Column: 12},
                          Map: []*compiler.MapLiteral{
                            &compiler.MapLiteral{
                              Pos: Position{Filename: "", Offset: 629, Line: 41, Column: 7},
                              Key: &"444",
                              Value: &compiler.Literal{
                                Pos: Position{Filename: "", Offset: 636, Line: 41, Column: 14},
                                Str: &"444",
                              },
                            },
                          },
                        },
                      },
                      &compiler.MapLiteral{
                        Pos: Position{Filename: "", Offset: 653, Line: 43, Column: 5},
                        Key: &"test",
                        Value: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 661, Line: 43, Column: 13},
                          Reference: &"u.name",
                        },
                      },
                    },
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 675, Line: 45, Column: 3},
                      Reference: &"return",
                    },
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 682, Line: 45, Column: 10},
                        Reference: &"u.name",
                      },
                      Operator: &"+",
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 691, Line: 45, Column: 19},
                          Str: &" ",
                        },
                        Operator: &"+",
                        Right: &compiler.Expression{
                          Left: &compiler.Literal{
                            Pos: Position{Filename: "", Offset: 697, Line: 45, Column: 25},
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
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 706, Line: 48, Column: 1},
      Name: "newState",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 724, Line: 48, Column: 19},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 724, Line: 48, Column: 19},
            Name: "value",
            Type: "string",
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 852, Line: 57, Column: 1},
      Name: "render",
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 925, Line: 61, Column: 1},
      Name: "Counter",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 942, Line: 61, Column: 18},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 942, Line: 61, Column: 18},
            Name: "a",
            Type: "int",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 949, Line: 61, Column: 25},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 949, Line: 61, Column: 25},
            Name: "b",
            Type: "int",
          },
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 956, Line: 61, Column: 32},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 968, Line: 62, Column: 3},
          FuncCall: &compiler.FuncCallStatement{
            Pos: Position{Filename: "", Offset: 968, Line: 62, Column: 3},
            Name: "return",
            Values: []*compiler.Literal{
              &compiler.Literal{
                Pos: Position{Filename: "", Offset: 981, Line: 63, Column: 5},
                Xml: &compiler.XmlLiteral{
                  Pos: Position{Filename: "", Offset: 981, Line: 63, Column: 5},
                  Name: "View",
                  EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                  Children: []*compiler.XmlLiteral{
                    &compiler.XmlLiteral{
                      Pos: Position{Filename: "", Offset: 994, Line: 64, Column: 7},
                      Name: "View",
                      EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                      Children: []*compiler.XmlLiteral{
                        &compiler.XmlLiteral{
                          Pos: Position{Filename: "", Offset: 1009, Line: 65, Column: 9},
                          Name: "Text",
                          EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                          Value: "Hello",
                          Close: "Text",
                        },
                      },
                      Close: "View",
                    },
                    &compiler.XmlLiteral{
                      Pos: Position{Filename: "", Offset: 1066, Line: 69, Column: 7},
                      Name: "View",
                      EndName: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                      Children: []*compiler.XmlLiteral{
                        &compiler.XmlLiteral{
                          Pos: Position{Filename: "", Offset: 1081, Line: 70, Column: 9},
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
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 1151, Line: 78, Column: 1},
      Name: "adder",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 1166, Line: 78, Column: 16},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 1166, Line: 78, Column: 16},
            Name: "a",
            Type: "int",
          },
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 1173, Line: 78, Column: 23},
          ValueField: &compiler.ValueField{
            Pos: Position{Filename: "", Offset: 1173, Line: 78, Column: 23},
            Name: "b",
            Type: "int",
          },
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 1180, Line: 78, Column: 30},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1191, Line: 79, Column: 3},
          If: &compiler.IfStatement{
            Pos: Position{Filename: "", Offset: 1191, Line: 79, Column: 3},
            Condition: []*compiler.Expression{
              &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1194, Line: 79, Column: 6},
                  Reference: &"a",
                },
                Operator: &"<",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 1198, Line: 79, Column: 10},
                    Reference: &"b",
                  },
                },
              },
            },
            Result: []*compiler.Statement{
              &compiler.Statement{
                Pos: Position{Filename: "", Offset: 1206, Line: 80, Column: 5},
                ReturnStatement: &compiler.ReturnStatement{
                  Pos: Position{Filename: "", Offset: 1206, Line: 80, Column: 5},
                  Expression: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 1213, Line: 80, Column: 12},
                      Reference: &"a",
                    },
                    Operator: &"-",
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 1217, Line: 80, Column: 16},
                        Reference: &"b",
                      },
                      Operator: &"+",
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 1221, Line: 80, Column: 20},
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
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1229, Line: 82, Column: 3},
          ReturnStatement: &compiler.ReturnStatement{
            Pos: Position{Filename: "", Offset: 1229, Line: 82, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 1236, Line: 82, Column: 10},
                Reference: &"a",
              },
              Operator: &"+",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1240, Line: 82, Column: 14},
                  Reference: &"b",
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
      Pos: Position{Filename: "", Offset: 1245, Line: 85, Column: 1},
      Name: "for equal or not equal",
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1279, Line: 86, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1279, Line: 86, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 1287, Line: 86, Column: 11},
                Int: &1,
              },
              Operator: &"==",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1292, Line: 86, Column: 16},
                  Int: &1,
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1296, Line: 87, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1296, Line: 87, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 1304, Line: 87, Column: 11},
                Int: &2,
              },
              Operator: &"==",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1309, Line: 87, Column: 16},
                  Int: &2,
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1313, Line: 88, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1313, Line: 88, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 1321, Line: 88, Column: 11},
                Int: &4,
              },
              Operator: &"==",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1326, Line: 88, Column: 16},
                  Int: &4,
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 1330, Line: 89, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 1330, Line: 89, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 1338, Line: 89, Column: 11},
                Int: &3,
              },
              Operator: &"!=",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 1343, Line: 89, Column: 16},
                  Int: &13,
                },
              },
            },
          },
        },
      },
    },
  },
}
