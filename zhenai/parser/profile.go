package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
	"strconv"
)

// \d就是所有的数字0-9
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var nativePlaceRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	profile.Age = extractInt(contents, ageRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Weight = extractInt(contents, weightRe)

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Car = extractString(contents, carRe)
	profile.Education = extractString(contents, educationRe)
	profile.NativePlace = extractString(contents, nativePlaceRe)
	profile.House = extractString(contents, houseRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Constellation = extractString(contents, constellationRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(contents)
	str := ""
	if len(match) >= 2 {
		str = string(match[1])
	}
	number, _ := strconv.Atoi(str)
	return number
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}