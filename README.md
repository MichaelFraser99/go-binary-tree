# Go Binary Tree
This package contains a simple binary tree implementation in Go

## Usage
The package uses the following interface allowing for custom, complex, types to be used as the value of the tree nodes
```go
type C interface {
	New() C // Creates a new instance of the defined implementation
	GetContents() *Contents
	Compare(newValue any) bool // Should return true for greater than, and false for less than
	Equals(value any) bool
}
```

## Example Implementation
The below code is a sample implementation of the C interface for an integer value
```go
type IntComparison struct {
	Contents
}

func (c *IntComparison) Compare(newValue any) bool {
	return c.Value.(int) < newValue.(int)
}

func (c *IntComparison) Equals(newValue any) bool {
	return c.Value.(int) == newValue.(int)
}

func (c *IntComparison) New() C {
	return &IntComparison{}
}
```
The above implementation can then be used as follows:
```go
tree := New(1, &IntComparison{})
tree.Add(2)
tree.Add(2)
tree.Add(0)
ascList := tree.AscList()
// Yields: [0 1 2 2]

descList := tree.DescList()
// Yields: [2 2 1 0]
```

The package also enables the implementation of more complex types as so:
```go
type ComplexObject struct {
    id          uuid.UUID
    priority    int
    createdDate time.Time
    name        string
    description string
}

func (c ComplexObject) String() string {
    return fmt.Sprintf("{id: %s, priority: %d, createdDate: %d, name: %s, description: %s}\n", c.id.String(), c.priority, c.createdDate.Unix(), c.name, c.description)
}

/*
   *
   Given the scenario where we have a complex object with the following ordering rules:
   - an object can have a priority which defines the order in the list. The lower priority is deemed greater than the higher priority
   - if two objects have the same priority, the one that was created earlier is placed first

   each object also has a UUID, so two objects are only equal if they have the same UUID
*/
func (c *ComplexObjectComparison) Compare(newValue any) bool {
    if c.Value.(ComplexObject).priority == newValue.(ComplexObject).priority {
        // Was the new object created before the old one? If yes, its greater
        return newValue.(ComplexObject).createdDate.Before(c.Value.(ComplexObject).createdDate)
    } else {
        return newValue.(ComplexObject).priority < c.Value.(ComplexObject).priority
    }
}

func (c *ComplexObjectComparison) Equals(newValue any) bool {
    return c.Value.(ComplexObject).priority == newValue.(ComplexObject).priority && c.Value.(ComplexObject).createdDate.Equal(newValue.(ComplexObject).createdDate)
}

func (c *ComplexObjectComparison) New() C {
    return &ComplexObjectComparison{}
}
```
Which can then be used as follows:
```go
func example() {
	
}
objP5N1 := ComplexObject{
    id:          uuid.New(),
    priority:    5,
    createdDate: time.Now(),
    name:        "Object priority 5, 1",
    description: "This is the first object with priority 5",
}

tree := New(objP5N1, &ComplexObjectComparison{})

objP3N1 := ComplexObject{
    id:          uuid.New(),
    priority:    3,
    createdDate: time.Now(),
    name:        "Object priority 3, 1",
    description: "This is the first object with priority 3",
}

objP3N2 := ComplexObject{
    id:          uuid.New(),
    priority:    3,
    createdDate: time.Now(),
    name:        "Object priority 3, 2",
    description: "This is the second object with priority 3",
}

objP7N1 := ComplexObject{
    id:          uuid.New(),
    priority:    7,
    createdDate: time.Now(),
    name:        "Object priority 7, 1",
    description: "This is the first object with priority 7",
}

tree.Add(objP3N1)
tree.Add(objP3N2)
tree.Add(objP7N1)

ascList := tree.AscList()
fmt.Println(ascList)
/* Yields:
[
    {id: 95ac322c-b69a-4644-946f-2af41abd0520, priority: 7, createdDate: 1687096388, name: Object priority 7, 1, description: This is the first object with priority 7}
    {id: b88ebee3-4d5c-46f0-bee6-c5784b112e2b, priority: 5, createdDate: 1687096388, name: Object priority 5, 1, description: This is the first object with priority 5}
    {id: 713d0fcc-728a-4078-a143-0e330bd057f8, priority: 3, createdDate: 1687096388, name: Object priority 3, 2, description: This is the second object with priority 3}
    {id: eecd289c-d866-41f4-adf7-62fd80f7a18c, priority: 3, createdDate: 1687096388, name: Object priority 3, 1, description: This is the first object with priority 3}
    {id: eecd289c-d866-41f4-adf7-62fd80f7a18c, priority: 3, createdDate: 1687096388, name: Object priority 3, 1, description: This is the first object with priority 3}
]
*/
```
## Functionality
Once you have your tree constructed, you can use the following functions to interact with it:
```go
tree.Add(value) // Adds a value to the tree
tree.Remove(value) // Removes one instance of value from the tree. Returns error if this would result in an empty tree
tree.AscList() // Returns a list of the values in the tree in ascending order
tree.DescList() // Returns a list of the values in the tree in descending order
tree.Find(value) // Returns the specified value if the tree contains it, otherwise nil