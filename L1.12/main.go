package main

import "fmt"

func main(){

    // Создаем множество строк
    stringSet := make(map[string]struct{})

    // Добавляем строки в множество
    for _, v := range []string{"cat", "cat", "dog", "cat", "tree"} {
        stringSet[v] = struct{}{}
    }
    
    for v := range stringSet {
        fmt.Println(v)
    }
}
