package buildtags

import (
	"errors"
	"strings"

	"github.com/samber/lo"
)

// Parse函数用于解析给定的标签字符串，并返回
// 一个清理过的字符串切片。同时支持逗号和空格作为分隔符，
// 但不支持混合使用，若混合使用则会返回错误。
func X解析(标签字符串 string) ([]string, error) {
	if 标签字符串 == "" {
		return nil, nil
	}

	separator := ""
	if strings.Contains(标签字符串, ",") {
		separator = ","
	}
	if strings.Contains(标签字符串, " ") {
		if separator != "" {
			return nil, errors.New("cannot use both space and comma separated values with `-tags` flag")
		}
		separator = " "
	}
	if separator == "" {
// 我们未能找到任何分隔符，因此将整个字符串作为用户标签使用
// 否则最终我们将得到一个包含标签字符串中每个单字符的列表，
// 例如：`t,e,s,t`
		return []string{标签字符串}, nil
	}

	var userTags []string
	for _, tag := range strings.Split(标签字符串, separator) {
		thisTag := strings.TrimSpace(tag)
		if thisTag != "" {
			userTags = append(userTags, thisTag)
		}
	}
	return userTags, nil
}

// Stringify 将给定的标签切片转换为与 go build -tags 标志兼容的字符串

// ff:
// tags:
func Stringify(tags []string) string {
	tags = lo.Map(tags, func(tag string, _ int) string {
		return strings.TrimSpace(tag)
	})
	return strings.Join(tags, ",")
}
