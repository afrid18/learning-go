package iteration

const repeatCountConst = 5

func Repeat(c string, repeatCount int) string {
  if repeatCount == 0 {
    repeatCount = repeatCountConst
  }
  result := ""
  for i := 0; i < repeatCount; i++ {
    result += c
  }
  return result
}
