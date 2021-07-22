/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 23:55 7月
 **/
package model

type Profile3 struct {
	Name     string
	Gender   string
	Age      int
	Marriage string
	Hukou    string // 居住地
	Income   string // 月收入
	Height   int
}

// 页面修改后的解析办法
type Profile2 struct {
	Name     string
	//Gender   string
	Age      string
	Marriage string
	Hukou    string // 工作地
	Income   string // 月收入
	Height   string

	Xingzuo    string
	Weight     string
	Education  string // 学历
	Occupation string // 职业

}

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hukou      string
	Xingzuo    string
	House      string
	Car        string
}
