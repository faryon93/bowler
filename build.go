package main

import (
    "github.com/faryon93/bowler/bowlerfile"
    "github.com/faryon93/bowler/version"

    "os"
    "os/exec"
    "strings"
    "fmt"
    "strconv"
)


// ----------------------------------------------------------------------------------
//  Konstanten
// ----------------------------------------------------------------------------------

// TODO: in ~/.bowlerrc auslagern
const ASSETS_OUTPUT_FILE = "assets__.go"


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func taskBuild(buildFile *bowlerfile.Bowlerfile) {
    //-----------------------------------------------------------------------------------
    /// Initial Checks

    // Check GO version
    BeginStepMessage("Checking GO version")
    installedVersion := version.InstalledGoVersion()
    if (installedVersion == nil) {
        EndStepMessageStr(TYPE_FAILED, "error fetching installed GO version")
        os.Exit(-1)
    }

    // the installed GO version is too old
    if (installedVersion.IsOlderThan(buildFile.MinGoVersion)) {
        EndStepMessageStr(TYPE_FAILED, "needing GO version " + buildFile.MinGoVersion.String() + " but found " + installedVersion.String())
        os.Exit(-2)

    // the installed Go version is newer or equals the required version -> carry on
    } else {
        EndStepMessageStr(TYPE_OKAY, "found " + installedVersion.String() + ", required: " + buildFile.MinGoVersion.String())
    }


    //-----------------------------------------------------------------------------------
    // Do stuff

    // create working directory
    BeginStepMessage("Creating working directory")
    err := createPackageBase(buildFile)
    if (err != nil) {
        // directory already exists
        if (strings.Contains(err.Error(), "exists")) {
            EndStepMessageStr(TYPE_SKIPPED, "")

        // error occoured while creating directory
        }else {
            EndStepMessage(err)
        }
		
    // Working direcotory successfully created
    } else {
        EndStepMessageStr(TYPE_OKAY, "")
    }
	
    // create output artifact directory
    BeginStepMessage("Creating artifact directory")
    err = os.MkdirAll("bin", 0777)
    EndStepMessage(err)	
	
    // build assets
    BeginStepMessage("Building assets")
    if (len(buildFile.Assets) > 0) {
        err, o := buildAssets(buildFile)
        if (err == nil) {
            EndStepMessageStr(TYPE_OKAY,  "included " + strconv.Itoa(len(buildFile.Assets)) + " assets")

        // building of assets failed
        } else {
            EndStepMessageStr(TYPE_FAILED, err.Error())
            fmt.Println(o)
            os.Exit(-1)
        }

    // the user has not specified assets to build
    } else {
        EndStepMessageStr(TYPE_SKIPPED, "No assets to build")
    }

    // fetch project dependencies
    BeginStepMessage("Fetching project dependencies")
    err, o := fetchDependencies(buildFile)
    if (err != nil) {
        EndStepMessageStr(TYPE_FAILED, err.Error())
        fmt.Println(o)
        os.Exit(-1)
    } else {
        EndStepMessageStr(TYPE_OKAY, "")
    }

    // Build the binary
    BeginStepMessage("Building project " + buildFile.Name)
    err, o = executeBuild(buildFile)
    if (err != nil) {
        EndStepMessageStr(TYPE_FAILED, err.Error())
        fmt.Println(o)
        os.Exit(-1)
    } else {
        EndStepMessageStr(TYPE_OKAY, "")
    }
}

func createPackageBase(config *bowlerfile.Bowlerfile) (error) {
    folders := strings.Split(config.Package, "/")

    // put package base path together
    basePath := ""
    for i := 0; i < (len(folders) - 1); i++ {
        basePath += "/" + folders[i]
    }

    // create nessessary folders
    err := os.MkdirAll(".bowler/src" + basePath, 0755)
    if (err != nil) {
        return err
    }

    // create symlic link for package
    err = os.Symlink(strings.Repeat("../", len(folders) + 1), ".bowler/src/" + config.Package)
    if (err != nil) {
        return err
    }

    return nil
}

func fetchDependencies(project *bowlerfile.Bowlerfile) (error, string) {
    pwd, _ := os.Getwd()

    // prepare env for go get command
    command := exec.Command("go", "get", project.Package)
    command.Env = []string{
        "GOBIN=" + pwd +"/bin",
        "GOPATH=" + pwd + "/.bowler/",
        "PATH=" + os.Getenv("PATH"),
        "HOME=" + os.Getenv("HOME")}

    // run command 
    out, err := command.CombinedOutput()
    return err, string(out)
}

func executeBuild(project *bowlerfile.Bowlerfile) (error, string) {
    pwd, _ := os.Getwd()

    // exectue go version command
    // TODO: use -ldflags "-X ..." to implement a automatically updated version file
    command := exec.Command("go", "build", "-o", "bin/" + project.Name, project.Package)
    command.Env = []string{
        "GOBIN=" + pwd +"/bin",
        "GOPATH=" + pwd + "/.bowler/",
        "PATH=" + os.Getenv("PATH")}

    out, err := command.CombinedOutput()
    return err, string(out)
}

func buildAssets(project *bowlerfile.Bowlerfile) (error, string) {
    pwd, _ := os.Getwd()
    command := exec.Command("go-bindata", "-o", ASSETS_OUTPUT_FILE, strings.Join(project.Assets, " "))
    command.Env = []string{
        "GOBIN=" + pwd +"/bin",
        "GOPATH=" + pwd + "/.bowler/",
        "PATH=" + os.Getenv("PATH")}

    out, err := command.CombinedOutput()
    return err, string(out)
}