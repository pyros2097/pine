&ast.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Entries: []*ast.Entry{
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Package: &"hello",
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 15, Line: 3, Column: 1},
      Imports: []*ast.ImportStatement{
        &ast.ImportStatement{
          Pos: Position{Filename: "", Offset: 15, Line: 3, Column: 1},
          Name: "time",
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 27, Line: 4, Column: 1},
      Imports: []*ast.ImportStatement{
        &ast.ImportStatement{
          Pos: Position{Filename: "", Offset: 27, Line: 4, Column: 1},
          Name: "ui",
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 38, Line: 6, Column: 1},
      Structs: []*ast.Struct{
        &ast.Struct{
          Pos: Position{Filename: "", Offset: 38, Line: 6, Column: 1},
          Name: "Person",
          Traits: []*ast.ClassTrait{
            &ast.ClassTrait{
              Pos: Position{Filename: "", Offset: 51, Line: 6, Column: 14},
              ID: "Component",
            },
          },
          Fields: []*ast.ClassField{
            &ast.ClassField{
              Pos: Position{Filename: "", Offset: 66, Line: 7, Column: 5},
              Name: "name",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 70, Line: 7, Column: 9},
                Name: "String",
              },
            },
            &ast.ClassField{
              Pos: Position{Filename: "", Offset: 83, Line: 8, Column: 5},
              Name: "dob",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 86, Line: 8, Column: 8},
                Name: "Time",
              },
            },
            &ast.ClassField{
              Pos: Position{Filename: "", Offset: 97, Line: 9, Column: 5},
              Name: "createdAt",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 106, Line: 9, Column: 14},
                Name: "Time",
              },
            },
            &ast.ClassField{
              Pos: Position{Filename: "", Offset: 117, Line: 10, Column: 5},
              Name: "updatedAt",
              Type: &ast.Type{
                Pos: Position{Filename: "", Offset: 126, Line: 10, Column: 14},
                Name: "Time",
              },
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 138, Line: 13, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 138, Line: 13, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 138, Line: 13, Column: 1},
            Name: "age",
            ReturnType: &ast.ReturnType{
              Pos: Position{Filename: "", Offset: 148, Line: 13, Column: 11},
              Name: "Int",
            },
            Body: []*ast.Statement{
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 159, Line: 14, Column: 5},
                ReturnStatement: &ast.ReturnStatement{
                  Pos: Position{Filename: "", Offset: 159, Line: 14, Column: 5},
                  Liter: 0,
                },
              },
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 173, Line: 17, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 173, Line: 17, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 173, Line: 17, Column: 1},
            Name: "fullname",
            ReturnType: &ast.ReturnType{
              Pos: Position{Filename: "", Offset: 188, Line: 17, Column: 16},
              Name: "String",
            },
            Body: []*ast.Statement{
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 202, Line: 18, Column: 5},
                ReturnStatement: &ast.ReturnStatement{
                  Pos: Position{Filename: "", Offset: 202, Line: 18, Column: 5},
                  Liter: "${name} full",
                },
              },
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 229, Line: 21, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 229, Line: 21, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 229, Line: 21, Column: 1},
            Name: "render",
            Body: []*ast.Statement{
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 246, Line: 22, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 246, Line: 22, Column: 5},
                  Name: "age",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 253, Line: 22, Column: 12},
                    Literal: 35,
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 260, Line: 23, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 260, Line: 23, Column: 5},
                  Name: "text",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 268, Line: 23, Column: 13},
                    Literal: "123",
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 278, Line: 24, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 278, Line: 24, Column: 5},
                  Name: "arr",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 285, Line: 24, Column: 12},
                    Literal: [1, 2, 3],
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 299, Line: 25, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 299, Line: 25, Column: 5},
                  Name: "mymap",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 308, Line: 25, Column: 14},
                    Map: &ast.MapStatement{
                      Pos: Position{Filename: "", Offset: 308, Line: 25, Column: 14},
                      Value: &ast.MapLiteral{
                        Pos: Position{Filename: "", Offset: 340, Line: 27, Column: 9},
                        Key: &"555",
                        Value: &ast.MapValue{
                          Pos: Position{Filename: "", Offset: 347, Line: 27, Column: 16},
                          Literal: &ast.AssignmentLiteral{
                            Pos: Position{Filename: "", Offset: 347, Line: 27, Column: 16},
                            Map: &ast.MapStatement{
                              Pos: Position{Filename: "", Offset: 347, Line: 27, Column: 16},
                              Value: &ast.MapLiteral{
                                Pos: Position{Filename: "", Offset: 361, Line: 28, Column: 13},
                                Key: &"ttt",
                                Value: &ast.MapValue{
                                  Pos: Position{Filename: "", Offset: 368, Line: 28, Column: 20},
                                  Literal: &ast.AssignmentLiteral{
                                    Pos: Position{Filename: "", Offset: 368, Line: 28, Column: 20},
                                    Literal: "aaa",
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
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 394, Line: 31, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 394, Line: 31, Column: 5},
                  Name: "fullName",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 406, Line: 31, Column: 17},
                    Call: &ast.FuncCallStatement{
                      Pos: Position{Filename: "", Offset: 406, Line: 31, Column: 17},
                      Name: "getFullName",
                    },
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 424, Line: 32, Column: 5},
                Assignment: &ast.AssignmentStatement{
                  Pos: Position{Filename: "", Offset: 424, Line: 32, Column: 5},
                  Name: "user",
                  Value: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 432, Line: 32, Column: 13},
                    Call: &ast.FuncCallStatement{
                      Pos: Position{Filename: "", Offset: 432, Line: 32, Column: 13},
                      Name: "getUser",
                      Values: []*ast.Literal{
                        1,
                        2,
                        3,
                      },
                    },
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 485, Line: 34, Column: 5},
                FuncCall: &ast.FuncCallStatement{
                  Pos: Position{Filename: "", Offset: 485, Line: 34, Column: 5},
                  Name: "loadData",
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 500, Line: 35, Column: 5},
                FuncCall: &ast.FuncCallStatement{
                  Pos: Position{Filename: "", Offset: 500, Line: 35, Column: 5},
                  Name: "loadData",
                  Values: []*ast.Literal{
                    1,
                    age,
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 551, Line: 37, Column: 5},
                If: &ast.IfStatement{
                  Pos: Position{Filename: "", Offset: 551, Line: 37, Column: 5},
                  Condition: "abc",
                  End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 570, Line: 39, Column: 5},
                For: &ast.ForStatement{
                  Pos: Position{Filename: "", Offset: 570, Line: 39, Column: 5},
                  Infinite: &ast.ForInfiniteStatement{
                    Pos: Position{Filename: "", Offset: 570, Line: 39, Column: 5},
                    End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 586, Line: 41, Column: 5},
                For: &ast.ForStatement{
                  Pos: Position{Filename: "", Offset: 586, Line: 41, Column: 5},
                  Conditional: &ast.ForConditionalStatement{
                    Pos: Position{Filename: "", Offset: 586, Line: 41, Column: 5},
                    Identifier: "flag",
                    End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                  },
                },
              },
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 642, Line: 47, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 642, Line: 47, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 642, Line: 47, Column: 1},
            Name: "create_data",
            Parameters: []*ast.MethodParameter{
              &ast.MethodParameter{
                Pos: Position{Filename: "", Offset: 658, Line: 47, Column: 17},
                ID: "delta",
                Type: &ast.Type{
                  Pos: Position{Filename: "", Offset: 663, Line: 47, Column: 22},
                  Name: "Time",
                },
              },
              &ast.MethodParameter{
                Pos: Position{Filename: "", Offset: 671, Line: 47, Column: 30},
                ID: "frame",
                Type: &ast.Type{
                  Pos: Position{Filename: "", Offset: 676, Line: 47, Column: 35},
                  Name: "RenderFrame",
                },
              },
            },
            ReturnType: &ast.ReturnType{
              Pos: Position{Filename: "", Offset: 691, Line: 47, Column: 50},
              Name: "bool",
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 704, Line: 50, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 704, Line: 50, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 704, Line: 50, Column: 1},
            Name: "test_render_fail",
            Body: []*ast.Statement{
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 731, Line: 51, Column: 5},
                AssertStatement: &ast.AssertStatement{
                  Pos: Position{Filename: "", Offset: 731, Line: 51, Column: 5},
                  LHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 738, Line: 51, Column: 12},
                    Literal: 1,
                  },
                  Operator: &ast.Operator{
                    Pos: Position{Filename: "", Offset: 740, Line: 51, Column: 14},
                    Value: "==",
                  },
                  RHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 743, Line: 51, Column: 17},
                    Literal: 2,
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 749, Line: 52, Column: 5},
                AssertStatement: &ast.AssertStatement{
                  Pos: Position{Filename: "", Offset: 749, Line: 52, Column: 5},
                  LHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 756, Line: 52, Column: 12},
                    Literal: abc,
                  },
                  Operator: &ast.Operator{
                    Pos: Position{Filename: "", Offset: 760, Line: 52, Column: 16},
                    Value: "==",
                  },
                  RHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 763, Line: 52, Column: 19},
                    Literal: cde,
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 771, Line: 53, Column: 5},
                AssertStatement: &ast.AssertStatement{
                  Pos: Position{Filename: "", Offset: 771, Line: 53, Column: 5},
                  LHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 778, Line: 53, Column: 12},
                    Call: &ast.FuncCallStatement{
                      Pos: Position{Filename: "", Offset: 778, Line: 53, Column: 12},
                      Name: "getValue",
                    },
                  },
                  Operator: &ast.Operator{
                    Pos: Position{Filename: "", Offset: 789, Line: 53, Column: 23},
                    Value: "==",
                  },
                  RHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 792, Line: 53, Column: 26},
                    Literal: 25,
                  },
                },
              },
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 799, Line: 54, Column: 5},
                AssertStatement: &ast.AssertStatement{
                  Pos: Position{Filename: "", Offset: 799, Line: 54, Column: 5},
                  LHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 806, Line: 54, Column: 12},
                    Call: &ast.FuncCallStatement{
                      Pos: Position{Filename: "", Offset: 806, Line: 54, Column: 12},
                      Name: "totalDiscount",
                    },
                  },
                  Operator: &ast.Operator{
                    Pos: Position{Filename: "", Offset: 822, Line: 54, Column: 28},
                    Value: "==",
                  },
                  RHS: &ast.AssignmentLiteral{
                    Pos: Position{Filename: "", Offset: 825, Line: 54, Column: 31},
                    Literal: 300.9,
                  },
                },
              },
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
    &ast.Entry{
      Pos: Position{Filename: "", Offset: 1248, Line: 78, Column: 1},
      Funs: []*ast.Fun{
        &ast.Fun{
          Pos: Position{Filename: "", Offset: 1248, Line: 78, Column: 1},
          FunDecl: ast.FunDecl{
            Pos: Position{Filename: "", Offset: 1248, Line: 78, Column: 1},
            Name: "main",
            Body: []*ast.Statement{
              &ast.Statement{
                Pos: Position{Filename: "", Offset: 1263, Line: 79, Column: 5},
                FuncCall: &ast.FuncCallStatement{
                  Pos: Position{Filename: "", Offset: 1263, Line: 79, Column: 5},
                  Name: "render",
                  Values: []*ast.Literal{
                    Person,
                  },
                },
              },
            },
            End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
          },
        },
      },
    },
  },
}
