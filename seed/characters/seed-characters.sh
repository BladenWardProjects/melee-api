#!/bin/bash

CHARACTERS=(
    "captain_falcon"
    "donkey_kong"
    "dr._mario"
    "falco"
    "fox"
    "ganondorf"
    "ice_climbers"
    "kirby"
    "bowser"
    "link"
    "luigi"
    "mario"
    "marth"
    "mewtwo"
    "mr._game_&_watch"
    "ness"
    "peach"
    "pichu"
    "pikachu"
    "jigglypuff"
    "roy"
    "samus"
    "sheik"
    "yoshi"
    "young_link"
    "zelda"
)

for CHARACTER in "${CHARACTERS[@]}"
do
    echo "Creating character $CHARACTER..."
    ./get-character.sh "$CHARACTER" > "$CHARACTER.yaml"
done
