package hello

import time
import ui
import dom

// primitive Dom
// end

// primitive String
//     fun hello()
//     end
// end

class Person(Component)
    name: String
    dob: Time
    createdAt: Time
    updatedAt: Time

    fun age(): Int
        return 0
    end

    fun fullname(): String
        return "${name} full"
    end

    fun render()
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
        loadData()
        loadData(1, age)
        // loadData(fetch("123"))
        if abc
        end
        for
        end
        for flag
        end
        // for num in nums
        // end
    end

    fun create_data(delta: Time, frame: RenderFrame) : bool
    end
end

class PersonTest(Test)

    fun render_fail()
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
end

class Counter
    var count: I32

    fun render(): Node
        return Row(
            flex: 0.5,
            chilren: Array(
                Text("123")
            ),
            Text("123")
        )
    end
end

class Counter
    var count: I32
    return Row(
        flex: 0.5,
        chilren: Array(
            Text("123")
        ),
        Text("123")
    )
end

class Main
    fun main()
        // Dom.render(Person)
    end
end