
package main;

import "winnow/actions";
import "winnow/console";
import "winnow/structs";
import "os";
import "strings";



func showUsage() {

	console.Info("");
	console.Info("Tholian Winnow");
	console.Info("SBOM and Vulnerability Analyzer");
	console.Info("");
	console.Log("");
	console.Log("Usage: tholian-winnow [Action] [SBOM]");
	console.Log("");
	console.Log("Available Actions:");
	console.Log("");
	console.Log("    Action  | Description                                             |");
	console.Log("    --------|---------------------------------------------------------|");
	console.Log("    analyze | Analyzes the SBOM and returns a list of vulnerabilities |");
	console.Log("");
	console.Log("Examples:");
	console.Log("");
	console.Log("    tholian-winnow analyze /tmp/cyclonedx-report.json;");
	console.Log("");

}



func isFile (file string) bool {

	var result bool = false;

	// TODO: More thorough validation via
	// actually parsing the file, and if no errors
	// then return true

	if strings.HasSuffix(file, ".json") {
		result = true;
	} else if strings.HasSuffix(file, ".xml") {
		result = true;
	}

	return result;

}


func main() {

	var action string = "";
	var file   string = "";

	if len(os.Args) == 3 {

		action = strings.ToLower(os.Args[1]);
		tmp   := strings.ToLower(os.Args[2]);

		if isFile(tmp) {
			file = tmp;
		} else {
			file = "";
		}

	}


	console.Clear();

	console.Info("");
	console.Info("Tholian Winnow");
	console.Info("");

	console.Log("");
	console.Log("tholian-winnow: Command-Line Arguments:");
	console.Log("{");
	console.Log("    \"action\": \"" + action + "\",");
	console.Log("    \"file\":   \"" + file   + "\"");
	console.Log("}");


	if action != "" && file != "" {

		// TODO: This should be a better path?
		// TODO: Does it make sense to go:embed all vulnerabilities?

		cwd, _ := os.Getwd();

		var database = structs.NewDatabase(cwd + "../external/vulnerabilities/database");

		// TODO: sbom = structs.NewBOM(path/to/file.json);

		if action == "analyze" {
			actions.Analyze(&database);
		}

	} else {

		showUsage();
		os.Exit(1);

	}

}

