&compiler.Module{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Name: "main",
  Enums: []*compiler.Enum{
    &compiler.Enum{
      Pos: Position{Filename: "", Offset: 13, Line: 3, Column: 1},
      Name: "Direction",
      Value: []string{
        "Left",
        "Right",
      },
    },
    &compiler.Enum{
      Pos: Position{Filename: "", Offset: 48, Line: 8, Column: 1},
      Name: "Position",
      Value: []string{
        "Top",
        "Bottom",
      },
    },
  },
  Types: []*compiler.Type{
    &compiler.Type{
      Pos: Position{Filename: "", Offset: 82, Line: 13, Column: 1},
      Name: "User",
      Fields: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 96, Line: 14, Column: 3},
          Name: "name",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 111, Line: 15, Column: 3},
          Name: "age",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 125, Line: 16, Column: 3},
          Name: "gender",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 142, Line: 17, Column: 3},
          Name: "address",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 160, Line: 18, Column: 3},
          Name: "createdAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 180, Line: 19, Column: 3},
          Name: "updatedAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 200, Line: 20, Column: 3},
          Name: "deletedAt",
          Value: "string",
        },
      },
    },
  },
  Functions: []*compiler.Function{
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 221, Line: 23, Column: 1},
      Type: "proc",
      Name: "adder",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 232, Line: 23, Column: 12},
          Name: "a",
          Value: "int",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 240, Line: 23, Column: 20},
          Name: "b",
          Value: "int",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 248, Line: 23, Column: 28},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 259, Line: 24, Column: 3},
          If: &compiler.IfStatement{
            Pos: Position{Filename: "", Offset: 259, Line: 24, Column: 3},
            Condition: []*compiler.Expression{
              &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 262, Line: 24, Column: 6},
                  Reference: &"a",
                },
                Operator: &"<",
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 266, Line: 24, Column: 10},
                    Reference: &"b",
                  },
                },
              },
            },
            Result: []*compiler.Statement{
              &compiler.Statement{
                Pos: Position{Filename: "", Offset: 274, Line: 25, Column: 5},
                ReturnStatement: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 274, Line: 25, Column: 5},
                    Reference: &"return",
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 281, Line: 25, Column: 12},
                      Reference: &"a",
                    },
                    Operator: &"-",
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 285, Line: 25, Column: 16},
                        Reference: &"b",
                      },
                      Operator: &"+",
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 289, Line: 25, Column: 20},
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
          Pos: Position{Filename: "", Offset: 297, Line: 27, Column: 3},
          ReturnStatement: &compiler.Expression{
            Left: &compiler.Literal{
              Pos: Position{Filename: "", Offset: 297, Line: 27, Column: 3},
              Reference: &"return",
            },
            Right: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 304, Line: 27, Column: 10},
                Reference: &"a",
              },
              Operator: &"+",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 308, Line: 27, Column: 14},
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
      Pos: Position{Filename: "", Offset: 313, Line: 30, Column: 1},
      Name: "for equal or not equal",
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 347, Line: 31, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 347, Line: 31, Column: 3},
            Expression: &compiler.Expression{
              Left: &compiler.Literal{
                Pos: Position{Filename: "", Offset: 354, Line: 31, Column: 10},
                Int: &1,
              },
              Operator: &"==",
              Right: &compiler.Expression{
                Left: &compiler.Literal{
                  Pos: Position{Filename: "", Offset: 359, Line: 31, Column: 15},
                  Int: &1,
                },
                Right: &compiler.Expression{
                  Left: &compiler.Literal{
                    Pos: Position{Filename: "", Offset: 363, Line: 32, Column: 3},
                    Reference: &"assert",
                  },
                  Right: &compiler.Expression{
                    Left: &compiler.Literal{
                      Pos: Position{Filename: "", Offset: 370, Line: 32, Column: 10},
                      Int: &2,
                    },
                    Operator: &"==",
                    Right: &compiler.Expression{
                      Left: &compiler.Literal{
                        Pos: Position{Filename: "", Offset: 375, Line: 32, Column: 15},
                        Int: &2,
                      },
                      Right: &compiler.Expression{
                        Left: &compiler.Literal{
                          Pos: Position{Filename: "", Offset: 379, Line: 33, Column: 3},
                          Reference: &"assert",
                        },
                        Right: &compiler.Expression{
                          Left: &compiler.Literal{
                            Pos: Position{Filename: "", Offset: 386, Line: 33, Column: 10},
                            Int: &4,
                          },
                          Operator: &"==",
                          Right: &compiler.Expression{
                            Left: &compiler.Literal{
                              Pos: Position{Filename: "", Offset: 391, Line: 33, Column: 15},
                              Int: &4,
                            },
                            Right: &compiler.Expression{
                              Left: &compiler.Literal{
                                Pos: Position{Filename: "", Offset: 395, Line: 34, Column: 3},
                                Reference: &"assert",
                              },
                              Right: &compiler.Expression{
                                Left: &compiler.Literal{
                                  Pos: Position{Filename: "", Offset: 402, Line: 34, Column: 10},
                                  Int: &3,
                                },
                                Operator: &"!=",
                                Right: &compiler.Expression{
                                  Left: &compiler.Literal{
                                    Pos: Position{Filename: "", Offset: 407, Line: 34, Column: 15},
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
                },
              },
            },
          },
        },
      },
    },
  },
}
