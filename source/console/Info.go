
package console;

import "fmt";
import "strings";



func Info (message string) {

	if strings.Contains(message, "\n") {

		var lines = strings.Split(message, "\n");

		for l := 0; l < len(lines); l++ {
			fmt.Println("\u001b[42m " + lines[l] + "\u001b[K");
		}

		fmt.Println("\u001b[0m");

	} else {
		fmt.Println("\u001b[42m " + message + "\u001b[K\u001b[0m");
	}

}

