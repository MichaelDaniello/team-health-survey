package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("Welcome to the Nic Cage Team Survey")
	fmt.Println("-------------------------------------------------")

	label := []string{"1. Team members are passionate and unguarded in their discussion of issues",
		"2. Team members call out one another’s deficiencies or unproductive behaviors",
		"3. Team members know what their peers are working on and how they contribute to the collective good of the team",
		"4. Team members quickly and genuinely apologize to one another when they say or do something inappropriate or possibly damaging to the team",
		"5. Team members willingly make sacrifices (such as budget, turf, head count) in their departments or areas of expertise for the good of the team",
		"6. Team members openly admit their weaknesses and mistakes",
		"7. Team meetings are compelling, not boring",
		"8. Team members leave meetings confident that their peers are completely committed to the decisions that were agreed on, even if they were in initial disagreement",
		"9. Morale is significantly affected by the failure to achieve team goals",
		"10.During team meetings, the most important—and difficult—issues are put on the table to be resolved",
		"11. Team members are deeply concerned about the prospect of letting down their peers",
		"12. Team members know about one another’s personal lives and are comfortable discussing them",
		"13. Team members end discussions with clear and specific resolutions and action plans",
		"14. Team members challenge one another about their plans and approaches",
		"15. Team members are slow to seek credit for their own contributions, but quick to point out those of others"}
	total := 0

	for i := 0; i < len(label); i++ {
		prompt := promptui.Select{
			Label: label[i],
			Items: []string{"Usually", "Sometimes", "Rarely"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v.\n", err)
			return
		}
		fmt.Printf("You choose %q.\n", result)
		switch {
		case result == "Usually":
			total += 3
		case result == "Sometimes":
			total += 2
		default:
			total++
		}
		fmt.Printf("Your total is %v.\n", total)
	}

	// fmtPrintf("You choose %q\n", result)

}
