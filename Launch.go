package main

import (
	"os/exec"
)

func main() {
	// exec.Command("C:\\Program Files\\OpenVPN\\bin\\openvpn-gui.exe", "--config connect D:\\carbar\\vpn\\dev\\client.ovpn").Start()
	// exec.Command("rundll32", "url.dll,FileProtocolHandler", "C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	// exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe").Start()
	// exec.Command("C:\\Program Files\\Docker\\Docker\\Docker Desktop.exe").Start()
	exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\OpenVPN\\bin\\openvpn-gui.exe").Start()
	// exec.Command("C:\\Users\\karan\\AppData\\Local\\Microsoft\\WindowsApps\\TheDebianProject.DebianGNULinux_76v4gfsz19hv4\\debian.exe").Start()
	// exec.Command("C:\\Program Files\\Git\\git-bash.exe").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2019.1.3\\bin\\phpstorm64.exe", "D:\\carbar\\CarbarNew").Start()
	// exec.Command("C:\\Program Files\\JetBrains\\DataGrip 2019.1.3\\bin\\datagrip64.exe").Start()
	exec.Command("C:\\Users\\karan\\AppData\\Local\\slack\\slack.exe").Start()
	exec.Command("C:\\Users\\karan\\AppData\\Local\\SourceTree\\SourceTree.exe").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://mail.google.com/mail/u/0/#inbox").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://mail.google.com/mail/u/1/#inbox").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://mail.google.com/mail/u/2/#inbox").Start()
	// exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "https://carbar.atlassian.net/secure/RapidBoard.jspa?rapidView=1&projectKey=DE&view=planning.nodetail").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2019.1.3\\bin\\phpstorm64.exe", "D:\\myflightsonline\\server").Start()
	exec.Command("C:\\xampp_new_1\\xampp-control.exe").Start()
}
