package main

import(
"os/exec"
)

func main(){
	// exec.Command("rundll32", "url.dll,FileProtocolHandler", "C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\Git\\git-bash.exe").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2018.2.5\\bin\\phpstorm64.exe","D:\\carbar\\CarbarNew").Start()
	exec.Command("C:\\Program Files\\JetBrains\\DataGrip 2018.2.4\\bin\\datagrip64.exe").Start()
	exec.Command("C:\\Users\\karan\\AppData\\Local\\slack\\slack.exe").Start()
	exec.Command("C:\\Users\\karan\\AppData\\Local\\SourceTree\\app-3.0.8\\SourceTree.exe").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/0/#inbox").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/1/#inbox").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/2/#inbox").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://carbar.atlassian.net/secure/RapidBoard.jspa?rapidView=1&projectKey=DE&view=planning.nodetail").Start()
}