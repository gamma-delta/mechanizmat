package main

import (
    // Local packages
    "mechanizmat/key"
    "mechanizmat/commands"
    "mechanizmat/game_classes"

    // System packages
    "fmt"
    "os" //No idea what this is for
    "os/signal" //Something about waiting?
    "syscall" //Who knows
    "github.com/bwmarrin/discordgo"
)

func main() {
    fmt.Println("Loading game...")
    game_classes.Mechanizmat.Load()

    fmt.Println("Creating session...")
    session, err := discordgo.New(key.key)
    if err != nil {
        fmt.Println("!!ERROR creating session: ", err)
        return
    }
    fmt.Println("Session created.")
    
    //Register command switchboards
    session.AddHandler(commands.CommandSwitchboard)
    fmt.Println("Added handlers.")

    fmt.Println("Connecting to Discord...")
    err = session.Open()
    if err != nil {
        fmt.Println("!!ERROR opening connection: ", err)
        return
    }
    fmt.Printf("Logged in as %s#%s\n", session.State.User.Username, session.State.User.Discriminator)

    fmt.Println("Mechanizmat is ready to rumble.\n")

    //Wait here or something?
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    //Cleanly close
    fmt.Println("Saving and closing...")
    game_classes.Mechanizmat.Save()
    session.Close()
}