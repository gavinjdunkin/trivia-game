package main

import (
	"log"
	"net/http"
	"sync"
	"encoding/json"
	"flag"
	
    "io/ioutil"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// Game represents the state of the trivia game
type Game struct {
	users map[string]int
	host *websocket.Conn
	question_idx int
	question_count int
	questions Questions
	game_code string
	// Add fields to represent game state, such as questions, scores, etc.
	sync.Mutex

}
type Questions struct {
	Questions []Question `json:"questions"`
}


type Question struct {
	Question string `json:"question`
	Options []string `json:"options"`
	Correct_Answer string `json:"correct_answer"`
}

var game = Game{
	users: make(map[string]int),
	question_idx: 0,

	question_count: 0,
	game_code: "aaaa",
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleHost(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conn.WriteJSON(map[string]interface{}{"type": "code", "code": "aaaa"})
	game.host = conn
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Handle the join game event
	var username string
	if err := conn.ReadJSON(&username); err != nil {
		log.Println(err)
		return
	}
	game.users[username] = 0
	log.Println(username)


	// Add the user to the game state
	// Broadcast the updated state to all clients

	for {
		/**game_state := struct {
			options []string `json:"options"`
			score int `json:"score"`
			place int `json:"place"`
			time int `json:"time"`
		}{
			options: game.questions[game.question_idx].options,
			score: game.users[username],
			place: 1,
			time: 1000,
		}*/
		// Send the JSON data as a message
		//log.Println(game.questions)
		err = game.host.WriteJSON(map[string]interface{}{"type": "question", "question": game.questions.Questions[game.question_idx].Question})
		if err != nil {
			log.Println()
			log.Println(err)
			return
		}


		err = conn.WriteJSON(map[string]interface{}{"type": "question", "options":  game.questions.Questions[game.question_idx].Options})
		if err != nil {
			log.Println(err)
			return
		}
	
		// Handle messages from the client (e.g., answer question)
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(message)
		var result = false
		if string(message) == game.questions.Questions[game.question_idx].Correct_Answer {
			game.users[username] = game.users[username] + 1
			result = true

		}
		conn.WriteJSON(map[string]interface{}{"type": "result", "correct": result, "score": game.users[username]})
		game.host.WriteJSON(map[string]interface{}{"type":"result", "leaderboard": game.users})
		time.Sleep(3 * time.Second)
		game.question_idx++
	
	}
}

var addr = flag.String("addr", ":3000", "http service address")

func main() {
	flag.Parse()
	// Open our jsonFile
    jsonFile, _ := os.Open("questions.json")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "player.html")
	})

	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "host.html")

	})
	http.HandleFunc("/hostws", handleHost)
	defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
	var questions Questions

	json.Unmarshal(byteValue, &questions)
	game.questions = questions

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
