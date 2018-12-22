package owned_classes

import (
    "fmt"
    "math"
    "github.com/ulule/deepcopier"
)

type Mecha struct {
    Name       string          `json:"name"`
    Items      ItemCollection  `json:"items"` //Item string -> Count
    Components []Component    `json:"components"`
}

func (m *Mecha) Tick() {
    //This happens in REVERSE. (last component ticks first)
    for c := len(m.Components)-1; c >= 0; c-- {
        m.Components[c].Tick(&m.Items)
    }
}

func (m *Mecha) TickTest(max int) string {
    //Simulates the result of ticking all the way down to 0.
    //Set max < 0 to go until the mecha is no longer functioning.
    copy_mecha := &Mecha{}
    output := ""
    deepcopier.Copy(m).To(copy_mecha)
    for c := 1; c != max && copy_mecha.IsFunctioning(); c++ {
        //Run as long as c hasn't reached max OR the mecha dies
        output += copy_mecha.Info(true) + "\n"
        copy_mecha.Tick()
    }
    fmt.Println(output)
    output += copy_mecha.Info(true)
    return output
}

func (m *Mecha) Info(long_info bool) string {
    var is_on string
    if m.IsOn {is_on = "O"} else {is_on = "_"}
    var is_functioning string
    if m.IsFunctioning {is_functioning = "F"} else {is_functioning = "_"}
    czech := fmt.Sprintf("%s: S(%d) W(%d) A(%d) %s%s", m.Name, m.Speed(), m.Weight(), m.Armor())
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

func (m *Mecha) Speed() int {
    speed := 0
    for _, comp := range m.Components {
        speed += comp.E_Speed()
    }
    output := int(50*math.Sqrt(float64(6*speed)) / 3) - int(m.Weight() / 2)
    return int(output)
}

func (m *Mecha) Weight() int {
    output := 0
    for _, comp := range m.Components {
        output += comp.E_Weight()
    }
    for item, count := range m.Items {
        output += GetItem(item).Weight * count
    }
    return output
}

func (m *Mecha) Armor() int {
    if len(m.Components) == 0 {
        return 0 //also there's something wrong
    }
    return m.Components[0].E_Armor() //only outermost matters
}

func (m *Mecha) IsFunctioning() bool { //As long as one heart functions, the whole mech does.
    functioning := false
    for _, comp := range m.Components {
        if comp.IsHeart && comp.IsOn && comp.IsFunctioning {
            functioning = true
        }
    }
    return functioning
}