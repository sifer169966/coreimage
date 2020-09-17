package playground

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func AppendData(_p1, _p2 string) {
	var base [][]string
	var condition [][]string
	pNumber := strings.Split(_p1, ",")
	nNumber := strings.Split(_p2, ",")
	pDataSets := assignNumber(pNumber)
	nDataSets := assignNumber(nNumber)
	for i := 0; i < len(pDataSets); i++ {
		base = append(base, convertToString(pDataSets[i]))
	}
	mergestring := prepareData(base)
	fmt.Print("BASE: ")
	lists := convertToNumber(mergestring)
	sort.Ints(lists)
	fmt.Println(lists)

	for i := 0; i < len(nDataSets); i++ {
		condition = append(condition, convertToString(nDataSets[i]))
	}
	mergeCond := prepareData(condition)
	condlists := convertToNumber(mergeCond)
	fmt.Print("CONDITION: ")
	sort.Ints(condlists)
	fmt.Println(condlists)

	fmt.Print("AFTER REMOVED: ")
	removedData := removeDuplicateData(lists, condlists)
	sort.Ints(removedData)
	fmt.Print("LASTDATA: ")
	fmt.Println(removedData)
	// for i := 0; i < len(nDataSets); i++ {
	// 	condition = append(condition, nDataSets[i]...)
	// }

	// fmt.Print("CONDITION: ")
	// fmt.Println(condition)
	// r := removeDuplicateData(base, condition)
	// fmt.Print("Last: ")
	// fmt.Println(r)

	// for i := 0; i < len(buffer1); i++ {
	// 	output = append(output, convertToString(buffer1[i]))
	// }
	// // fmt.Println(pNumber[0])
	// // fmt.Println(pNumber[1])
	// fmt.Print("Output is: ")
	// for i := 0; i < len(output); i++ {
	// 	oo = append(oo, "["+output[i][0]+"-"+output[i][len(output[i])-1]+"]")
	// }
	// fmt.Print("condition is: ")
	// fmt.Println(condition)
	// fmt.Println(oo)
}

func prepareData(ss [][]string) []string {
	var r []string
	if len(ss) == 1 {
		return ss[0]
	}
	for i := 0; i < len(ss); i++ {
		if i != len(ss)-1 {
			buffer := make([][]string, 0)
			buffer = append(buffer, appendCategory(ss[i], ss[i+1]))
			for _, v := range buffer[0] {
				r = append(r, v)
			}
		} else {
			fmt.Println("ELSE")
		}
	}
	return r
}
func convertToNumber(s []string) []int {
	var buffer []int
	for i := 0; i < len(s); i++ {
		if n, err := strconv.Atoi(s[i]); err == nil {
			buffer = append(buffer, n)
		} else {
			fmt.Println(s, "is not an integer.")
		}
	}
	return buffer
}
func appendCategory(a []string, b []string) []string {

	check := make(map[string]int)
	d := append(a, b...)
	res := make([]string, 0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}

	return res
}
func removeDuplicateData(v []int, cond []int) []int {
	var buffer []int
	for i := 0; i < len(v); i++ {
		for _, val := range cond {
			if v[i] == val {
				v[i] = v[len(v)-1]
				v = v[:len(v)-1]
				buffer = v
				fmt.Print("BUFFER: ")
				fmt.Println(buffer)
				sort.Ints(buffer)
			}
		}
	}
	return buffer
}
func removeSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
func convertToString(n []int) []string {
	buffer := make([]string, 0)
	for _, v := range n {
		buffer = append(buffer, strconv.Itoa(v))
	}
	return buffer
}
func assignNumber(sets []string) [][]int {
	var newArray [][]int
	for i := 0; i < len(sets); i++ {
		newArrayBuffer := make([]int, 0)
		// sFirst := parseToInteger(deleteSpecialChar(s[i]))
		// sLast := parseToInteger(deleteSpecialChar(s[len(s)-1]))
		splited := strings.Split(sets[i], "-")
		first := make([]int, 0)
		last := parseToInteger(deleteSpecialChar(splited[len(splited)-1]))
		splited = splited[:len(splited)-1]
		fmt.Print("Splited is: ")
		fmt.Println(splited)
		first = append(first, parseToInteger(deleteSpecialChar(splited[0])))
		for j := 0; j < len(first); j++ {
			if first[j]+1 < last {
				first = append(first, first[j]+1)
			}
		}
		fmt.Print("First is: ")
		fmt.Println(first)
		for _, v := range first {
			newArrayBuffer = append(newArrayBuffer, v)
		}
		newArrayBuffer = append(newArrayBuffer, last)
		newArray = append(newArray, newArrayBuffer)
		newArrayBuffer = newArrayBuffer[:0]
		fmt.Println(newArray)
		//newArray = append(newArray, newArrayBuffer)
	}
	fmt.Print("New array is: ")
	fmt.Println(newArray)
	return newArray
}
func parseToInteger(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(s, "is not an integer.")
	}
	return n
}
func deleteSpecialChar(s string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(s, "")
}
