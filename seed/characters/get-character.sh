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

    sed "s/$CHARACTER_UPPER/$CHARACTER\n\tname: $CHARACTER\n/g" | \

    sed "s/^Ground Attacks/\tmoves:\n\t\tground:\n/g" | \
    sed "s/^Jab$/\t\t\tjab:/g" | \
    sed -r "s/Active Frames ([[:digit:]]+)-([[:digit:]]+)$/\t\t\t\tstart: \1\n\t\t\t\tend: \2/g" | \
    sed -r "s/^([[:digit:]]+) Total Frames$/\t\t\t\tframes: \1/g" | \
    sed -r "s/^IASA Frame ([[:digit:]]+).*$/\t\t\t\tiasa: \1/g" | \
    sed -r "s/^Shield Stun ([[:digit:]]+).*$/\t\t\t\tshield_stun: \1/g" | \
    sed -r "s/^([[:digit:]]+)\.[[:digit:]]( \/ )?([[:digit:]]+)?(\.[[:digit:]])? % Base Damage/\t\t\t\tbase_damage: \[\1, \3\]/g" | \
    sed -r "s/, \]$/]/g" | \

    sed -r "s/^Aerial Attacks$/\t\t\taerial:\n/g" | \
    sed -r "s/^([[:digit:]]+) Frames Landing Lag$/\t\t\t\tlanding_lag: \1/g" | \
    sed -r "s/^([[:digit:]]+)( Frames)? L-Cancel Lag.*$/\t\t\t\tlcancel_lag: \1/g" | \
    sed -r "s/^Won't Auto Cancel Frames [[:digit:]]+-([[:digit:]]+)$/\t\t\t\tauto_cancel: \1/g" | \

    sed -r "s/^Special Attacks$/\t\t\tspecial:\n/g" | \

    sed -r "s/^Grabs$/\t\t\tgrab:\n/g" | \

    sed -r "s/^Throws$/\t\t\tthrow:\n/g" | \
    sed -r "s/^Dodges\/Rolls$/\t\t\tdodge:\n/g"


    

echo $CHARACTER_UPPER
