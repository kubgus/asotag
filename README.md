```
           _____  ____ _______       _____
    /\    / ____|/ __ \__   __|/\   / ____|
   /  \  | (___ | |  | | | |  /  \ | |  __
  / /\ \  \___ \| |  | | | | / /\ \| | |_ |
 / ____ \ ____) | |__| | | |/ ____ \ |__| |
/_/    \_\_____/ \____/  |_/_/    \_\_____|
```

# ASOTAG

> *A systems-oriented terminal adventure game.*

## Contents 📄

- [Overview 👀](#overview-) - Information about this project
- [For players 🎮](#for-players-) - How to download and play
- [For developers 💻](#for-developers-) - How to compile and run
- [For contributors 🌟](#for-contributors-) - How to contribute

## Overview 👀

ASOTAG is a game/framework that aims to achieve as much gameplay as possible
from the least amount of game mechanics. It's designed to favor <ins>systems</ins>
over _storylines_ or _scripted events_. It is a sort of an experiment to see
if a game can be created almost solely with <ins>emergent game mechanics</ins>.

### Two concepts

The entire game is based almost entirely based on just **2 concepts**:
- <ins>Entities</ins> - every object in the game world is represented by
an entity. Entities usually have a defined behavior that is executed once
per round on their turn. Most entities execute their turn based on
a set of rules defined in their script. If players are in the game,
their turn usually includes some text inputs to enable interaction
with the game. All interactions between entities (esp. players) should be
provided by items.
Static entities are also possible to allow for locations in the world.
(*chests, ore deposits, crafting stations, etc.*)
The preferred terms for this distinction are
**active** and **non-active** entities.
- <ins>Items</ins> allow entities (esp. players) to interact with other entites
and therefore the rest of the world. (i.e. dealing damage with weapons,
opening chests and mining deposits with tools, applying potions to
themselves or other entities, etc.) The rule of thumb with items it that
they should be universal for all applicable entities.

### Game world

The game world is defined as a finite (usually square) **grid** of **tiles**.
A tile can be occupied by any number of entities (**occupants**).
All occupants of a single tile should be able to interact.
Entities on the 8 surrounding tiles are conventionally considered the tile's
**neighbors**.

### Game loop

The game loop consists of **rounds** and **turns**. Any number of rounds can
take place, usually until a certain win/lose condition is met. A single round
is made up of a number of turns corresponding to the number of
alive active entities present in the world.

### Emergence in action

- **Swords:** Swords are items that can be used *by* an entity (user)
*on* an entity (target) that occupies the same tile to deal damage.
One play-tester of this game was about to be killed by goblins and there
was nothing they could do to stop it. That's when they decided that
"they ain't gonna be killed by no goblin" and instead, use the sword on
themselves, killing themselves and saving their pride.
- ***Fire spreading:*** Fire is an entity that damages all
occupants of its present tile and has a chance to spawn the fire entity
on each of its neighboring tiles. This simple set of rules just created
a new game mechanic that is able to interact with the entire ecosystem.

*(Titles marked with italic have not been implemented yet)*

---

### Terminal user interface (TUI)

The TUI usually consists of scrollable text that describe what is happening
**from the perspective of the world**. This can be extended with some
graphical (ASCII art) elements
(like the map that is available with cheat codes).

### Language choice

**Go** was chosen for this project mostly as a challenge and learning
experience. The codebase doesn't currently use any external dependencies.

### License

This project is licensed under the
[MIT License](LICENSE.md).

```
Copyright (c) 2026 Jakub Guštafik (@kubgus)
```

---

## For players 🎮

To **play** the game, first download the corresponding version of the
game for your system from the
[Releases](https://github.com/kubgus/asotag/releases)
section:

| OS | Architecture / Chipset | Download File Pattern |
| :- | :- | :- |
| **Windows** (7, 8, 10, 11) | Intel / AMD (Standard) | `..._windows_amd64.zip` |
| | Snapdragon / ARM | `..._windows_arm64.zip` |
| **macOS** (Sonoma, Sequoia, Tahoe) | Intel | `..._darwin_amd64.zip` |
| | Apple Silicon (M1–M5) | `..._darwin_arm64.zip` |
| **Linux** (Arch, Debian, etc.) | Intel / AMD (Standard) | `..._linux_amd64.zip` |
| | Raspberry Pi / ARM64 | `..._linux_arm64.zip` |

Then use your system's archive tool to **unzip** the downloaded file
and run the executable inside.

### Troubleshooting

Your firewall/antivirus software might flag the executable program as
suspicious and prevent it from running. This is because freshly compiled
programs usually need a bit of time to get recognized and verified as safe.
**Never feel pressured to run programs you don't trust on your machine!**
You can view the source code of the program in this repository to verify it
or even [compile it yourself](https://go-tutorial.com/build-and-run).
Still, here are ways you can bypass the "program not recognized" popups on
different systems:

- Windows - press the `Run anyway` button at the bottom of the popup.
(you might need to click something like "Show more" first)
- MacOS - go to `System Settings > Privacy & Security` and scroll to the bottom
to find the program that has been blocked from running and click `Open anyway`.

### How to play

**Your first round:**

- You can start by entering the `inventory` command to view your
inventory so you know what items you're working with.
- Continue by entering the `look` command and choosing a direction, which
will show you the nearest entity to you in that direction. There is a limit
on how many times you can look around per turn, but it will never end your
turn.
- If you find something interesting, you can enter `move` to move in that
direction and end your turn. You can also `wait` if you don't know where to
move.
- When you end your turn, you will see other entities' turns displayed in
chronological order.
- After a while, you might encounter an enemy in your `Nearby` list. This
means you can interact with it. Enter `use` to attack it with a weapon.

**Executing commands:**

- All available commands are always visible on the screen displayed in
<ins>yellow font</ins>. This makes the game very beginner friendly and
self-explanatory.
- If a command has any additional arguments, the game will prompt you to enter
them separately after running the initial command.
(yellow font rule applies as well)
- Some commands automatically end your turn while others don't. We
encourage you to explore the commands to learn how they behave.
- For more advanced players, you can also chain commands and arguments
by separating them with a space.
This means instead of entering `look` and then `west` when prompted for
the direction, you can just enter `look west`. This also works
with commands as a sort of "premove", so you can also enter `wait wait` to
wait for 2 turns.
- When entering numbers to select inventory items and nearby entities
(occupying the same tile as you), you can refer to the `Nearby` list
where entities are always sorted in the correct order. (starting with `0`)
The order of your inventory also doesn't change unless you use `bundle`.

**Crafting and bundling:**

- Crafting may be a little overwhelming at first, but it integrates very
well with the system and is very intuitive after you learn it.
- In order to craft, you must first encounter a crafting station.
(like the `Workbench`)
- Use the `examine` command on the station to see its crafting combinations.
(this won't end your turn)
- You can craft items by `use`ing the ingredients on the workbench.
However, most combinations require multiple items. That's when we need to
bundle the items.
- Use `bundle` to enter the bundling interface. Here you can select items from
your inventory either one by one or separated with a space. The bundling
interface can be exited by entering `x`. This will give you a bundle
that represents the items you selected. You can now `use` that bundle
on the crafting station to fulfill the recipe.
- *Note:* You can also use `bundle` and enter a number corresponding to the
bundle in your inventory to unbundle it and get your items back.
- *Note 2:* You can exit the interface by selecting another bundle instead
of typing `x` to add items to that bundle. Bundles cannot be bundled.
(non-stackable)

### We want your feedback!

We highly encourage you to play the game and leave your feedback
(*feature requests, bugs, etc.*)
in the [Issues section](https://github.com/kubgus/asotag/issues).
This will help us develop the game further and hopefully also give you
a fun experience.

## For developers 💻

```sh
# Clone and enter
git clone https://github.com/kubgus/asotag
cd asotag

# Run the game
go run .
```

```sh
# Enable scripts
chmod +x _scripts/*.sh

# Run the game
./_scripts/dev.sh

# Build for different/multiple platforms
./_scripts/build-all.sh
./_scripts/build-windows.sh
./_scripts/build-linux.sh
```

The codebase should be regularly formatted by running `go fmt ./...`,
ideally at least before every commit.

## For contributors 🌟

**Commits** should follow this pattern:

```yml
feat: add example item
tune: edit chest drop chances
fix: remove extra newline in player move action
refactor: reorder for loops in potion code for better cache performance
chore: add github workflows
```

**Tags** (as well as **tag messages**) follow the `SEMVER` pattern:

```html
v<MAJOR>.<MINOR>.<PATCH>
```

The project automatically builds for all major platforms when a new **tag**
gets pushed to the repo.

Refer to the [License section](#license) for information about the license.
