package main


import "testing"

func Test_Div(t *testing.T) {
 tests := []struct {
  name string
  a    int
  b    int
  want int
 }{
  {"test 1", 1, 0, 0}, // примеры входных данных для проверки
  {"test 2", 0, 1, 0},
  {"test 3", -3, 2, -1},
  {"test 4", -6, -3, 2},
  {"test 5", 2, 3, 0},
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   if got, _ := div(tt.a, tt.b); got != tt.want { // проверка операции, при делении выводится 2 аргумента, при др операциях 1 аргумент
    t.Errorf("Sum() = %v, want %v", got, tt.want)
   }
  })
 }
}