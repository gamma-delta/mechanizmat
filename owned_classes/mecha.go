package owned_classes

import (
    "fmt"
    "math"
)

type Mecha struct {
    Name       string         `json:"name"`
    Items      ItemCollection `json:"items"` //Item string -> Count
    Components []Component    `json:"components"`
    IsFunctioning bool        `json:"is_functioning"`
}

func (m *Mecha) Tick() {
    //This happens in REVERSE. (last component ticks first)
    still_functions := false
    for c := len(m.Components)-1; c >= 0; c-- {
        one_support := m.Components[c].Tick(&m.Items)
        if !still_functions && one_support {
            //we only need one to function
            still_functions = true
        }
    }
    m.IsFunctioning = still_functions
}

func (m Mecha) TickTest(max int) string {
    //Simulates the result of ticking all the way down to 0.
    //Set max < 0 to go until the mecha is no longer functioning.
    copy_mecha := m
    output := ""
    for c := 1; c != max && copy_mecha.IsFunctioning; c++ {
        //Run as long as c hasn't reached max OR the mecha dies
        output += copy_mecha.Info(true) + "\n-----------\n"
        copy_mecha.Tick()
    }
    output += copy_mecha.Info(true)
    fmt.Println(output)
    return output
}

func (m Mecha) Info(long_info bool) string {
    is_functioning := "_"
    if m.IsFunctioning {is_functioning = "F"}
    czech := fmt.Sprintf("%s: S(%d) W(%d) A(%d) %s", 
        m.Name, m.Speed(), m.Weight(), m.Armor(), is_functioning)
    if long_info {
        output := fmt.Sprintf("Info for mecha: %s %s\n", m.Name, czech)
        output += fmt.Sprintf("Components x%d:\n", len(m.Components))
        for c, comp := range m.Components {
            output += fmt.Sprintf("\t#%d: %s %s\n", c+1, comp.Name, comp.Info())
        }
        output += fmt.Sprintf("Items x%d:\n", m.Items.Size())
        for item_id, count := range m.Items {
            output += fmt.Sprintf("\t%dx %s\n", count, GetItem(item_id).Name)
        }
        return output
    } else {
        return czech
    }
}

func (m Mecha) Speed() int {
    speed := 0
    for _, comp := range m.Components {
        speed += comp.E_Speed()
    }
    output := int(50*math.Sqrt(float64(6*speed)) / 3) - int(m.Weight() / 2)
    return int(output)
}

func (m Mecha) Weight() int {
    output := 0
    for _, comp := range m.Components {
        output += comp.E_Weight()
    }
    for item, count := range m.Items {
        output += GetItem(item).Weight * count
    }
    return output
}

func (m Mecha) Armor() int {
    if len(m.Components) == 0 {
        return 0 //also there's something wrong
    }
    return m.Components[0].E_Armor() //only outermost matters
}