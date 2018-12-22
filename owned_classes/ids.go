package owned_classes

import (
    "log" //oh no, unknown item!
)

var AllItems map[string]Item
var AllComponents map[string]Component

func init() {
    AllItems = map[string]Item {
        "GOLD": Item{Weight: 5, Name: "Gold"},
        "WATER": Item{Weight: 1, Name: "Water"},
        "FUEL": Item{Weight: 2, Name: "Fuel"},
        "STEAM": Item{Weight: 0, Name: "Steam"},
        "PETROS": Item{Weight: 3, Name: "Petros"},
    }

    AllComponents = map[string]Component {
        "BOILER_1": Component{Name: "Xin-Hankook General Use Boiler", 
            S_Weight: 40, S_Armor: 30,
            ReqdItems: ItemCollection{
                "WATER": 10,
                "FUEL": 2, 
            }, ProdItems: ItemCollection{"STEAM": 10},
            IsHeart: true,
            IsOn: true, IsFunctioning: true,
        },

        "ARMOR_1": Component{Name: "Ottoman Plating", 
            S_Weight: 500, S_Armor: 50,
            IsOn: true, IsFunctioning: true,
        },

        "TREADS_1": Component{Name: "Whitney Tread MK I",
            S_Weight: 150, S_Armor: 20, 
            D_Speed: 100,
            ReqdItems: ItemCollection{"STEAM": 20},
            IsOn: true, IsFunctioning: true,
        },
    }
}

func GetItem(id string) Item {
    item, ok := AllItems[id]
    if ok {
        return item
    } else {
        log.Fatalf("!!ERROR: Unknown item '%s'!\n", id)
        return Item{} //doesn't matter, as we're panicking anyway
    }
}

func GetItems(ids ...string) []Item {
    var output []Item
    for _, id := range ids {
        output = append(output, GetItem(id))
    }
    return output
}

func GetComponent(id string) Component {
    comp, ok := AllComponents[id]
    if ok {
        return comp
    } else {
        log.Fatalf("!!ERROR: Unknown component %s!\n", id)
        return Component{}
    }
}

func GetComponents(ids ...string) []Component {
    var output []Component
    for _, id := range ids {
        output = append(output, GetComponent(id))
    }
    return output
}