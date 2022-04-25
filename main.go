package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func swap(sl []int, i, mini int) {
	k := sl[i]
	sl[i] = sl[mini]
	sl[mini] = k
}

func sort(part []int, wg *sync.WaitGroup) {
	for i := 0; i < len(part); i++ {
		mini := len(part) - 1
		for j := i; j < len(part); j++ {
			if part[j] < part[mini] {
				mini = j
			}
		}
		swap(part, i, mini)
	}
	wg.Done()
}

func merge(parts [][]int) []int {
	var i0, i1 int = 0, 0
	var res []int
	for i0 < len(parts[0]) && i1 < len(parts[1]) {
		if parts[0][i0] <= parts[1][i1] {
			res = append(res, parts[0][i0])
			i0++
		} else {
			res = append(res, parts[1][i1])
			i1++
		}
	}
	for i0 < len(parts[0]) {
		res = append(res, parts[0][i0])
		i0++
	}
	for i1 < len(parts[1]) {
		res = append(res, parts[1][i1])
		i1++
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	var sl []int
	var i int = 0
	for str[i] != '\n' {
		el, _ := strconv.Atoi(string(str[i]))
		sl = append(sl, el)
		i++
	}
	var parts = make([][]int, 4)

	parts[0] = sl[:len(sl)/4]
	parts[1] = sl[len(sl)/4 : len(sl)/2]
	parts[2] = sl[len(sl)/2 : 3*len(sl)/4]
	parts[3] = sl[3*len(sl)/4:]
	var wg sync.WaitGroup
	wg.Add(4)
	for i = 0; i < 4; i++ {
		go sort(parts[i], &wg)
	}
	wg.Wait()
	var res []int
	preRes := make([][]int, 2)
	preRes[0] = merge(parts[:2])
	preRes[1] = merge(parts[2:])
	res = merge(preRes)
	fmt.Println(res)
}
