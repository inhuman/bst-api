package bst

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/inhuman/bst-api/interfaces"
	"github.com/rs/zerolog"
)

type TreeNode struct {
	Key   int                 `json:"key"`
	Value interface{}         `json:"value"`
	Left  interfaces.TreeNode `json:"left"`
	Right interfaces.TreeNode `json:"right"`
}

func (n *TreeNode) GetRight() interfaces.TreeNode {
	return n.Right
}

func (n *TreeNode) SetRight(right interfaces.TreeNode) {
	n.Right = right
}

func (n *TreeNode) GetLeft() interfaces.TreeNode {
	return n.Left
}

func (n *TreeNode) SetLeft(left interfaces.TreeNode) {
	n.Left = left
}

func (n *TreeNode) GetValue() interface{} {
	return n.Value
}

func (n *TreeNode) SetValue(value interface{}) {
	n.Value = value
}

func (n *TreeNode) GetKey() int {
	return n.Key
}

func (n *TreeNode) SetKey(key int) {
	n.Key = key
}

func (n *TreeNode) UnmarshalJSON(data []byte) error {

	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if m["left"] != nil {

		nl := TreeNode{}

		if err := json.Unmarshal(m["left"], &nl); err != nil {
			return err
		}

		n.Left = &nl
	}

	if m["right"] != nil {
		nr := TreeNode{}
		if err := json.Unmarshal(m["right"], &nr); err != nil {
			return err
		}

		n.Right = &nr
	}

	if m["key"] != nil {
		if err := json.Unmarshal(m["key"], &n.Key); err != nil {
			return err
		}
	}

	if m["value"] != nil {
		if err := json.Unmarshal(m["value"], &n.Value); err != nil {
			return err
		}
	}

	return nil
}

type Container struct {
	Root   interfaces.TreeNode
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

func (con *Container) GetRoot() interfaces.TreeNode {
	return con.Root
}

func (con *Container) Insert(n interfaces.TreeNode, key int, value interface{}) error {

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing insert key: %d, value: %v", key, value)).
		Send()

	switch {

	case n.GetKey() == key:
		con.Logger.Debug().
			Str("event", fmt.Sprintf("key %d exists, return error", key)).
			Send()

		return fmt.Errorf("key exists: %d", key)

	case n.GetKey() > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.GetKey())).
			Send()

		if n.GetLeft() == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("left node not exists, create new one with key: %d, value: %v", key, value)).
				Send()

			n.SetLeft(&TreeNode{
				Key:   key,
				Value: value,
			})
			return nil
		}

		return con.Insert(n.GetLeft(), key, value)

	case n.GetKey() < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.GetKey())).
			Send()

		if n.GetRight() == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("right node not exists, create new one with key: %d, value: %v", key, value)).
				Send()

			n.SetRight(&TreeNode{
				Key:   key,
				Value: value,
			})
			return nil
		}

		return con.Insert(n.GetRight(), key, value)

	default:
		return errors.New("unreachable statement")
	}
}

func (con *Container) Find(n interfaces.TreeNode, key int) interface{} {

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing search key: %d", key)).
		Send()

	switch {
	case n.GetKey() == key:
		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal current node key, return value: %v", key, n.GetValue())).
			Send()

		return n.GetValue()

	case n.GetKey() > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.GetKey())).
			Send()

		if n.GetLeft() == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("left node not exists, no record for key %d", key)).
				Send()

			return nil
		}

		return con.Find(n.GetLeft(), key)

	case n.GetKey() < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.GetKey())).
			Send()

		if n.GetRight() == nil {

			con.Logger.Debug().
				Str("event",
					fmt.Sprintf("right node not exists, no record for key %d", key)).
				Send()

			return nil
		}

		return con.Find(n.GetRight(), key)

	default:
		return errors.New("unreachable statement")
	}
}

func (con *Container) Delete(n interfaces.TreeNode, key int) interfaces.TreeNode {

	con.Logger.Debug().
		Str("event", fmt.Sprintf("processing delete key: %d", key)).
		Send()

	switch {

	case n.GetKey() > key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d smaller than current node key %d, processing left node", key, n.GetKey())).
			Send()

		n.SetLeft(con.Delete(n.GetLeft(), key))

	case n.GetKey() < key:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d bigger than current node key %d, processing right node", key, n.GetKey())).
			Send()

		n.SetRight(con.Delete(n.GetRight(), key))

	case (n.GetKey() == key) && (n.GetLeft() != nil) && (n.GetRight() != nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has left and right children", key)).
			Send()

		minNode := con.minNode(n.GetRight())

		n.SetKey(minNode.GetKey())
		n.SetValue(minNode.GetValue())

		n.SetRight(con.Delete(n.GetRight(), n.GetKey()))

	case (n.GetKey() == key) && (n.GetLeft() != nil) && (n.GetRight() == nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has only left child", key)).
			Send()

		n = n.GetLeft()

	case (n.GetKey() == key) && (n.GetLeft() == nil) && (n.GetRight() != nil):

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node has only right child", key)).
			Send()

		n = n.GetRight()

	default:

		con.Logger.Debug().
			Str("event",
				fmt.Sprintf("key %d equal node key, node is child free", key)).
			Send()

		n = nil
	}

	return n
}

func (con *Container) minNode(n interfaces.TreeNode) interfaces.TreeNode {
	if n.GetLeft() == nil {
		return n
	}

	return con.minNode(n.GetLeft())
}
