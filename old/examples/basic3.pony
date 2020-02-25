import crysis
import react
import dom

// primitive Color(Enum):
//     fun gray() => 0
//     fun gray() => 0


    
class Person(Model, Component):
    name: String
    age: Int
    dob: Int
    createdAt: Date
    updatedAt: Date

    class Meta:
        name = models.String(min: 10, max: 100, regex: '/hello/')
        dob = models.Int(min: 0, max: 100)
    
    fun age(): Int
        return 0

    fun fullname():
        return "${name} full"

    fun render():
        pass

struct Person 


class App:

    fun render(): Node
        return Col(
            Col(
                Text('123'),
            ),
            Col(
                Text('123'),
            ),
        )



ReactDom.Render(App, 'root')