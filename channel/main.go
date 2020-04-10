package main

import (
  "fmt"
  "time"
)

func main() {
  
  ch:=make(chan string)
  nums:=[5]string{"one","two","three","four","five"}

  for _, n :=range nums{
    go isReady(n,ch)
  }

  for i:=0; i<len(nums);i++{
    fmt.Println(<-ch)

  }
  

}

  func isReady(num string,ch chan string){
    time.Sleep(time.Second*3)
    ch<-num+" clear"
  }
