package bst

import (
	"github.com/inhuman/bst-api/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonData = `{
  "key": 8,
  "value": "root",
  "left": {
    "key": 4,
    "value": "4",
    "left": {
      "key": 2,
      "value": "2",
      "left": {
        "key": 1,
        "value": "1"
      }
    },
    "right": {
      "key": 6,
      "value": "6",
      "right": {
        "key": 7,
        "value": "7"
      }
    }
  },
  "right": {
    "key": 10,
    "value": "10",
    "left": {
      "key": 9,
      "value": "9"
    },
    "right": {
      "key": 11,
      "value": "11"
    }
  }
}`

func TestNewBstContainer(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)
	assert.Equal(t, 8, con.Root.GetKey())
	assert.Equal(t, 4, con.Root.GetLeft().GetKey())
	assert.Equal(t, 2, con.Root.GetLeft().GetLeft().GetKey())
	assert.Equal(t, 6, con.Root.GetLeft().GetRight().GetKey())
	assert.Equal(t, 10, con.Root.GetRight().GetKey())
	assert.Equal(t, 11, con.Root.GetRight().GetRight().GetKey())
	assert.Equal(t, 9, con.Root.GetRight().GetLeft().GetKey())
}

func TestNewBstContainerEmptyTree(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer(nil, l)
	assert.NoError(t, err)
	assert.NotNil(t, con)
}

func TestNewBstContainerBadJson(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer([]byte("bad json"), l)
	assert.Error(t, err)
	assert.Equal(t, (*Container)(nil), con)
}

func TestTreeNode_Insert(t *testing.T) {

	jsonData := `{"key": 10}`
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	err = con.Insert(con.Root, 2, "2")
	assert.NoError(t, err)

	err = con.Insert(con.Root, 12, "12")
	assert.NoError(t, err)

	err = con.Insert(con.Root, 11, "11")
	assert.NoError(t, err)

	err = con.Insert(con.Root, 4, "4")
	assert.NoError(t, err)

	assert.Equal(t, 10, con.Root.GetKey())
	assert.Equal(t, 2, con.Root.GetLeft().GetKey())
	assert.Equal(t, 12, con.Root.GetRight().GetKey())
	assert.Equal(t, 11, con.Root.GetRight().GetLeft().GetKey())
	assert.Equal(t, 4, con.Root.GetLeft().GetRight().GetKey())
}

func TestTreeNode_InsertExists(t *testing.T) {

	jsonData := `{"key": 10}`
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	err = con.Insert(con.Root, 10, "10")
	assert.Error(t, err, "key exists")
}

func TestTreeNode_Find(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	val := con.Find(con.Root, 11)
	assert.Equal(t, "11", val)

	val2 := con.Find(con.Root, 9)
	assert.Equal(t, "9", val2)

	val3 := con.Find(con.Root, 100)
	assert.Equal(t, nil, val3)
}

func TestTreeNode_DeleteNodeWithTwoChildren(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	con.Delete(con.Root, 10)

	assert.Equal(t, nil, con.Find(con.Root, 10))
	assert.Equal(t, 11, con.Root.GetRight().GetKey())
	assert.Equal(t, 9, con.Root.GetRight().GetLeft().GetKey())
}

func TestTreeNode_DeleteNodeWithOneRightChild(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	con.Delete(con.Root, 6)
	assert.Equal(t, nil, con.Find(con.Root, 6))
	assert.Equal(t, 7, con.Root.GetLeft().GetRight().GetKey())
}

func TestTreeNode_DeleteNodeWithNOneLeftChild(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	con.Delete(con.Root, 2)
	assert.Equal(t, nil, con.Find(con.Root, 2))
}

func TestTreeNode_DeleteNodeWithNoChild(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	con.Delete(con.Root, 1)
	assert.Equal(t, nil, con.Find(con.Root, 1))
}

func TestTreeNode_DeleteRootNode(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	con.Delete(con.Root, 8)
	assert.Equal(t, nil, con.Find(con.Root, 8))
	assert.Equal(t, 4, con.Root.GetLeft().GetKey())
}
