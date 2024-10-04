package main


import "fmt"


func main() {
    var a, b int
    var oper string
    fmt.Println("Введите первое число:")
    fmt.Scan(&a)
    fmt.Println("Выберете операцию:")
    fmt.Scan(&oper)
    fmt.Println("Введите второе число:")
    fmt.Scan(&b)

    if oper == "+" { 
        sum(a, b)
    }
    
    if oper == "*"{
        mult(a, b)
    }
}

// функция сложения
func sum(a int, b int) int { 
    return a + b
}

// функция сложения
func mult(a int, b int) int { 
    return a * b
}
