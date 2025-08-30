# DSA
in this we are going to solve common DSA questions that interviewer asks in the interview.

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