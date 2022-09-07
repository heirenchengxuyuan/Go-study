package Hello

import "fmt"

func main() {
	var i int
	// 定义变量 数据类型放在最后
	i = 10
	fmt.Println("i =", i)

	var conferenceName = "Go conferance"
	const conferenceTickets = 50
	var remainingTickets = 50
	// go中如果没有使用到所定义变量 则会显示报错 若包没有被用也一样 const用来定义常量 var用来定义变量
	fmt.Println(conferenceName)
	fmt.Println("welcome to", conferenceName, "booking application")
	fmt.Println("we have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	//打印语句中使用变量
	fmt.Printf("welcome to %v conferenceName booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)

	//使用printf的格式化输出 %v 即占位符 后面要加上所占位的

}
