package main

import (
	"fmt"
	"log"
	"os"
)

// Glog or Logrus (famous)
// github.com/sirupsen/logrus
// - Trace, Debug, Info, Warn, Error, Fatal, Panic i.e log.Trace(), log.Debug() etc
func main() {
	logFile := "/media/kako77sub/DA182A9E182A79A11/Users/uchiha-kako/projects/projects/go-advanced-kodekloud/log.txt"
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer f.Close()
	// logrus.SetOutput(f)
	log.SetOutput(f)
	log.Println("Some logs happening here")
}
