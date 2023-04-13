package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("Hello, World!"));
}

func reverse(str string) string {
  reverseStr:= "";
  originStr := strings.Split(str, "");
  for i := len(str) -1; i >= 0; i--{
    reverseStr += originStr[i];
  }
  return reverseStr
}
