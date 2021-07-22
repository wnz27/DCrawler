/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 02:37 7月
 **/
package parser

import (
	"f27dc/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {

		result.Items = append(result.Items, "City: " + string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
			})
	}
	return result
}

