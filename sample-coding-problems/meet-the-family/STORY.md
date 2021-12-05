
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


