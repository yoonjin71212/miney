import requests
print("Command?:")
command = input()
if command=="generate":
    r = requests.post ('http://daegu.yjlee-dev.pe.kr:35000/create', json = { "server-name" : "Miney", "gamemode" : "creative", "force-gamemode" : True, "difficulty" : "easy", "allow-cheat" : True, "max-players" : 10, "online-mode" : True, "white-list" : False, "server-port" : 19132, "server-portv6" : 19133, "view-distance" : 32, "tick-distance" : 4, "player-idle-timeout" : 30, "max-threads" : 8, "level-name" : "Bedrock level", "level-seed" : "MineCraftX", "default-player-permission-level" : "operator", "texturepack-required" : False, "content-log-file-enabled" : True, "compression-threshold" : 20, "server-authoritative-movement" : "server-auth", "player-movement-score-threshold" : 0.85, "player-movement-distance-threshold" : 0.3, "player-movement-duration-threshold-in-ms" : 500, "correct-player-movement" : True, "server-authoritative-block-breaking" : True })
    print(r.json())
elif command=="stop":
    print("Tag?:")
    tag = input()
    r = requests.post ('http://daegu.yjlee-dev.pe.kr:35000/delete', data=tag)
