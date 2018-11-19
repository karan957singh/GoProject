package main

import(
"os/exec"
)

func main(){
	// exec.Command("rundll32", "url.dll,FileProtocolHandler", "C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\Sublime Text 3\\sublime_text.exe").Start()
	exec.Command("C:\\Program Files\\Git\\git-bash.exe").Start()
	exec.Command("C:\\Program Files\\JetBrains\\PhpStorm 2018.2.5\\bin\\phpstorm64.exe","D:\\my_projects\\react_project").Start()
	
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://youtube.com").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/0/#inbox").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://mail.google.com/mail/u/1/#inbox").Start()
	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","https://soundcloud.com").Start()
}