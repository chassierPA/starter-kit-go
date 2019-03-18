package output

import (
	"log"
	"os"
)

var (
	//Debug Logger var
	Debug *log.Logger
	//Info Logger var
	Info *log.Logger
	//Warning Logger var
	Warning *log.Logger
	//Error Logger var
	Error *log.Logger
)

//LogType type defined for type of log
type LogType int

var (
	// LTDebug Debug type
	LTDebug LogType = 1
	// LTInfo Infos type
	LTInfo LogType = 2
	// LTWarning Warning type
	LTWarning LogType = 3
	// LTError Error Type
	LTError LogType = 4
)

// Log static variable log initalized is true overwrise is fasle
var Log bool

// InitLog function used to init norm log
func InitLog() {
	debugHandle := os.Stdout
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr

	// print in purple and red
	Debug = log.New(debugHandle,
		"\033[91m[\033[95mDEBUG\033[91m]: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)

	// print in green
	Info = log.New(infoHandle,
		"\033[32mINFO: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)

	// print in yellow
	Warning = log.New(warningHandle,
		"\033[33mWARNING: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)

	//print in red
	Error = log.New(errorHandle,
		"\033[31mERROR: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)
	Log = true
}

/*
func main() {
	InitLog()

	DisplayStart([]string{"Pierre-Alexandre CHASSIER"}, "0.0.1")
	fmt.Println()
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
	Debug.Println("Debuging Information")
}
*/
