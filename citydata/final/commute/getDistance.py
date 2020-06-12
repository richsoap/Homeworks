import requests
import re
import pandas as pd
import time

airbnb = pd.read_csv('../mergeScore/airbnb.csv')
school = {'Longitude':-2.233771, 'Latitude':53.46679}
result = []
def getFloat(text):
    num = re.findall('\d*\.\d*公里', text)
    try:
        return num[0][:-2]
    except Exception:
        return -1
for index, row in airbnb.iterrows():
    kw = {'long1': school['Longitude'],
    'lat1': school['Latitude'],
    'long2': row['Longitude'],
    'lat2': row['Latitude']}
    response = requests.get('http://tool.yovisun.com/longlat/index.php?', params=kw)
    result.append({'distance': getFloat(response.text)})
    time.sleep(0.05)

df = pd.DataFrame(result)
df.to_csv('distance.csv')