package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var totals = map[string]int{
	"Absence of Trust":            0,
	"Fear of Conflict":            0,
	"Lack of Commitment":          0,
	"Avoidance of Accountability": 0,
	"Inattention to Results":      0,
}

func main() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("Welcome to the Nic Cage Team Survey")
	fmt.Println("-------------------------------------------------")

	label := []string{
		"1. Team members are passionate and unguarded in their discussion of issues",
		"2. Team members call out one another’s deficiencies or unproductive behaviors",
		"3. Team members know what their peers are working on and how they contribute to the collective good of the team",
		"4. Team members quickly and genuinely apologize to one another when they say or do something inappropriate or possibly damaging to the team",
		"5. Team members willingly make sacrifices (such as budget, turf, head count) in their departments or areas of expertise for the good of the team",
		"6. Team members openly admit their weaknesses and mistakes",
		"7. Team meetings are compelling, not boring",
		"8. Team members leave meetings confident that their peers are completely committed to the decisions that were agreed on, even if they were in initial disagreement",
		"9. Morale is significantly affected by the failure to achieve team goals",
		"10. During team meetings, the most important—and difficult—issues are put on the table to be resolved",
		"11. Team members are deeply concerned about the prospect of letting down their peers",
		"12. Team members know about one another’s personal lives and are comfortable discussing them",
		"13. Team members end discussions with clear and specific resolutions and action plans",
		"14. Team members challenge one another about their plans and approaches",
		"15. Team members are slow to seek credit for their own contributions, but quick to point out those of others",
	}

	for i := 1; i <= len(label); i++ {
		prompt := promptui.Select{
			Label: label[i-1],
			Items: []string{"Usually", "Sometimes", "Rarely"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v.\n", err)
			return
		}
		resultChecker(result, categoryGetter(i))
	}
	fmt.Println(
		"Your score for Absence of Trust is",
		totals["Absence of Trust"],
		"\nYour score for Fear of Conflict is",
		totals["Fear of Conflict"],
		"\nYour score for Lack of Commitment is",
		totals["Lack of Commitment"],
		"\nYour score for Avoidance of Accountability is",
		totals["Avoidance of Accountability"],
		"\nYour score for Inattention to Results is",
		totals["Inattention to Results"],
	)
	fmt.Println(
		"-------------------------------------------------",
		"A score of 3 to 5 is probably an indication that the dysfunction needs to be addressed.",
		"A score of 6 or 7 indicates that the dysfunction could be a problem.",
		"A score of 8 or 9 is a probable indication that the dysfunction is not a problem for your team.",
		"-------------------------------------------------\n",
		`Reguardless of your scores, it is important to keep in mind that every team needs constant work,
	 	because without it, even the best ones deviate toward dysfunction.`)
}

func resultChecker(result string, category string) {

	switch {
	case result == "Usually":
		totals[category] += 3
	case result == "Sometimes":
		totals[category] += 2
	default:
		totals[category]++
	}
}

func categoryGetter(questionNum int) string {
	fmt.Println(questionNum)
	switch {
	case questionNum == 4:
		return "Absence of Trust"
	case questionNum == 6:
		return "Absence of Trust"
	case questionNum == 12:
		return "Absence of Trust"
	case questionNum == 1:
		return "Fear of Conflict"
	case questionNum == 7:
		return "Fear of Conflict"
	case questionNum == 10:
		return "Fear of Conflict"
	case questionNum == 3:
		return "Lack of Commitment"
	case questionNum == 8:
		return "Lack of Commitment"
	case questionNum == 13:
		return "Lack of Commitment"
	case questionNum == 2:
		return "Avoidance of Accountability"
	case questionNum == 11:
		return "Avoidance of Accountability"
	case questionNum == 14:
		return "Avoidance of Accountability"
	case questionNum == 5:
		return "Inattention to Results"
	case questionNum == 9:
		return "Inattention to Results"
	case questionNum == 15:
		return "Inattention to Results"
	}
	return ""

}
