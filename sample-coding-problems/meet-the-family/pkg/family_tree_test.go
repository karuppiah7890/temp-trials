package pkg_test

import (
	"testing"

	"github.com/karuppiah7890/temp-trials/sample-coding-problems/meet-the-family/pkg"
)

func TestAddPerson(t *testing.T) {
	expectedFamilyCount := 1
	expectedName := "King Shan"
	expectedGender := pkg.Male
	expectedID := 1

	f := pkg.NewFamilyTree()
	f.AddPerson(expectedName, expectedGender)

	familyCount := f.GetFamilyCount()
	if familyCount != expectedFamilyCount {
		t.Fatalf("expected family count to be %v but got %v", expectedFamilyCount, familyCount)
	}

	person, err := f.GetPerson(expectedName)
	if err != nil {
		t.Fatalf("expected no error while getting person but got error - %v", err)
	}

	assertPerson(t, person, expectedName, expectedGender, expectedID)
}

func TestAddChildToMother(t *testing.T) {
	t.Run("Child Added Successfully", func(t *testing.T) {
		expectedFamilyCount := 2
		expectedChildName := "Amba"
		expectedChildGender := pkg.Female
		expectedChildID := 2

		motherName := "Queen Anga"

		f := pkg.NewFamilyTree()

		f.AddPerson(motherName, pkg.Female)

		err := f.AddChildToMother(motherName, expectedChildName, expectedChildGender)
		if err != nil {
			t.Fatalf("expected no error while adding child to mother but got error - %v", err)
		}

		familyCount := f.GetFamilyCount()
		if familyCount != expectedFamilyCount {
			t.Fatalf("expected family count to be %v but got %v", expectedFamilyCount, familyCount)
		}

		child, err := f.GetPerson(expectedChildName)
		if err != nil {
			t.Fatalf("expected no error while getting person but got error - %v", err)
		}

		assertPerson(t, child, expectedChildName, expectedChildGender, expectedChildID)

		childRelations := child.GetRelations()
		parents := childRelations["Parent"]
		numberOfParents := len(parents)
		expectedNumberOfParents := 1

		if numberOfParents != expectedNumberOfParents {
			t.Fatalf("expected child to have %v parent(s) but got %v", expectedNumberOfParents, numberOfParents)
		}

		parent := parents[0]
		expectedParentName := motherName
		parentName := parent.GetName()
		expectedParentGender := pkg.Female
		parentGender := parent.GetGender()

		if parentName != expectedParentName {
			t.Fatalf("expected parent to have name '%v' but got '%v'", expectedParentName, parentName)
		}

		if parentGender != expectedParentGender {
			t.Fatalf("expected parent to have gender '%v' but got '%v'", expectedParentGender, parentGender)
		}

		mother, err := f.GetPerson(motherName)
		if err != nil {
			t.Fatalf("expected no error while getting person but got error - %v", err)
		}

		children := mother.GetRelations()["Child"]
		numberOfChildren := len(children)
		expectedNumberOfChildren := 1

		if numberOfChildren != expectedNumberOfChildren {
			t.Fatalf("expected mother to have %v children but got %v", expectedNumberOfChildren, numberOfChildren)
		}

		child = children[0]
		childName := child.GetName()
		childGender := child.GetGender()

		if childName != expectedChildName {
			t.Fatalf("expected child to have name '%v' but got '%v'", expectedChildName, childName)
		}

		if childGender != expectedChildGender {
			t.Fatalf("expected child to have gender '%v' but got '%v'", expectedChildGender, childGender)
		}
	})

	t.Run("Mother not found", func(t *testing.T) {
		childName := "Amba"
		childGender := pkg.Female
		motherName := "Queen Anga"
		expectedErrorMessage := "Mother named 'non-existent' not found"

		f := pkg.NewFamilyTree()

		f.AddPerson(motherName, pkg.Female)

		err := f.AddChildToMother("non-existent", childName, childGender)
		if err == nil {
			t.Fatal("expected error while adding child to non existent mother but got no error")
		}

		if err, ok := err.(pkg.PersonNotFoundError); ok {
			if err.Error() != expectedErrorMessage {
				t.Fatalf("expected error message to be '%v' but got '%v'", expectedErrorMessage, err.Error())
			}
		} else {
			t.Fatalf("expected error to be of type PersonNotFoundError but it is of type %T. Error - %v", err, err)
		}
	})

	t.Run("Person not mother", func(t *testing.T) {
		childName := "Amba"
		childGender := pkg.Female
		name := "John"
		expectedErrorMessage := "Person named 'John' is not a mother"

		f := pkg.NewFamilyTree()

		f.AddPerson(name, pkg.Male)

		err := f.AddChildToMother(name, childName, childGender)
		if err == nil {
			t.Fatal("expected error while adding child to a person who is not a mother but got no error")
		}

		if err, ok := err.(pkg.PersonNotMotherError); ok {
			if err.Error() != expectedErrorMessage {
				t.Fatalf("expected error message to be '%v' but got '%v'", expectedErrorMessage, err.Error())
			}
		} else {
			t.Fatalf("expected error to be of type PersonNotMotherError but it is of type %T. Error - %v", err, err)
		}
	})
}

func assertPerson(t *testing.T, person *pkg.Person, expectedName string, expectedGender pkg.Gender, expectedID int) {
	name := person.GetName()
	if name != expectedName {
		t.Fatalf("expected person name to be %v but got %v", expectedName, name)
	}

	gender := person.GetGender()
	if gender != expectedGender {
		t.Fatalf("expected person gender to be %v but got %v", expectedGender, gender)
	}

	id := person.GetID()
	if id != expectedID {
		t.Fatalf("expected person ID to be %v but got %v", expectedID, id)
	}
}
