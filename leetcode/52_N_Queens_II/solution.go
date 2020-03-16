package main

import (
    "fmt"
    "sync"
    )

func goDeep(target, n int, record *[]int, resultChan chan int) {
  if target == n {
    resultChan <- 1
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

func generateResult(resultChan chan int, n, seed int, wait *sync.WaitGroup) {
  defer wait.Done()
  record := make([]int, n, n)
  record[0] = seed
  goDeep(1, n, &record, resultChan)
}

func solveNQueen(n int) int {
result := 0
  resultChan := make(chan int)
  var wait sync.WaitGroup

  wait.Add(n)
  for i := 0;i < n;i ++ {
    go generateResult(resultChan, n, i, &wait)
  }

  go func() {
    wait.Wait()
    close(resultChan)
  }();

  for delta := range resultChan {
    result += delta
  }

  return result
}

func main() {
  result := solveNQueen(8)
  fmt.Println(result)
}
