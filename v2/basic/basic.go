package basic

func Sum(vals ...int) int {
  res := 0;
  for _, v := range vals {
    res += v
  }
  return res
}
