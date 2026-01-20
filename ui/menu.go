package ui

import (
	"bufio"
	"control_users/model"
	"control_users/service"
	"control_users/utils"
	"fmt"
	"strings"
)

func Menu(service *service.UserService, scanner *bufio.Scanner) {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1)add users switch status and name.\n2)Display users switch true status.\n3)Update status by name.\n4)Found user by name.5)Display all users\n6)Delete user by name\n7)End program.")
		scanner.Scan()
		x := strings.TrimSpace(scanner.Text())
		switch x {
		case "1":
			fmt.Println("input qautity users: ")
			num, err := utils.ReadInt(scanner)
			if err != nil {
				fmt.Println("invalid qautity users")
				continue
			}
			for i := 0; i < num; i++ {
				fmt.Println("input please name")
				name, err := utils.ReadLine(scanner)
				if err != nil {
					fmt.Println("invalid input name")
					i--
					continue
				}
				fmt.Println("input status")
				status, err := utils.ReadBool(scanner)
				if err != nil {
					fmt.Println("invalid input status")
					i--
					continue
				}
				u := model.User{Name: name, Status: status}
				service.AddUser(u)
			}
		case "2":
			trueStatus := service.FilterUsersByStatus()
			if len(trueStatus) <= 0 {
				fmt.Println("users switch status true not found")
			}
			for _, j := range trueStatus {
				fmt.Println(j)
			}
		case "3":
			fmt.Println("input name")
			name, err := utils.ReadLine(scanner)
			if err != nil {
				fmt.Println("invalid input")
				continue
			}
			_, err = service.FindUsersByName(name)
			if err != nil {
				fmt.Println("user not found")
				continue
			}
			fmt.Println("input new status")
			newStatus, err := utils.ReadBool(scanner)
			if err != nil {
				fmt.Println("invalid input for update status")
				continue
			}
			err = service.UpdateUserStatus(name, newStatus)
			if err != nil {
				fmt.Println("error update user status")
			}
			fmt.Printf("status user %s update to %t\n", name, newStatus)
		case "4":
			fmt.Println("input name")
			name, err := utils.ReadLine(scanner)
			if err != nil {
				fmt.Println("invalid input")
				continue
			}
			user, err := service.FindUsersByName(name)
			if err != nil {
				fmt.Println("user not found")
				continue
			}
			fmt.Println(user)
		case "5":
			users := service.ListUsers()
			for i, j := range users {
				fmt.Printf("user #%d name: %s status: %t\n", i+1, j.Name, j.Status)
			}
		case "6":
			fmt.Println("input name user for deletion")
			name, err := utils.ReadLine(scanner)
			if err != nil {
				fmt.Println("invalid input name user for deletion")
				continue
			}
			err = service.DeleteUser(name)
			if err != nil {
				fmt.Println("failed to delete user")
			} else {
				fmt.Println("user успешно удален")
			}
		case "7":
			fmt.Println("End program")
			return
		default:
			fmt.Println("Неизвестная командна")
			continue
		}
	}
}
