package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	srcFolder := "/Users/karan/data/docs"
	destFolder := "/Volumes/Private_01/DOCS_Backup" + time.Now().Format("2006_01_02_15_04_05")
	cmd := exec.Command("cp", "-rf", srcFolder, destFolder)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	fmt.Println("Completed")
}
