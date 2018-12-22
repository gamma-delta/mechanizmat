package game_classes

import (

    "github.com/bwmarrin/discordgo"
    "fmt"
    "encoding/json"
    "path/filepath"
    "io/ioutil"
)

type Game struct {
    Players map[string]*Player `json:"players"` //User ID to player object
}

var Mechanizmat Game

func init() {
    Mechanizmat.Players = make(map[string]*Player)
}

func (game *Game) GetPlayer(user *discordgo.User) *Player {
    player, ok := game.Players[user.ID]
    if ok {
        return player
    } else {
        return nil
    }
}

func (game *Game) Register(user *discordgo.User) (bool) {
    _, ok := game.Players[user.ID]
    if ok {
        return false //doesn't work
    } else {
        game.Players[user.ID] = NewPlayer(user.Username)
        return true
    }
}

func (game *Game) Unregister(user *discordgo.User) bool {
    _, ok := game.Players[user.ID]
    if ok {
        delete(game.Players, user.ID)
        return true //managed to delete
    } else {
        return false
    }
}

func (game *Game) Save() {
    json_bytes, err := json.Marshal(game)
    if err != nil {
        fmt.Println("!!ERROR marshalling game to JSON:", err)
        return
    }
    abs_path, _ := filepath.Abs("storage/game_data.json")

    err = ioutil.WriteFile(abs_path, json_bytes, 0644)
    if err != nil {
        fmt.Println("!!ERROR writing to game file:", err)
        return
    }
    fmt.Println("Game saved successfully.")
}

func (game *Game) Load() {
    abs_path, _ := filepath.Abs("storage/game_data.json")
    raw_data, err := ioutil.ReadFile(abs_path)
    if err != nil {
        fmt.Println("!!ERROR loading game file:", err)
        return
    }

    err = json.Unmarshal(raw_data, &game)
    if err != nil {
        fmt.Println("!!ERROR marshalling game data to struct:", err)
        return
    }
    fmt.Println("Game loaded successfully.")
}