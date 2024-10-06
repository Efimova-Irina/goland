package main


import "fmt"


func main() {
    a, b, oper := vparam()
    var res int
    var err error

    switch oper {
        case "+": 
            res = sum(a, b)
        case "*":
           res = mult(a, b)  
        case "-":
            res = minus(a, b)  
        case "/":
            res, err = div(a, b)
            if err != nil {
                fmt.Printf("error: %v", err)
            }
        default:
            znak(oper)            
    
    }
    fmt.Print(a, oper, b, "=", res)   
}

// функция ввода данных
func vparam() (int, int, string) {
    var a, b int
    var oper string

    fmt.Print("Введите первое число: ")
    fmt.Scanln(&a)

    fmt.Print("Введите второе число: ")
    fmt.Scanln(&b)

    fmt.Print("Введите оператор (+, -, *, /): ")
    fmt.Scanln(&oper)

    return a, b, oper
}


// функция сложения
func sum(a int, b int) int { 
    return a + b
}

// функция умножения
func mult(a int, b int) int { 
    return a * b
}
// функция вычитания
func minus(a int, b int) int { 
    return a - b
}

// функция деления
func div(a int, b int) (int, error) {
    if b == 0 {
       return 0, fmt.Errorf("Делить на ноль нельзя!")
     }
   
     return a / b, nil
}

// функция определения нужной операции
func znak(oper string) error {
    return fmt.Errorf("Читать не умеешь!!! Нет такой операции!")
}

