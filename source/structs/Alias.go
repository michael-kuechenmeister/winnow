
package structs;

import "strings";



type Alias struct {

	Name   string `json:"name"`;
	Vendor string `json:"vendor"`;
	Type   string `json:"type"`;

}

func NewAlias (cpe string) Alias {

	var alias Alias;


	if strings.HasPrefix(cpe, "cpe:") {

		cpe = strings.ReplaceAll(cpe, "\\:", "%SEPARATOR%");

		var tmp = strings.Split(strings.TrimSpace(cpe), ":");

		for t := 0; t < len(tmp); t++ {
			tmp[t] = strings.ReplaceAll(tmp[t], "%SEPARATOR%", ":");
		}

		if len(tmp) == 4 && tmp[0] == "cpe" {

			var cpe_part    = strings.TrimSpace(tmp[1]);
			var cpe_vendor  = strings.TrimSpace(tmp[2]);
			var cpe_product = strings.TrimSpace(tmp[3]);

			if cpe_part == "/a" {
				alias.SetType("software");
			} else if cpe_part == "/h" {
				alias.SetType("hardware");
			} else if cpe_part == "/o" {
				alias.SetType("system");
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				alias.SetName(cpe_product);
			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				alias.SetVendor(cpe_vendor);
			}

		} else if len(tmp) == 5 && tmp[0] == "cpe" {

			var cpe_part    = strings.TrimSpace(tmp[1]);
			var cpe_vendor  = strings.TrimSpace(tmp[2]);
			var cpe_product = strings.TrimSpace(tmp[3]);
			// var cpe_version = strings.TrimSpace(tmp[4]);

			if cpe_part == "/a" {
				alias.SetType("software");
			} else if cpe_part == "/h" {
				alias.SetType("hardware");
			} else if cpe_part == "/o" {
				alias.SetType("system");
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				alias.SetName(cpe_product);
			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				alias.SetVendor(cpe_vendor);
			}

		} else if len(tmp) == 13 && tmp[0] == "cpe" && tmp[1] == "2.3" {

			var cpe_part    = strings.TrimSpace(tmp[2]);
			var cpe_vendor  = strings.TrimSpace(tmp[3]);
			var cpe_product = strings.TrimSpace(tmp[4]);

			if cpe_part == "a" {
				alias.SetType("software");
			} else if cpe_part == "h" {
				alias.SetType("hardware");
			} else if cpe_part == "o" {
				alias.SetType("system");
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				alias.SetName(cpe_product);
			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				alias.SetVendor(cpe_vendor);
			}

		}

	}


	return alias;

}



func IsIdenticalAlias (a Alias, b Alias) bool {

	var result bool = false;

	if a.Name == b.Name && a.Vendor == b.Vendor {
		result = true;
	}

	return result;

}

func IsAlias (alias Alias) bool {

	var result bool = true;

	if alias.Name == "" {
		result = false;
	}

	if alias.Vendor == "" {
		result = false;
	}

	return result;

}

func (alias *Alias) Matches (search Alias) bool {

	var matches_name   bool = false;
	var matches_vendor bool = false;

	if alias.Name == search.Name {
		matches_name = true;
	}

	if alias.Vendor == "" || search.Vendor == "" {
		matches_vendor = true;
	} else if alias.Vendor == search.Vendor {
		matches_vendor = true;
	}

	return matches_name && matches_vendor;

}

func (alias *Alias) SetName (value string) {
	alias.Name = value;
}

func (alias *Alias) SetVendor (value string) {
	alias.Vendor = value;
}

func (alias *Alias) SetType (value string) {

	if value == "hardware" {
		alias.Type = "hardware";
	} else if value == "software" {
		alias.Type = "software";
	} else if value == "system" {
		alias.Type = "system";
	}

}

