#!/bin/bash

CHARACTERS=(
    "captain_falcon" "donkey_kong" "dr._mario" "falco" "fox" "ganondorf"
    "ice_climbers" "kirby" "bowser" "link" "luigi" "mario" "marth"
    "mewtwo" "mr._game_&_watch" "ness" "peach" "pichu" "pikachu"
    "jigglypuff" "roy" "samus" "sheik" "yoshi" "young_link" "zelda"
)

## BUG: MAYBE: The character name at the top of the file might be throwing the
##      Yaml parser off.

for CHARACTER in "${CHARACTERS[@]}"
do
    echo "Creating character $CHARACTER..."
    ./seed/characters/get-character.sh "$CHARACTER" | yq -o=json > "seed/characters/$CHARACTER.json"
done
