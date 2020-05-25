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
      Type: "method",
      Name: "fullName",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 237, Line: 23, Column: 17},
          Name: "u",
          Value: "User",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 246, Line: 23, Column: 26},
        Name: "string",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 260, Line: 24, Column: 3},
          Assignment: &compiler.AssignmentStatement{
            Pos: Position{Filename: "", Offset: 260, Line: 24, Column: 3},
            Name: "name",
            Expression: u.name map :=  return u.name + " " + u.age,
          },
        },
      },
    },
    &compiler.Function{
      Pos: Position{Filename: "", Offset: 401, Line: 35, Column: 1},
      Type: "proc",
      Name: "Counter",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 414, Line: 35, Column: 14},
          Name: "a",
          Value: "int",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 422, Line: 35, Column: 22},
          Name: "b",
          Value: "int",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 430, Line: 35, Column: 30},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 441, Line: 36, Column: 3},
          FuncCall: &compiler.FuncCallStatement{
            Pos: Position{Filename: "", Offset: 441, Line: 36, Column: 3},
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
      Pos: Position{Filename: "", Offset: 624, Line: 52, Column: 1},
      Type: "proc",
      Name: "adder",
      Parameters: []compiler.TypeField{
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 635, Line: 52, Column: 12},
          Name: "a",
          Value: "int",
        },
        compiler.TypeField{
          Pos: Position{Filename: "", Offset: 643, Line: 52, Column: 20},
          Name: "b",
          Value: "int",
        },
      },
      ReturnType: &compiler.ReturnType{
        Pos: Position{Filename: "", Offset: 651, Line: 52, Column: 28},
        Name: "int",
      },
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 662, Line: 53, Column: 3},
          If: &compiler.IfStatement{
            Pos: Position{Filename: "", Offset: 662, Line: 53, Column: 3},
            Condition: []*compiler.Expression{
              a < b,
            },
            Result: []*compiler.Statement{
              &compiler.Statement{
                Pos: Position{Filename: "", Offset: 677, Line: 54, Column: 5},
                ReturnStatement: &compiler.ReturnStatement{
                  Pos: Position{Filename: "", Offset: 677, Line: 54, Column: 5},
                  Expression: a - b + b,
                },
              },
            },
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 700, Line: 56, Column: 3},
          ReturnStatement: &compiler.ReturnStatement{
            Pos: Position{Filename: "", Offset: 700, Line: 56, Column: 3},
            Expression: a + b,
          },
        },
      },
    },
  },
  Tests: []*compiler.Test{
    &compiler.Test{
      Pos: Position{Filename: "", Offset: 716, Line: 59, Column: 1},
      Name: "for equal or not equal",
      Statements: []*compiler.Statement{
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 750, Line: 60, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 750, Line: 60, Column: 3},
            Expression: 1 == 1,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 767, Line: 61, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 767, Line: 61, Column: 3},
            Expression: 2 == 2,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 784, Line: 62, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 784, Line: 62, Column: 3},
            Expression: 4 == 4,
          },
        },
        &compiler.Statement{
          Pos: Position{Filename: "", Offset: 801, Line: 63, Column: 3},
          AssertStatement: &compiler.AssertStatement{
            Pos: Position{Filename: "", Offset: 801, Line: 63, Column: 3},
            Expression: 3 != 13,
          },
        },
      },
    },
  },
}
