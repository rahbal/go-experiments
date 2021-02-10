package main

import "fmt"

var s = []uint64{0, 1}

func fib(x uint64) uint64 {
  if x < uint64(len(s)) {
     return s[x]
  }
  zz := fib(x-1) + fib(x-2)
  s = append(s, zz)
  return s[x]
}

func ProductFib(prod uint64) [3]uint64 {
  
  var p uint64 = 0
  var x uint64 = 2
  for p < prod {
    x++
    p = fib(x)*fib(x-1)
    // p = fib(x)
    fmt.Println(p, x)
  }
  if p == prod {
    return [3]uint64{s[x-1], s[x], 1}
  } else {
    return [3]uint64{s[x-1], s[x], 0}
  }
  
}


func main() {

  fmt.Println(ProductFib(4895))
  
}

