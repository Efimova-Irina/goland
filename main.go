package main


import "fmt"


func main() {
    var a, b int
    var oper string
    fmt.Println("Введите первое число:")
    fmt.Scan(&a)
    fmt.Println("Выберете операцию: +, -, *, / ")
    fmt.Scan(&oper)
    fmt.Println("Введите второе число:")
    fmt.Scan(&b)

    switch oper {
        case "+": 
            sum(a, b)
        case "*":
            mult(a, b)  
        case "-":
            minus(a, b)  
        case "/":
            div(a, b)
        default:
            znak(oper)            
    
    }
    
}


// функция сложения
func sum(a int, b int) { 
    fmt.Println(a, "+", b, "=", a + b)
}

// функция сложения
func mult(a int, b int) { 
    fmt.Println(a, "*", b, "=", a * b)
}
// функция вычитания
func minus(a int, b int) { 
    fmt.Println(a, "-", b, "=", a - b)
}

// функция деления
func div(a int, b int) { 
   if b == 0 {
    fmt.Println("Делить на ноль нельзя!")
   } else {
    fmt.Printf("%x / %x = %f", a, b, a / b)
    }
}

// функция определения нужной операции
func znak(oper string) {
    fmt.Println("Читать не умеешь!!! Нет такой операции!")
}

