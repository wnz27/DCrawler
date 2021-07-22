/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 23:25 7月
 **/
package parser

import (
	"f27dc/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])  // 整个copy， m 共享一个地址，每层循环复用这个m 使用解析函数不是立即，
		// 而是后面引擎run的时候，那时候的跑的m都是一样的
		result.Items = append(result.Items, "User: " + string(m[2]))
		// todo 最简单的办法在这里直接把人的信息拿出来
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				// 函数式编程，闭包的使用
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile2(c, name)
				},
			})
	}
	return result
}



