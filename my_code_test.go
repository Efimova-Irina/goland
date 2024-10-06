package main


import "testing"

func Test_Sum(t *testing.T) {
 tests := []struct {
  name string
  a    int
  b    int
  want int
 }{
  {"test 1", 0, 0, 0},
  {"test 2", 0, 1, 1},
  //    и тд
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   if got := sum(tt.a, tt.b); got != tt.want {
    t.Errorf("Sum() = %v, want %v", got, tt.want)
   }
  })
 }
}