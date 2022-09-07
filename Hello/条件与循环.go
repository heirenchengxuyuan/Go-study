package Hello

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conferance"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	//bookings := []string{}
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T，conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//go语言中所用变量必须在定义变量语句的下面行
	for remainingTickets > 0 && len(bookings) < 50 {
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

		//if userTickets > remainingTickets {
		//	fmt.Printf("we only have %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTickets)
		//	break
		//	//break 会让余下的代码也不运行
		//	continue
		//	//continue 会跳出这一个循环 但是会进入到下一个循环迭代
		//}

		if userTickets <= remainingTickets {

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

			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. come back next year")
				break
				//以上是循环语句 下面是布尔语句
				// var noTicketsRemaining bool = remainingTickets == 0
				//或者写成 noTicketsRemaining := remainingTickets == 0
				// if noTicketsRemaining{
				//fmt.Println("our conference is booked out. come back next year")
				//break}
			}
		} else {
			fmt.Printf("we only have %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
}
