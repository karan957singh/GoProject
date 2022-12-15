package main

import (
	"os/exec"
)

func main() {
	// exec.Command("C:\\Program Files (x86)\\BraveSoftware\\Brave-Browser\\pplication\\brave.exe").Start()
	// exec.Command("rundll32", "url.dll,FileProtocolHandler", "C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	// exec.Command("C:\\Users\\karan\\AppData\\Local\\Microsoft\\WindowsApps\\debian.exe").Start()
	// exec.Command("C:\\Program Files\\Git\\git-bash.exe").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2019.1.3\\bin\\phpstorm64.exe", "D:\\myflightsonline").Start()
	// exec.Command("C:\\Program Files\\JetBrains\\DataGrip 2019.1.3\\bin\\datagrip64.exe").Start()
	exec.Command("C:\\Users\\karan\\AppData\\Local\\SourceTree\\SourceTree.exe").Start()

	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://youtube.com").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://mail.google.com/mail/u/0/#inbox").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/1/#inbox").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://soundcloud.com").Start()

	exec.Command("C:\\Users\\karan\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe", "D:\\myflightsonline\\server").Start()
	exec.Command("C:\\xampp_new_1\\xampp-control.exe").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2019.1.3\\bin\\phpstorm64.exe", "D:\\myflightsonline\\client").Start()
}
