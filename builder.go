type Person struct {
    Name string
    Age  int
    Job  string
}

type PersonBuilder struct {
    person *Person
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{
        person: &Person{},
    }
}

func (b *PersonBuilder) SetName(name string) *PersonBuilder {
    b.person.Name = name
    return b
}

func (b *PersonBuilder) SetAge(age int) *PersonBuilder {
    b.person.Age = age
    return b
}

func (b *PersonBuilder) SetJob(job string) *PersonBuilder {
    b.person.Job = job
    return b
}

func (b *PersonBuilder) Build() *Person {
    return b.person
}

// Usage
func main() {
    personBuilder := NewPersonBuilder()
    person := personBuilder.
        SetName("John Doe").
        SetAge(30).
        SetJob("Software Engineer").
        Build()

    fmt.Println(person) // Output: {John Doe 30 Software Engineer}
}
