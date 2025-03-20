package split

import (
	"strings"
)

// Split 基础函数，等待测试
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i]) // 从找到的索引处之前的数据追加到 result 切片中
		s = s[i+1:]                    // 将字符串整理，从找到切割索引之后的所有字符串保存到原字符串中，等待下次的搜寻
		i = strings.Index(s, sep)      // 改变循环变量，通过再次寻找切割的字符所在的索引
	}
	result = append(result, s) // 将最后不包含切割字符的剩余字符追加到结果中
	return
}

func SplitMoreChar(s, sep string) (result []string) {
	// 修复切片动态扩容对内容的调用
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	//fmt.Println("当前字符串是： ", s, " 当前要找的分隔符是： ", sep, " 找到的索引是： ", i)

	for i > -1 {
		if i != 0 {
			// 分隔符不在首位的时候，把到首个分隔符之前的追加进去
			// 分隔符在首位的，要跳过
			result = append(result, s[:i])
		}
		//result = append(result, s[:i])
		// 从找到的索引处之前的数据追加到 result 切片中
		s = s[i+len(sep):]        // 将字符串整理，从找到切割索引之后的所有字符串保存到原字符串中，等待下次的搜寻
		i = strings.Index(s, sep) // 改变循环变量，通过再次寻找切割的字符所在的索引
	}
	result = append(result, s) // 将最后不包含切割字符的剩余字符追加到结果中
	return
}
