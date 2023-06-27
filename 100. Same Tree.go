package leetcode

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func isSameTree(p *TreeNode1, q *TreeNode1) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if (*p).Val == (*q).Val {
		if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
			if (*p).Val == (*q).Val {
				return true
			}
		}
	}

	return false
}
