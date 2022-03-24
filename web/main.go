package main

import (
	"fmt"
	"hangman/hangman"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Save struct {
	Filename []string
	Name     string
}

func main() {
	game := template.Must(template.ParseFiles("html/game.html"))
	gameWin := template.Must(template.ParseFiles("html/gameWin.html"))
	gameLose := template.Must(template.ParseFiles("html/gameLose.html"))
	chooseSave := template.Must(template.ParseFiles("html/save/chooseSave.html"))
	successSave := template.Must(template.ParseFiles("html/save/successSave.html"))
	noSave := template.Must(template.ParseFiles("html/save/noSave.html"))
	rank := template.Must(template.ParseFiles("html/rank.html"))
	var code int
	var GameInfo hangman.HangManData
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "html/login.html")
		case "POST":
			if len(r.FormValue("Username")) > 0 {
				GameInfo.Username = hangman.ToUpper(r.FormValue("Username"))
			} else {
				GameInfo.Username = ".Anonymous"
			}
			http.ServeFile(w, r, "html/index.html")
		}
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			if r.FormValue("STOP") == "STOP" {
				http.ServeFile(w, r, "html/save/makeSave.html")
			} else if len(r.FormValue("PLAY")) != 0 {
				switch r.FormValue("PLAY") {
				case "EASY":
					GameInfo.Word = hangman.Chooseword("dictionary/words.txt")
					GameInfo.Difficulty = 0.33
				case "MEDIUM":
					GameInfo.Word = hangman.Chooseword("dictionary/words2.txt")
					GameInfo.Difficulty = 0.67
				case "HARD":
					GameInfo.Word = hangman.Chooseword("dictionary/words3.txt")
					GameInfo.Difficulty = 1
				}
				GameInfo.ToFind, GameInfo.Life, GameInfo.Nmax, GameInfo.Nl, GameInfo.Sl = hangman.CreateWord(GameInfo.Word), 10, 100, 10, 0
				GameInfo, code = hangman.Game(GameInfo)
				game.Execute(w, GameInfo)
			} else if GameInfo.Input = r.FormValue("input"); len(GameInfo.Input) > 0 {
				GameInfo, code = hangman.Game(GameInfo)
				switch code {
				case 0:
					gameWin.Execute(w, GameInfo)
					GameInfo = hangman.HangManData{}
				case 1:
					game.Execute(w, GameInfo)
				case 2:
					game.Execute(w, GameInfo)
				case 3:
					gameLose.Execute(w, GameInfo)
					GameInfo = hangman.HangManData{}
				}
			} else {
				game.Execute(w, GameInfo)
			}
		}
	})
	http.HandleFunc("/loadSave", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			if r.FormValue("SAVE") == "SAVE" && GameInfo.Username != ".Anonymous" {
				var filename Save
				direction := "saves/" + GameInfo.Username
				if _, err := os.Stat(direction); !os.IsNotExist(err) {
					files, _ := ioutil.ReadDir(direction)
					var filelist []string
					for _, file := range files {
						tmp := file.Name()
						tmp = tmp[:len(tmp)-4]
						filelist = append(filelist, tmp)
					}
					filename.Filename = filelist
					chooseSave.Execute(w, filename)
				} else {
					noSave.Execute(w, GameInfo)
				}
			} else if r.FormValue("SAVE") == "SAVE" {
				http.ServeFile(w, r, "html/login.html")
			} else if input := r.FormValue("input"); len(input) > 0 {
				GameInfo = hangman.Decode(input, GameInfo.Username)
				GameInfo, code = hangman.Game(GameInfo)
				game.Execute(w, GameInfo)
			}
		}
	})
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			if namesave := r.FormValue("input"); len(namesave) > 0 && GameInfo.Username != ".Anonymous" {
				GameInfo.Input = ""
				hangman.Code(GameInfo, namesave)
				savename := Save{Name: namesave}
				successSave.Execute(w, savename)
			} else {
				http.ServeFile(w, r, "html/login.html")
			}
		}
	})
	http.HandleFunc("/rank", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			scoreboard := hangman.AfficheRank()
			rank.Execute(w, scoreboard)
		}
	})
	css, img := http.FileServer(http.Dir("./css")), http.FileServer(http.Dir("./images"))
	http.Handle("/css/", http.StripPrefix("/css/", css))
	http.Handle("/images/", http.StripPrefix("/images/", img))
	fmt.Printf("Starting server http://localhost:8080/\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
