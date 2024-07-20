package main

import "fmt"

func main() {
    table := []int{1, 2, 3}
    ac := 93
    FoldInt(Add, table, ac)
    FoldInt(Mul, table, ac)
    FoldInt(Sub, table, ac)
    fmt.Println()

    table = []int{0}
    FoldInt(Add, table, ac)
    FoldInt(Mul, table, ac)
    FoldInt(Sub, table, ac)
}
func Add(n1 , n2 int) int {
    return n1 + n2
}
func Mul(n1 , n2 int) int {
    return n1 * n2
}
func Sub(n1 , n2 int) int {
    return n1 - n2
}
func FoldInt(f func(int, int) int, a []int, n int) {
    result := n
    for _, v := range a {
        result = f(result, v)
    }
    fmt.Println(result)
}