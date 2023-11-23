Comp426 - 001 Project 3
Completed by Andrew Janedy
Date Completed: 11/22/23

This assignment has met all of the requirements except for
the differing maps and tiles.  There is only one map that
is used for all three levels that consists of 7 tiles. 
Though they are animated, the animations seen around the 
screen is animated maps, not animated sprites.

The map is 15x15

2 layers, the black is the floor and the fire is the 
boundary

There are no "object layers" but I think this requirement
was removed

There are 12 different animations, player, rainbow man, 
bees, golem, two player attacks, rainbow man attack, 
bee attack, 3 power ups, a key, plus the map animations
makes 13, maybe.

The player is 8 direction as is rainbow man.  

The player can interact with the quest keeper on the first
level, enemies on all three levels, as well as the power ups
and keys that spawn on levels 1 and 2.

I've done my best to keep game entities on the map, enemies
will only follow the player off the map if the bug that
I cannot fix is encountered that allows the player to 
pass through the barrier tiles.  (Use the teleport to get
back in I suppose)

Player stats can be improved through power ups, health pots,
speed boosts, mana pots.  These entities are obvious when
dropped and their impact is reflected in the HUD at the 
bottom left

Enemies attack player when in range, level one enemies
follow player between a certain range, aggression can be 
lost if far enough away

I have not yet implemented "wearables", items that the 
can only have a finite amount of, instead, attributes
are increased through power ups, scaling is unlimited
but each enemy only drops one power up at random

Player advances map by collecting the key

Move with WASD
Shoot fireball with left click
Massive explosion with Q
Teleport with E

All skills are on a cool down which is represented in the 
lower left in the HUD

At the start of the game travel to the creepy yellow
quest giver to get your instructions