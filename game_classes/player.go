package game_classes

import (
    "mechanizmat/owned_classes"

    "fmt"
)

type Player struct {
    Name       string  `json:"name"`
    Mechas     []*owned_classes.Mecha       `json:"mechas"`
    Components []*owned_classes.Component   `json:"components"`
    Items      owned_classes.ItemCollection `json:"items"` 
}

func NewPlayer(name string) *Player {
    new_player := &Player{Name: name}
    new_player.GiveStarterKit()
    return new_player
}

func (p *Player) Info() string {
    output := fmt.Sprintf("Info for %s\n", p.Name)

    //mechas
    output += fmt.Sprintf("Mechas x%d:\n", len(p.Mechas))
    for c, mecha := range p.Mechas {
        output += fmt.Sprintf("\t#%d: %s\n", c+1, mecha.Info(false))
    }

    //Components
    output += fmt.Sprintf("Components x%d:\n", len(p.Components))
    for c, comp := range p.Components {
        output += fmt.Sprintf("\t#%d: %s\n", c+1, comp.Info())
    }

    //Items
    output += fmt.Sprintf("Items x%d:\n", p.Items.Size())
    for item_id, count := range p.Items {
        output += fmt.Sprintf("\t%dx %s\n", count, owned_classes.GetItem(item_id).Name)
    }
    return output
}

func (p *Player) FetchMecha(index int) *owned_classes.Mecha {
    if index - 1 >= len(p.Mechas) || index < 0 {
        return nil
    } else {
        return p.Mechas[index]
    }
}

func (p *Player) GiveStarterKit() {
    p.Mechas = append(p.Mechas, &owned_classes.Mecha{
        Name: "Luftwaffe Training Mecha",
        Components: owned_classes.GetComponents("ARMOR_1", "TREADS_1", "TREADS_1", "BOILER_1"),
        Items: owned_classes.ItemCollection{
            "WATER": 200,
            "FUEL": 20,
        },
        IsFunctioning: true,
    })

    p.Items = owned_classes.ItemCollection{
        "WATER": 1000,
        "FUEL": 200,
        "GOLD": 80,
        "PETROS": 50,
    }
}