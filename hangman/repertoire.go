package hangman

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type HangManData struct {
	Pendu      string
	Word       string
	ToFind     string
	Life       int
	Stockage   []string
	Input      string
	Username   string
	Text       string
	Difficulty float64
	Score      int
	Sl         int
	Nmax       float64
	Nl         float64
}

type RankData struct {
	Username string `json : "Username"`
	Score    int    `json : "Score"`
}
type ScoreBoard struct {
	User []string
}

func Game(x HangManData) (HangManData, int) {
	x.Input = ToUpper(x.Input)
	tmp := x.ToFind
	if len(x.Input) > 0 {
		if len(x.Input) >= 2 {
			if x.Input == x.Word {
				x.Score = Score(x)
				if x.Username != ".Anonymous" {
					OpenRank(x.Username, x)
				}
				x.Life, x.Text = 0, "You win"
				x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
				return x, 0
			} else {
				if x.Life > 2 {
					x.Life -= 2
					x.Text = "Wrong word"
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 2
				} else {
					x.Life, x.Text = 0, "You lost"
					x.Score = Score(x)
					if x.Username != ".Anonymous" {
						OpenRank(x.Username, x)
					}
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 3
				}
			}
		} else if Memelettre(x.Input, &x.Stockage) {
			if VerifLettre(x.Input, &x.Word, &x.ToFind) {
				temp := ""
				for _, letter := range x.ToFind {
					temp += string(letter)
				}
				if temp == x.Word {
					x.ToFind = tmp
					x.Score = Score(x)
					if x.Username != ".Anonymous" {
						OpenRank(x.Username, x)
					}
					x.Life, x.Text = 0, "You win"
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 0
				} else {
					x.Text = "You find a letter"
					x.Sl = int(float64(x.Sl) + (float64(x.Nl) * x.Difficulty * (float64(x.Life) / 10)) + 0.9)
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 1
				}
			} else {
				x.Life -= 1
				if x.Life == 0 {
					x.Text = "You lost"
					x.Score = Score(x)
					if x.Username != ".Anonymous" {
						OpenRank(x.Username, x)
					}
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 3
				} else {
					x.Text = "Wrong letter"
					x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
					return x, 2
				}
			}
		} else {
			x.Text = "Letter already used"
			x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
			return x, 1
		}
	}
	x.Pendu = "../images/Pendu/pendu_" + strconv.Itoa(x.Life) + ".png"
	return x, -1
}

func ToUpper(s string) string {
	temp := ""
	for _, lettre := range s {
		if lettre == 'à' || lettre == 'â' || lettre == 'ä' {
			temp += "A"
		} else if lettre == 'é' || lettre == 'è' || lettre == 'ê' || lettre == 'ë' {
			temp += "E"
		} else if lettre == 'ï' || lettre == 'î' {
			temp += "I"
		} else if lettre == 'ô' || lettre == 'ö' {
			temp += "O"
		} else if lettre == 'ù' || lettre == 'û' || lettre == 'ü' {
			temp += "U"
		} else if lettre == 'ÿ' {
			temp += "Y"
		} else if lettre == 'ç' {
			temp += "C"
		} else {
			temp += string(lettre)
		}
	}
	word, result := []rune(temp), ""
	for index := range temp {
		if word[index] >= 97 && word[index] <= 122 {
			word[index] = word[index] - 32
			result += string(word[index])
		} else if word[index] >= 65 && word[index] <= 90 {
			result += string(word[index])
		}
	}
	return result
}

func CreateWord(word string) string {
	wordtofind := []string{}
	for k := 0; k < len([]rune(word)); k++ {
		if word[k] == '-' {
			wordtofind = append(wordtofind, "-")
		} else {
			wordtofind = append(wordtofind, "_")
		}
	}
	for i := 0; i < (len(word)/2 - 1); i++ {
		tempr := AleatoireNbr(len(word) - 1)
		if wordtofind[tempr] == "_" {
			wordtofind[tempr] = string([]rune(word)[tempr])
		} else {
			i--
		}
	}
	myrep := ""
	for _, letter := range wordtofind {
		myrep += letter
	}
	return myrep
}

func Chooseword(s string) string {
	f, _ := os.OpenFile(s, os.O_RDWR, 0644)
	scanner, listemots := bufio.NewScanner(f), []string{}
	for scanner.Scan() {
		listemots = append(listemots, scanner.Text())
	}
	nombrealea := AleatoireNbr(len(listemots) - 1)
	return ToUpper(listemots[nombrealea])
}

func AleatoireNbr(n int) int {
	aleaint := 0
	if n >= 1 {
		rand.Seed(time.Now().UnixNano())
		aleaint = rand.Intn(n) + 1
	}
	return aleaint
}

func VerifLettre(input string, mot *string, motcache *string) bool {
	bon, tmp := false, ""
	for index, lettre := range *mot {
		if input == string(lettre) && input != string((*motcache)[index]) {
			tmp += input
			bon = true
		} else {
			tmp += string((*motcache)[index])
		}
	}
	*motcache = tmp
	return bon
}

func Memelettre(input string, stockage *[]string) bool {
	for _, lettre := range *stockage {
		if input == lettre {
			return false
		}
	}
	*stockage = append(*stockage, input)
	return true
}

func Code(data HangManData, namesave string) {
	path := "saves/" + data.Username
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	path += "/" + namesave + ".txt"
	b, _ := json.Marshal(data)
	f, _ := os.Create(path)
	f.Write(b)
	f.Close()
}

func Decode(savename string, username string) HangManData {
	path := "saves/"
	if len(username) > 0 {
		path += username
	} else {
		path += ".Anonymous"
	}
	savename = path + "/" + savename + ".txt"
	f, err := os.OpenFile(savename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := os.ReadFile(savename)
	dataSave := HangManData{}
	_ = json.Unmarshal(b, &dataSave)
	f.Close()
	DeleteSave(savename)
	files, _ := ioutil.ReadDir(path)
	if len(files) < 1 {
		DeleteSave(path)
	}
	dataSave.Username = username
	return dataSave
}

func DeleteSave(savename string) {
	if _, err := os.Stat(savename); err == nil {
		os.Remove(savename)
	}
}

func OpenRank(username string, x HangManData) {
	if username == ".Anonymous" {
		return
	}
	f, err := os.OpenFile("score/rank.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	scanner, tableau, test := bufio.NewScanner(f), []RankData{}, true
	for scanner.Scan() {
		dRank := RankData{}
		_ = json.Unmarshal(scanner.Bytes(), &dRank)
		tableau = append(tableau, dRank)
	}
	f.Close()
	for index, player := range tableau {
		if player.Username == username {
			if x.Score > tableau[index].Score {
				tableau[index].Score = x.Score
			}
			test = false
			break
		}
	}
	if test {
		newuser := RankData{username, x.Score}
		tableau = append(tableau, newuser)
	}
	for index := range tableau {
		minindex := index
		for i := index; i < len(tableau); i++ {
			if tableau[minindex].Score < tableau[i].Score {
				minindex = i
			}
		}
		tableau[index], tableau[minindex] = tableau[minindex], tableau[index]
	}
	f, _ = os.OpenFile("score/rank.txt", os.O_RDWR|os.O_TRUNC, 0777)
	for index := range tableau {
		b, _ := json.Marshal(tableau[index])
		f.Write(b)
		f.WriteString("\n")
	}
	f.Close()
}

func AfficheRank() ScoreBoard {
	var tableautmp []string
	var tableau []string
	f, _ := os.OpenFile("score/rank.txt", os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dRank := RankData{}
		_ = json.Unmarshal(scanner.Bytes(), &dRank)
		tmp := dRank.Username + " : "
		long := 15 - len(tmp)
		for i := 0; i < long; i++ {
			tmp += " "
		}
		tmp += strconv.Itoa(dRank.Score)
		tableautmp = append(tableautmp, tmp)
	}
	f.Close()
	for index, valeur := range tableautmp {
		if index < 5 {
			tableau = append(tableau, valeur)
		} else {
			break
		}
	}
	return ScoreBoard{tableau}
}

func Score(x HangManData) int {
	Cv, Tmax, T := float64(x.Life)/10, float64(int(len(x.Word)-(len(x.Word)/2)+1)), float64(0)
	for _, lettre := range x.ToFind {
		if lettre == 95 {
			T++
		}
	}
	Nm := ((T*(x.Nmax-x.Nl)/(Tmax-1) + x.Nl - (x.Nmax-x.Nl)/(Tmax-1)) * Cv * x.Difficulty) + 0.9 + float64(x.Sl)
	Score := int(Nm)
	return Score
}
