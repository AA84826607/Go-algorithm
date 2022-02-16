package leecode_test

import (
	"fmt"
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *TreeNode {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	left := root.Left
	left.Left = &TreeNode{Val: 3}
	left.Right = &TreeNode{Val: 4}
	return &root
}

// func show(node *TreeNode,res *TreeNode, root *TreeNode){
//	res.Val=node.Val
//	if node.Left !=nil{
//		res.Right = &TreeNode{}
//		res=res.Right
//		show(node.Left,res,root)
//	}
//
//	if node.Right != nil{
//		res.Right = &TreeNode{}
//		res=res.Right
//		show(node.Right,res,root)
//	}
// }
// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	i := len(nums1)
	if i%2 == 0 {
		return (float64(nums1[i/2]) + float64(nums1[(i/2)+1])) / 2
	}
	return float64(nums1[i/2])
}

// leetcode submit region end(Prohibit modification and deletion)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	num1 := []int{1, 2, 3, 5, 8}
	num2 := []int{2, 4, 6, 7, 10}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// 我进行反着加
	tail := 0
	v := x
	for v != 0 {
		tail = tail*10 + v%10
		v = v / 10
	}
	return tail == x
}

func TestPalindromeNumber(t *testing.T) {
	fmt.Printf("是不是回文:%v", isPalindrome(1010101))
}

func longestPalindrome(s string) string {
	l := len(s)
	begin := 0
	maxlen := 0
	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if maxlen < j-i+1 && Judge_str(s[i:j+1]) {
				maxlen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxlen]
}

// 判断是不是回文串
func Judge_str(s string) bool {
	l := len(s)
	mid := l/2 + 1
	if l%2 == 0 {
		mid = l / 2
	}
	for i := 0; i < mid; i++ {
		for s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func TestLongestPalindromicSubstring(t *testing.T) {
	// fmt.Printf("是不是回文:%v", Judge_str("aafaafaa"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	// 我进行反着加
	tail := 0
	v := x
	for v != 0 {
		tail = tail*10 + v%10
		v = v / 10
	}
	return tail
}

// func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
//	root := NewTree()
//	flatten(root)
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(ins []int) *ListNode {
	root := ListNode{}
	head := &root
	for k, v := range ins {
		head.Val = v
		if k != len(ins)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return &root
}

// leetcode submit region end(Prohibit modification and deletion)

func reserve(h1 *ListNode) *ListNode {
	var prev, cur *ListNode = nil, h1
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func middleNode(head *ListNode) *ListNode {
	// 快慢指针
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l2 := mid.Next
	mid.Next = nil
	rev := reserve(l2)
	mergeList(head, rev)
	showList(head)
}

func showList(l *ListNode) {
	for l != nil {
		fmt.Println("node is ", l.Val)
		l = l.Next
	}
}

func TestReorderList(t *testing.T) {
	ins := []int{1, 3, 4, 5, 6}
	list := NewList(ins)
	reorderList(list)
}

// leetcode submit region begin(Prohibit modification and deletion)
// 暴力破解法
// func maxArea(height []int) int {
// 	max :=0
// 	l := len(height)
// 	for i:=0; i<l-1;i++ {
// 		for j:=i+1;j<l;j ++{
// 			h :=min(height[i],height[j])
// 			if max < h *(j -i){
// 				max = h *(j -i)
// 			}
// 		}
// 	}
// 	return max
// }
// 双指针方法
func maxArea(height []int) int {
	max := 0
	t := len(height) - 1
	h := 0
	for h < t {
		hg, isLeft := min(height[h], height[t])
		if max < sq(t-h, hg) {
			max = sq(t-h, hg)
		}
		if isLeft {
			h++
		} else {
			t++
		}
	}
	return max
}

func sq(d, h int) int {
	return d * h
}

func min(i, j int) (int, bool) {
	if i > j {
		return j, true
	}
	return i, false
}

// leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {
	height := []int{8, 4, 5, 6, 25, 6}
	fmt.Println(maxArea(height))
}

// 回溯算法
// leetcode submit region begin(Prohibit modification and deletion)
var all [][]int

func permute(nums []int) [][]int {
	all = [][]int{}
	used := make([]int, 0, len(nums))
	sort.Ints(nums)
	backTrack(len(nums), nums, []int{}, used, 0)
	return all
}

func backTrack(length int, output []int, path []int, used []int, useNum int) {
	if length == 0 {
		// 要新开辟一个int数组，新的地址 保证all内的数据不会被修改
		p := make([]int, len(path))
		copy(p, path)
		all = append(all, p)
	}
	for i := 0; i < length; i++ {
		if len(used) < useNum+1 || used[useNum] != output[i] {
			curNum := output[i]
			if len(used) < useNum+1 {
				used = append(used, curNum)
			} else {
				used[useNum] = curNum
			}
			path = append(path, curNum)
			output = append(output[0:i], output[i+1:]...)
			backTrack(len(output), output, path, used, useNum+1)
			output = append(output[:i], append([]int{curNum}, output[i:]...)...)
			path = path[:len(path)-1]
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func TestPermutations(t *testing.T) {
	fmt.Println(permute([]int{1, 1, 2}))
}

// 112题
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 展示这个Tree
func ShowTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("node is %d", root.Val)
	ShowTree(root.Left)
	ShowTree(root.Right)
	fmt.Println(".")
}

// [5,4,8,11,null,13,4,7,2,null,null,null,1]
// 使用堆构建二叉树
func NewTreeAccordingHeap(heap []int) *TreeNode {
	root := &TreeNode{}
	CreateTree(root, heap, 0)
	return root
}

func CreateTree(root *TreeNode, heap []int, num int) {
	root.Val = heap[num]
	if len(heap) > 2*num+1 {
		if heap[2*num+1] != -1 {
			root.Left = &TreeNode{}
			CreateTree(root.Left, heap, 2*num+1)
		}
	}
	if len(heap) > 2*num+2 {
		if heap[2*num+2] != -1 {
			root.Right = &TreeNode{}
			CreateTree(root.Right, heap, 2*num+2)
		}
	}
}

var isTrue = false

func hasPathSum(root *TreeNode, targetSum int) bool {
	isTrue = false
	backTrackPath(root, 0, targetSum)
	return isTrue
}

func backTrackPath(root *TreeNode, preSum int, targetSum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == targetSum {
			isTrue = true
		}
		return
	}
	backTrackPath(root.Left, preSum+root.Val, targetSum)
	backTrackPath(root.Right, preSum+root.Val, targetSum)
}

// leetcode submit region end(Prohibit modification and deletion)

func TestPathSum(t *testing.T) {
	// fmt.Println(hasPathSum(NewTree(),6))
	root := NewTreeAccordingHeap([]int{1, 2, 3})
	fmt.Println(hasPathSum(root, 22))
}

// var node []int
// func inorderTraversal(root *TreeNode)[]int {
// 	node =[]int{}
// 	inoTraversal(root)
// 	return node
// }
//
// func inoTraversal(root *TreeNode){
// 	if root == nil{
// 		return
// 	}
// 	inoTraversal(root.Left)
// 	node = append(node, root.Val)
// 	inoTraversal(root.Right)
// }

// 中序遍历
func inorderTraversal(root *TreeNode)(nodes []int) {
	inoTraversal(root, &nodes)
	return nodes
}

func inoTraversal(root *TreeNode, node *[]int)  {
	if root == nil {
		return
	}

	inoTraversal(root.Left, node)
	*node = append(*node, root.Val)
	inoTraversal(root.Right, node)
}

// leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	root := NewTreeAccordingHeap([]int{1, 2, 3})
	fmt.Println(inorderTraversal(root))
}

var nodes []*TreeNode
func generateTrees(n int) []*TreeNode {
	node :=make([]int,0,n)
	for i:=1;i<=n;i++{
		node = append(node, i)
	}
	createTree(node,&TreeNode{})
	return nodes
}

func createTree(node []int,tNode *TreeNode,){
	if len(node) == 0 {
		nodes = append(nodes, tNode)
	}
	for k := range node {
		curNum:=node[k]
		tNode.Val = curNum
		if len(node) ==1 {
			node = []int{}
		}else{
			node = append(node[0:k], node[k+1:]...)
		}
		tNode.Val=curNum
		tNode.Left= &TreeNode{}
		tNode.Right = &TreeNode{}
		createTree(node,tNode.Left)
		createTree(node,tNode.Right)
		node = append(append(node[0:k], curNum), node[k:]...)
	}
}

func TestUniqueBinarySearchTreesIi(t *testing.T){
	n :=generateTrees(2)
	fmt.Println(n)
	fmt.Println("hahaha")
}