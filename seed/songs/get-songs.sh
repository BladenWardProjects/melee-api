#!/bin/bash

# Song list for Super Smash Bros. Melee
# TODO: Fix the captialization of the origins and durings
curl https://nintendo.fandom.com/wiki/List_of_Super_Smash_Bros._Melee_songs -k -s | grep "<li><b>"| sed -r -e "s/<li><b>(.*)<\/li>/\1/g" -e "s/^(<ol>)?\s*(.*)<\/b> &#8211;(.*)/Title: \2 ---\3/g" -e "s/A.*<i><(a|span).*>(.*)<\/(a|span)>(.*)<\/i>/From: \2 --- \4/g" -e "s/Plays (in)?(on)?( ?the)?(.*)/ PlaysDuring:\4/g" -e "s/(A|a)cts as an(.*)/PlaysDuring:\2/g" -e "s/<(a|span).*>(.*)<\/(a|span)>/\2/g" -e "s/From: (.*) ---.*PlaysDuring:/From: \1 --- PlaysDuring:/g" -e "s/&#160;/ /g" -e "s/:w/: W/g" -e "s/\&amp;/\&/g" -e "s/^(.*)<\/b>.*;/Title: \1 ---/g" -e "s/--- \./---/g" -e "s/---From/--- From/g" -e "s/Flat Zone (.*)--- ,.*/Flat Zone \1--- PlaysDuring: Flat Zone stage/g" -e "s/Pokémon Stadium ---.*/Pokémon Stadium --- From: Pokémon Red, Pokémon Blue, and Pokémon Green --- PlaysDuring: Pokémon Stadium stage/g" -e "s/(Poké Floats --- From:) .* ---/\1 Pokémon Red, Blue, and Green ---/g" -e "s/(From: Earthbound ---).*/\1 PlaysDuring: Fourside stage/g" -e "s/(Icicle Mountain ---).*/\1 From: Ice Climber --- PlaysDuring: Icicle Mountain stage/g" -e "s/(Super Mario Bros\. 3 --- .* ---).*/\1 PlaysDuring: Alternate track to the Yoshi's Island stage, as well as Mario's victory music in 1P Mode/g" -e "s/(Mario's Victory.*PlaysDuring:).*/\1 when Mario, Luigi, Bowser, Dr. Mario or Princess Peach win a battle/g" -e "s/(Zelda Team Victory.*PlaysDuring:).*/\1 when Link, Young Link, Princess Zelda, Sheik, or Ganondorf win a battle/g" -e "s/(Fox's Victory.*PlaysDuring:).*/\1 when Fox McCloud or Falco Lombardi win a battle/g" -e "s/(Pokémon Victory.*From:).*/\1 Pokémon Red, Blue, and Green --- PlaysDuring: when Pikachu, Pichu, Jigglypuff or Mewtwo win a battle/g" -e "s/(Fire Emblem Team Victory.*PlaysDuring:).*/\1 when Marth or Roy win a battle/g" -e "s/(Mr. Game & Watch's Victory.*---) ,.*/\1 PlaysDuring: when Mr. Game \& Watch wins a battle/g" -e "s/(Metal Battle.*Melee ---).*/\1 PlaysDuring: when fighting a Metal character in Classic and Adventure Modes/g" -e "s/(Menu 2 ---).*/\1 From: Super Smash Bros. Melee --- PlaysDuring: alternate menu theme/g" -e "s/(Multi-Man Melee 2 ---).*(PlaysDuring:.*)/\1 From: Super Smash Bros. Melee --- \2/g" -e "s/(Title:) (From: Kirby Super Star.*---).*/\1 All-Star Intro --- \2 PlaysDuring: All-star/g" -e "s/(Classic Intro.*---).*/\1 PlaysDuring: 1P Mode Intro theme/g" -e "/Stage Clear 2/d" -e "s/(Hammer.*---).*/\1 PlaysDuring: when someone gets a Hammer item/g" -e "s/<\/?ol>//g" -e "s/Melee\./Melee/g" -e "s/\.$//g" -e "s/(Fountain of Dreams ---).*/\1 From: Kirby Super Star --- PlaysDuring: Fountain of Dreams stage/g" -e "s/Title: From:/Title: Adventure Intro --- From:/g" -e "s/when/When/g" -e "s/alternate/Alternate/g" -e "s/tournament/Tournament/g" -e "s/\"/'/g" -e "s/Title: (.*) --- From: /- title: \"\1\"\nFrom: /g" -e "s/From: (.*) ---/  origin: \"\1\"\n/g" -e "s/PlaysDuring: (.*)/ plays_during: \"\1\"/g" -e "1s/^/songs:\n/" | yq -o=json > seed/songs/songs.json
