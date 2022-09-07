package Hello

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conferance"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T，conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//%T 即数据类型
	//go 会自动判断数据类型 不说明也可使用 但是为了代码的健壮性应当作好定义
	//且设置好一个类型会保护变量 不让它得到不该得到的值

	//var bookings = [50]string{"Nana", "Nicole", "Peter"}
	//数组大小即其能有多少元素 数组就像数组列表或元素列表

	//var bookings [50]string
	//bookings[0] = "Nana"
	//bookings[10] = "Nicole"

	//go语言中所用变量必须在定义变量语句的下面行
	for {
		var userName string
		var email string
		var userTickets uint

		//uint 是指正数
		fmt.Println("Enter your first name:")
		fmt.Scan(&userName)
		//询问用户名 &即指针 传递的是变量的内存地址

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		var bookings [50]string
		//基本上无需开始时指定大小 会在新元素被编辑时自动扩展
		bookings[0] = userName + " "
		//直接写变量即可

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))
		//打印数组的长度

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		userNames := []string{}
		//go中下划线用于标识未使用变量
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			userNames = append(userNames, names[0])

		}
		fmt.Printf("The names of bookings are: %v\n", userNames)
		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}

//以上是数组 如果用slice则是 (会在新元素被编辑时自动扩展)
//var bookings []string{}
//bookings = append(bookings, userName+ " ")
//fmt.Printf("The whole slice: %v\n", bookings)
//fmt.Printf("The first value: %v\n", bookings[0])
//fmt.Printf("slice type: %T\n", bookings)
//fmt.Printf("slice length: %v\n", len(bookings))
