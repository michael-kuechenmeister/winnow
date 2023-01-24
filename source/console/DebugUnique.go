
package console;

import "fmt";
import "strings";



var DEBUG_CACHE []string;

func DebugUnique (message string) {

	var found bool = false;

	for d := 0; d < len(DEBUG_CACHE); d++ {

		if DEBUG_CACHE[d] == message {
			found = true;
			break;
		}

	}

	if found == false {

		DEBUG_CACHE = append(DEBUG_CACHE, message);

		if strings.Contains(message, "\n") {

			var lines = strings.Split(message, "\n");

			for l := 0; l < len(lines); l++ {
				fmt.Println("\u001b[41m " + lines[l] + "\u001b[K");
			}

			fmt.Println("\u001b[0m");

		} else {
			fmt.Println("\u001b[41m " + message + "\u001b[K\u001b[0m");
		}

	}

}

