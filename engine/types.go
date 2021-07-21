/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 02:38 7月
 **/
package engine

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
