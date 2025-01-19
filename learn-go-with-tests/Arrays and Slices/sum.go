package main

func Sum(arr []int) int{
  total := 0;
  for _, number := range arr {
    total += number
  }
  return total
}
