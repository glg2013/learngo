package main

import "fmt"

/*
练习题
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	//users2 = []string{
	//	"Matthew",
	//}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("获得金币的详细名单： ", distribution)
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	// 1.遍历用户名单，统计一下，每个人的名字值多少金币
	for _, roleName := range users {
		for _, word := range roleName {
			//fmt.Printf("%T\n", word)
			//fmt.Println(word)
			switch word {
			case 'e', 'E':
				distribution[roleName] += 1
				coins -= 1
			case 'i', 'I':
				distribution[roleName] += 2
				coins -= 2
			case 'o', 'O':
				distribution[roleName] += 3
				coins -= 3
			case 'u', 'U':
				distribution[roleName] += 4
				coins -= 4
			default:
			}
		}
	}
	//fmt.Println(coins, "剩余的就这么多了")
	return coins
}
