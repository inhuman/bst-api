package interfaces

type Container interface {
	Insert(n TreeNode, key int, value interface{}) error
	Find(n TreeNode, key int) interface{}
	Delete(n TreeNode, key int) TreeNode
	GetRoot() TreeNode
}
type TreeNode interface {
	GetRight() TreeNode
	SetRight(right TreeNode)
	GetLeft() TreeNode
	SetLeft(left TreeNode)
	GetValue() interface{}
	SetValue(value interface{})
	GetKey() int
	SetKey(key int)
}
