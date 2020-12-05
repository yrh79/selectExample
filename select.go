/*
 * @Author: Jame Yu
 * @Date: 2020-12-04 20:08:54
 * @LastEditors: Jame Yu
 * @LastEditTime: 2020-12-04 20:14:03
 * @FilePath: /selectExample/select.go
 */
package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	for {
		prompt := promptui.Select{
			Label: "Select Operations:",
			Items: []string{"Start", "Restart", "Stop", "Exit"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "Exit" {
			fmt.Println("Exit.")
			return
		}

		fmt.Printf("You choose %q\n", result)
	}

}
