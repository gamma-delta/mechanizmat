package owned_classes

import (
    "fmt"
)

type Component struct {
    Name string `json:"name"`

    IsOn          bool //Can be disabled by player
    IsFunctioning bool //Only set by self

    //Static SWA are always on.
    S_Speed  int `json:"static_speed"`
    S_Weight int `json:"static_weight"`
    S_Armor  int `json:"static_armor"`

    //Delta SWA become 0 when the component is off or not functioning
    D_Speed  int `json:"delta_speed"`
    D_Weight int `json:"delta_weight"`
    D_Armor  int `json:"delta_armor"`

    //Items reqd and produced
    ReqdItems ItemCollection `json:"required_items"`
    ProdItems ItemCollection `json:"produced_items"`

    IsHeart bool `json:"is_heart"` //as long as one heart still is OK in a mecha, it survives
}

func (c *Component) Tick(mech_items *ItemCollection) bool {
    if c.IsOn { //don't bother checking if I'm off
        if mech_items.CanContain(c.ReqdItems) {
            c.IsFunctioning = true
            mech_items.Subtract(c.ReqdItems)
            mech_items.Add(c.ProdItems)
        } else { //oh no, I no longer have function
            c.IsFunctioning = false // T_T
        }
    }
    if c.IsOn && c.IsFunctioning && c.IsHeart {
        return true
    } else {
        return false
    }
}

func (c Component) Info() string {
    is_on, is_functioning := "_", "_"
    if (c.IsOn) {is_on = "O"}
    if (c.IsFunctioning) {is_functioning = "F"}
    czech := fmt.Sprintf("%s%s ", is_on, is_functioning)

    //Speed
    if c.D_Speed == 0 { //no extra speed
        if c.S_Speed != 0 { //but I still should display S_speed
            czech += fmt.Sprintf("S(%d) ", c.S_Speed) //very short form
        }
    } else { //d_speed, so longer form
        czech += fmt.Sprintf("S(%d|%d, %d) ", c.E_Speed(), c.S_Speed, c.D_Speed)
    }

    //Weight. Wish I could iterate this...
    if c.D_Weight == 0 {
        if c.S_Weight != 0 {
            czech += fmt.Sprintf("W(%d) ", c.S_Weight)
        }
    } else { 
        czech += fmt.Sprintf("W(%d|%d, %d) ", c.E_Weight(), c.S_Weight, c.D_Weight)
    }

    //and armor
    if c.D_Armor == 0 {
        if c.S_Armor != 0 {
            czech += fmt.Sprintf("A(%d) ", c.S_Armor)
        }
    } else { //d_speed, so longer form
        czech += fmt.Sprintf("A(%d|%d, %d) ", c.E_Armor(), c.S_Armor, c.D_Armor)
    }

    return czech


}

//Effective SWA is whatever the S_SWA + D_SWA is.
func (c Component) E_Speed() int {
    output := c.S_Speed
    if c.IsOn && c.IsFunctioning {
        output += c.D_Speed
    }
    return output
}

func (c Component) E_Weight() int {
    output := c.S_Weight
    if c.IsOn && c.IsFunctioning {
        output += c.D_Weight
    }
    return output
}

func (c Component) E_Armor() int {
    output := c.S_Armor
    if c.IsOn && c.IsFunctioning {
        output += c.D_Armor
    }
    return output
}