package bst

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
)

type Container struct {
	Root   *TreeNode
	Logger *zerolog.Logger
}

func NewBstContainer(jsonData []byte, logger *zerolog.Logger) (*Container, error) {

	root := TreeNode{}

	if jsonData != nil {
		err := json.Unmarshal(jsonData, &root)
		if err != nil {
			return nil, err
		}
	}

	con := &Container{
		Root:   &root,
		Logger: logger,
	}

	return con, nil
}

type TreeNode struct {
	Key   int         `json:"key"`
	Value interface{} `json:"value"`
	Left  *TreeNode   `json:"left"`
	Right *TreeNode   `json:"right"`
}

func (con *Container) Insert(n *TreeNode, key int, value interface{}) error {

	switch {

	case n.Key == key:
		return fmt.Errorf("key exists: %d", key)

	case n.Key > key:
		if n.Left == nil {
			n.Left = &TreeNode{
				Key:   key,
				Value: value,
			}
			return nil
		}

		return con.Insert(n.Left, key, value)

	case n.Key < key:
		if n.Right == nil {
			n.Right = &TreeNode{
				Key:   key,
				Value: value,
			}
			return nil
		}

		return con.Insert(n.Right, key, value)

	default:
		return errors.New("unreachable statement")
	}
}

func (con *Container) Find(n *TreeNode, key int) interface{} {

	switch {
	case n.Key == key:
		return n.Value

	case n.Key > key:
		if n.Left == nil {
			return nil
		}

		return con.Find(n.Left, key)

	case n.Key < key:
		if n.Right == nil {
			return nil
		}

		return con.Find(n.Right, key)

	default:
		return errors.New("unreachable statement")
	}
}

func (con *Container) Delete(node *TreeNode, key int) *TreeNode {

	switch {

	case node == nil:
		return nil

	case node.Key > key:
		node.Left = con.Delete(node.Left, key)

	case node.Key < key:
		node.Right = con.Delete(node.Right, key)

	case (node.Left != nil) && (node.Right != nil):
		minNode := con.minNode(node.Right)

		node.Key = minNode.Key
		node.Value = minNode.Value

		node.Right = con.Delete(node.Right, node.Key)

	case (node.Left != nil) && (node.Right == nil):
		node = node.Left

	case (node.Left == nil) && (node.Right != nil):
		node = node.Right

	default:
		node = nil
	}

	return node
}

func (con *Container) minNode(n *TreeNode) *TreeNode {
	if n.Left == nil {
		return n
	}

	return con.minNode(n.Left)
}
