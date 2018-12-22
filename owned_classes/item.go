package owned_classes

type Item struct {
    Name   string `json:"name"`
    Weight int    `json:"weight"` //Weights are in g/cm^3-ish
}