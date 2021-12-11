package pkg

// TODO(karuppiah7890): Check if it's okay to export Person struct. Or we could actually have a NewPerson
// function, like a constructor for Person and not export Person and instead use 'person'

// TODO(karuppiah7890): Check if it's okay to put Person struct in the same package as FamilyTree struct.
// This question comes up because when they both are in the same package, and all their fields are not exported
// for outside packages, but they are still accessible within this package for read and write. Should a method or
// function in person.go file be able to modify the fields in FamilyTree struct with the write access it has.
// But yeah, Person and FamilyTree structs are closely related

type Person struct {
	id        int
	name      string
	gender    Gender
	relations map[string][]*Person
}

func (p *Person) GetID() int {
	return p.id
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetGender() Gender {
	return p.gender
}

func (p *Person) GetRelations() map[string][]*Person {
	return p.relations
}
