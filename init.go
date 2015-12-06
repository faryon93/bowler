package main

import (
    "github.com/faryon93/bowler/bowlerfile"
    "github.com/faryon93/bowler/version"
    "fmt"
    "os"
    "bufio"
    "strings"
    "path/filepath"
)


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func taskInit() {
    // No Bowlerfile in directory, create a new one
    if _, err := os.Stat(BUILD_FILE_NAME); err != nil {
        // ask user for the necessary information
        projectName := AskUser("project name", "")
        projectDescription := AskUser("project description", "")
        basePackage := AskUser("base package", "")
        goVersion := AskUser("required GO version", "1.4.2")

        // TODO: plausibility checks for entered data

        // Save the newly created Bowlerfile
        err = (&bowlerfile.Bowlerfile{
            Name: projectName,
            Description: projectDescription,
            Package: basePackage,

            MinGoVersion: version.FromString(goVersion),

            Assets: []string{},
        }).Save(BUILD_FILE_NAME)

        path, _ := filepath.Abs(BUILD_FILE_NAME)
        if (err == nil) {
            fmt.Println("Successfully created Bowlerfile: " + path)
        } else {
            fmt.Println("Error creating Bowlerfile: " + err.Error())
        }

    } else {
        fmt.Println("Bowlerfile already exists. Nothing to init here!")
    }
}

func AskUser(question string, defaultAnswer string) (string) {
    // print the question and default answer if provided
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\t" + question)
    if (len(defaultAnswer) > 0) {
        fmt.Print(" [" + defaultAnswer + "]")
    }
    fmt.Print(": ")

    // read the answer from user
    answer, _ := reader.ReadString('\n')

    // if only \n is ready back the user wants the default value
    // TODO: what happens with \r\n on windows?
    if (len(answer) == 1) {
        answer = defaultAnswer
    }

    // remove newlines and return the answer
    return strings.Replace(answer, "\n", "", -1)
}