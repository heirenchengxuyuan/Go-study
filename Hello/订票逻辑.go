package Hello

import "fmt"

func main() {
	var conferenceName string = "Go conferance"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T，conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//%T 即数据类型
	//go 会自动判断数据类型 不说明也可使用 但是为了代码的健壮性应当作好定义
	//且设置好一个类型会保护变量 不让它得到不该得到的值

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
	//直接写变量即可

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
