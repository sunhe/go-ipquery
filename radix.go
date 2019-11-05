package ipquery

import (
	"errors"
)

type radixTreeNode struct {
	left  *radixTreeNode
	right *radixTreeNode
	//parent *radixTreeNode
	value interface{}
}

func radixTreeCreate() *radixTreeNode {
	return new(radixTreeNode)
}

func (rt *radixTreeNode) insert(ip []byte, mask []byte, len int, value interface{}) error {
	node := rt
	for i := 0; i < len; i++ {
		for bit := byte(0x80); bit != 0 && bit&mask[i] != 0; bit >>= 1 {
			if bit&ip[i] == 0 {
				if node.left == nil {
					node.left = new(radixTreeNode)
				}
				node = node.left
			} else {
				if node.right == nil {
					node.right = new(radixTreeNode)
				}
				node = node.right
			}
		}
	}

	if node.value != nil {
		return errors.New("There is inflict!")
	}

	node.value = value
	return nil
}

func (rt *radixTreeNode) query(ip []byte, len int) interface{} {
	var value interface{}

	node := rt
	for i := 0; i < len; i++ {
		for bit := byte(0x80); bit != 0; bit >>= 1 {
			if bit&ip[i] == 0 {
				node = node.left
			} else {
				node = node.right
			}
			if node != nil {
				if node.value != nil {
					value = node.value
				}
			} else {
				goto found
			}
		}
	}

found:
	return value
}

func (rt *radixTreeNode) destroy() {
	rt.left = nil
	rt.right = nil
	rt.value = nil
}
