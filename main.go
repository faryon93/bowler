package main

import (
    "github.com/faryon93/bowler/bowlerfile"
    "os"
    "fmt"
    "path/filepath"
    "time"
)


// ----------------------------------------------------------------------------------
//  Konstanten
// ----------------------------------------------------------------------------------

/** Name des Buildfiles. */
const BUILD_FILE_NAME = "Bowlerfile"


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func main() {
    if (len(os.Args) >= 2) {
        startTime := time.Now()

        // the user requests us to build the project
        if (os.Args[1] == "build") {
            bowlerfile := loadBowlerfile()

            fmt.Println("Executing task 'build':")
            taskBuild(bowlerfile)

        // clean the project directory
        } else if (os.Args[1] == "clean") {
            bowlerfile := loadBowlerfile()

            // exectute clean task
            fmt.Println("Executing task 'clean':")	
            BeginStepMessage("Cleaning project root")
            err := taskClean(bowlerfile)
            EndStepMessage(err)

        } else if (os.Args[1] == "install") {
            bowlerfile := loadBowlerfile()
            fmt.Println("Executing task 'install':")  
            taskBuild(bowlerfile)
            taskInstall(bowlerfile)

        // initilaize a new project
        } else if (os.Args[1] == "init") {
            fmt.Println("Executing task 'init':")
            taskInit()

        } else if (os.Args[1] == "version") {
            fmt.Println("Bowler version 0.2 (rev3f4ab50e1)")
            os.Exit(0)

        // i don't know what the user want form me :O
        } else {
            fmt.Println("ERROR: invalid command: " + os.Args[1])
            os.Exit(-1)
        }

        fmt.Printf("\nExecution finished in %s\n", time.Since(startTime))

    // the subcommand was not provided -> error	
    } else {
        fmt.Println("ERROR: missing command")
    }
}

func loadBowlerfile() (*bowlerfile.Bowlerfile) {
    path, err := filepath.Abs(BUILD_FILE_NAME)
    if (err == nil) {
        fmt.Println("Using buildfile " + path)
        fmt.Println()
    }

    // Load the Bowlerfile
    buildFile, err := bowlerfile.Load(BUILD_FILE_NAME)
    if (err != nil) {
        fmt.Printf("Could not open Bowlerfile: %s\n", err)
        os.Exit(-3)
    }

    return buildFile
}