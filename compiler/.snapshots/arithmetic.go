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
          Pos: Position{Filename: "", Offset: 110, Line: 15, Column: 3},
          Name: "age",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 123, Line: 16, Column: 3},
          Name: "gender",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 139, Line: 17, Column: 3},
          Name: "address",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 156, Line: 18, Column: 3},
          Name: "createdAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 175, Line: 19, Column: 3},
          Name: "updatedAt",
          Value: "string",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 194, Line: 20, Column: 3},
          Name: "deletedAt",
          Value: "string",
        },
      },
    },
  },
  Functions: []*compiler.Function{
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 214, Line: 23, Column: 1},
      Type: "method",
      Name: "fullName",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 230, Line: 23, Column: 17},
          Name: "u",
          Value: "User",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 239, Line: 23, Column: 26},
        Name: "string",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 251, Line: 24, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 251, Line: 24, Column: 3},
            Name: "name",
            Expression: u.name map :=  return u.name + " " + u.age,
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 392, Line: 35, Column: 1},
      Type: "proc",
      Name: "Counter",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 405, Line: 35, Column: 14},
          Name: "a",
          Value: "int",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 412, Line: 35, Column: 21},
          Name: "b",
          Value: "int",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 419, Line: 35, Column: 28},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 428, Line: 36, Column: 3},
          FuncCall: &compiler.FuncCallStatement{
            Pos: Position{Filename: "", Offset: 428, Line: 36, Column: 3},
            Name: "return",
            Values: []*compiler.Literal{
                React.createElement(
    View,
    {},
    React.createElement(
        View,
        {},
        React.createElement(
                Text,
                {},
        )
,     )
,     React.createElement(
        View,
        {},
        React.createElement(
                Text,
                {},
        )
,     )
,   )
,
            },
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 611, Line: 52, Column: 1},
      Type: "proc",
      Name: "adder",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 622, Line: 52, Column: 12},
          Name: "a",
          Value: "int",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 629, Line: 52, Column: 19},
          Name: "b",
          Value: "int",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 637, Line: 52, Column: 27},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 646, Line: 53, Column: 3},
          If: &compiler.IfStatement{
            Pos: Position{Filename: "", Offset: 646, Line: 53, Column: 3},
            Condition: []*compiler.Expression{
              a < b,
            },
            Result: []*compiler.Statement{
              &compiler.Statement{
                Pos: Position{Filename: "", Offset: 661, Line: 54, Column: 5},
                ReturnStatement: &compiler.ReturnStatement{
                  Pos: Position{Filename: "", Offset: 661, Line: 54, Column: 5},
                  Expression: a - b + b,
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 684, Line: 56, Column: 3},
          ReturnStatement: &compiler.ReturnStatement{
            Pos: Position{Filename: "", Offset: 684, Line: 56, Column: 3},
            Expression: a + b,
          },
        },
      },
    },
  },
  Tests: []*compiler.Test{
    &compiler.Test{
      Pos: Position{Filename: "", Offset: 700, Line: 59, Column: 1},
      Name: "for equal or not equal",
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 734, Line: 60, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 734, Line: 60, Column: 3},
            Expression: 1 == 1,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 751, Line: 61, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 751, Line: 61, Column: 3},
            Expression: 2 == 2,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 768, Line: 62, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 768, Line: 62, Column: 3},
            Expression: 4 == 4,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 785, Line: 63, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 785, Line: 63, Column: 3},
            Expression: 3 != 13,
          },
        },
      },
    },
  },
}
