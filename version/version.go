package version

import (
	"strings"
	"strconv"
	"os/exec"
	"regexp"
)


// ----------------------------------------------------------------------------------
//  Typen
// ----------------------------------------------------------------------------------

type Version struct {
	/** Parsete Version. */
	versionInfos []int
}


// ----------------------------------------------------------------------------------
//  Konstruktoren
// ----------------------------------------------------------------------------------

func FromString(version string) (*Version) {
	// split the parts of the versions
	var versions []int
	for _, piece := range strings.Split(version, ".") {
		// convert string to int and append to array
		partInt, err := strconv.Atoi(piece)
		if (err == nil) {
			versions = append(versions, partInt)
		}
		
	}

	return &Version { versionInfos: versions }
}


// ----------------------------------------------------------------------------------
//  Auskunfsmethoden
// ----------------------------------------------------------------------------------

// TODO: Not working with different version length!
func (this *Version) IsOlderThan(other *Version) (bool) {
	equals := false

	// comapre all parts of the version
	for index, piece := range this.versionInfos {
		// if one piece is newer we are finished
		if (piece < other.versionInfos[index]) {
			return true

		} else if (piece == other.versionInfos[index]) {
			equals = true
		}
	}

	// non of the parts was smaller than other
	// so this version is older
	return !equals
}

func (this *Version) String() string {
	version := ""
	for index, piece := range this.versionInfos {
		version += strconv.Itoa(piece)
		if (index < (len(this.versionInfos) - 1)) {
			version += "."
		}
	}

	return version
}

func InstalledGoVersion() (*Version) {
	// exectue go version command
	out, err := exec.Command("go", "version").Output()
	if (err != nil) {
		return nil
	}

	// Installierte GO Version rausfischen
	versionRegex := regexp.MustCompile(`go([\d|.]+)`)
	versionString := versionRegex.FindAllStringSubmatch(string(out), -1)

	// commands output returned a valid text
	if (len(versionString) > 0) {
		return FromString(versionString[0][1])

	// regex pattern was not found in
	// the commands output
	} else {
		return nil
	}
}