import folium.plugins as plugins
import pandas as pd
import folium

def drawSiteLocation(loc, name):
    basemap = folium.Map(location=loc)
    folium.Marker(location=loc)
    basemap.add_child(folium.LatLngPopup())
    basemap.save(name)
loc=[370942.319/10000, 413034.349/10000]
drawSiteLocation(loc, "1.html")
loc=[413034.349/10000, 370942.319/10000]
drawSiteLocation(loc, "2.html")