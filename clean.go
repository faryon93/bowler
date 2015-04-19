package main

import (
	"git.1vh.de/maximilian.pachl/bowler/bowlerfile"

	"os"
	"strings"
)


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func clean(project *bowlerfile.Bowlerfile) (error) {
	// remove output directory
	err := os.RemoveAll("bin")
	if (err != nil) {
		return err
	}

	// remove source symlink
	err = os.Remove(".bowler/src/" + project.Package)
	if (err != nil) {
		return err
	}

	// remove bowler working directory
	err = os.RemoveAll(".bowler")
	if (err != nil) {
		return err
	}

	// remove assets file
	err = os.Remove(ASSETS_OUTPUT_FILE)
	if (err != nil && !strings.Contains(err.Error(), "no such")) {
		return err
	}

	return nil
}