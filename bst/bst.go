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

	l := logger.With().Str("source", "bst").Logger()
	l.Debug().Str("event", "create new bst container").Send()

	root := TreeNode{}

	if jsonData != nil {
		l.Debug().
			Str("event", "fill bst from json data").
			Send()

		err := json.Unmarshal(jsonData, &root)
		if err != nil {
			return nil, err
		}
	}

	con := &Container{
		Root:   &root,
		Logger: &l,
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

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing insert key: %d, value: %v", key, value)).
		Send()

	switch {

	case n.Key == key:
		con.Logger.Debug().
			Str("event", fmt.Sprintf("key %d exists, return error", key)).
			Send()

		return fmt.Errorf("key exists: %d", key)

	case n.Key > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.Key)).
			Send()

		if n.Left == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("left node not exists, create new one with key: %d, value: %v", key, value)).
				Send()

			n.Left = &TreeNode{
				Key:   key,
				Value: value,
			}
			return nil
		}

		return con.Insert(n.Left, key, value)

	case n.Key < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.Key)).
			Send()

		if n.Right == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("right node not exists, create new one with key: %d, value: %v", key, value)).
				Send()

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

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing search key: %d", key)).
		Send()

	switch {
	case n.Key == key:
		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal current node key, return value: %v", key, n.Value)).
			Send()

		return n.Value

	case n.Key > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.Key)).
			Send()

		if n.Left == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("left node not exists, no record for key %d", key)).
				Send()

			return nil
		}

		return con.Find(n.Left, key)

	case n.Key < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.Key)).
			Send()

		if n.Right == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("right node not exists, no record for key %d", key)).
				Send()

			return nil
		}

		return con.Find(n.Right, key)

	default:
		return errors.New("unreachable statement")
	}
}

func (con *Container) Delete(n *TreeNode, key int) *TreeNode {

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing delete key: %d", key)).
		Send()

	switch {

	case n.Key > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.Key)).
			Send()

		n.Left = con.Delete(n.Left, key)

	case n.Key < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.Key)).
			Send()

		n.Right = con.Delete(n.Right, key)

	case (n.Key == key) && (n.Left != nil) && (n.Right != nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has left and right children", key)).
			Send()

		minNode := con.minNode(n.Right)

		n.Key = minNode.Key
		n.Value = minNode.Value

		n.Right = con.Delete(n.Right, n.Key)

	case (n.Key == key) && (n.Left != nil) && (n.Right == nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has only left child", key)).
			Send()

		n = n.Left

	case (n.Key == key) && (n.Left == nil) && (n.Right != nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has only right child", key)).
			Send()

		n = n.Right

	default:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node is child free", key)).
			Send()

		n = nil
	}

	return n
}

func (con *Container) minNode(n *TreeNode) *TreeNode {
	if n.Left == nil {
		return n
	}

	return con.minNode(n.Left)
}
