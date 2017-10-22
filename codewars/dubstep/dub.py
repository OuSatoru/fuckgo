import re

def song_decoder(song):
    wub = re.compile(r"(WUB)+")
    return re.sub(wub, " ", song).strip()

print(song_decoder("WUBWEWUBAREWUBWUBTHEWUBCHAMPIONSWUBMYWUBFRIENDWUB"))
