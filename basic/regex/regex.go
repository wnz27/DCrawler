/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 01:39 7月
 **/
package main

import (
	"fmt"
	"reflect"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com@abc.com
email1 is eew@def.org
email2 is kkk@qq.com
email3 is ddd@amc.com.cn
`

func main() {
	//re, err := regexp.Compile("ccmouse@gmail.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindString(text)  // 只会匹配第一个
	fmt.Println(match)

	match2 := re.FindAllString(text, -1)  // -1 找所有匹配
	fmt.Println(match2, reflect.TypeOf(match2))

	match3 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match3, reflect.TypeOf(match3))
}


