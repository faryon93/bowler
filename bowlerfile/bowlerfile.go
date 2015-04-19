package bowlerfile

import "git.1vh.de/maximilian.pachl/bowler/version"
import "io/ioutil"
import "encoding/json"
import "strings"


// ----------------------------------------------------------------------------------
//  Typen
// ----------------------------------------------------------------------------------

type Bowlerfile struct {
	Name		string
	Description	string
	Package		string

	MinGoVersion *version.Version

}

type bowlerfile struct {
	Project struct {	
		Name		string	`json:"name"`
		Description	string	`json:"description"`
		Package		string	`json:"package"`
	} `json:"project"`

	Go struct {
		MinVersion	string 	`json:"min-version"`
	} `json:"go"`
}


// ----------------------------------------------------------------------------------
//  Konstruktoren
// ----------------------------------------------------------------------------------

func Load(filePath string) (*Bowlerfile, error) {
	// read the Bowlerfile
	buffer, err := ioutil.ReadFile(filePath)

	//decode the Bowlerfile file
	var decoded bowlerfile
	decoder := json.NewDecoder(strings.NewReader(string(buffer)))
	err = decoder.Decode(&decoded)
	if (err != nil) {
		return nil, err
	}

	// return newly created object
	return &Bowlerfile{
		Name: decoded.Project.Name,
		Description: decoded.Project.Description,
		Package: decoded.Project.Package,
		MinGoVersion: version.FromString(decoded.Go.MinVersion)}, nil
}