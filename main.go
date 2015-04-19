package main

import "git.1vh.de/maximilian.pachl/bowler/bowlerfile"
import "os"
import "fmt"


// ----------------------------------------------------------------------------------
//  Konstanten
// ----------------------------------------------------------------------------------

/** Name des Buildfiles. */
const BUILD_FILE_NAME = "Bowlerfile"


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func main() {
	// Load the Bowlerfile
	BeginStepMessage("Loading Bowlerfile")
	buildFile, err := bowlerfile.Load(BUILD_FILE_NAME)
	EndStepMessage(err)

	if (len(os.Args) >= 2) {
		// the user requests us to build the project
		if (os.Args[1] == "build") {
			build(buildFile)	

		// clean the project directory
		} else if (os.Args[1] == "clean") {
			BeginStepMessage("Cleaning project root")
			err = clean(buildFile)
			EndStepMessage(err)

		// i don't know what the user want form me :O
		} else {
			fmt.Println("ERROR: invalid command: " + os.Args[1])
		}

	// the subcommand was not provided -> error	
	} else {
		fmt.Println("ERROR: missing command")
	}
}