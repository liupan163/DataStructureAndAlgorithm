package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//println("result bool is ======>", isAnagram("anagram", "nagaram"))
	//println("result string is ======>", removeDuplicates("bbaacaab"))
	//println("result string is ======>", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//println("myPow string is ======>", myPow(0.00001, 2147483647))
	//println("calculateMinimumHP string is ======>", calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	a := &TreeNode{Val: 27, Left: nil, Right: nil}
	b := &TreeNode{Val: 34, Left: nil, Right: nil}
	c := &TreeNode{Val: 58, Left: nil, Right: nil}
	d := &TreeNode{Val: 50, Left: nil, Right: nil}
	e := &TreeNode{Val: 44, Left: nil, Right: nil}
	a.Right = b
	b.Right = c
	c.Left = d
	d.Left = e
	//println("calculateMinimumHP string is ======>", minDiffInBST(a))
	//ba := &TreeNode{Val: 1, Left: nil, Right: nil}
	//println("binaryTreePaths string is ======>", strings.Join(binaryTreePaths(ba), ""))
	//println("binaryTreePaths string is ======>", len(binaryTreePaths(ba)))
	//println("lengthOfLongestSubstring is ======>", lengthOfLongestSubstring("abcabcbb"))
	println("topKFrequent is ======>", topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	b10 := &TreeNode{Val: 10, Left: nil, Right: nil}
	b5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	b15 := &TreeNode{Val: 15, Left: nil, Right: nil}
	b6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	b3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	b7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	b13 := &TreeNode{Val: 13, Left: nil, Right: nil}
	b18 := &TreeNode{Val: 18, Left: nil, Right: nil}
	b20 := &TreeNode{Val: 20, Left: nil, Right: nil}
	b1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	b10.Left = b5
	b10.Right = b15
	b15.Left = b6
	b15.Right = b20
	println("topKFrequent is ======>", isValidBST(b10))
	b10.Left = b5
	b10.Right = b15
	b5.Left = b3
	b5.Right = b7
	b15.Left = b13
	b15.Right = b18
	b3.Left = b1
	b7.Left = b6
	println("rangeSumBST is ======>", rangeSumBST(b10, 6, 10))
	println("containsNearbyAlmostDuplicate is ======>", containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	println("countSmaller is ======>", countSmaller([]int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41})[3])
	println("findCircleNum is ======>", findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))
	println("letterCasePermutation is ======>", len(letterCasePermutation("a1b2")))
}

func letterCasePermutation(S string) []string {
	var resultList = make([]string, 0)
	for index, _ := range S {
		if len(resultList) == 0 {
			if S[index] >= 'a' && S[index] <= 'z' {
				resultList = append(resultList, S[index])
				resultList = append(resultList, exchangeChar(S[index], false))
			} else if S[index] >= 'A' && S[index] <= 'Z' {
				resultList = append(resultList, S[index])
				resultList = append(resultList, exchangeChar(S[index], false))
			} else {
				resultList = append(resultList, S[index])
			}
			continue
		}
		var newResult = make([]byte, 0)
		for _, item := range resultList {
			if S[index] >= 'a' && S[index] <= 'z' {
				newResult = append(item, S[index])
				newResult = append(item, exchangeChar(S[index], false))
			} else if S[index] >= 'A' && S[index] <= 'Z' {
				newResult = append(newResult, S[index])
				newResult = append(newResult, exchangeChar(S[index], false))
			} else {
				newResult = append(newResult, S[index])
			}
		}
		resultList = newResult
	}
	return resultList
}

func exchangeChar(c byte, isUpper bool) byte {
	if isUpper {
		c = c + 32
	} else {
		c = c - 32
	}
	return c
}

type Trie struct {
	value map[string]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (trie *Trie) Insert(word string) {
	var root = trie
	for _, w := range word {
		if v, ok := root.value[string(w)]; ok {
			root = v
		} else {
			newTrie := make(map[string]*Trie, 1)
			newTrie[string(w)] = nil
			root.value[string(w)] = &Trie{
				value: newTrie,
				isEnd: false,
			}
		}
	}
	root.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	var node = trie
	for _, w := range word {
		if _, ok := node.value[string(w)]; ok {
			node = node.value[string(w)]
			continue
		} else {
			return false
		}
	}
	return node.isEnd
}

func eventualSafeNodes(graph [][]int) []int {
	var resultList = make([]int, 0)
	var finalList = make([]int, 0)
	for index := len(graph) - 1; index >= 0; index-- {
		var tempValue = graph[index]
		if tempValue == nil || len(tempValue) == 0 {
			resultList = append(resultList, index)
			finalList = append(finalList, index)
		} else {
			if len(tempValue) == 1 {
				var value = tempValue[0]
				for _, finalPort := range finalList {
					if value == finalPort {
						resultList = append(resultList, index)
					}
				}
			}
		}
	}
	sort.Ints(resultList)
	return resultList
}

func findCircleNum(M [][]int) int {
	var resultList = make([][]int, 0)
	for x := 0; x < len(M); x++ {
		var rowList = make([]int, 0)
		rowList = append(rowList, x)
		for y := 0; y < len(M[0]); y++ {
			if M[x][y] == 1 {
				rowList = append(rowList, y)
			}
		}
		resultList = append(resultList, rowList)
	}
test:
	for {
		var isFinishLoop = true
	tempLoop:
		for i := 0; i < len(resultList)-1; i++ {
			list1 := resultList[i]
			list2 := resultList[i+1]
			for _, value1 := range list1 {
				for _, value2 := range list2 {
					if value1 == value2 {
						resultList[i] = append(resultList[i], resultList[i+1]...)
						resultList = append(resultList[:i+1], resultList[i+2:]...)
						isFinishLoop = false
						break tempLoop
					}
				}
			}
		}
		if isFinishLoop {
			break test
		}
	}
	return len(resultList)
}

type KthLargest struct {
	capacity int
	data     []int
	len      int
}

//func Constructor(k int, nums []int) KthLargest {
//	var kthLargest = KthLargest{}
//	kthLargest.capacity = k
//	kthLargest.data = make([]int, 0)
//	for _, num := range nums {
//		kthLargest.Add(num)
//	}
//	return kthLargest
//}

func (kth *KthLargest) Add(val int) int {
	if kth.len < kth.capacity {
		kth.data = append(kth.data, val)
		kth.len = kth.len + 1
		return 0
	} else if kth.len == kth.capacity {
		kth.data = append(kth.data, val)
		n := len(kth.data) - 1
		for i := n / 2; i >= 1; i-- {
			heapify(kth.data, i)
		}
	} else {
		if val > kth.data[0] {
			kth.data[0] = val
			heapify(kth.data, 1)
		}
	}
	return kth.data[1]
}

func heapify(data []int, index int) {
	n := len(data) - 1
	for {
		minPos := index
		if index*2 <= n && data[index*2] < data[minPos] {
			minPos = index * 2
		}
		if index*2+1 <= n && data[index*2+1] < data[minPos] {
			minPos = index*2 + 1
		}
		if minPos == index {
			break
		}
		data[minPos], data[index] = data[index], data[minPos]
		index = minPos
	}
}

func numIslands(grid [][]byte) int {
	var resultIslandCount = 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == 1 {
				resultIslandCount = resultIslandCount + 1
				checkAroundPositions(grid, y, x)
			}
		}
	}
	return resultIslandCount
}

func checkAroundPositions(grid [][]byte, y, x int) {
	grid[y][x] = 0
	if x+1 < len(grid[0]) && grid[y][x+1] == 1 {
		checkAroundPositions(grid, y, x+1)
	}
	if y+1 < len(grid) && grid[y+1][x] == 1 {
		checkAroundPositions(grid, y+1, x)
	}
}

func countSmaller(nums []int) []int {
	var recordCountMap = make(map[int]int, 0)
	var recordValueIndex = make(map[int]int, 0)
	var resultList = make([]int, len(nums))
	for index := len(nums) - 1; index >= 0; index-- {
		var value = nums[index]
		if recordCountMap[value] == 0 {
			recordValueIndex[value] = index
		}
		recordCountMap[value] = recordCountMap[value] + 1
		var tempCount = 0
		for v, _ := range recordValueIndex {
			if value > v {
				var recordCount = recordCountMap[v]
				tempCount = tempCount + recordCount
			}
		}
		resultList[index] = tempCount
		fmt.Printf("%v || %v\n", value, resultList)
	}
	fmt.Printf("%v", resultList)
	return resultList
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	var isExist = false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if absValue(nums[j], nums[i]) <= t {
				isExist = true
			}
		}
	}
	return isExist
}

func absValue(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

var resultInt = 0

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val >= L && root.Val <= R {
		resultInt = resultInt + root.Val
	}
	rangeSumBST(root.Left, L, R)
	rangeSumBST(root.Right, L, R)
	return resultInt
}

var resultPaths = make([]string, 0)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var pathRecord = strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		resultPaths = append(resultPaths, pathRecord)
	}
	if root.Left != nil {
		iterateTreePath(pathRecord, root.Left)
	}
	if root.Right != nil {
		iterateTreePath(pathRecord, root.Right)
	}
	return resultPaths
}
func iterateTreePath(pathRecord string, node *TreeNode) {
	if node == nil {
		return
	}
	pathRecord = pathRecord + "->" + strconv.Itoa(node.Val)
	if node.Left != nil {
		iterateTreePath(pathRecord, node.Left)
	}
	if node.Right != nil {
		iterateTreePath(pathRecord, node.Right)
	}
	if node.Left == nil && node.Right == nil {
		resultPaths = append(resultPaths, pathRecord)
		return
	}
}

var resultValue []int

func calculateMinimumHP(dungeon [][]int) int {
	//var capacity = len(dungeon) * len(dungeon[0])
	resultValue = make([]int, 0)
	var minInitialValue = 0.0
	var accumulateValue = 0
	if 0 == len(dungeon) || 0 == len(dungeon[0]) {
		minInitialValue = 1
		return int(minInitialValue)
	}
	x := 0
	y := 0
	doIterate(dungeon, minInitialValue, accumulateValue, x, y)
	sort.Ints(resultValue)
	//for index, value := range resultValue {
	//	println("index =>", index, "||value==>", value)
	//}
	//println("index =>", resultValue[0])
	//println("minValue =>", 1-resultValue[0])
	return resultValue[0]
}

func doIterate(dungeon [][]int, minInitialValue float64, accumulateValue, x, y int) {
	if x == len(dungeon[0])-1 && y == len(dungeon)-1 {
		accumulateValue = accumulateValue + dungeon[x][y]
		if accumulateValue <= 0 {
			minInitialValue = math.Max(minInitialValue, float64(1-accumulateValue))
			if minInitialValue == float64(1-accumulateValue) {
				resultValue = append(resultValue, int(minInitialValue))
			}
		}
		return
	}
	if x == len(dungeon[0])-1 {
		for index := y; index < len(dungeon); index++ {
			accumulateValue = accumulateValue + dungeon[x][index]
			if accumulateValue <= 0 {
				minInitialValue = math.Max(minInitialValue, float64(1-accumulateValue))
				if minInitialValue == float64(1-accumulateValue) {
					resultValue = append(resultValue, int(minInitialValue))
				}
			}
		}
		return
	}
	if y == len(dungeon)-1 {
		for index := x; index < len(dungeon[0]); index++ {
			accumulateValue = accumulateValue + dungeon[x][index]
			if accumulateValue <= 0 {
				minInitialValue = math.Max(minInitialValue, float64(1-accumulateValue))
				if minInitialValue == float64(1-accumulateValue) {
					resultValue = append(resultValue, int(minInitialValue))
				}
			}
		}
		return
	}
	if x != len(dungeon[0])-1 && y != len(dungeon)-1 {
		var value = dungeon[x][y]
		accumulateValue = accumulateValue + value
		if accumulateValue <= 0 {
			minInitialValue = math.Max(minInitialValue, float64(1-accumulateValue))
			if minInitialValue == float64(1-accumulateValue) {
				resultValue = append(resultValue, int(minInitialValue))
			}
		}
		doIterate(dungeon, minInitialValue, accumulateValue, x, y+1)
		doIterate(dungeon, minInitialValue, accumulateValue, x+1, y)
	}
}

var minValue = 1<<32 - 1

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		var result = root.Val - root.Left.Val
		minValue = compareMinValue(result, minValue)
		minFunc(root.Left)
	}
	if root.Right != nil {
		var result = root.Right.Val - root.Val
		minValue = compareMinValue(result, minValue)
		minFunc(root.Right)
	}
	return minValue
}
func minFunc(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		var result = node.Val - node.Left.Val
		minValue = compareMinValue(result, minValue)
		minFunc(node.Left)
	}
	if node.Right != nil {
		var result = node.Right.Val - node.Val
		minValue = compareMinValue(result, minValue)
		minFunc(node.Right)
	}
}
func compareMinValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	if n > 0 {
		return iterate(x, n)
	} else {
		return 1 / iterate(x, int(math.Abs(float64(n))))
	}
}

func iterate(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	var tempValue = iterate(x, n/2)
	if n%2 == 1 {
		return tempValue * tempValue * x
	} else {
		return tempValue * tempValue
	}
}

func groupAnagrams(strs []string) [][]string {
	var strLength = len(strs)
	var resultMap = make(map[string][]string, strLength)
	for _, value := range strs {
		var tempValue = strings.Split(value, "")
		sort.Strings(tempValue)
		println("add value  ===>" + value)
		if _, ok := resultMap[strings.Join(tempValue, "")]; !ok {
			resultMap[strings.Join(tempValue, "")] = []string{value}
		} else {
			resultMap[strings.Join(tempValue, "")] = append(resultMap[strings.Join(tempValue, "")], value)
		}
	}
	var resultList = make([][]string, 0)
	for _, value := range resultMap {
		if value == nil {
			println("resultMap value is nil===>")
			continue
		}
		println(len(value))
		var rowList = make([]string, 0)
		for _, wordValue := range value {
			if wordValue == "" {
				continue
			}
			rowList = append(rowList, wordValue)
			println("wordValue===>" + wordValue + " || rowList===>" + strings.Join(rowList, ""))
		}
		resultList = append(resultList, rowList)
	}
	println(len(resultList))
	return resultList
}

func removeDuplicates(S string) string {
	var tempString = strings.Split(S, "")
	var isFinish = false
	for {
		isFinish = true
		for index, value := range tempString {
			if index+1 <= len(tempString)-1 && value == tempString[index+1] {
				if index-1 < 0 { //	0下标位，0和1位
					tempString = tempString[index+2:]
				} else {
					tempString = append(tempString[:index], tempString[index+2:]...)
				}
				isFinish = false
				break
			}
			println("tempString is :" + strings.Join(tempString, ""))
		}
		if isFinish {
			break
		}
	}
	return strings.Join(tempString, "")
}

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	var symbolMap = make(map[int]int32, len(s))
	for index, value := range t {
		symbolMap[index] = value
	}
	for _, value := range s {
		for key, mapValue := range symbolMap {
			if mapValue == value {
				delete(symbolMap, key)
				break
			}
		}
	}
	if len(symbolMap) > 0 {
		return false
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var sSlice = strings.Split(s, "")
	var maxLength = 0
	if len(sSlice) == 1 {
		return 1
	}
	for index, value := range sSlice {
		var recordMap = make(map[string]string, 0)
		recordMap[value] = value
		for i := index + 1; i < len(sSlice); i++ {
			if _, ok := recordMap[sSlice[i]]; ok {
				var tempMaxLength = i - index
				maxLength = compareMaxValue(maxLength, tempMaxLength)
				break
			} else {
				recordMap[sSlice[i]] = sSlice[i]
				var tempMaxLength = i - index + 1
				maxLength = compareMaxValue(maxLength, tempMaxLength)
			}
		}
	}
	return maxLength
}
func compareMaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topKFrequent(words []string, k int) []string {
	var recordMap = make(map[string]int, 0)
	var resultList = make([]string, 0)
	for _, value := range words {
		if _, ok := recordMap[value]; ok {
			recordMap[value] = recordMap[value] + 1
		} else {
			recordMap[value] = 1
		}
	}
	var reverseMap = make(map[int]string, 0)
	var countList = make([]int, 0)
	for index, value := range recordMap {
		reverseMap[value] = index
		countList = append(countList, value)
	}
	//sort.Reverse(sort.IntSlice(countList))
	sort.Sort(sort.Reverse(sort.IntSlice(countList)))
	for i := 0; i < k; i++ {
		resultList = append(resultList, reverseMap[countList[i]])
	}
	return resultList
}

func isValidBST(root *TreeNode) bool {
	var isTrue = false
	if root == nil {
		return true
	}
	isTrue = verifyLeftTree(root.Left, root.Val, root.Val) && verifyRightTree(root.Right, root.Val, root.Val)
	return isTrue
}

func verifyLeftTree(root *TreeNode, upper, limitValue int) bool {
	var isTrue = true
	if root == nil {
		return true
	}
	if root.Val > limitValue || root.Val > upper {
		isTrue = false
		return isTrue
	}
	if root.Left != nil && root.Val < root.Left.Val {
		isTrue = false
		return isTrue
	}
	if root.Right != nil && root.Val > root.Right.Val {
		isTrue = false
		return isTrue
	}
	if root.Left != nil {
		isTrue = isTrue && verifyLeftTree(root.Left, root.Val, limitValue)
	}
	if root.Right != nil {
		isTrue = isTrue && verifyRightTree(root.Right, root.Val, limitValue)
	}
	return isTrue
}
func verifyRightTree(root *TreeNode, upper, limitValue int) bool {
	var isTrue = true
	if root == nil {
		return true
	}
	if root.Val < limitValue || root.Val < upper {
		isTrue = false
		return false
	}
	if root.Left != nil && root.Val < root.Left.Val {
		isTrue = false
		return isTrue
	}
	if root.Right != nil && root.Val > root.Right.Val {
		isTrue = false
		return isTrue
	}
	if root.Left != nil {
		isTrue = isTrue && verifyLeftTree(root.Left, root.Val, limitValue)
	}
	if root.Right != nil {
		isTrue = isTrue && verifyRightTree(root.Right, root.Val, limitValue)
	}
	isTrue = verifyLeftTree(root.Left, root.Val, limitValue)
	isTrue = verifyRightTree(root.Right, root.Val, limitValue)
	return isTrue
}
