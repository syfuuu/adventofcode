package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a, b int) int {
  if a < b {
    return b - a
  }
  return a - b
}

func calcDiff(l1, l2 []int) int {
  sort.Ints(l1) 
  sort.Ints(l2) 

  res := 0
  for i := 0; i < len(l1); i++ {
    res += abs(l1[i], l2[i])
  }
  
  return res
}

func calcSimarity(l1, l2 []int) int {
  var res int
  m1, m2 := make(map[int]int), make(map[int]int)
  for i := 0; i < len(l1); i++ {
    m1[l1[i]] += 1
    m2[l2[i]] += 1
  }

  for f1, f2 := range m1 {
    res += (f1 * m2[f1]) * f2
  }

  return res
}

func main() {
  // read input
  data, err := os.Open("input.txt") // or ioutil.ReadFile for older versions
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
  defer data.Close()

  l1, l2 := make([]int, 0, 1000), make([]int, 0, 1000)
  
  scanner := bufio.NewScanner(data)
  for scanner.Scan() {
    line := strings.Fields(scanner.Text())
    n1, err := strconv.Atoi(line[0])
    if err != nil {
      fmt.Println("Error converting string to int:", err)
      return
    }

    n2, err := strconv.Atoi(line[1])
    if err != nil {
      fmt.Println("Error converting string to int:", err)
      return
    }

    l1 = append(l1, n1)
    l2 = append(l2, n2)
  }


  // fmt.Println(calcDiff(l1, l2))

  fmt.Println(calcSimarity(l1, l2))
}
