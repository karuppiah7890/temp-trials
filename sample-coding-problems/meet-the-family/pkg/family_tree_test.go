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
