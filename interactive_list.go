package interactivelist

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/digvijay-tech/interactive_list/utilities"
)

func SimpleSelector(items []string, title, description string) (selectedItem string) {
	if len(items) < 1 {
		log.Fatalln("items cannot be empty")
	}

	currenPos := 0

	for {
		printList(items, currenPos, title, description)
		oldState := utilities.EnableRawMode()

		byteArr := make([]byte, 3)
		os.Stdin.Read(byteArr)

		utilities.DisableRawMode(oldState)

		if byteArr[0] == 3 {
			break
		}

		if byteArr[0] == 13 {
			break
		}

		if byteArr[0] == 27 && byteArr[1] == 91 {
			switch byteArr[2] {
			case 65:
				if currenPos <= 0 {
					currenPos = len(items) - 1
				} else {
					currenPos -= 1
				}
				continue
			case 66:
				if currenPos >= len(items)-1 {
					currenPos = 0
				} else {
					currenPos += 1
				}
				continue
			}
		}
	}

	return items[currenPos]
}

func printList(items []string, currPos int, title, description string) {
	utilities.Clear()

	if strings.TrimSpace(title) != "" {
		fmt.Println(title)
	}

	if strings.TrimSpace(description) != "" {
		fmt.Println(description)
		fmt.Printf("\n")
	}

	for i, v := range items {
		if i == currPos {
			fmt.Printf("> %s\n", v)
		} else {
			fmt.Printf("  %s\n", v)
		}
	}
}
