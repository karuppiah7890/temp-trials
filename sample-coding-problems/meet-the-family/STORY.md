Person

- has a name
- has a gender, for the sake of the problem, only male and female
- has relationships with other people

Relationship

- has a name
- there are two people involved in any relationship

Different kinds of relationships

- Parent
  - Mother
  - Father
- Child
  - Son
  - Daughter
- Siblings
  - Brother
  - Sister
- Spouse
  - Husband
  - Wife
- Paternal-Uncle
- Maternal-Uncle
- Paternal-Aunt
- Maternal-Aunt
- Sister-In-Law
- Brother-In-Law

Paternal = Father's family tree side

Maternal = Mother's family tree side

Valid Commands

- ADD_CHILD
  - Format - ADD_CHILD "Mother’s-Name" "Child's-Name" "Gender"
  - Example - ADD_CHILD Chitra Aria Female
  - Output Examples
    - CHILD_ADDITION_SUCCEEDED
    - PERSON_NOT_FOUND
      - if given mother name is not found in family tree
    - CHILD_ADDITION_FAILED
      - if given mother name is not a female but a male
      - what if there's error in gender - assuming valid values are only male and female? probably it won't happen
- GET_RELATIONSHIP
  - Format - GET_RELATIONSHIP "Name" “Relationship"
  - Example - GET_RELATIONSHIP Aria Siblings
  - Output Examples
    - Jnki Ahit
    - PERSON_NOT_FOUND
      - if given name is not found in family tree
    - What if relation name is unknown? Hmm, probably it won't happen

---

https://github.com/geektrust/coding-problem-artefacts/tree/master/Go

https://github.com/geektrust/coding-problem-artefacts/blob/master/sample-io/Family

https://github.com/geektrust/coding-problem-artefacts/blob/master/sample-io/Family/BD-PS1/output2.txt

---

[TODO] [Level-1]

- Initialize the existing family tree
- Take test file path (location) as input from the command line argument
- Read the contents of the test file from the test file path
- Parse the test file contents to get the various commands and command inputs / arguments
- Pass the commands and command inputs to the family tree engine to execute them and get results and/ errors
- Get the results and/ errors and print them in proper format

[TODO] [Level-2]

- Think about data structures and models / entities to use

---

Modelling -

Person has 0 or more relations. Each relation is with a person and has a relation name / relationship name as it is a relationship

Person can be a structure / struct (golang), with a name which is a string

We can use pointers to

Modelling relationships -

We can use map / list (array, slice) / struct or a combination of them

Relations can be a map with key as relationship name like child, spouse etc and with value as list of people who correspond to that relationship

If say Sarah has two children, it would look like (pseudo code)

```go
var Ray, Kay, Bay Person = Person{}, Person{}, Person{}

Sarah := Person{
    Name: "Sarah",
    Relations: map[string]*Person{
        "Child": [ &Ray, &Kay ],
        "Spouse": [ &Bay ],
    }
}
```

`Child` can also be `Children`, which makes sense if there's always more than one or just as a common relationship name even when there's only one child

Other ways to model it could be like

```go
var Ray, Kay, Bay Person = Person{}, Person{}, Person{}

Sarah := Person{
    Name: "Sarah",
    Relations: []Relation{
        { Relationship: "Child", Person: &Ray },
        { Relationship: "Child", Person: &Kay },
        { Relationship: "Spouse", Person: &Bay },
    }
}
```

```go
var Ray, Kay Person = Person{}, Person{}

Sarah := Person{
    Name: "Sarah",
    Relations: []Relation{
        { Relationship: "Child", Persons: [ &Ray, &Kay ] },
        { Relationship: "Spouse", Persons: [ &Bay ] },
    }
}
```

I'm planning to go ahead with the first idea for now, as I don't see much difference among the different ideas except for possible inefficiencies in traversing the relations when using list of relations instead of map and slightly extra memory used for storing data when storing same relationship for example child across multiple structs when there are multiple people in that relationship, like having two children and hence two structs for it in the second example

When we add a relation, to help with traversal for processing later, we would need to add double sided relationship. For example, when adding a child, we need to put a relationship from mother to child saying "child is this person", and in child too we need to add a relationship saying "parent is this person", so that we can traverse relationships both ways very easily. For example, to get paternal aunt of a person, we can look at their parents, and look at the male parent(s), the father(s) that is, and look at their sibilings who are female which we can find directly if there's a direct relationship on the person for siblings or go to the parent of the person (the father) and then look at all their children and then find the female(s) and then get their names

To generalize things, I'm planning to avoid any fields name `child`, `spouse`, `parent` etc and instead use them in field value, say in relationship name field

There are many relationships among the family. We could store all the relationships in each person's structure and then use it to find the answers, or derive the relationship between two people on the go based on existing basic relationships like - child, parent, spouse, and then everything else can be easily derived from that set of relationships

There's processing that needs to be done to find people in relationship, based on basic relationships. For example we just thought about how to find paternal aunt by doing some traversing in the family tree. Every relationship would need some traversal. We have some predefined set of relationships in the problem so we will define traversals / actions / processing only for those relationships. And when extra relationships are needed, those could be easily added to the list, and existing ones can be modified or reused too as much as possible. It's best to write loosely coupled code as much as possible, to keep things extensible, but with some basics concrete

We have an extra thing to take care of -

"Given a ‘name’ and a ‘relationship’, you should output the people corresponding to the relationship in the order in
which they were added to the family tree. Assume the names of the family members are unique."

Since we need ordering, I think we could have a number at a global level, that could be used to assign a unique number to every person as their ID. The global number could be incremented (by 1) every time a person is added so that each person gets a unique number and since it's incremented (by 1) every time, the person who gets added later will have a bigger number, so when printing the output, we will print people in ascending order of their person ID

At the top level / global level, there would be a family tree struct, which contains a list of all the people in the family or maybe just the people at the root of the family tree - the king shan and queen anga and also the global number used for assigning unique person ID

Since we can assume that the family members names are unique, the family tree could look like this

```go
FamilyTree {
    FamilyCount: 2,
    Family: map[string]*Person{
        "Ankit": &Person{Name: "Ankit", ID: 1, Relations: ...},
        "Parag": &Person{Name: "Parag", ID: 2, Relations: ...},
    }
}
```

Family Count (total folks in family) can be used for the unique number generation, or some other similar field name could be used

In the above case, we store all the people in the family in the map, which makes it easy to quickly find a person in the family. The key is the name assuming it's unique, or else we could use person ID but that won't be useful when finding people with name which is important for our operations. But we still need unique ID for sorting in the output

The name in the struct is important though it's duplicate from key name. It's important because when traversing the family tree, we will get pointers to person struct, and that's just the value in the map, and we would need people's name in a lot of cases, so we need it in the struct as we can't find key (person name) from value, that's not helpful but instead tedious and unnecessary and is absurd
