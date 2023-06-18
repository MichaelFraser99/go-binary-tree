package ordered_tree

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func TestOrderedNode_AddInt(t *testing.T) {
	tree := New(1, &IntComparison{})
	require.NotNil(t, tree)

	require.Nil(t, tree.Add(2))
	assert.Equal(t, 1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Contents.GetContents().Count)

	require.Nil(t, tree.Add(2))
	assert.Equal(t, 1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Count)

	require.Nil(t, tree.Add(0))
	assert.Equal(t, 1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Count)
	assert.Equal(t, 0, tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Left.Contents.GetContents().Count)

	ascList := tree.AscList()
	fmt.Println(ascList)
	require.NotNil(t, ascList)
	assert.Len(t, ascList, 4)
	assert.Equal(t, 0, ascList[0])
	assert.Equal(t, 1, ascList[1])
	assert.Equal(t, 2, ascList[2])
	assert.Equal(t, 2, ascList[3])

	descList := tree.DescList()
	fmt.Println(descList)
	require.NotNil(t, descList)
	assert.Len(t, descList, 4)
	assert.Equal(t, 2, descList[0])
	assert.Equal(t, 2, descList[1])
	assert.Equal(t, 1, descList[2])
	assert.Equal(t, 0, descList[3])

}

type StringComparison struct {
	Contents
}

func (c *StringComparison) Compare(newValue any) bool {
	return c.Value.(string) < newValue.(string)
}

func (c *StringComparison) Equals(newValue any) bool {
	return c.Value.(string) == newValue.(string)
}

func (c *StringComparison) New() C {
	return &StringComparison{}
}

func TestOrderedNode_AddString(t *testing.T) {
	tree := New("m", &StringComparison{})
	require.NotNil(t, tree)

	assert.NotNil(t, tree.Add(2)) //cannot add a miss-matched type

	require.Nil(t, tree.Add("c"))
	assert.Equal(t, "m", tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, "c", tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Left.Contents.GetContents().Count)

	require.Nil(t, tree.Add("c"))
	assert.Equal(t, "m", tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, "c", tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Left.Contents.GetContents().Count)

	require.Nil(t, tree.Add("z"))
	assert.Equal(t, "m", tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, "c", tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Left.Contents.GetContents().Count)
	assert.Equal(t, "z", tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Contents.GetContents().Count)

	require.Nil(t, tree.Add("d"))
	assert.Equal(t, "m", tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, "c", tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Left.Contents.GetContents().Count)
	assert.Equal(t, "z", tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Contents.GetContents().Count)
	assert.Equal(t, "d", tree.Left.Right.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Left.Right.Contents.GetContents().Count)

	ascList := tree.AscList()
	require.NotNil(t, ascList)
	fmt.Println(ascList)
	assert.Len(t, ascList, 5)
	assert.Equal(t, "c", ascList[0])
	assert.Equal(t, "c", ascList[1])
	assert.Equal(t, "d", ascList[2])
	assert.Equal(t, "m", ascList[3])
	assert.Equal(t, "z", ascList[4])

	descList := tree.DescList()
	require.NotNil(t, descList)
	fmt.Println(descList)
	assert.Len(t, descList, 5)
	assert.Equal(t, "z", descList[0])
	assert.Equal(t, "m", descList[1])
	assert.Equal(t, "d", descList[2])
	assert.Equal(t, "c", descList[3])
	assert.Equal(t, "c", descList[4])
}

type ComplexObjectComparison struct {
	Contents
}

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

func TestOrderedNode_AddComplexType(t *testing.T) {
	objP5N1 := ComplexObject{
		id:          uuid.New(),
		priority:    5,
		createdDate: time.Now(),
		name:        "Object priority 5, 1",
		description: "This is the first object with priority 5",
	}

	tree := New(objP5N1, &ComplexObjectComparison{})
	require.NotNil(t, tree)

	assert.NotNil(t, tree.Add(2))   //cannot add a miss-matched type
	assert.NotNil(t, tree.Add("a")) //cannot add a miss-matched type

	objP3N1 := ComplexObject{
		id:          uuid.New(),
		priority:    3,
		createdDate: time.Now(),
		name:        "Object priority 3, 1",
		description: "This is the first object with priority 3",
	}

	require.Nil(t, tree.Add(objP3N1))
	assert.Equal(t, objP5N1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, objP3N1, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Contents.GetContents().Count)

	require.Nil(t, tree.Add(objP3N1))
	assert.Equal(t, objP5N1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, objP3N1, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Count)

	objP3N2 := ComplexObject{
		id:          uuid.New(),
		priority:    3,
		createdDate: time.Now(),
		name:        "Object priority 3, 2",
		description: "This is the second object with priority 3",
	}

	require.Nil(t, tree.Add(objP3N2))
	assert.Equal(t, objP5N1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, objP3N1, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Count)
	assert.Equal(t, objP3N2, tree.Right.Left.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Left.Contents.GetContents().Count)

	objP7N1 := ComplexObject{
		id:          uuid.New(),
		priority:    7,
		createdDate: time.Now(),
		name:        "Object priority 7, 1",
		description: "This is the first object with priority 7",
	}

	require.Nil(t, tree.Add(objP7N1))
	assert.Equal(t, objP5N1, tree.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Contents.GetContents().Count)
	assert.Equal(t, objP3N1, tree.Right.Contents.GetContents().Value)
	assert.Equal(t, 2, tree.Right.Contents.GetContents().Count)
	assert.Equal(t, objP3N2, tree.Right.Left.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Right.Left.Contents.GetContents().Count)
	assert.Equal(t, objP7N1, tree.Left.Contents.GetContents().Value)
	assert.Equal(t, 1, tree.Left.Contents.GetContents().Count)

	ascList := tree.AscList()
	require.NotNil(t, ascList)
	fmt.Println(ascList)
	assert.Len(t, ascList, 5)
	assert.Equal(t, objP7N1, ascList[0])
	assert.Equal(t, objP5N1, ascList[1])
	assert.Equal(t, objP3N2, ascList[2])
	assert.Equal(t, objP3N1, ascList[3])
	assert.Equal(t, objP3N1, ascList[4])

	descList := tree.DescList()
	require.NotNil(t, descList)
	fmt.Println(descList)
	assert.Len(t, descList, 5)
	assert.Equal(t, objP3N1, descList[0])
	assert.Equal(t, objP3N1, descList[1])
	assert.Equal(t, objP3N2, descList[2])
	assert.Equal(t, objP5N1, descList[3])
	assert.Equal(t, objP7N1, descList[4])
}
