package main

/*
int sum(int a, int b) {
  return a+b;
}
*/
import "C"
// C代码和引用类库不能有空格
func main() {
	println(C.sum(1, 1))
}
