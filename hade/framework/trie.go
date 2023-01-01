package framework

import (
	"errors"
	"log"
	"strings"
)

// 代表树结构

type Tree struct {
	root *node // 根节点
}

type node struct {
	isLast  bool
	segment string            // uri中的字符串
	handler ControllerHandler // 控制器
	childs  []*node           // 这个节点下的子节点
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))

	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {
	segements := strings.SplitN(uri, "/", 2)

	segement := segements[0]
	log.Println(segement)
	if !isWildSegment(segement) {
		segement = strings.ToUpper(segement)
	}

	cnodes := n.filterChildNodes(segement)

	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	if len(segements) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}

		return nil
	}

	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segements[1])
		if tnMatch != nil {

			return tnMatch

		}
	}

	return nil

}

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	log.Println(handler)
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segement := range segments {

		if !isWildSegment(segement) {
			segement = strings.ToUpper(segement)
		}

		isLast := index == len(segments)-1

		var objNode *node
		childNodes := n.filterChildNodes(segement)

		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segement {
					objNode = cnode
					break
				}

			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segement
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}

			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}

	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)

	if matchNode == nil {
		return nil
	}

	return matchNode.handler
}
