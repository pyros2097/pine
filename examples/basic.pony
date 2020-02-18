package hello

import time
import ui

trait Component
    render()
end

class Person implements Component
    name: String
    dob: Time
    createdAt: Time
    updatedAt: Time
end

fun age() -> Int
    return 0
end

fun fullname(firstName, lastName string) -> String
    return "${firstName} ${lastName}"
end

func render() {
    age := 35
    text := "123"
    arr := [1, 2, 3]
    mymap := {
        "123": "123",
        "555": {
            "ttt": "aaa"
        }
    }
    fullName := getFullName()
    user := getUser(1, 2, 3)
    // v := (price * age * 100)
    loadData()
    loadData(1, age)
    // loadData(fetch("123"))
    if abc {

    }
    for i := 0; i < 4; i++ {

    }
    for flag {
    }
    // for num in nums
    // end
}

fun create_data(delta: Time, frame: RenderFrame) -> bool
end

fun test_render_fail()
    assert 1 == 2
    assert abc == cde
    assert getValue() == 25
    assert totalDiscount() == 300.9
    // given:, when:, then:

    // expect:
    //     Math.max(1, 3) == 3
    //     Math.max(7, 4) == 7
    //     Math.max(0, 0) == 0

    // expect:
    //     Math.max(a, b) == c
    // where:
    //     a := [5, 3]
    //     b := [1, 9]
    //     c := [5, 9]

    // expect:
    //     Math.max(a, b) == c
    // where:
    //     a | b | c
    //     1 | 3 | 3
    //     7 | 4 | 7
    //     0 | 0 | 0
end

fun main()
    render(Person)
end