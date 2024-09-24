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
CHARACTER_URL=$CHARACTER
CHARACTER_UPPER=$(echo "$CHARACTER" | tr '[:lower:]' '[:upper:]')

case "$CHARACTER" in
"captain_falcon")
CHARACTER_URL="captain_falcon" CHARACTER_UPPER="CAPTAIN FALCON"
;;
"donkey_kong")
CHARACTER_URL="donkey_kong" CHARACTER_UPPER="DONKEY KONG"
;;
"dr._mario")
CHARACTER_URL="dr._mario" CHARACTER_UPPER="DR. MARIO"
;;
"ice_climbers")
CHARACTER_URL="ice_climbers" CHARACTER_UPPER="ICE CLIMBERS"
;;
"mr._game_&_watch")
CHARACTER_URL="mr._game_&_watch" CHARACTER_UPPER="MR. GAME &amp; WATCH"
;;
"young_link")
CHARACTER_URL="young_link" CHARACTER_UPPER="YOUNG LINK"
;;
esac

curl -k -s https://meleeframedata.com/$CHARACTER_URL | \
sed -E "s/<\/*div( class=[\"'].*[\"'])?>//g" | \
grep -v \-e "<" -e "{" -e "}" -e ";" -e "function" | \
sed "s/^[[:space:]]*//g" | sed "/^[[:space:]]*$/d" | grep -v "Notes" | \
sed "s/Damage/Damage\n/g" | sed -r '/^[[:digit:]]+\.0$/N;s/\n/ /' | sed -n '/\+/q;p' | \

# FIX: The ampersand in the name of Mr. Game & Watch does not play nice with sed
# so this is the workaround
if [ "$CHARACTER" = "mr._game_&_watch" ]; then
{ printf "name: mr._game_&_watch\n" ; cat; }
else
sed "s/$CHARACTER_UPPER/name: $CHARACTER\n/g"
fi | \

sed -r "s/^Ground Attacks/ground:\n/g" | sed -r "s/^([[:digit:]]+) Frame Startup.*$/    start: \1/g" | \
sed -r "s/^Active Frames [[:digit:]]+-([[:digit:]]+)$/    end: \1/g" | sed -r "s/^([[:digit:]]+) Total Frames$/    frames: \1/g" | \
sed -r "s/^IASA Frame ([[:digit:]]+).*$/    iasa: \1/g" | \
sed -r "s/^Shield Stun ([[:digit:]]+).*$/    shield_stun: \1/g" | sed -r "s/^Shield Stun  Frames$/    shield_stun: 0/g" | \
sed -r "s/^([[:digit:]]+)\.[[:digit:]]( \/ )?([[:digit:]]+)?(\.[[:digit:]])? ?% Base Damage/    base_damage: \1\n    weak_damage: \3/g" | \
sed -r "/^    weak_damage: $/d" | \

sed -r "s/^Aerial Attacks$/aerial:\n/g" | sed -r "s/^([[:digit:]]+) Frames Landing Lag.*$/    landing_lag: \1/g" | \
sed -r "s/^([[:digit:]]+)( Frames)? L-Cancel Lag.*$/    lcancel_lag: \1/g" | \
sed -r "s/^Won't Auto Cancel Frames [[:digit:]]+-([[:digit:]]+)$/    auto_cancel: \1/g" | \

sed -r "s/^Special Attacks$/special:\n/g" | sed -r "s/^([[:digit:]]+) Frames LFS Lag.*$/    landing_fall_special: \1/g" | \

sed -r "s/^Grabs$/grab:\n/g" | \

sed -r "s/^Throws$/throw:\n/g" | sed -r "s/^([[:digit:]]+).*% Damage.*$/    base_damage: \1/g" | \

sed -r "s/^Dodges\/Rolls$/dodge:\n/g" | sed -r "s/^Inv\. ?Frames [[:digit:]]+-([[:digit:]]+).*$/    end: \1/g" | \
sed -r "s/^Landing Fall Special Lag: ([[:digit:]]+) Frames.*$/    landing_fall_special: \1/g" | \

sed -r "s/^Jab$/  - name: jab/g" | sed -r "s/^Jab 2$/  - name: jab2/g" | sed -r "s/^Jab 3$/  - name: jab3/g" | \
sed -r "s/^Rapid Jab$/  - name: rapid_jab/g" | sed -r "s/^Forward Tilt$/  - name: forward_tilt/g" | \
sed -r "s/^Up Tilt$/  - name: up_tilt/g" | sed -r "s/^Down Tilt$/  - name: down_tilt/g" | \
sed -r "s/^Dash Attack$/  - name: dash_attack/g" | sed -r "s/^Forward Smash$/  - name: forward_smash/g" | \
sed -r "s/^Up Smash$/  - name: up_smash/g" | sed -r "s/^Down Smash$/  - name: down_smash/g" | \

sed -r "s/^Neutral Air$/  - name: neutral_air/g" | sed -r "s/^Forward Air$/  - name: forward_air/g" | \
sed -r "s/^Back Air$/  - name: back_air/g" | sed -r "s/^Up Air$/  - name: up_air/g" | sed -r "s/^Down Air$/  - name: down_air/g" | \

sed -r "s/^Neutral B$/  - name: neutral_b/g" | sed -r "s/^Aerial Neutral B$/  - name: aerial_neutral_b/g" | \
sed -r "s/^Aerial Side B$/  - name: aerial_side_b/g" | sed -r "s/^Aerial Up B$/  - name: aerial_up_b/g" | \
sed -r "s/^Side B$/  - name: side_b/g" | sed -r "s/^Up B$/  - name: up_b/g" | sed -r "s/^Down B$/  - name: down_b/g" | \
sed -r "s/^Aerial Down B$/  - name: aerial_down_b/g" | \

sed -r "s/^Standing Grab$/  - name: standing_grab/g" | sed -r "s/^Dash Grab$/  - name: dash_grab/g" | \

sed -r "s/^Forward Throw$/  - name: forward_throw/g" | sed -r "s/^Back Throw$/  - name: back_throw/g" | \
sed -r "s/^Down Throw$/  - name: down_throw/g" | sed -r "s/^Up Throw$/  - name: up_throw/g" | \
sed -r "s/^Active Frames -/ACTIVE_FRAMES_HERE/g" | \

sed -r "s/^Spot Dodge$/  - name: spot_dodge/g" | sed -r "s/^Backward Roll$/  - name: backward_roll/g" | \
sed -r "s/^Forward Roll$/  - name: forward_roll/g" | sed -r "s/^Air Dodge$/  - name: air_dodge/g" | \

sed -r "s/^Weight: ([[:digit:]]+).*$/weight: \1/g" | \
sed -r "s/^Fast Fall Speed: ([[:digit:]]+\.[[:digit:]]+).*\$/fastfall_speed: \1/g" | \
sed -r "s/^Dash Speed: ([[:digit:]]+\.[[:digit:]]+).*$/dash_speed: \1/g" | \
sed -r "s/^Run Speed: ([[:digit:]]+\.[[:digit:]]+).*$/run_speed: \1/g" | \
sed -r "s/^Wavedash Length \(Rank\): ([[:digit:]]+).*$/wavedash_length_rank: \1/g" | \
sed -r "s/^PLA Intangibility Frames: ([[:digit:]]+).*$/galint: \1/g" | \
sed -r "s/^Jump Squat: ([[:digit:]]+).*$/jump_squat: \1/g" | \
sed -r "s/^Yes$/true/g" | sed -r "s/^No$/false/g" | sed -r "s/^(true|false)$/walljump: \1/g" | \

sed -r "/^Miscellaneous Info$/d" | sed -r "/^Wall Jump:.*$/d" | sed -r "/^Active Frames 32-$/d" | sed -r "/^[[:space:]]*$/d" | \
sed -r "/^Total Frames$/d" | sed -r "/shield_stun: 0/d" | sed -r "/^.*ACTIVE_FRAMES_HERE$/d" | \

# FIX: Character specific alterations
case "$CHARACTER" in
    "sheik")
        sed -r "s/^None Frames Landing Lag$/    landing_lag: 20/g" | sed -r "s/^-1 L-Cancel Lag$/    lcancel_lag: 10/g"
        ;;
    "young_link")
        sed -r "s/^.*(dash_attack).*$/  - name: dash_attack\n    start: 7\n    end: 12\n    frames: 53\n    iasa: 40\n    shield_stun: 6\n    base_damage: 11\n    weak_damage: 10/g" | \
            sed -r "/% Base Damage/d" | sed -r "/^Frame Startup$/d"
        ;;
    "peach")
        sed -r "/^.*- name: forward_smash$/,+1d" | \
        sed -r "s/^.*% Base Damage$/    frames: 47/g" | sed -r "/^Frame Startup$/d"
        ;;
    "marth"|"roy")
        sed -r "/^.*- name: side_b/d" | sed -r "/^.*% Base Damage$/d" | sed -r "/^Frame Startup$/d"
        ;;
    "mr._game_&_watch")
        sed -r "/^.*- name: side_b/d" | sed -r "/^.*% Base Damage$/d" | sed -r "/^Frame Startup$/d"
        ;;
    "pikachu")
        sed -r "/^.*Active Frames 13-.*$/d" | sed -r "/^Frame Startup$/d"
        ;;
    "jigglypuff")
        sed -r "s/^.*% Base Damage$/    start: 1\n    end: 20\n    base_damage: 10\n    frames: 59/g" | sed -r "/^Frame Startup$/d"
        ;;
    "yoshi")
        sed -r "s/^.*% Base Damage$/    start: 1\n    end: 20\n    base_damage: 4\n    frames: 59/g" | sed -r "/^Frame Startup$/d"
        ;;
    *)
        sed -r "/^Frame Startup$/d"
        ;;
esac
