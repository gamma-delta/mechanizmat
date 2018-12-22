package owned_classes

type ItemCollection map[string]int

func (i ItemCollection) Size() int {
    output := 0
    for _, count := range i {
        output += count
    }
    return output
}

func (i ItemCollection) CanContain(other ItemCollection) bool {
    //Returns true if every item count in "other" fits in "i"
    for my_id, my_count := range i {
        if my_count < other[my_id] { //if it's too big
            return false
        }
    }
    return true //made it!
}

func (i ItemCollection) Add(other ItemCollection) {
    for my_id, _ := range i {
        i[my_id] += other[my_id]
    }
}

func (i ItemCollection) Subtract(other ItemCollection) bool {
    //Subtracts each item in other from self.
    //Returns "true" if successful
    if !i.CanContain(other) {
        return false //too small!
    } else {
        for my_id, _ := range i {
            i[my_id] -= other[my_id]
        }
        return true
    }
}