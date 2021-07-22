/**
 * @project DCrawler
 * @Author 27
 * @Description //TODO
 * @Date 2021/7/22 23:43 7月
 **/
package parser

import (
	"f27dc/engine"
	"f27dc/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(.*album\.zhenai\.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`.*album\.zhenai\.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hukou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xingzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
const originStr = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`

// 匹配的顺序为  未婚 31岁  天秤座(09.23-10.22) 168cm 65kg 工作地:阿坝小金 月收入:3-5千 生产领班 中专
// 婚姻 年龄 星座 身高 体重 工作地 收入 职业 学历
// 0  1     2    3    4   5    6    7   8

func ParseProfile2(contents []byte, name string, gender string) engine.ParseResult {
	re := regexp.MustCompile(originStr)
	matches := re.FindAllSubmatch(contents, -1)
	profile := model.Profile2{}
	profile.Name = name
	profile.Gender = gender
	l := len(matches)
	for i := 0; i < l; i++ {
		m := matches[i]
		if i == 0 {
			profile.Marriage = string(m[1])
		}
		if i == 1 {
			profile.Age = string(m[1])
		}
		if i == 2 {
			profile.Xingzuo = string(m[1])
		}
		if i == 3 {
			profile.Height = string(m[1])
		}
		if i == 4 {
			profile.Weight = string(m[1])
		}
		if i == 5 {
			profile.Hukou = string(m[1])
		}
		if i == 6 {
			profile.Income = string(m[1])
		}
		if i == 7 {
			profile.Occupation = string(m[1])
		}
		if i == 8 {
			profile.Education = string(m[1])
		}

	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
