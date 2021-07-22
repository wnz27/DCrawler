/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/23 00:14 7月
 **/
package parser

import (
	"f27dc/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "安静的雪")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile)
	profile.Name = "安静的雪"

	expected := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		Xingzuo:     "牡羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		House:      "已购房",
		Hukou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
	}
	if profile != expected {
		t.Errorf("expected %v; but got %v", expected, profile)
	}
}

func TestParseProfile2(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data2.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile2(contents, "安静的雪", "女")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 element; but was %v", result.Items)
	}
	profile := result.Items[0].(model.Profile2)
	profile.Name = "安静的雪"

	expected := model.Profile2{
		Age:        "31岁",
		Height:     "168cm",
		Weight:     "65kg",
		Income:     "月收入:3-5千",
		Gender:     "女",
		Name:       "安静的雪",
		Xingzuo:     "天秤座(09.23-10.22)",
		Occupation: "生产领班",
		Marriage:   "未婚",
		Hukou:      "工作地:阿坝小金",
		Education:  "中专",
	}
	if profile != expected {
		t.Errorf("expected %v; but got %v", expected, profile)
	}
}
