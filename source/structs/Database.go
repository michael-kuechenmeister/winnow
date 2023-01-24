
package structs;

import "encoding/json";
import "os";
import "strings";



type Database struct {

	Folder string `json:"folder"`;

}



func NewDatabase (folder string) Database {

	var database Database;

	if strings.HasSuffix(folder, "/") {
		folder = folder[0:len(folder) - 1];
	}

	stat, err := os.Stat(folder);

	if err == nil && stat.IsDir() {

		database.Folder = folder;

	} else {

		err1 := os.MkdirAll(folder,                    0750);
		err2 := os.MkdirAll(folder + "/product",       0750);
		err3 := os.MkdirAll(folder + "/vulnerability", 0750);

		if err1 == nil && err2 == nil && err3 == nil {
			database.Folder = folder;
		}

	}

	return database;

}

func (database *Database) Reset () {

	stat1, err1 := os.Stat(database.Folder + "/product");

	if err1 == nil && stat1.IsDir() {

		entries, err := os.ReadDir(database.Folder + "/product");

		if err == nil {

			for e := 0; e < len(entries); e++ {

				var entry = entries[e].Name();

				if strings.HasSuffix(entry, ".json") {

					tmp := strings.Split(entry[0:len(entry) - 5], ":");

					if len(tmp) == 3 {

						product := database.ReadProduct(tmp[0], tmp[1], tmp[2]);

						if product.IsEdited == false {
							os.Remove(database.Folder + "/product/" + entries[e].Name());
						}

					}

				}

			}

		}

	}

	stat2, err2 := os.Stat(database.Folder + "/vulnerability");

	if err2 == nil && stat2.IsDir() {

		entries, err := os.ReadDir(database.Folder + "/vulnerability");

		if err == nil {

			for e := 0; e < len(entries); e++ {

				var entry = entries[e].Name();

				if strings.HasSuffix(entry, ".json") {

					identifier    := entry[0:len(entry) - 5];
					vulnerability := database.ReadVulnerability(identifier);

					if vulnerability.IsEdited == false {
						os.Remove(database.Folder + "/vulnerability/" + entries[e].Name());
					}

				}

			}

		}

	}

}

func (database *Database) ReadProduct (typ string, vendor string, name string) Product {

	product      := NewProduct("");
	buffer, err1 := os.ReadFile(database.Folder + "/product/" + typ + ":" + vendor + ":" + name + ".json");

	if err1 == nil && len(buffer) > 2 {

		err2 := json.Unmarshal(buffer, &product);

		if err2 == nil {
			// Do Nothing
		} else {
			product.SetName("");
			product.SetVendor("");
			product.SetType("");
		}

	} else {
		product.SetName(name);
		product.SetVendor(vendor);
		product.SetType(typ);
	}

	return product;

}

func (database *Database) ReadProducts () []Product {

	var result []Product;

	stat, err1 := os.Stat(database.Folder + "/product");

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(database.Folder + "/product");

		if err2 == nil {

			for e := 0; e < len(entries); e++ {

				entry := entries[e].Name();

				if strings.HasSuffix(entry, ".json") {

					tmp := strings.Split(entry[0:len(entry) - 5], ":");

					if len(tmp) == 3 {

						product := database.ReadProduct(tmp[0], tmp[1], tmp[2]);

						if IsProduct(product) {
							result = append(result, product);
						}

					}

				}

			}

		}

	}

	return result;

}

func (database *Database) ReadVulnerability (identifier string) Vulnerability {

	vulnerability := NewVulnerability();
	buffer, err1  := os.ReadFile(database.Folder + "/vulnerability/" + identifier + ".json");

	if err1 == nil && len(buffer) > 2 {

		err2 := json.Unmarshal(buffer, &vulnerability);

		if err2 == nil {
			// Do Nothing
		} else {
			vulnerability.SetIdentifier(identifier);
			vulnerability.SetSeverity("none");
			vulnerability.SetState("invalid");
		}

	} else {
		vulnerability.SetIdentifier(identifier);
		vulnerability.SetSeverity("none");
		vulnerability.SetState("invalid");
	}

	return vulnerability;

}

func (database *Database) ReadVulnerabilities () []Vulnerability {

	var result []Vulnerability;

	stat, err1 := os.Stat(database.Folder + "/vulnerability");

	if err1 == nil && stat.IsDir() {

		entries, err2 := os.ReadDir(database.Folder + "/vulnerability");

		if err2 == nil {

			for e := 0; e < len(entries); e++ {

				var entry = entries[e].Name();

				if strings.HasSuffix(entry, ".json") {

					identifier    := entry[0:len(entry) - 5];
					vulnerability := database.ReadVulnerability(identifier);

					if IsVulnerability(vulnerability) {
						result = append(result, vulnerability);
					}

				}

			}

		}

	}

	return result;

}

