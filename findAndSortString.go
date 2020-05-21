package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
  new_str := []string{}
  for _, v := range array1 {
    for _, k := range array2 {
      if strings.Contains(k, v) {
        if FindString(new_str, v) {
          continue
        }
        new_str = append(new_str, v)
        sort.Strings(new_str)
      }
    }
  }
  return new_str
}

func FindString(a []string, x string) bool {
        for _, n := range a {
                if x == n {
                        return true
                }
        }
        return false
}
