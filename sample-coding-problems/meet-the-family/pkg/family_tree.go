// TODO(karuppiah7890): Try to change the name of the package and the package directory 
// from pkg to something like family_tree or familytree or similar. pkg is too generic
package pkg

import "fmt"

// TODO(karuppiah7890): Check if it's okay to not export a complete struct type
// directly but instead through NewXYZ function which is like a constructor
type familyTree struct {
	familyCount int
	family      map[string]*Person
}

func NewFamilyTree() familyTree {
	return familyTree{
		familyCount: 0,
		family:      make(map[string]*Person),
	}
}

func (f *familyTree) AddPerson(name string, gender Gender) {
	f.familyCount++
	person := Person{
		id:     f.familyCount,
		name:   name,
		gender: gender,
	}
	f.family[person.name] = &person
}

func (f *familyTree) GetFamilyCount() int {
	return f.familyCount
}

func (f *familyTree) GetPerson(name string) (*Person, error) {
	person, ok := f.family[name]

	if !ok {
		return nil, fmt.Errorf("Person named '%s' not found", name)
	}

	return person, nil
}
