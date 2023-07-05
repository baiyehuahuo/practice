package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由
	part     string  // 路由中的一部分
	children []*node // 子结点
	isWild   bool    // 是否精确匹配  true 为模糊匹配 false 为精确匹配
}

// 匹配第一个符合要求的子结点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 匹配所有符合要求的子结点
func (n *node) matchChildren(part string) []*node {
	var ans []*node
	for _, child := range n.children {
		if child.part == part || child.isWild {
			ans = append(ans, child)
		}
	}
	return ans
}

// 插入一个模式
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern // 插入到叶结点 将 pattern 写入节点中
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} // 拓展，继续向下插入
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 搜索模式
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" { // 未匹配到叶结点
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1) // 对符合要求的子结点继续搜索，DFS 深度优先搜索
		if result != nil {
			return result
		}
	}

	return nil
}
