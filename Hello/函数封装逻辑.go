package Hello

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go conferance"
var remainingTickets uint = 50
var bookings []string

//

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	//这里便是创建了greetusers 函数 并在main中调用

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isvalueName, isvalueEmail, isvalueTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		//要将返回值放置等号前 ：=

		if isvalueName && isvalueEmail && isvalueTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			//调用函数

			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. come back next year")
				break
			}
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
			}
		}
	}
}

func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("Welcome to our conference")
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings []string) []string {
	//括号内是输入参数 括号外为输出参数

	firstNames := []string{}
	//go中下划线用于标识未使用变量
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
	//fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isvalueName := len(firstName) >= 2 && len(lastName) >= 2
	isvalueEmail := strings.Contains(email, "@")
	isvalueTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isvalueName, isvalueEmail, isvalueTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
