package pkg

type Relationship struct {
	Name                string
	ReverseRelationship *Relationship
}

var Parent = Relationship{Name: "Parent"}
var Child = Relationship{Name: "Child"}
var Sibling = Relationship{Name: "Sibling"}
var Spouse = Relationship{Name: "Spouse"}

func init() {
	Parent.ReverseRelationship = &Child
	Child.ReverseRelationship = &Parent
	Sibling.ReverseRelationship = &Sibling
	Spouse.ReverseRelationship = &Spouse
}

// TODO(karuppiah7890): Do we need Sibling relationship? As in, however as of now - when we add people, we just add children
// and also add spouse during family tree initialization. But other than that, siblings and other relationships can be found out
// in other ways. Like, to find sibling - if parent and child relationships are correctly known, then that's enough!

// TODO(karuppiah7890): Find a way to define relationships which can be derived from existing relationships.
// For example, Paternal Uncle - Find father, and then find their siblings who are male. We probably need a
// function here, which is pretty flexible. Anything else would be more rigid - like a set of instructions in
// string etc like "go to father", "find siblings" etc. Instead using full on code is better. So, we have to
// define how to find a person in a relationship using functions or something like that. Something like
// Relationship Finder or just Get Relationship according to the problem but that name doesn't make exact
// sense. It could be called 'Find People in this Relationship' or 'Find People who have this Relationship with this Person',
// pretty long names, haha
