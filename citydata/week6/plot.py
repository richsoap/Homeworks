import folium.plugins as plugins
import pandas as pd
import seaborn as sns
import folium
import matplotlib.pyplot as plt

siteList = ['1418A', '3015A', '3133A', '3014A', '1419A']
siteRecord = {}
df = pd.DataFrame()
siteFile = 'sites.csv'
dataFile = 'merged_daily_chongqing.csv'
cityname = '重庆'

# startdates = [20180604, 20180611, 20180618]
startdates = [20190610, 20190617]
#startdates = [20191202, 20191209, 20191216]
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

def readData():
    prepareSiteRecord()
    data = pd.read_csv(dataFile)
    if startdate != 0:
        data = data[(data['type'] == polution) & (data['date'] >= startdate) & (data['date'] < enddate)]
    else:
        data = data[data['type'] == polution]
    data = data.fillna(1) # for heatmap with time, value should in (0, 1], Maybe use average is better
    data = data.sort_values(['date', 'hour'])
    result = pd.DataFrame()
    count = 0
    preval = {}
    for _, row in data.iterrows():
        count += 1
        print(count)
        for site in siteRecord:
            se = pd.Series()
            se['site'] = site
            se['date'] = row['date']
            se['hour'] = row['hour']
            if row[site] == 1 and site in preval:
                se['value'] = preval[site]
            else:
                se['value'] = row[site]
            preval[site] = se['value']
            se['index'] = count
            result = result.append(se, ignore_index=True)
    return result

def mainfunc():
    daycsv = readData()
    print(daycsv)
    sns.lineplot(x="index", y="value", hue="site", data=daycsv)
    plt.show()

for sdate in startdates:
    startdate = sdate
    enddate = sdate + 7
    siteMapPath = 'site_{}.html'.format(cityname)
    heatMapPath = 'heat_{}_{}_{}.html'.format(cityname, polution, startdate)
    heatMapWithTimePath = 'heat_time_{}_{}_{}.html'.format(cityname, polution, startdate)
    mainfunc()