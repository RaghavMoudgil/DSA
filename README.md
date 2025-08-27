# DSA
in this we are going to solve common DSA questions that interviewer asks in the interview.
Find pairs with a given sum

Find two elements in an array that sum up to a target value.

go
func findPair(arr []int, target int) (int, int, bool) {
    m := make(map[int]int)
    for _, v := range arr {
        if idx, ok := m[target - v]; ok {
            return idx, v, true
        }
        m[v] = v
    }
    return 0, 0, false
}
Maximum subarray sum (Kadane’s Algorithm)

Find the contiguous subarray with the maximum sum.

go
func maxSubArray(nums []int) int {
    maxSoFar, maxEndingHere := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        maxEndingHere = max(nums[i], maxEndingHere + nums[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    return maxSoFar
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
Strings
Reverse a string

Convert string to runes and swap elements to handle Unicode.

go
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
Check if a string is a palindrome

go
func isPalindrome(s string) bool {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
Linked Lists
Reverse a linked list

go
type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
Detect cycle in linked list (Floyd’s cycle detection)

go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
Trees
Traversal - Inorder

go
func inorder(root *TreeNode) []int {
    var res []int
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }
        dfs(node.Left)
        res = append(res, node.Val)
        dfs(node.Right)
    }
    dfs(root)
    return res
}
Validate binary search tree

go
func isValidBST(root *TreeNode) bool {
    var prev *int
    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil { return true }
        if !inorder(node.Left) { return false }
        if prev != nil && node.Val <= *prev { return false }
        prev = &node.Val
        return inorder(node.Right)
    }
    return inorder(root)
}
Graphs
Breadth-First Search (BFS)

go
func bfs(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    queue := []int{start}
    var result []int
    visited[start] = true

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        result = append(result, node)

        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
    return result
}
Stacks and Queues
Valid parentheses check

go
func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if c == '(' || c == '[' || c == '{' {
            stack = append(stack, c)
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
Dynamic Programming
Fibonacci with Memoization

go
func fib(n int) int {
    memo := make([]int, n+1)
    for i := range memo { memo[i] = -1 }
    var dp func(k int) int
    dp = func(k int) int {
        if k == 0 || k == 1 { return k }
        if memo[k] != -1 { return memo[k] }
        memo[k] = dp(k-1) + dp(k-2)
        return memo[k]
    }
    return dp(n)
}
--