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

		//以下为验证语句
		isvalueName := len(userName) >= 2
		isvalueEmail := strings.Contains(email, "@")
		isvalueTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		//isvalidCity := city == "xiaogan" || city == "fuling"
		//不用写ture 默认

		if isvalueName && isvalueEmail && isvalueTicketNumber {

			remainingTickets = remainingTickets - userTickets

			//var bookings [50]string
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

			}
			//起到决策树的作用
		} else {
			if !isvalueName {
				fmt.Println("Name you entered is too short")
			}
			if !isvalueEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isvalueTicketNumber {
				fmt.Printf("we only have %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTickets)
				continue
				//多个if嵌套则每一个都验证 如果是else if 语句则第一个通过其余将不验证
			}
		}
	}

	//
	//city := "London"
	//
	//switch city {
	//case "New York":
	//	// execute code ……
	//case "Singapore","XiaoGan":
	//	// ……
	//case "Berlin":
	////……
	//default:
	//	fmt.Print("No valid city selected")

	//}
}
