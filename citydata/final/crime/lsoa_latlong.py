import os
import pandas as pd

result = []
lsoaprefix = 'lsoa:'
latprefix = '    geo:lat '
longprefix = '    geo:long '
with open("./data/lsoa.ttl/lsoa.ttl", 'r') as f:
    lsoaname = ""
    lat = 0
    lon = 0
    for line in f:
        if line.startswith(lsoaprefix):
            lsoaname = line[len(lsoaprefix):-1]
        elif line.startswith(latprefix):
            lat = float(line[len(latprefix):-2])
        elif line.startswith(longprefix):
            lon = float(line[len(longprefix):-2])
            result.append({'lsoa': lsoaname, 'lat': lat, 'lon': lon})
result = pd.DataFrame(result)
result.to_csv("lsoa_position.csv")
