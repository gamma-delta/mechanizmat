// commands.go
//Holds all the commands players & admins can run.
//Player commands start with $ 
//Admin commands start with ^ 

package commands

import (
    "mechanizmat/game_classes"

    "fmt"
    "strings"
    "strconv"
    "github.com/bwmarrin/discordgo"
    "github.com/schollz/closestmatch"
)

type input_command func(*discordgo.Session, *discordgo.Message, []string) //NOTE these functions must not return anything, only send messages

var player_map = map[string]input_command{
    "ping": ping,
    "register": register,
    "unregister": unregister,
    "info": info,
    "tick_test": tick_test,
}

var admins = map[string]bool{ //Maps allow hashing, so fast access
    "273636822565912578": true,
}

var admin_map = map[string]input_command{
    "ping": a_ping,
    "user_info": a_user_info,
    "game_info": a_game_info,
    "save": a_save,
    "load": a_load,
}

func CommandSwitchboard(s *discordgo.Session, m *discordgo.MessageCreate) {
    //IGNORES
    if m.Author.ID == s.State.User.ID {return} //Ignore message if I wrote it
    prefix := string(m.Content[0])
    if prefix != "$" && prefix != "^" {return}
    //You should already restrict its scope to whatever channel you want it running on.

    //Alright, we got a working message

    //Start typing!
    if ok, _ := comesFromDM(s, m); ok {
        channel, _ := s.UserChannelCreate(m.Author.ID)
        s.ChannelTyping(channel.ID)
    } else {
        s.ChannelTyping(m.ChannelID)
    }

    split_message := strings.Fields(strings.ToLower(m.Content))
    cmd := split_message[0][1:] //everything but the first char
    args := split_message[1:] //all the other args

    if prefix == "$" {
        //this is a plain ol' command
        if _, ok := player_map[cmd]; ok {
            fmt.Printf("%s#%s ran command $%s%s\n", m.Author.Username, m.Author.Discriminator, cmd, args)
            player_map[cmd](s, m.Message, args) //run the command
            fmt.Print("\n")
        } else {
            no_command(s, m.Message, cmd)
        }
    } else if prefix == "^" { //better double check
        //this is an ADMIN command
        if _, ok := admins[m.Author.ID]; !ok {
            fmt.Printf("Non-admin %s#%s tried to run admin command ^%s%s!\n", m.Author.Username, m.Author.Discriminator, cmd, args)
            send(s, m.Message, fmt.Sprintf("AAAA @everyone %s tried to run an admin command!", m.Author.Username))
            return
        } 

        if _, ok := admin_map[cmd]; ok {
            fmt.Printf("%s#%s ran admin command ^%s%s\n", m.Author.Username, m.Author.Discriminator, cmd, args)
            admin_map[cmd](s, m.Message, args) //run the command
            fmt.Print("\n")
        } else {
            a_no_command(s, m.Message, cmd)
        }
    }
}

//magic internet code
func comesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
    channel, err := s.State.Channel(m.ChannelID)
    if err != nil {
        if channel, err = s.Channel(m.ChannelID); err != nil {
            return false, err
        }
    }

    return channel.Type == discordgo.ChannelTypeDM, nil
}

func send(s *discordgo.Session, m *discordgo.Message, msg string) {
    split_message := split_long(msg)
    for _, msg := range split_message {
        s.ChannelMessageSend(m.ChannelID, msg)
    }
}

func pm(s *discordgo.Session, m *discordgo.Message, msg string) {
    channel, _ := s.UserChannelCreate(m.Author.ID)
    split_message := split_long(msg)
    for _, msg := range split_message {
        s.ChannelMessageSend(channel.ID, msg)
    }
}

func split_long(msg string) []string {
    //Returns the message chunked into blocks. Almost 100% magic internet code
    line_size := 10
    var coarse_split [][]string
    fine_split := strings.FieldsFunc(msg, func(input rune) bool {
        return input == '\n'
    })
    for line_size < len(fine_split) {
        fine_split, coarse_split = fine_split[line_size:], 
            append(coarse_split, fine_split[0:line_size:line_size])
    }
    coarse_split = append(coarse_split, fine_split)
    
    var conglomerate []string
    for _, batch := range coarse_split {
        long_line := ""
        for _, line := range batch {
            long_line += line + "\n"
        }
        conglomerate = append(conglomerate, long_line)
    }

    needs_code_tags := false
    var output []string
    for _, chunk := range conglomerate {
        if needs_code_tags {
            chunk = "```" + chunk
        }
        if strings.Count(chunk, "```") % 2 == 1 {
            //if there's an odd number of ```
            needs_code_tags = true
            chunk += "```"
        }
        output = append(output, chunk)
    }
    return output
}

//----------------Define player commands----------------------------

func ping(s *discordgo.Session, m *discordgo.Message, args []string) {
    send(s, m, "Pong!")
}

func register(s *discordgo.Session, m *discordgo.Message, args []string) {
    success := game_classes.Mechanizmat.Register(m.Author)
    var msg string
    if success {
        msg = fmt.Sprintf("Thank you for registering, %s.\n", m.Author.Username)
        msg += "You've recieved a welcome package. Try `$info self`!"
    } else {
        msg = fmt.Sprintf("%s, you've already registered!", m.Author.Username)
    }
    pm(s, m, msg)
}

func unregister(s *discordgo.Session, m *discordgo.Message, args []string) {
    success := game_classes.Mechanizmat.Unregister(m.Author)
    var msg string
    if success {
        msg = fmt.Sprintf("You have unregistered, %s.", m.Author.Username)
    } else {
        msg = fmt.Sprintf("%s, you haven't registered yet!", m.Author.Username)
    }
    send(s, m, msg)
}

func info(s *discordgo.Session, m *discordgo.Message, args []string) {
    msg := "```"
    is_error := false
    err_msg := "Incorrect usage! Try:\n"
    err_msg += "\t-`$info self` to get an overview of yourself\n"
    err_msg += "\t-`$info mecha X` to get information about mecha #X"

    if len(args) == 0 {
        is_error = true
    } else {
        switch args[0] {
        case "self":
            msg += game_classes.Mechanizmat.GetPlayer(m.Author).Info()

        case "mecha":
            if len(args) != 2 {
                is_error = true
            } else {
                if i, err := strconv.Atoi(args[1]); err != nil {
                    is_error = true
                } else {
                    target_mecha := game_classes.Mechanizmat.GetPlayer(m.Author).FetchMecha(i-1)
                    if target_mecha == nil {
                        msg += "You don't have that many mechas!"
                    } else {
                        msg += target_mecha.Info(true)
                    }
                }
            }

        default:
            is_error = true
        }
    }
    if (is_error) {
        pm(s, m, err_msg)
    } else {
        msg += "```"
        pm(s, m, msg)
    }
    
}

func tick_test(s *discordgo.Session, m *discordgo.Message, args []string) {
    if len(args) != 1 {
        pm(s, m, "Incorrect usage!")
    }
    if i, err := strconv.Atoi(args[0]); err != nil {
        pm(s, m, "That's not a number!")
    } else {
        target_mecha := game_classes.Mechanizmat.GetPlayer(m.Author).FetchMecha(i-1)
        if target_mecha == nil {
            pm(s, m, "You don't have that many mechas!")
        } else {
            msg := fmt.Sprintf("```TEST\n%s```", target_mecha.TickTest(-1))
            pm(s, m, msg)
        }
    }
}

//----------------Define admin commands-----------------------------
func a_ping(s *discordgo.Session, m *discordgo.Message, args []string) {
    send(s, m, "Admin pong!")
}

func a_user_info(s *discordgo.Session, m *discordgo.Message, args []string) {
    if len(args) >= 1 {
        pm(s, m, fmt.Sprintf("`%#v`", m.Author))
    } else {
        fmt.Printf("%#v\n", m.Author)
    }
}

func a_player_info(s *discordgo.Session, m *discordgo.Message, args []string) {
    player := game_classes.Mechanizmat.GetPlayer(m.Author)
    if player != nil {
        if len(args) >= 1 {
            pm(s, m, fmt.Sprintf("`%#v`", player))
        } else {
            fmt.Printf("`%#v`\n", player)
        }
    } else {
        pm(s, m, "You haven't registered yet!")
    }
}

func a_game_info(s *discordgo.Session, m *discordgo.Message, args []string) {
    if len(args) >= 1 {
        pm(s, m, fmt.Sprintf("`%#v`", game_classes.Mechanizmat))
    } else {
        fmt.Printf("%#v\n", game_classes.Mechanizmat)
    }
}

func a_save(s *discordgo.Session, m *discordgo.Message, args []string) {
    game_classes.Mechanizmat.Save()
}

func a_load(s *discordgo.Session, m *discordgo.Message, args []string) {
    game_classes.Mechanizmat.Load()
}



//The commandn'ts
func no_command(s *discordgo.Session, m *discordgo.Message, cmd string) {
    possible_commands := make([]string, len(player_map))
    c := 0
    for cmd_name := range player_map {
        possible_commands[c] = cmd_name
        c++
    }
    cm := closestmatch.New(possible_commands, []int{2})
    closest := cm.Closest(cmd)

    msg := fmt.Sprintf("`$%s` isn't a known command!", cmd)
    if closest != "" {msg += fmt.Sprintf("\nDid you mean `$%s`?", closest)}
    send(s, m, msg)
}

func a_no_command(s *discordgo.Session, m *discordgo.Message, cmd string) {
    possible_commands := make([]string, len(admin_map))
    c := 0
    for cmd_name := range admin_map {
        possible_commands[c] = cmd_name
        c++
    }
    cm := closestmatch.New(possible_commands, []int{2})
    closest := cm.Closest(cmd)

    msg := fmt.Sprintf("`^%s` isn't a known admin command!", cmd)
    if closest != "" {msg += fmt.Sprintf("\nDid you mean `^%s`?", closest)}
    send(s, m, msg)
}