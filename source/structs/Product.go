
package structs;

import "strings";



func toArchitecture (cpe_version string, cpe_update string, cpe_target_hardware string) string {

	var architecture = "any";

	if strings.Contains(cpe_version, "32-bit") {
		architecture = "x86";
	} else if strings.Contains(cpe_version, "64-bit") {
		architecture = "x86_64";
	} else if strings.Contains(cpe_version, "x32") {
		architecture = "x86";
	} else if strings.Contains(cpe_version, "x64") {
		architecture = "x86_64";
	} else if strings.Contains(cpe_version, "x86") {
		architecture = "x86";
	} else if strings.Contains(cpe_version, "x86_64") {
		architecture = "x86_64";
	}

	if strings.Contains(cpe_update, "x32") {
		architecture = "x86";
	} else if strings.Contains(cpe_update, "x64") {
		architecture = "x86_64";
	} else if strings.Contains(cpe_update, "x86") {
		architecture = "x86";
	} else if strings.Contains(cpe_update, "x86_64") {
		architecture = "x86_64";
	}

	if architecture == "any" && IsArchitecture(cpe_target_hardware) {
		architecture = cpe_target_hardware;
	}

	return architecture;

}



type Product struct {

	Name         string  `json:"name"`;
	Version      string  `json:"version"`;
	Architecture string  `json:"architecture"`;
	Vendor       string  `json:"vendor"`;
	Type         string  `json:"type"`;
	IsEdited     bool    `json:"is_edited"`;

	Aliases      []Alias `json:"aliases"`;

	// Useless CPE Properties
	// Update          string `json:"update"`;
	// Edition         string `json:"edition"`;
	// Language        string `json:"language"`;
	// SoftwareEdition string `json:"sw_edition"`;
	// TargetSoftware  string `json:"target_sw"`;
	// TargetHardware  string `json:"target_hw"`;
	// Other           string `json:"other"`;

}

func NewProduct (cpe string) Product {

	var product Product;

	product.Version      = "any";
	product.Architecture = "any";
	product.Vendor       = "";
	product.Type         = "any";
	product.Aliases      = make([]Alias, 0);


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
				product.SetType("software");
			} else if cpe_part == "/h" {
				product.SetType("hardware");
			} else if cpe_part == "/o" {
				product.SetType("system");
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				product.SetName(cpe_product);
			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				product.SetVendor(cpe_vendor);
			}

		} else if len(tmp) == 5 && tmp[0] == "cpe" {

			var cpe_part    = strings.TrimSpace(tmp[1]);
			var cpe_vendor  = strings.TrimSpace(tmp[2]);
			var cpe_product = strings.TrimSpace(tmp[3]);
			var cpe_version = strings.TrimSpace(tmp[4]);

			if cpe_part == "/a" {
				product.SetType("software");
			} else if cpe_part == "/h" {
				product.SetType("hardware");
			} else if cpe_part == "/o" {
				product.SetType("system");
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				product.SetName(cpe_product);
			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				product.SetVendor(cpe_vendor);
			}

			if cpe_version != "" && cpe_version != "*" && cpe_version != "-" {
				product.SetVersion(cpe_version);
			}

		} else if len(tmp) == 13 && tmp[0] == "cpe" && tmp[1] == "2.3" {

			var cpe_part             = strings.TrimSpace(tmp[2]);
			var cpe_vendor           = strings.TrimSpace(tmp[3]);
			var cpe_product          = strings.TrimSpace(tmp[4]);
			var cpe_version          = strings.TrimSpace(tmp[5]);
			var cpe_update           = strings.TrimSpace(tmp[6]);
			// var cpe_edition          = strings.TrimSpace(tmp[7]);
			// var cpe_language         = strings.TrimSpace(tmp[8]);
			// var cpe_software_edition = strings.TrimSpace(tmp[9]);
			// var cpe_target_software  = strings.TrimSpace(tmp[10]);
			var cpe_target_hardware  = strings.TrimSpace(tmp[11]);
			// var cpe_other            = strings.TrimSpace(tmp[12]);
			var architecture         = toArchitecture(cpe_version, cpe_update, cpe_target_hardware);

			if cpe_part == "a" {
				product.SetType("software");
			} else if cpe_part == "h" {
				product.SetType("hardware");
			} else if cpe_part == "o" {
				product.SetType("system");
			}

			if strings.HasSuffix(cpe_version, "_x86_64") {
				cpe_version = cpe_version[0:len(cpe_version) - 7];
			} else if strings.HasSuffix(cpe_version, "_x86") {
				cpe_version = cpe_version[0:len(cpe_version) - 4];
			} else if strings.HasSuffix(cpe_version, "_64-bit") {
				cpe_version = cpe_version[0:len(cpe_version) - 7];
			} else if cpe_update == "32-bit" {
				cpe_version = "";
			} else if cpe_version == "64-bit" {
				cpe_version = "";
			}

			if strings.HasPrefix(cpe_update, "x86_64_") {
				cpe_update = cpe_update[7:];
			} else if strings.HasPrefix(cpe_update, "x86_") {
				cpe_update = cpe_update[4:];
			} else if cpe_update == "32-bit" {
				cpe_update = "";
			} else if cpe_update == "64-bit" {
				cpe_update = "";
			}

			if cpe_product != "" && cpe_product != "*" && cpe_product != "-" {
				product.SetName(cpe_product);
			}

			if cpe_version != "" && cpe_version != "*" && cpe_version != "-" {

				if cpe_update != "" && cpe_update != "*" && cpe_update != "-" {
					product.SetVersion(cpe_version + "-" + cpe_update);
				} else {
					product.SetVersion(cpe_version);
				}

			}

			if cpe_vendor != "" && cpe_vendor != "*" && cpe_vendor != "-" {
				product.SetVendor(cpe_vendor);
			}

			if architecture != "any" {
				product.SetArchitecture(architecture);
			}

		}

	}


	return product;

}



func IsIdenticalProduct (a Product, b Product) bool {

	var result bool = false;

	if a.Name == b.Name && a.Version == b.Version && a.Architecture == b.Architecture && a.Vendor == b.Vendor && a.Type == b.Type {
		result = true;
	}

	return result;

}

func IsProduct (product Product) bool {

	var result bool = true;

	if product.Name != "" {

		if product.Vendor == "" {
			result = false;
		}

	} else {
		result = false;
	}

	return result;

}

func (product *Product) Matches (search Product) bool {

	var matches_name    bool = false;
	var matches_version bool = false;
	var matches_vendor  bool = false;
	var matches_type    bool = false;

	if product.Name == search.Name {
		matches_name = true;
	}

	if product.Version == "any" || search.Version == "any" {
		matches_version = true;
	} else {

		// TODO: Version Comparison (like in Software.Matches())

	}

	if product.Vendor == "" || search.Vendor == "" {
		matches_vendor = true;
	} else if product.Vendor == search.Vendor {
		matches_vendor = true;
	}

	if product.Type == "any" || search.Type == "any" {
		matches_type = true;
	} else if product.Type == search.Type {
		matches_type = true;
	}

	for a := 0; a < len(product.Aliases); a++ {

		var alias = product.Aliases[a];

		if alias.Name == search.Name && alias.Vendor == search.Vendor && alias.Type == search.Type {
			matches_name   = true;
			matches_vendor = true;
			matches_type   = true;
			break;
		}

	}

	return matches_name && matches_version && matches_vendor && matches_type;

}

func (product *Product) MatchesAlias (search Alias) bool {

	var matches_name   bool = false;
	var matches_vendor bool = false;
	var matches_type   bool = false;

	if product.Name == search.Name {
		matches_name = true;
	}

	if product.Vendor == search.Vendor {
		matches_vendor = true;
	}

	if product.Type == search.Type {
		matches_type = true;
	}

	for a := 0; a < len(product.Aliases); a++ {

		var alias = product.Aliases[a];

		if alias.Matches(search) == true {
			matches_name   = true;
			matches_vendor = true;
			matches_type   = true;
			break;
		}

	}

	return matches_name && matches_vendor && matches_type;

}

func (product *Product) SetName (value string) {

	value = strings.TrimSpace(value);
	value = strings.ToLower(value);
	value = strings.ReplaceAll(value, ":", "-");
	value = strings.ReplaceAll(value, "_", "-");

	product.Name = value;

}

func (product *Product) SetVersion (value string) {

	if value == "all" || value == "any" || value == "*" {
		product.Version = "any";
	} else if value != "" {
		product.Version = value;
	}

}

func (product *Product) SetArchitecture (value string) {

	if value == "all" || value == "any" || value == "*" {
		product.Architecture = "any";
	} else if value == "i386" || value == "i686" || value == "x32" || value == "x86" {
		product.Architecture = "x86";
	} else if value == "amd64" || value == "ia64" || value == "x64" || value == "x86_64" {
		product.Architecture = "x86_64";
	} else if value == "armel" || value == "armv6" {
		product.Architecture = "armv6";
	} else if value == "armhf" || value == "armv7" || value == "armv7h" {
		product.Architecture = "armv7";
	} else if value == "aarch64" || value == "armv8" {
		product.Architecture = "armv8";
	} else if value == "sparc" {
		product.Architecture = "sparc";
	} else if value == "sparc64" {
		product.Architecture = "sparc64";
	}

}

func (product *Product) SetVendor (value string) {

	value = strings.TrimSpace(value);
	value = strings.ToLower(value);
	value = strings.ReplaceAll(value, ":", "-");
	value = strings.ReplaceAll(value, "_", "-");

	product.Vendor = value;

}

func (product *Product) SetType (value string) {

	if value == "hardware" {
		product.Type = "hardware";
	} else if value == "software" {
		product.Type = "software";
	} else if value == "system" {
		product.Type = "system";
	}

}

func (product *Product) AddAlias (value Alias) {

	var found bool = false;

	for a := 0; a < len(product.Aliases); a++ {

		var other = product.Aliases[a];

		if IsIdenticalAlias(other, value) {
			found = true;
			break;
		}

	}

	if found == false {
		product.Aliases = append(product.Aliases, value);
	}

}

func (product *Product) RemoveAlias (value Alias) {

	var index int = -1;

	for a := 0; a < len(product.Aliases); a++ {

		var other = product.Aliases[a];

		if IsIdenticalAlias(other, value) {
			index = a;
			break;
		}

	}

	if index != -1 {
		product.Aliases = append(product.Aliases[:index], product.Aliases[index + 1:]...);
	}

}

func (product *Product) SetAliases (value []Alias) {

	var filtered []Alias;

	for v := 0; v < len(value); v++ {

		var alias = value[v];

		if IsAlias(alias) {
			filtered = append(filtered, alias);
		}

	}

	product.Aliases = filtered;

}

