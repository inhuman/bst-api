package bst

import (
	"encoding/json"
	"fmt"
	"github.com/inhuman/bst-api/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

var jsonData = `{
  "key": 8,
  "value": "Root",
  "left": {
    "key": 4,
    "value": "4",
    "left": {
      "key": 2,
      "value": "2"
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
	assert.Equal(t, 8, con.Root.Key)
	assert.Equal(t, 4, con.Root.Left.Key)
	assert.Equal(t, 2, con.Root.Left.Left.Key)
	assert.Equal(t, 6, con.Root.Left.Right.Key)
	assert.Equal(t, 10, con.Root.Right.Key)
	assert.Equal(t, 11, con.Root.Right.Right.Key)
	assert.Equal(t, 9, con.Root.Right.Left.Key)
}

func TestNewBstContainerEmptyTree(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer(nil, l)
	assert.NoError(t, err)
	assert.NotNil(t, con)
}

func TestTreeNode_Insert(t *testing.T) {

	jsonData := `{"key": 10}`
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	err = Insert(con.Root, 2, "2")
	assert.NoError(t, err)

	err = Insert(con.Root, 12, "12")
	assert.NoError(t, err)

	err = Insert(con.Root, 11, "11")
	assert.NoError(t, err)

	err = Insert(con.Root, 4, "4")
	assert.NoError(t, err)

	assert.Equal(t, 10, con.Root.Key)
	assert.Equal(t, 2, con.Root.Left.Key)
	assert.Equal(t, 12, con.Root.Right.Key)
	assert.Equal(t, 11, con.Root.Right.Left.Key)
	assert.Equal(t, 4, con.Root.Left.Right.Key)
}

func TestTreeNode_InsertExists(t *testing.T) {

	jsonData := `{"key": 10}`
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	err = Insert(con.Root, 10, "10")
	assert.Error(t, err, "key exists")
}

func TestTreeNode_Find(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	val := Find(con.Root, 11)
	assert.Equal(t, "11", val)

	val2 := Find(con.Root, 9)
	assert.Equal(t, "9", val2)

	val3 := Find(con.Root, 100)
	assert.Equal(t, nil, val3)

	val4 := Find(con.Root, 1)
	assert.Equal(t, nil, val4)
}

func TestTreeNode_DeleteNodeWithTwoChildren(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	Delete(con.Root, 10)

	assert.Equal(t, nil, Find(con.Root, 10))
	assert.Equal(t, 11, con.Root.Right.Key)
	assert.Equal(t, 9, con.Root.Right.Left.Key)
}

func TestTreeNode_DeleteNodeWithOneChild(t *testing.T) {
	l := log.NewLogger()

	con, err := NewBstContainer([]byte(jsonData), l)
	assert.NoError(t, err)

	Delete(con.Root, 6)
	assert.Equal(t, nil, Find(con.Root, 6))
	assert.Equal(t, 7, con.Root.Left.Right.Key)
}

func TestGenerateBinaryTree(t *testing.T) {

	l := log.NewLogger()

	con, err := NewBstContainer(nil, l)
	assert.NoError(t, err)

	num := 100000

	rootKey := num / 2
	rootValue := fmt.Sprintf("%d", rootKey)

	con.Root.Key = rootKey
	con.Root.Value = rootValue

	for i := 1; i < num; i++ {

		n := rand.Intn(num)

		if Find(con.Root, n) == nil {
			err := Insert(con.Root, n, fmt.Sprintf("%d", n))
			assert.NoError(t, err)
		}
	}

	file, err := json.Marshal(con.Root)
	assert.NoError(t, err)

	err = ioutil.WriteFile("generated.json", file, 0644)
	assert.NoError(t, err)

}
