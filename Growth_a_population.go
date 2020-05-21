package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
  fp := p0 + int(float64(p0) * (percent/100)) + aug
  if fp >= p {
    return 1
  } else {
    return 1 + NbYear(fp, percent, aug, p)
  }
}
