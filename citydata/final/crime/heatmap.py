import pandas as pd

totaldata = pd.read_csv("greater-manchester-street-merge.csv")

def GetLocationMap(data):
    return data[['Month', 'Crime type', 'Longitude', 'Latitude']]

MapData = GetLocationMap(totaldata)
MapData.to_csv("crime_loction.csv")
