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

curl -k -s https://meleeframedata.com/$CHARACTER | \
    sed -E "s/<\/*div( class=[\"'].*[\"'])?>//g" | \
    grep -v \-e "<" -e "{" -e "}" -e ";" -e "function" | \
    sed "s/^[[:space:]]*//g" | \
    sed "/^[[:space:]]*$/d" | \
    grep -v "Notes" | \
    sed "s/Damage/Damage\n/g" | \
    sed '/\.0$/N;s/\n/ /' | \
    sed -n '/\+/q;p' > characters/$CHARACTER.info.template
