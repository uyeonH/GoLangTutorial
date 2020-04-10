package main

import (
  "fmt"
)

func main() {
  
  ch:=make(chan bool)
  nums:=[2]string{"one","two"}

  for _, n :=range nums{
    go isReady(n,ch)
  }

  fmt.Println(<-ch)
  fmt.Println(<-ch)

  //deadlock
  //fmt.Println(<-ch)

  //result:=<-ch
  //fmt.Println(result)

}

  func isReady(nums string,ch chan bool){
    fmt.Println(nums)
    ch<-true
  }