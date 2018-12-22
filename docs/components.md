# COMPONENTS

Every single mecha, from the zeppelins of our own German Luftwaffe to the Xin-Hankook two-leg walkers, can be considered as a number of components stuck together. Components are what make a mecha run, walk, shoot, make steam, hover, or any other function. Every turn in a battle, each component might consume some items, produce some items, and/or increase to your mecha's stats.

## Stats

Just like mechas, Components have a Speed, Weight, and Armor stat, known as the *static stats* or *s-stats*, for reason's you'll see in a moment. They also have additional *delta stats* or *d-stats* that only apply when the component is switched on and has enough items. The *s-stat* plus the *d-stat* for a given stat is known as the *effective stat* or the *e-stat*; this is the number you have to worry most about.

Examples:

> The most common model of American-made treads, the Whitney Tread MK I, has the following stats:
> - S-speed: 0
> - S-weight: 40
> - S-armor: 20
> - D-speed: 30
> - D-weight: 0
> - D-armor: 5
>
> Therefore, when the component is on and functioning:
>
> - E-speed: 30 (0 + 30)
> - E-weight: 40 (40 + 0)
> - E-armor: 25 (20 + 5)
>
> When the component isn't on or isn't functioning:
>
> - E-speed: 0 (0 + 0)
> - E-weight: 40 (40 + 0)
> - E-armor: 20 (20 + 0)

> A Luftwaffe standard-issue steam-balloon has the following stats:
> - S-speed: 0
> - S-weight: 5
> - S-armor: 10
> - D-speed: 0
> - D-weight: -25
> - D-armor: -5
>
> Therefore, when the component is on and functioning:
>
> - E-speed: 0 (0 + 0)
> - E-weight: -20 (5 + -25)
> - E-armor: 5 (10 + -5)
>
> When the component isn't on or isn't functioning:
>
> - E-speed: 0 (0 + 0)
> - E-weight: 5 (5 + 0)
> - E-armor: 10 (10 + 0)

Obviously, it would be exhausting to write these stats longhand, even before the items required and produced by the component were also included. Therefore, a shorthand was developed by Czech Technical University known as Czech notation for writing all the information about a mecha. The notation for the stats of a mech consists of the following pattern repeated for each stat (here shown just for the Speed stat):

> `S(e-speed|s-speed, d-speed)`

If the S, D, and E-stats for a given stat are all 0, it can be omitted. Similarly, if the D-stat for a given stat is 0, the S-stat can be written alone.

The Czech notations for the two components previously listed are provided below as an example.

> Whitney Tread MK I: `S(30|0, 30) W(40) A(20|0, 20)`
>
> Standard Luftwaffe steam-balloon: `W(-20|5, -25) A(5|10, -5)`

This shorthand is much simpler to read and write. The development of this simple notation is widely regarded as one of the leading causes for the Mechanical Revolution, for reasons outside the scope of this article.

## Enabling / Disabling and Functioning / Not Functioning
Components can be manually turned off. A component that is off doesn't consume or make any materials, or in its d-stat into its e-stat.

When a component runs out of any materials, it stops functioning. A component that isn't functioning also doesn't produce materials or provide its d-stat. However, it will automatically start functioning again when all the materials it needs are provided again.

In Czech notation, a component being on is represented with "O", while off is "\_". A functioning component is "F", while a component that isn't functioning is "\_". This is usually written in front of the stats, but sometimes it is written after consumed and produced items (see below).

Examples:
> - "OF" = on and functioning
> - "\_F" = functioning, but not on (if it's turned back on, it will work)
> - "O\_" = on, but not functioning (it ran out of materials, presumably)
> - "\_\_" = neither on nor functioning (this is fairly abnormal, as the component had to be manually turned off after it ran out of materials.)

## Consuming and Producing Items

Most components require items to function; some components also produce items. The specifics for the order and timing of this production and consumption during a battle are more thouroughly described in the Mecha section of this paper. Simply put, after an action is taken a mechanism tries to consume all the items required to function, then produces all the items it can produce.

Examples:

> A Xin-Hankook General Use Boiler consumes:
>
> - 10 Water
> - 2 Fuel
>
> ...and produces:
>
> - 20 Steam
>
> The aforementioned Luftwaffe steam-balloon consumes:
>
> - 30 Steam
>
> ...and doesn't produce anything.
>
> Basic armor plating from the Ottoman Empire generally doesn't produce or consume anything.
>
> Very few known components consume nothing and produce items.

As writing this longhand is also time-consuming, Czech notation also includes a shorthand for production and consumption of items. The Czech notation for the Xin-Hankook boiler is as follows:

> `REQ[10WA 2FU] PRO[20ST]`
> 
> See the section on items for abbreviations.

As before, if a component doesn't require or produce any materials, that information may be omitted instead of writing all zeroes. Thus, the Czech notation for the Luftwaffe steam-balloon is as follows:

> `REQ[30ST]`

Usually, the Czech notations for stats and materials are written together. The full Czech notations for all four components are written below. Associating each notation with each component is left as an exercise for the reader.

> `OF W(-20|5, -25) A(5|10, -5) REQ[30ST]`
>
> `OF W(500) A(50)`
>
> `OF S(100|0, 100) W(150) A(20) REQ[20ST]`
>
> `OF W(40) A(30) REQ[10WA 2FU] PRO[20ST]`

## Ticking
"Ticking" is what happens after an action is taken in battle. When a component ticks, it tries to consume all the materials it can, stops or starts functioning accordingly, and makes materials if it can.