package pkg

// TODO(karuppiah7890): Check if it's okay to export Gender type. Or we could actually just use the
// constants like Male and Female for Gender and not export Gender and instead use 'gender'. But yeah,
// Person, Gender are closely tied and there's also FamilyTree struct where gender is used

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) IsMale() bool {
	return g == Male
}

func (g Gender) IsFemale() bool {
	return g == Female
}

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown Gender"
	}
}
