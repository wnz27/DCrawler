/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 01:08 7月
 **/
package main

import (
	"f27dc/engine"
	"f27dc/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// debug
	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/1396998601")
	//if err != nil {
	//	panic(err)
	//}
	//res := parser.ParseProfile2(contents, "TEST")
	//fmt.Printf("-->%s", res.Items)

}

