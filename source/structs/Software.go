
package structs;

import "winnow/console";
import "bytes";
import "strings";



func IsArchitecture (architecture string) bool {

	var ARCHITECTURES = []string{
		"any",
		"x86",
		"x86_64",
		"armv6",
		"armv7",
		"armv8",
		"sparc",
		"sparc64",
	};

	var result bool = false;

	for a := 0; a < len(ARCHITECTURES); a++ {

		var other = ARCHITECTURES[a];
		if other == architecture {
			result = true;
			break;
		}

	}

	return result;

}

func IsManager (manager string) bool {

	var MANAGERS = []string{

		"apk",       // Alpine
		"apt",       // Debian, Ubuntu
		"dnf",       // Alma, RedHat Enterprise
		"dpkg",      // Debian, Ubuntu
		"pacman",    // Arch
		"qpkg",      // QNX
		"yum",       // Amazon, Fedora
		"zypper",    // OpenSUSE, SUSE Enterprise

		"cargo",      // Rust
		"chocolatey", // C#
		"cocoapods",  // Cocoa / iOS
		"composer",   // PHP
		"conan",      // C++
		"conda",      // Python
		"cran",       // R
		"gem",        // Ruby
		"golang",     // Go
		"hackage",    // Haskell
		"hex",        // Erlang
		"maven",      // Java
		"msi",        // Microsoft Installer
		"npm",        // node.js
		"nuget",      // C#
		"pear",       // PHP
		"phar",       // PHP
		"pypi",       // Python
		"pub",        // Dart

		"docker",     // Docker
		"generic",

	};

	var result bool = false;

	for m := 0; m < len(MANAGERS); m++ {

		var other = MANAGERS[m];

		if other == manager {
			result = true;
			break;
		}

	}

	return result;

}



type Software struct {

	Name         string `json:"name"`;
	Version      string `json:"version"`;
	Architecture string `json:"architecture"`;
	Manager      string `json:"manager"`;
	Vendor       string `json:"vendor"`;

}

func NewSoftware () Software {

	var software Software;

	software.Version      = "any";
	software.Architecture = "any";
	software.Manager      = "any";
	software.Vendor       = "";

	return software;

}



func IsIdenticalSoftware (a Software, b Software) bool {

	var result bool = false;

	if a.Name == b.Name && a.Version == b.Version && a.Architecture == b.Architecture && a.Manager == b.Manager && a.Vendor == b.Vendor {
		result = true;
	}

	return result;

}

func IsSoftware (software Software) bool {

	var result bool = true;

	if software.Name != "" {

		if software.Vendor == "" {
			result = false;
		}

	} else {
		result = false;
	}

	return result;

}

func (software *Software) Matches (search Software) bool {

	var matches_name         bool = false;
	var matches_version      bool = false;
	var matches_architecture bool = false;
	var matches_manager      bool = false;
	var matches_vendor       bool = false;

	if software.Name == search.Name {
		matches_name = true;
	}

	if software.Version == "any" || search.Version == "any" {
		matches_version = true;
	} else {

		// TODO: software.Version = "<= 1.2.3"
		// TODO: software.Version = "< 1.2.3"
		// TODO: search.Version = "<= 1.2.3"
		// TODO: search.Version = "< 1.2.3"

	}

	if software.Architecture == "any" || search.Architecture == "any" {
		matches_architecture = true;
	} else if software.Architecture == search.Architecture {
		matches_architecture = true;
	}

	if software.Manager == "any" || search.Manager == "any" {
		matches_manager = true;
	} else if software.Manager == search.Manager {
		matches_manager = true;
	}

	if software.Vendor == "" || search.Vendor == "" {
		matches_vendor = true;
	} else if software.Vendor == search.Vendor {
		matches_vendor = true;
	}

	return matches_name && matches_version && matches_architecture && matches_manager && matches_vendor;

}

func (software *Software) ToLog () string {

	var buffer bytes.Buffer;

	buffer.WriteString("Hardware({\n");
	buffer.WriteString("    Name: \""         + software.Name         + "\",\n");
	buffer.WriteString("    Version: \""      + software.Version      + "\",\n");
	buffer.WriteString("    Architecture: \"" + software.Architecture + "\",\n");
	buffer.WriteString("    Manager: \""      + software.Manager      + "\",\n");
	buffer.WriteString("    Vendor: \""       + software.Vendor       + "\"\n");
	buffer.WriteString("})");

	return buffer.String();

}

func (software *Software) Log () {

	var buffer = software.ToLog();
	var chunks = strings.Split(buffer, "\n");

	for c := 0; c < len(chunks); c++ {
		console.Log(chunks[c]);
	}

}

func (software *Software) SetArchitecture (value string) {

	if value == "all" || value == "any" || value == "*" {
		software.Architecture = "any";
	} else if value == "i386" || value == "i686" || value == "x32" || value == "x86" {
		software.Architecture = "x86";
	} else if value == "amd64" || value == "ia64" || value == "x64" || value == "x86_64" {
		software.Architecture = "x86_64";
	} else if value == "armel" || value == "armv6" {
		software.Architecture = "armv6";
	} else if value == "armhf" || value == "armv7" || value == "armv7h" {
		software.Architecture = "armv7";
	} else if value == "aarch64" || value == "armv8" {
		software.Architecture = "armv8";
	} else if value == "sparc" {
		software.Architecture = "sparc";
	} else if value == "sparc64" {
		software.Architecture = "sparc64";
	}

}

func (software *Software) SetManager (value string) {

	if IsManager(value) {
		software.Manager = value;
	}

}

func (software *Software) SetName (value string) {
	software.Name = strings.TrimSpace(value);
}

func (software *Software) SetVendor (value string) {
	software.Vendor = value;
}

func (software *Software) SetVersion (value string) {

	if value == "all" || value == "any" || value == "*" {
		software.Version = "any";
	} else if value != "" {
		software.Version = value;
	}

}

