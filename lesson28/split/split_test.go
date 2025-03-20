package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%v, got:%v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want:%v, got:%v", want, got)
	}
}

func TestMoreSpli(t *testing.T) {
	got := SplitMoreChar("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestGroupSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}

	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	// 遍历切片，注意执行测试用例
	for _, tc := range tests {
		got := SplitMoreChar(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%#v, got:%#v", tc.want, got)
		}
	}

}

func TestGroupMapSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	//for name, tc := range tests {
	//	got := SplitMoreChar(tc.input, tc.sep)
	//	if !reflect.DeepEqual(got, tc.want) {
	//		t.Errorf("name:%s expected:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
	//	}
	//}

	// 使用子测试
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SplitMoreChar(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 前面的是函数的测试
// 以下的是基准测试
//func BenchmarkSplitMoreChar(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SplitMoreChar("沙河有沙又有河", "沙")
//	}
//}

// 应该没有这种了
//	func BenchmarkSplitMoreChar(b *testing.B) {
//		for i := 0; i < b.N; i++ {
//			SplitMoreChar("a:b:c", ":")
//		}
//	}
//func BenchmarkSplitMoreChar1(b *testing.B)  { SplitMoreChar("a:b:c", ":") }
//func BenchmarkSplitMoreChar2(b *testing.B)  { SplitMoreChar("a:b:c", ":") }
//func BenchmarkSplitMoreChar3(b *testing.B)  { SplitMoreChar("a:b:c", ":") }
//func BenchmarkSplitMoreChar10(b *testing.B) { SplitMoreChar("a:b:c", ":") }
//func BenchmarkSplitMoreChar20(b *testing.B) { SplitMoreChar("a:b:c", ":") }

// 假设我们有一个需要测试性能的函数
//func Fibonacci(n int) int {
//	if n <= 1 {
//		return n
//	}
//	return Fibonacci(n-1) + Fibonacci(n-2)
//}
//
//// 基准测试函数
//func BenchmarkFibonacci(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Fibonacci(20) // 测试 Fibonacci(20) 的性能
//	}
//}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(20)
	}
}
