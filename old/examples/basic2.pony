import crysis
import react
import dom

class Person(Model, Component)
    name: String
    age: Int
    dob: Int
    createdAt: Date
    updatedAt: Date

    class Meta:
        name = models.String(min: 10, max: 100, regex: '/hello/')
        dob = models.Int(min: 0, max: 100)
    
    fun age():
        return 0
    end

    fun fullname()
        return "%s %s" % (firstname, lastname)
    end

    fun render()
    end
end
