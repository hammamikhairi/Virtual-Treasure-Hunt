package serveragent

import (
	scl "blood/SQLClient"
	utils "blood/Utils"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"net/http"
)

type ServerAgent struct {
	sqlc *scl.SQLClient
}

func ServerAgentInit() ServerAgent {
	return ServerAgent{
		scl.OpenConn(),
	}
}

type OKStatus struct {
	Msg string
}

func (ag ServerAgent) GetPlayers(w http.ResponseWriter, req *http.Request) {
	utils.Log("players")

	// if req.Method != http.MethodPost {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	rows := ag.sqlc.GetRows(
		fmt.Sprintf(
			"select * from %s order by LEVEL DESC,  UPDATED_AT DESC",
			scl.TBN_PLAYERS,
		),
	)
	defer rows.Close()

	players := []Player{}

	for rows.Next() {
		player := Player{}
		rows.Scan(&player.Player_Id, &player.Email, &player.Player, &player.Level, &player.Updated_At)
		players = append(players, player)
	}

	// ag.sqlc.PushRows([][]any{logRow}, scl.TBN_TRANS)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(players)
}

type Player struct {
	Player_Id  string    `json:"player_id"`
	Email      string    `json:"email"`
	Player     string    `json:"username"`
	Level      int       `json:"level"`
	Updated_At time.Time `json:"updated_at"`
}

func (ag *ServerAgent) getPlayerData(playerId string) *Player {
	rows := ag.sqlc.GetRows(
		fmt.Sprintf(
			"select * from %s where PLAYER_ID = '%s'",
			scl.TBN_PLAYERS,
			playerId,
		),
	)
	defer rows.Close()
	rows.Next()

	var player *Player = &Player{}
	if err := rows.Scan(&player.Player_Id, &player.Email, &player.Player, &player.Level, &player.Updated_At); err != nil {
		fmt.Println(err)
		return nil
	}

	return player
}

func (ag *ServerAgent) levelUpPlayer(player *Player) error {
	ag.sqlc.Query(
		fmt.Sprintf(
			"UPDATE %s set LEVEL = %d, UPDATED_AT = DATEADD(hour, 1, GETDATE())  where PLAYER_ID = '%s'",
			scl.TBN_PLAYERS,
			player.Level+1,
			player.Player_Id,
		),
	)
	return nil
}

func (ag *ServerAgent) GetPlayer(w http.ResponseWriter, req *http.Request) {
	playerId := req.URL.Query().Get("playerId")
	utils.Log("get user <" + playerId + ">")

	player := ag.getPlayerData(playerId)
	fmt.Println(playerId)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if player == nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(player)
}

func (ag *ServerAgent) LevelUp(w http.ResponseWriter, req *http.Request) {
	playerId, flag, level := req.URL.Query().Get("playerId"), req.URL.Query().Get("flag"), req.URL.Query().Get("level")
	utils.Log("leveling up <" + playerId + "> with <" + flag + ">")

	player := ag.getPlayerData(playerId)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if val, _ := SOLUTIONS[level]; val == flag && fmt.Sprint(player.Level) == level {
		w.WriteHeader(http.StatusOK)
		ag.levelUpPlayer(player)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

var FILES = map[string]string{
	"1": "first.txt",
}

const FILES_DIR string = "./Files"

func (ag *ServerAgent) ServeFile(w http.ResponseWriter, req *http.Request) {
	level := req.URL.Query().Get("level")
	levelFilePath := FILES_DIR + "/" + levels[level].PUB_Level.Files[0]
	w.Header().Set("Access-Control-Allow-Origin", "*")
	file, err := os.Open(levelFilePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Unable to get file info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	http.ServeContent(w, req, fileInfo.Name(), fileInfo.ModTime(), file)
}

type PUB_Level struct {
	Message string
	Files   []string
	Title   string
	Images  []string
	Audio   string
}

type Level struct {
	PUB_Level
	Flag string
}

var SOLUTIONS = map[string]string{
	"1": "b5fJ2kRdLp1uH8W",
	"2": "4xhGzTnS3s2vA9K",
	"3": "m6NpCjL1bVz5aE7",
	"4": "KHAIRI",
	"5": "tKqX5bJy8WfPzL9",
	"6": "1qfzkbi6xhdjykl",
	"7": "b2WzK6cP8jRfH1T",
	"8": "5cH8nZ6xaMTjU9F",
	"9": "a8Jx6zT9uGyDcN3",
}

var levels = map[string]Level{
	"1": {
		PUB_Level: PUB_Level{
			Title:   "5c",
			Message: "The Flag Should be somewhere Here$$flag{b5fJ2kRdLp1uH8W}",
		},
		Flag: "b5fJ2kRdLp1uH8W",
	},
	"2": {
		PUB_Level: PUB_Level{
			Title:   "H8",
			Message: "Thats a lot of characters!",
			Files:   []string{"base64.txt"},
		},
		Flag: "4xhGzTnS3s2vA9K",
	},
	"3": {
		PUB_Level: PUB_Level{
			Title:   "nZ",
			Message: "What if the image you're looking at is not just an image? What if it's waiting to be decoded using a language of its own?",
			Images:  []string{"imc_logo.png"},
		},
		Flag: "m6NpCjL1bVz5aE7",
	},

	"4": {
		PUB_Level: PUB_Level{
			Title:   "6x",
			Message: "The hint is in this page (this time the flag is shorter and not inside a 'flag{}'):",
			Images:  nil,
			Files:   nil,
		},
		Flag: "KHAIRI",
	},
	"5": {
		PUB_Level: PUB_Level{
			Title:   "aM",
			Message: "Looks like this image is carrying some extra weight! It might've swallowed the other half... I wonder how you can extract the latter...",
			Files:   []string{"qr.png"},
		},
		Flag: "tKqX5bJy8WfPzL9",
	},
	"6": {
		PUB_Level: PUB_Level{
			Title:   "Tj",
			Message: "Decode the audio or video to get the flag.",
			Audio:   "Hmmmm.mp3",
		},
		Flag: "flag6",
	},
	"7": {
		PUB_Level: PUB_Level{
			Title:   "U9",
			Message: "You deserve a break. The flag will be found in one of my socials, I wonder which one... (khairi hammami)",
			Images:  nil,
			Files:   nil,
		},
		Flag: "flag8",
	},
	"8": {
		PUB_Level: PUB_Level{
			Title:   "F",
			Message: "Don't worry about the flag, it's probably enjoying its vacation on the previous levels. You know, the one you probably didn't care about...",
			Images:  nil,
			Files:   nil,
		},
		Flag: "flag9",
	},
	"9": {
		PUB_Level: PUB_Level{
			Title:   "Final",
			Message: "Oh no, a hacker stole your flag! Good luck getting it back!",
			Files:   []string{"system_logs.zip"},
		},
		Flag: "qsd",
	},
}

func (ag *ServerAgent) GetLevel(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	requestedLevel := req.URL.Query().Get("level")
	json.NewEncoder(w).Encode(levels[requestedLevel].PUB_Level)
}