package main

import (
	"fmt"
)

func main() {
	fmt.Println(ValidBraces("[({})](]"))
}

func stringInSlice(a string, list []string) (bool, int) { //返回判断结果，也返回所在的index,入参定义是【】string不定长slice，调用时传入也应该是不定长，见21，22
	for pos, b := range list {
		if b == a {
			return true, pos
		}
	}
	return false, 0
}

func ValidBraces(str string) bool {
	/*
		  可以考虑用map字典来存储brace的对应关系
		  var m = map[string]string{
		    "{": "}",
		    "(": ")",
		    "[": "]",
		}
	*/
	open_list := []string{"(", "[", "{"}
	close_list := []string{")", "]", "}"}
	var stack []string      //定义一个空的不定长slice
	for _, i := range str { //遍历string的每一个字符，type为rune，之后要string(i)再把每个字符转换为string
		if c, _ := stringInSlice(string(i), open_list); c == true { //在if条件判断里加short declaration复制变量，然后再根据变量值做判断
			fmt.Println(string(i))
			stack = append(stack, string(i)) //用append([]string, string)追加入栈
		} else if _, f := stringInSlice(string(i), close_list); f >= 0 {
			if len(stack) > 0 && open_list[f] == stack[len(stack)-1] {
				fmt.Println(string(i))
				stack = stack[:len(stack)-1] //删除stack数组最后一个元素，也就是栈的后入先出
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

/*
注意if else的写法，和slice和数组的声明方式
*/
