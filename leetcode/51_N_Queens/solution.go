package main

import (
    "fmt"
    "sync"
    "strings"
    )

func num2string(l, n int) string {
  var sb strings.Builder
  for i := 0;i < n;i ++ {
    if l == i {
      sb.WriteString("Q")
    } else {
      sb.WriteString(".")
    }
  }
  return sb.String()
}

func goDeep(target, n int, record *[]int, resultChan chan []string) {
  if target == n {
    result := make([]string, n, n)
    for index, val := range *record {
      result[index] = num2string(val, n)
    }
    resultChan <- result
    return
  }
  for i := 0;i < n;i ++ {
    avaiable := true
    for rindex, rval := range *record {
      if rindex >= target {
        break
      } else if rval == i ||  rval - i == target - rindex || i - rval == target-rindex {
        avaiable = false;
        break;
      }
    }
    if avaiable {
      (*record)[target] = i
      goDeep(target + 1, n, record, resultChan)
    }
  }
}

func generateResult(resultChan chan []string, n, seed int, wait *sync.WaitGroup) {
  defer wait.Done()
  record := make([]int, n, n)
  record[0] = seed
  goDeep(1, n, &record, resultChan)
}

func solveNQueen(n int) [][]string {
  var result [][]string
  resultChan := make(chan []string)
  var wait sync.WaitGroup

  wait.Add(n)
  for i := 0;i < n;i ++ {
    go generateResult(resultChan, n, i, &wait)
  }

  go func() {
    wait.Wait()
    close(resultChan)
  }();

  for line := range resultChan {
    result = append(result, line)
  }

  return result
}

func main() {
  result := solveNQueen(8)
  fmt.Println("Solution: ", len(result))
  for _, graph := range result {
    for _, line := range graph {
      fmt.Println(line)
    }
    fmt.Println()
  }
}
