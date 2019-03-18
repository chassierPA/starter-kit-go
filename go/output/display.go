package output

//DisplayStart display on screen a text when the programme start
func DisplayStart(fullName []string, version string) {
	Info.Println("\033[31mDEV(S): \033[32m", fullName)
	Info.Println("\033[31mVERSION: \033[32m", version)
	Info.Println("\033[34m###################################################################################")
	Info.Println("\033[34m#######\033[37m          STARTING        \033[34m##################################################")
	Info.Println("\033[34m###################################################################################")
	Info.Println("\033[34m##################################################################################")
	Info.Println("\033[34m##################################################################################\033[0m")
}
