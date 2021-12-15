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

type PersonNotFoundError string

func (e PersonNotFoundError) Error() string {
	return string(e)
}

type PersonNotMotherError string

func (e PersonNotMotherError) Error() string {
	return string(e)
}

// TODO(karuppiah7890): Think about initializing a Family Tree with at least one / exactly one person.
// This person would be kinda the root of the family. We then won't need AddPerson to add the root person
// of the family. AddPerson has a problem where a Person could be added with no relationship to any person
// in the family which doesn't make sense, and it's helpful to just add the root person (male/female). We can
// use something like AddSpouse or generic AddRelation to add other relations / relatives (people with relationship)
// after the root person is present in the family tree

func NewFamilyTree() familyTree {
	return familyTree{
		familyCount: 0,
		family:      make(map[string]*Person),
	}
}

func (f *familyTree) AddPerson(name string, gender Gender) {
	f.familyCount++
	person := Person{
		id:        f.familyCount,
		name:      name,
		gender:    gender,
		relations: make(map[string][]*Person),
	}
	f.family[person.name] = &person
}

func (f *familyTree) AddChildToMother(motherName string, childName string, childGender Gender) error {
	mother, err := f.GetPerson(motherName)

	if err != nil {
		return PersonNotFoundError(fmt.Sprintf("Mother named '%s' not found", motherName))
	}

	if !mother.gender.IsFemale() {
		return PersonNotMotherError(fmt.Sprintf("Person named '%s' is not a mother", motherName))
	}

	f.familyCount++
	child := &Person{
		id:        f.familyCount,
		name:      childName,
		gender:    childGender,
		relations: make(map[string][]*Person),
	}

	// TODO(karuppiah7890): This feels a bit complicated due to some premature generalization idea for any
	// relationship. We should make it specific and simple and give an idea of how to generalize it for later
	// but not do it
	mother.relations[Child.Name] = append(mother.relations[Child.Name], child)
	child.relations[Child.ReverseRelationship.Name] = append(child.relations[Child.ReverseRelationship.Name], mother)

	f.family[child.name] = child

	return nil
}

// TODO(karuppiah7890): AddSpouseToPerson(personName string, spouseName string, spouseGender Gender)

func (f *familyTree) AddSpouseToPerson(personName string, spouseName string, spouseGender Gender) error {
	person, err := f.GetPerson(personName)

	if err != nil {
		return PersonNotFoundError(fmt.Sprintf("Person named '%s' not found", personName))
	}

	// TODO(karuppiah7890): Adding a new person to a family seems like a usual process - incrementing family
	// count and constructing a person with ID and adding them. We need to get rid of this duplicate code here
	// and in other places
	f.familyCount++
	spouse := &Person{
		id:        f.familyCount,
		name:      spouseName,
		gender:    spouseGender,
		relations: make(map[string][]*Person),
	}

	// TODO(karuppiah7890): This feels a bit complicated due to some premature generalization idea for any
	// relationship. We should make it specific and simple and give an idea of how to generalize it for later
	// but not do it
	person.relations[Spouse.Name] = append(person.relations[Spouse.Name], spouse)
	spouse.relations[Spouse.ReverseRelationship.Name] = append(spouse.relations[Spouse.ReverseRelationship.Name], person)

	f.family[spouse.name] = spouse

	return nil
}

func (f *familyTree) GetFamilyCount() int {
	return f.familyCount
}

func (f *familyTree) GetPerson(name string) (*Person, error) {
	person, ok := f.family[name]

	// TODO(karuppiah7890): Use PersonNotFoundError error type here
	if !ok {
		return nil, fmt.Errorf("Person named '%s' not found", name)
	}

	return person, nil
}
