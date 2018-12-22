# MECHAS
Mechas are the new revolution, the next step of humanity's development. Most of the scientific, economic, and mechanical breakthroughs of recent years have been a result of or responsible for the advancement of mechas.

Despite this grand description, it is not that complicated to operate, service, or even create a new mecha. This is due to a curious phenomonon; nearly every single manufacturer of components across the world have agreed, with no input from some greater body, on a standard on how these components can connect to each other. The amazing result is a mecha can be made with a boiler from Xin-Hankook, American treads, and a Czech battering ram, and still function properly.

No matter how complex or simple a mecha, it can be reduced to 3 statistics ("stats") and a list of components in order. This article provides an outline of these considerations.

## Stats
Stats are provided by the components and items in a mecha. They can be considered a sum or combination of the e-stats of all the components in a mecha, which also have these stats. A detailed description of each follows.

- SPEED: How far the mecha can move per turn.
- WEIGHT: How heavy the mecha is. This affects the speed.
- ARMOR: The chance an attack against the mecha will destroy the outermost part.

### Speed
Speed is how far a mecha can move per turn. It also affects certain attacks, such as battering rams.

The speed of a mecha is the sum of the e-speed stats of each component, tempered by the mecha's weight. Speed can be roughly calculated as such:

![(50\*sqrt(6\*espeed))/3 - weight/2](http://bit.ly/2PSnMfo "(50\*sqrt(6\*espeed))/3 - weight/2")

where _espeed_ is the sum of the e-speeds of each component and _weight_ is the weight of the mecha.

### Weight
Weight is a measure of how heavy a mecha is. It is usually best to keep this as low as possible, as a high weight makes it harder to escape from a battle. However, if a mech's weight becomes less than -5, it will float away. (Luftwaffe zepplins are kept around zero by an intricate system of checks and balances outside the scope of this document.)

Weight is simply the sum of all the weights of all the components and items a mech holds.

### Armor
Armor is, roughly, how likely it is an attack will destroy a component.

Mechas by themselves don't strictly have an armor rating, although it is convenient to think that way. A mecha's armor stat is equal to the armor stat of the outermost component. When that component is destroyed, the armor stat changes to the armor stat of the new outermost component, and so on.

## Lists of Components
Mechas are usually notated as a list of components along with the stats. These components are everything one needs to know to construct a mecha. The order in which these components are written does matter.
- The topmost component is the outermost one, and the one that will be destroyed first.
- After a turn is taken, components tick from bottom to top, or innermost to outermost.

As a result of this, the last component in a list is almost always a boiler of some sort. (Remember that when a mecha has no more working boilers, it loses the battle.)