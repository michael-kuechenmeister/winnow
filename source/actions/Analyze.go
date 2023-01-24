
package actions;

import "winnow/console";
import "winnow/structs";



func Analyze (database *structs.Database, /* bom *structs.BOM */) bool {

	var result bool = false;

	console.Info("");
	console.Info("> Update");
	console.Info("");

	// TODO: Validate BOM and show errors/vulnerabilities

	return result;

}

