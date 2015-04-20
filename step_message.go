package main 

import (
    "fmt"
    "github.com/fatih/color"
    "os"
)


// ----------------------------------------------------------------------------------
//  Konstanten
// ----------------------------------------------------------------------------------

const TYPE_OKAY = 0
const TYPE_FAILED = 1
const TYPE_SKIPPED = 2


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func BeginStepMessage(message string) {
    fmt.Print("\t[....] " + message)
}

func EndStepMessage(err error) {
    if (err == nil) {
        EndStepMessageStr(TYPE_OKAY, "")
    } else {
        EndStepMessageStr(TYPE_FAILED, err.Error())
        os.Exit(-1)
    }
}

func EndStepMessageStr(t int, message string) {
    // create colors for 
    red := color.New(color.FgRed).SprintFunc()
    green := color.New(color.FgGreen).SprintFunc()
    yellow := color.New(color.FgYellow).SprintFunc()

    // save cursor position, got back to begin
    fmt.Print("\033[s\r\t[")

    // print message type with color
    if (t == TYPE_OKAY) {
        fmt.Print(green(" ok "))
    } else if (t == TYPE_FAILED) {
        fmt.Print(red("FAIL"))		
    } else if (t == TYPE_SKIPPED) {
        fmt.Print(yellow("skip"))	
    }

    // if a message was supplied print it
    if (message == "") {
        fmt.Printf("]\033[u\n")	
    } else {
        fmt.Printf("]\033[u: %s\n", message)	
    }	
}