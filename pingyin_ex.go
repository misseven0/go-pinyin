package pinyin

import (
	"strings"
)

// 转换函数 首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 转换函数 全部大写
func ToUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s)
}
func ConvertAll(text string, args *Args) [][]string {
	pinyins := Convert(text, args)
	return CartesianProduct(pinyins)
}

func SlugAll(text string, a Args) []string {
	pinyins := ConvertAll(text, &a)
	slugs := make([]string, 0, len(pinyins))
	for _, v := range pinyins {
		pSlug := strings.Join(v, a.Separator)
		slugs = append(slugs, pSlug)
	}
	// pinyinsSlugs = Uniq(pinyinsSlugs)
	return slugs
}

func Uniq[T comparable](s []T) []T {
	if len(s) <= 1 { //only 1 element
		return s
	}
	result := make([]T, 0, len(s))
	seen := make(map[T]struct{}, len(s))
	for _, item := range s {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}
	return result
}

// 排列组合 笛卡尔积
func CartesianProduct(slices [][]string) [][]string {
	// 如果输入的二维切片为空，返回空切片
	if len(slices) == 0 {
		return [][]string{}
	}
	// 初始时，结果为第一个子切片
	result := make([][]string, len(slices[0]))
	for i, item := range slices[0] {
		result[i] = []string{item}
	}
	// 从第二个子切片开始计算笛卡尔积
	for i := 1; i < len(slices); i++ {
		var tmp [][]string
		// 计算当前子切片与之前结果的笛卡尔积
		for _, item := range slices[i] {
			for _, res := range result {
				tmp = append(tmp, append(append([]string{}, res...), item))
			}
		}
		// 更新结果为当前子切片与之前结果的笛卡尔积
		result = tmp
	}
	return result
}
