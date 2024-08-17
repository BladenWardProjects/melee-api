#!/bin/bash

# This script downloads the character data from meleeframedata.com
# and formats it into YAML files.
# URL: https://meleeframedata.com/fox

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <character>"
    exit 1
fi

CHARACTER=$1
CHARACTER_UPPER=$(echo "$CHARACTER" | tr '[:lower:]' '[:upper:]')

curl -k -s https://meleeframedata.com/$CHARACTER | \
    sed -E "s/<\/*div( class=[\"'].*[\"'])?>//g" | \
    grep -v \-e "<" -e "{" -e "}" -e ";" -e "function" | \
    sed "s/^[[:space:]]*//g" | \
    sed "/^[[:space:]]*$/d" | \
    grep -v "Notes" | \
    sed "s/Damage/Damage\n/g" | \
    sed '/\.0$/N;s/\n/ /' | \
    sed -n '/\+/q;p' | \

    sed "s/$CHARACTER_UPPER/$CHARACTER:\n  name: $CHARACTER\n/g" | \

    sed -r "s/^Ground Attacks/  moves:\n    ground:\n/g" | \
    sed -r "s/^([[:digit:]]+) Frame Startup.*$/        start: \1/g" | \
    sed -r "s/^Active Frames [[:digit:]]+-([[:digit:]]+)$/        end: \1/g" | \
    sed -r "s/^([[:digit:]]+) Total Frames$/        frames: \1/g" | \
    sed -r "s/^IASA Frame ([[:digit:]]+).*$/        iasa: \1/g" | \
    sed -r "s/^Shield Stun ([[:digit:]]+).*$/        shield_stun: \1/g" | \
    sed -r "s/^Shield Stun  Frames$/        shield_stun: 0/g" | \
    sed -r "s/^([[:digit:]]+)\.[[:digit:]]( \/ )?([[:digit:]]+)?(\.[[:digit:]])? % Base Damage/        base_damage: \[\1, \3\]/g" | \
    sed -r "s/, \]$/]/g" | \

    sed -r "s/^Aerial Attacks$/    aerial:\n/g" | \
    sed -r "s/^([[:digit:]]+) Frames Landing Lag.*$/        landing_lag: \1/g" | \
    sed -r "s/^([[:digit:]]+)( Frames)? L-Cancel Lag.*$/        lcancel_lag: \1/g" | \
    sed -r "s/^Won't Auto Cancel Frames [[:digit:]]+-([[:digit:]]+)$/        auto_cancel: \1/g" | \

    sed -r "s/^Special Attacks$/    special:\n/g" | \
    sed -r "s/^([[:digit:]]+) Frames LFS Lag.*$/        landing_fall_special: \1/g" | \

    sed -r "s/^Grabs$/    grab:\n/g" | \

    sed -r "s/^Throws$/    throw:\n/g" | \
    sed -r "s/^([[:digit:]]+).*% Damage.*$/        damage: \1/g" | \

    sed -r "s/^Dodges\/Rolls$/    dodge:\n/g" | \
    sed -r "s/^Inv\. ?Frames [[:digit:]]+-([[:digit:]]+).*$/        end: \1/g" | \
    sed -r "s/^Landing Fall Special Lag: ([[:digit:]]+) Frames.*$/        landing_fall_special: \1/g" | \

    sed -r "s/^Jab$/      jab:/g" | \
    sed -r "s/^Jab 2$/      jab2:/g" | \
    sed -r "s/^Jab 3$/      jab3:/g" | \
    sed -r "s/^Forward Tilt$/      forward_tilt:/g" | \
    sed -r "s/^Up Tilt$/      up_tilt:/g" | \
    sed -r "s/^Down Tilt$/      down_tilt:/g" | \
    sed -r "s/^Dash Attack$/      dash_attack:/g" | \
    sed -r "s/^Forward Smash$/      forward_smash:/g" | \
    sed -r "s/^Up Smash$/      up_smash:/g" | \
    sed -r "s/^Down Smash$/      down_smash:/g" | \

    sed -r "s/^Neutral Air$/      neutral_air:/g" | \
    sed -r "s/^Forward Air$/      forward_air:/g" | \
    sed -r "s/^Back Air$/      back_air:/g" | \
    sed -r "s/^Up Air$/      up_air:/g" | \
    sed -r "s/^Down Air$/      down_air:/g" | \

    sed -r "s/^Neutral B$/      neutral_b:/g" | \
    sed -r "s/^Aerial Neutral B$/      aerial_neutral_b:/g" | \
    sed -r "s/^Aerial Side B$/      aerial_side_b:/g" | \
    sed -r "s/^Aerial Up B$/      aerial_up_b:/g" | \
    sed -r "s/^Side B$/      side_b:/g" | \
    sed -r "s/^Up B$/      up_b:/g" | \
    sed -r "s/^Down B$/      down_b:/g" | \

    sed -r "s/^Standing Grab$/      standing_grab:/g" | \
    sed -r "s/^Dash Grab$/      dash_grab:/g" | \

    sed -r "s/^Forward Throw$/      forward_throw:/g" | \
    sed -r "s/^Back Throw$/      back_throw:/g" | \
    sed -r "s/^Down Throw$/      down_throw:/g" | \
    sed -r "s/^Up Throw$/      up_throw:/g" | \

    sed -r "s/^Spot Dodge$/      spot_dodge:/g" | \
    sed -r "s/^Backward Roll$/      backward_roll:/g" | \
    sed -r "s/^Forward Roll$/      forward_roll:/g" | \
    sed -r "s/^Air Dodge$/      air_dodge:/g" | \

    sed -r "s/^Weight: ([[:digit:]]+).*$/  weight: \1/g" | \
    sed -r "s/^Fast Fall Speed: ([[:digit:]]+\.[[:digit:]]+).*\$/  fastfall_speed: \1/g" | \
    sed -r "s/^Dash Speed: ([[:digit:]]+\.[[:digit:]]+).*$/  dash_speed: \1/g" | \
    sed -r "s/^Run Speed: ([[:digit:]]+\.[[:digit:]]+).*$/  run_speed: \1/g" | \
    sed -r "s/^Wavedash Length \(Rank\): ([[:digit:]]+).*$/  wavedash_length_rank: \1/g" | \
    sed -r "s/^PLA Intangibility Frames: ([[:digit:]]+).*$/  galint: \1/g" | \
    sed -r "s/^Jump Squat: ([[:digit:]]+).*$/  jump_squat: \1/g" | \
    sed -r "s/^Yes$/true/g" | \
    sed -r "s/^No$/false/g" | \
    sed -r "s/^(true|false)$/  walljump: \1/g" | \

    sed -r "/^Miscellaneous Info$/d" | \
    sed -r "/^Wall Jump:.*$/d" | \
    sed -r "/^[[:space:]]*$/d"

# TODO: Add the missing character stats from SSBM Wiki
