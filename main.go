package main

import (
	"fmt"
	"pranav/admin"
	"pranav/staff"
)

func main() {
	var user1 = admin.Createnewadmin("pk", "kp", 1)
	var test = user1.Createuser("Pranav", "khairnar")
	var staff1 = staff.CreateNewStaff("jp", "kp", 5)
	staff1.CreateContact(&test, "Yash", "shah")

	fmt.Println(test)
}
