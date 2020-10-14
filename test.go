package main

import (
	"math"
	"sort"
	"strings"
)

func main() {
	//println("result bool is ======>", isAnagram("anagram", "nagaram"))
	//println("result string is ======>", removeDuplicates("bbaacaab"))
	//println("result string is ======>", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	println("myPow string is ======>", myPow(0.00001, 2147483647))
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
