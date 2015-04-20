package bowlerfile

import (
    "git.1vh.de/maximilian.pachl/bowler/version"
    "io/ioutil"
    "encoding/json"
    "strings"
)


// ----------------------------------------------------------------------------------
//  Typen
// ----------------------------------------------------------------------------------

type Bowlerfile struct {
    Name        string
    Description string
    Package     string

    MinGoVersion *version.Version

    Assets []string
}

// private Datatypes for Bowlerfile json parsing
type goConfig struct {
    MinVersion string `json:"min-version"`
}

type project struct {
    Name		string	`json:"name"`
    Description	string	`json:"description"`
    Package		string	`json:"package"`
}

type bowlerfile struct {
    Project project `json:"project"`
    Go goConfig `json:"go"`
    Assets []string `json:"assets"`
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

    // TODO: filter possible dangerous symbols in assets, like &&, to prevent shell command execution

    // return newly created object
    return &Bowlerfile{
        Name: decoded.Project.Name,
        Description: decoded.Project.Description,
        Package: decoded.Project.Package,
        MinGoVersion: version.FromString(decoded.Go.MinVersion),
        Assets: decoded.Assets}, nil
}


// ----------------------------------------------------------------------------------
//  Aendernde Funktionen
// ----------------------------------------------------------------------------------

func (this *Bowlerfile) Save(filePath string) (error) {
    // create json string from this
    buffer, err := json.MarshalIndent(bowlerfile{
        Project: project{
            Name: this.Name,
            Description: this.Description,
            Package: this.Package,
        },
        Go: goConfig {
            MinVersion: this.MinGoVersion.String(),
        },
        Assets: this.Assets,
    }, "", "\t")

    // check for json error
    if (err != nil) {
        return err
    }

    // write to Bowlerfile
    err = ioutil.WriteFile(filePath, buffer, 0755)
    return err
}