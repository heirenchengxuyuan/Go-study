package Hello

import "fmt"

func main() {
	var conferenceName string = "Go conferance"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T，conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//%T 即数据类型
	//go 会自动判断数据类型 不说明也可使用 但是为了代码的健壮性应当作好定义
	//且设置好一个类型会保护变量 不让它得到不该得到的值

	var userName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)
	//询问用户名 &即指针 传递的是变量的内存地址

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, email)
	//分清楚println和printf

	fmt.Println(remainingTickets)
	// 上面是打印出值 下面是打印变量的内存地址（哈希值）
	fmt.Println(&remainingTickets)
	userName = "Tom"

}
