package parser

import (
	"example.com/go-http-demo/crawler/engine"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)cm</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile = model.Profile{}
	profile.Name = name
	// 获取姓名

	// 获取年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	// 获取婚况
	profile.Marriage = extractString(contents, marriageRe)

	// 获取身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = height
	}

	// 获取收入
	profile.Income = extractString(contents, incomeRe)

	// 获取体重
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		profile.Weight = weight
	}

    // 获取性别
    gender, err := strconv.Atoi(extractString(contents, genderRe))
    if err != nil {
    	profile.Gender = gender
	}

	// 获取星座
	profile.Xinzuo = extractString(contents, xinzuoRe)

	// 获取职业
	profile.Occupation = extractString(contents, occupationRe)

	// 获取学历
	profile.Education = extractString(contents, educationRe)

	// 获取籍贯
	profile.Hokou = extractString(contents, hokouRe)

	// 获取住房条件
	profile.House = extractString(contents, houseRe)

	// 获取是否购车
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items : []interface{} {profile},
	}

	return result
}

func extractString (contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}