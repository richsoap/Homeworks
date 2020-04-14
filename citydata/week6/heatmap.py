import folium.plugins as plugins
import pandas as pd
import folium

#siteList = ['1418A', '3015A', '3133A', '3014A', '1419A']
siteList = []
siteRecord = {}
df = pd.DataFrame()
siteFile = 'sites.csv'
dataFile = 'merged_daily_chongqing.csv'
cityname = '重庆'

startdates = [20190610, 20190617]
startdate = 0
enddate = 0
polution = 'AQI'

siteMapPath = 'site_{}.html'.format(cityname)
heatMapPath = 'heat_{}_{}_{}.html'.format(cityname, polution, startdate)
heatMapWithTimePath = 'heat_time_{}_{}_{}.html'.format(cityname, polution, startdate)

def prepareSiteRecord():
    global df
    df = pd.read_csv(siteFile, encoding='gbk')
    df = df[df['city'] == cityname]
    df['longitude'] = df['longitude'].astype('float')
    df['latitude'] = df['latitude'].astype('float')
    if len(siteList) == 0:
        for _, row in df.iterrows():
            siteRecord[row['code']] = [row['latitude'], row['longitude']]
    else:
        for site in siteList:
            siteLoc = df[df['code'] == site]
            for _, row in siteLoc.iterrows():
                siteRecord[row['code']] = [row['latitude'], row['longitude']]

def assembleData(site, val):
    if len(siteRecord) == 0:
        prepareSiteRecord()
    return siteRecord[site] + [val]

def drawSiteLocation():
    basemap = folium.Map(location=[df['latitude'].mean(), df['longitude'].mean()])
    for site in siteRecord:
        folium.Marker(location=siteRecord[site]).add_to(basemap)
    basemap.add_child(folium.LatLngPopup())
    basemap.save(siteMapPath)

def drawHeatMap(data):
    basemap = folium.Map(location=[df['latitude'].mean(), df['longitude'].mean()])
    heat = plugins.HeatMap(data)
    heat.add_to(basemap)
    basemap.save(heatMapPath)

# 需要归一化
def drawHeatMapWithTime(data, r = 50):
    basemap = folium.Map(location=[df['latitude'].mean(), df['longitude'].mean()])
    heat = plugins.HeatMapWithTime(data, radius=r, speed_step=2, min_speed=1)
    heat.add_to(basemap)
    basemap.save(heatMapWithTimePath)

def readData():
    prepareSiteRecord()
    drawSiteLocation()
    data = pd.read_csv(dataFile)
    if startdate != 0:
        data = data[(data['type'] == polution) & (data['date'] >= startdate) & (data['date'] < enddate)]
    else:
        data = data[data['type'] == polution]
    data = data.sort_values(['date', 'hour'])
    data = data.fillna(1) # for heatmap with time, value should in (0, 1]
    return data

def mainfunc():
    daycsv = readData()
    maxCol = daycsv.max()
    maxVal = -1
    errsites = []
    for site in siteRecord:
        try:
            maxVal = max(maxVal, maxCol[site])
        except:
            print("{} is not existed, delete later".format(site))
            errsites.append(site)
    for errsite in errsites:
        del siteRecord[errsite]

    daydata = []
    for _, row in daycsv.iterrows():
        hourdata = []
        for site in siteRecord:
            hourdata.append(assembleData(site, row[site]/maxVal))
        daydata.append(hourdata)
    drawHeatMap(daydata[0])
    print(len(daydata))
    drawHeatMapWithTime(daydata)

for sdate in startdates:
    startdate = sdate
    enddate = sdate + 7
    siteMapPath = 'site_{}.html'.format(cityname)
    heatMapPath = 'heat_{}_{}_{}.html'.format(cityname, polution, startdate)
    heatMapWithTimePath = 'heat_time_{}_{}_{}.html'.format(cityname, polution, startdate)
    mainfunc()