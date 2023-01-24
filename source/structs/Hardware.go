
package structs;

import "winnow/console";
import "bytes";
import "strings";



type Hardware struct {

	Name         string `json:"name"`;
	Version      string `json:"version"`;
	Architecture string `json:"architecture"`;
	Serial       string `json:"serial"`;
	Vendor       string `json:"vendor"`;

}

func NewHardware () Hardware {

	var hardware Hardware;

	hardware.Version      = "any";
	hardware.Architecture = "any";
	hardware.Serial       = "any";
	hardware.Vendor       = "";

	return hardware;

}



func IsIdenticalHardware (a Hardware, b Hardware) bool {

	var result bool = false;

	if a.Name == b.Name && a.Version == b.Version && a.Architecture == b.Architecture && a.Serial == b.Serial && a.Vendor == b.Vendor {
		result = true;
	}

	return result;

}

func IsHardware (hardware Hardware) bool {

	var result bool = true;

	if hardware.Name != "" {

		if hardware.Vendor == "" {
			result = false;
		}

	} else {
		result = false;
	}

	return result;

}

func (hardware *Hardware) Matches (search Hardware) bool {

	var matches_name         bool = false;
	var matches_version      bool = false;
	var matches_architecture bool = false;
	var matches_serial       bool = false;
	var matches_vendor       bool = false;

	if hardware.Name == search.Name {
		matches_name = true;
	}

	if hardware.Version == "any" || search.Version == "any" {
		matches_version = true;
	} else {

		// TODO: hardware.Version is unclear. Board/PCB version?

	}

	if hardware.Architecture == "any" || search.Architecture == "any" {
		matches_architecture = true;
	} else if hardware.Architecture == search.Architecture {
		matches_architecture = true;
	}

	if hardware.Serial == "any" || search.Serial == "any" {
		matches_serial = true;
	} else if hardware.Serial == search.Serial {
		matches_serial = true;
	}

	if hardware.Vendor == "" || search.Vendor == "" {
		matches_vendor = true;
	} else if hardware.Vendor == search.Vendor {
		matches_vendor = true;
	}

	return matches_name && matches_version && matches_architecture && matches_serial && matches_vendor;

}

func (hardware *Hardware) ToLog () string {

	var buffer bytes.Buffer;

	buffer.WriteString("Hardware({\n");
	buffer.WriteString("    Name: \""         + hardware.Name         + "\",\n");
	buffer.WriteString("    Version: \""      + hardware.Version      + "\",\n");
	buffer.WriteString("    Architecture: \"" + hardware.Architecture + "\",\n");
	buffer.WriteString("    Serial: \""       + hardware.Serial       + "\",\n");
	buffer.WriteString("    Vendor: \""       + hardware.Vendor       + "\"\n");
	buffer.WriteString("})");

	return buffer.String();

}

func (hardware *Hardware) Log () {

	var buffer = hardware.ToLog();
	var chunks = strings.Split(buffer, "\n");

	for c := 0; c < len(chunks); c++ {
		console.Log(chunks[c]);
	}

}

func (hardware *Hardware) SetArchitecture (value string) {

	if value == "all" || value == "any" || value == "*" {
		hardware.Architecture = "any";
	} else if value == "i386" || value == "i686" || value == "x32" || value == "x86" {
		hardware.Architecture = "x86";
	} else if value == "amd64" || value == "ia64" || value == "x64" || value == "x86_64" {
		hardware.Architecture = "x86_64";
	} else if value == "armel" || value == "armv6" {
		hardware.Architecture = "armv6";
	} else if value == "armhf" || value == "armv7" || value == "armv7h" {
		hardware.Architecture = "armv7";
	} else if value == "aarch64" || value == "armv8" {
		hardware.Architecture = "armv8";
	} else if value == "sparc" {
		hardware.Architecture = "sparc";
	} else if value == "sparc64" {
		hardware.Architecture = "sparc64";
	}

}

func (hardware *Hardware) SetName (value string) {
	hardware.Name = strings.TrimSpace(value);
}

func (hardware *Hardware) SetSerial (value string) {
	hardware.Serial = strings.TrimSpace(value);
}

func (hardware *Hardware) SetVendor (value string) {
	hardware.Vendor = value;
}

func (hardware *Hardware) SetVersion (value string) {

	if value == "all" || value == "any" || value == "*" {
		hardware.Version = "any";
	} else if value != "" {
		hardware.Version = value;
	}

}

