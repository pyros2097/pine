&compiler.Ast{
  Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
  Entries: []*compiler.Entry{
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1},
      Package: &"hello",
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 15, Line: 3, Column: 1},
      Imports: []*compiler.ImportStatement{
        &compiler.ImportStatement{
          Pos: Position{Filename: "", Offset: 15, Line: 3, Column: 1},
          Name: "time",
        },
      },
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 27, Line: 4, Column: 1},
      Imports: []*compiler.ImportStatement{
        &compiler.ImportStatement{
          Pos: Position{Filename: "", Offset: 27, Line: 4, Column: 1},
          Name: "ui",
        },
      },
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 37, Line: 5, Column: 1},
      Imports: []*compiler.ImportStatement{
        &compiler.ImportStatement{
          Pos: Position{Filename: "", Offset: 37, Line: 5, Column: 1},
          Name: "dom",
        },
      },
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 132, Line: 15, Column: 1},
      Classes: []*compiler.Class{
        &compiler.Class{
          Pos: Position{Filename: "", Offset: 132, Line: 15, Column: 1},
          Name: "Person",
          Traits: []*compiler.ClassTrait{
            &compiler.ClassTrait{
              Pos: Position{Filename: "", Offset: 145, Line: 15, Column: 14},
              ID: "Component",
            },
          },
          Fields: []*compiler.ClassField{
            &compiler.ClassField{
              Pos: Position{Filename: "", Offset: 160, Line: 16, Column: 5},
              Name: "name",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 164, Line: 16, Column: 9},
                Name: "String",
              },
            },
            &compiler.ClassField{
              Pos: Position{Filename: "", Offset: 177, Line: 17, Column: 5},
              Name: "dob",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 180, Line: 17, Column: 8},
                Name: "Time",
              },
            },
            &compiler.ClassField{
              Pos: Position{Filename: "", Offset: 191, Line: 18, Column: 5},
              Name: "createdAt",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 200, Line: 18, Column: 14},
                Name: "Time",
              },
            },
            &compiler.ClassField{
              Pos: Position{Filename: "", Offset: 211, Line: 19, Column: 5},
              Name: "updatedAt",
              Type: &compiler.Type{
                Pos: Position{Filename: "", Offset: 220, Line: 19, Column: 14},
                Name: "Time",
              },
            },
          },
          Methods: []*compiler.ClassMethod{
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 232, Line: 21, Column: 5},
              Name: "age",
              ReturnType: &compiler.Type{
                Pos: Position{Filename: "", Offset: 241, Line: 21, Column: 14},
                Name: "Int",
              },
              Body: []*compiler.Statement{
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 255, Line: 22, Column: 9},
                  ReturnStatement: &compiler.ReturnStatement{
                    Pos: Position{Filename: "", Offset: 255, Line: 22, Column: 9},
                    Liter: 0,
                  },
                },
              },
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 277, Line: 25, Column: 5},
              Name: "fullname",
              ReturnType: &compiler.Type{
                Pos: Position{Filename: "", Offset: 291, Line: 25, Column: 19},
                Name: "String",
              },
              Body: []*compiler.Statement{
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 308, Line: 26, Column: 9},
                  ReturnStatement: &compiler.ReturnStatement{
                    Pos: Position{Filename: "", Offset: 308, Line: 26, Column: 9},
                    Liter: "${name} full",
                  },
                },
              },
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 343, Line: 29, Column: 5},
              Name: "render",
              Body: []*compiler.Statement{
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 364, Line: 30, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 364, Line: 30, Column: 9},
                    Name: "age",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 371, Line: 30, Column: 16},
                      Literal: 35,
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 382, Line: 31, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 382, Line: 31, Column: 9},
                    Name: "text",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 390, Line: 31, Column: 17},
                      Literal: "123",
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 404, Line: 32, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 404, Line: 32, Column: 9},
                    Name: "arr",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 411, Line: 32, Column: 16},
                      Literal: [1, 2, 3],
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 429, Line: 33, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 429, Line: 33, Column: 9},
                    Name: "mymap",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 438, Line: 33, Column: 18},
                      Map: &compiler.MapStatement{
                        Pos: Position{Filename: "", Offset: 438, Line: 33, Column: 18},
                        Value: &compiler.MapLiteral{
                          Pos: Position{Filename: "", Offset: 478, Line: 35, Column: 13},
                          Key: &"555",
                          Value: &compiler.MapValue{
                            Pos: Position{Filename: "", Offset: 485, Line: 35, Column: 20},
                            Literal: &compiler.AssignmentLiteral{
                              Pos: Position{Filename: "", Offset: 485, Line: 35, Column: 20},
                              Map: &compiler.MapStatement{
                                Pos: Position{Filename: "", Offset: 485, Line: 35, Column: 20},
                                Value: &compiler.MapLiteral{
                                  Pos: Position{Filename: "", Offset: 503, Line: 36, Column: 17},
                                  Key: &"ttt",
                                  Value: &compiler.MapValue{
                                    Pos: Position{Filename: "", Offset: 510, Line: 36, Column: 24},
                                    Literal: &compiler.AssignmentLiteral{
                                      Pos: Position{Filename: "", Offset: 510, Line: 36, Column: 24},
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
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 548, Line: 39, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 548, Line: 39, Column: 9},
                    Name: "fullName",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 560, Line: 39, Column: 21},
                      Call: &compiler.FuncCallStatement{
                        Pos: Position{Filename: "", Offset: 560, Line: 39, Column: 21},
                        Name: "getFullName",
                      },
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 582, Line: 40, Column: 9},
                  Assignment: &compiler.AssignmentStatement{
                    Pos: Position{Filename: "", Offset: 582, Line: 40, Column: 9},
                    Name: "user",
                    Value: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 590, Line: 40, Column: 17},
                      Call: &compiler.FuncCallStatement{
                        Pos: Position{Filename: "", Offset: 590, Line: 40, Column: 17},
                        Name: "getUser",
                        Values: []*compiler.Literal{
                          1,
                          2,
                          3,
                        },
                      },
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 615, Line: 41, Column: 9},
                  FuncCall: &compiler.FuncCallStatement{
                    Pos: Position{Filename: "", Offset: 615, Line: 41, Column: 9},
                    Name: "loadData",
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 634, Line: 42, Column: 9},
                  FuncCall: &compiler.FuncCallStatement{
                    Pos: Position{Filename: "", Offset: 634, Line: 42, Column: 9},
                    Name: "loadData",
                    Values: []*compiler.Literal{
                      1,
                      age,
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 693, Line: 44, Column: 9},
                  If: &compiler.IfStatement{
                    Pos: Position{Filename: "", Offset: 693, Line: 44, Column: 9},
                    Condition: "abc",
                    End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 720, Line: 46, Column: 9},
                  For: &compiler.ForStatement{
                    Pos: Position{Filename: "", Offset: 720, Line: 46, Column: 9},
                    Infinite: &compiler.ForInfiniteStatement{
                      Pos: Position{Filename: "", Offset: 720, Line: 46, Column: 9},
                      End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 744, Line: 48, Column: 9},
                  For: &compiler.ForStatement{
                    Pos: Position{Filename: "", Offset: 744, Line: 48, Column: 9},
                    Conditional: &compiler.ForConditionalStatement{
                      Pos: Position{Filename: "", Offset: 744, Line: 48, Column: 9},
                      Identifier: "flag",
                      End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
                    },
                  },
                },
              },
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 820, Line: 54, Column: 5},
              Name: "create_data",
              Parameters: []*compiler.MethodParameter{
                &compiler.MethodParameter{
                  Pos: Position{Filename: "", Offset: 836, Line: 54, Column: 21},
                  ID: "delta",
                  Type: &compiler.Type{
                    Pos: Position{Filename: "", Offset: 841, Line: 54, Column: 26},
                    Name: "Time",
                  },
                },
                &compiler.MethodParameter{
                  Pos: Position{Filename: "", Offset: 849, Line: 54, Column: 34},
                  ID: "frame",
                  Type: &compiler.Type{
                    Pos: Position{Filename: "", Offset: 854, Line: 54, Column: 39},
                    Name: "RenderFrame",
                  },
                },
              },
              ReturnType: &compiler.Type{
                Pos: Position{Filename: "", Offset: 869, Line: 54, Column: 54},
                Name: "bool",
              },
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 889, Line: 58, Column: 1},
      Classes: []*compiler.Class{
        &compiler.Class{
          Pos: Position{Filename: "", Offset: 889, Line: 58, Column: 1},
          Name: "PersonTest",
          Traits: []*compiler.ClassTrait{
            &compiler.ClassTrait{
              Pos: Position{Filename: "", Offset: 906, Line: 58, Column: 18},
              ID: "Test",
            },
          },
          Methods: []*compiler.ClassMethod{
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 917, Line: 60, Column: 5},
              Name: "render_fail",
              Body: []*compiler.Statement{
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 943, Line: 61, Column: 9},
                  AssertStatement: &compiler.AssertStatement{
                    Pos: Position{Filename: "", Offset: 943, Line: 61, Column: 9},
                    LHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 950, Line: 61, Column: 16},
                      Literal: 1,
                    },
                    Operator: &compiler.Operator{
                      Pos: Position{Filename: "", Offset: 952, Line: 61, Column: 18},
                      Value: "==",
                    },
                    RHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 955, Line: 61, Column: 21},
                      Literal: 2,
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 965, Line: 62, Column: 9},
                  AssertStatement: &compiler.AssertStatement{
                    Pos: Position{Filename: "", Offset: 965, Line: 62, Column: 9},
                    LHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 972, Line: 62, Column: 16},
                      Literal: abc,
                    },
                    Operator: &compiler.Operator{
                      Pos: Position{Filename: "", Offset: 976, Line: 62, Column: 20},
                      Value: "==",
                    },
                    RHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 979, Line: 62, Column: 23},
                      Literal: cde,
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 991, Line: 63, Column: 9},
                  AssertStatement: &compiler.AssertStatement{
                    Pos: Position{Filename: "", Offset: 991, Line: 63, Column: 9},
                    LHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 998, Line: 63, Column: 16},
                      Call: &compiler.FuncCallStatement{
                        Pos: Position{Filename: "", Offset: 998, Line: 63, Column: 16},
                        Name: "getValue",
                      },
                    },
                    Operator: &compiler.Operator{
                      Pos: Position{Filename: "", Offset: 1009, Line: 63, Column: 27},
                      Value: "==",
                    },
                    RHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 1012, Line: 63, Column: 30},
                      Literal: 25,
                    },
                  },
                },
                &compiler.Statement{
                  Pos: Position{Filename: "", Offset: 1023, Line: 64, Column: 9},
                  AssertStatement: &compiler.AssertStatement{
                    Pos: Position{Filename: "", Offset: 1023, Line: 64, Column: 9},
                    LHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 1030, Line: 64, Column: 16},
                      Call: &compiler.FuncCallStatement{
                        Pos: Position{Filename: "", Offset: 1030, Line: 64, Column: 16},
                        Name: "totalDiscount",
                      },
                    },
                    Operator: &compiler.Operator{
                      Pos: Position{Filename: "", Offset: 1046, Line: 64, Column: 32},
                      Value: "==",
                    },
                    RHS: &compiler.AssignmentLiteral{
                      Pos: Position{Filename: "", Offset: 1049, Line: 64, Column: 35},
                      Literal: 300.9,
                    },
                  },
                },
              },
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
    &compiler.Entry{
      Pos: Position{Filename: "", Offset: 1552, Line: 89, Column: 1},
      Classes: []*compiler.Class{
        &compiler.Class{
          Pos: Position{Filename: "", Offset: 1552, Line: 89, Column: 1},
          Name: "Main",
          Methods: []*compiler.ClassMethod{
            &compiler.ClassMethod{
              Pos: Position{Filename: "", Offset: 1567, Line: 90, Column: 5},
              Name: "main",
              End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
            },
          },
          End: Position{Filename: "", Offset: 0, Line: 0, Column: 0},
        },
      },
    },
  },
}
