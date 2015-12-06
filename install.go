package main

import (
    "github.com/faryon93/bowler/bowlerfile"
    "os"
    "io"
    "io/ioutil"
    "path/filepath"
)


// ----------------------------------------------------------------------------------
//  Funktionen
// ----------------------------------------------------------------------------------

func taskInstall(buildFile *bowlerfile.Bowlerfile) {
	BeginStepMessage("Installing project binary")

	// get the global GOPATH
	globalGoPath := os.Getenv("GOPATH")
	if (globalGoPath == "") {
		EndStepMessageStr(TYPE_FAILED, "Global GOPATH not set")
	}

	// Copy the binary to global GOPATH
	err := CopyFile(globalGoPath + "/bin/" + buildFile.Name, "bin/" + buildFile.Name, 0755)
	EndStepMessage(err)
}


// ----------------------------------------------------------------------------------
//  Hilfsfunktionen
// ----------------------------------------------------------------------------------

func CopyFile(dst, src string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "")
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err = tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}

	if err = os.Rename(tmp.Name(), dst); err != nil {
		os.Remove(tmp.Name())
		return err
	}

	return os.Chmod(dst, perm)
}