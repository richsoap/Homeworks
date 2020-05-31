import pandas as pd
import numpy as np
import math


DistanceUnit = 6371**2
def OneDistanceSqure(point, targetsdf):
  result = np.zeros(len(targetsdf))
  for index, row in targetsdf.iterrows():
    degree = point['Latitude']
    radians = degree * math.pi / 180
    deltaLatitude = (point['Latitude'] - row['Latitude']) * math.pi / 180
    deltaLongitude = (point['Longitude'] - row['Longitude']) * math.pi / 180
    result[index] = deltaLatitude**2 + (deltaLongitude*math.cos(radians))**2
  result = result * DistanceUnit
  return result

def OneDistance(point, targets):
  result = OneDistanceSqure(point, targets)
  return np.sqrt(result)

def LoadScore(crime, key):
  result = np.zeros(len(crime))
  for index, row in crime.iterrows():
    result[index] = row[key]
  return result

def GetCrimeScore(airbnb, crime):
  print('build radis data')
  result = np.zeros(len(airbnb))
  score = LoadScore(crime, 'score_0')
  meanscore = np.mean(score)
  k = 50/meanscore
  print('load score data')
  for index, row in airbnb.iterrows():
    if index % 100 == 0:
      print("{}/{}".format(index, len(airbnb)))
    distance = OneDistanceSqure(row, crime)
    minindex = distance == np.min(distance)
    values = score[minindex] * k
    result[index] = np.sum(values)
  return result

def GetBusDistance(airbnb, bus):
  result = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    if index % 10 == 0:
      print("{}/{}".format(index, len(airbnb)))
    distance = OneDistance(row, bus)
    result[index] = np.min(distance)
  return result

def GetBusScore(airbnb, bus):
  distance = GetBusDistance(airbnb, bus)
  #print(distance)
  allScore = 100*(0.1)/distance
  allScore[allScore>100] = 100
  return allScore

def GetRoomScore(airbnb):
  result = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    result[index] = row['bed_score'] * 100
  return result

def GerPrice(airbnb):
  result = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    result[index] = row['price']
  return result

def GetPriceScore(airbnb):
  price = GerPrice(airbnb)
  #print('price {}'.format(price))
  delta = np.max(price) - np.min(price)
  #print('delta {}'.format(delta))
  score = price - np.min(price)
  #print('score {}'.format(score))
  score = score / delta
  #print('score {}'.format(score))
  score = 1 - score
  return score * 100

def GetRating(airbnb):
  rating = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    rating[index] = row['rating']
  return rating

def GetSelfScore(airbnb, bus):
  busScore = GetBusScore(airbnbData, busData)
  print('load bus score')
  rating = np.zeros(len(airbnb))
  bed = np.zeros(len(airbnb))
  price = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    rating[index] = row['rating']
    bed[index] = row['bed_score']
    price[index] = row['price']
  return rating * busScore * bed / price

airbnbkeys = ['beds','bed_type','bed_score','rating','price','Latitude','Longitude','name','host_id']
def MergeScore(airbnb, scores):
  result = []
  for index, row in airbnb.iterrows():
    result.append(row)
    for key in scores:
      result[index][key] = scores[key][index]
  return pd.DataFrame(result)

print("start")
airbnbData = pd.read_csv("airbnb.csv")
crimeData = pd.read_csv("crime.csv")
busData = pd.read_csv("bus.csv")

scoreDict = {}
print('load data')
crimeScore = GetCrimeScore(airbnbData, crimeData)
scoreDict['crime'] = crimeScore
print("crime score: {}".format(len(crimeScore)))
busScore = GetBusScore(airbnbData, busData)
scoreDict['bus'] = busScore
print("bus score: {}".format(len(busScore)))
roomScore = GetRoomScore(airbnbData)
scoreDict['room'] = roomScore
print('room score: {}'.format(len(roomScore)))
priceScore = GetPriceScore(airbnbData)
scoreDict['price'] = priceScore
print('price score: {}'.format(len(priceScore)))
rating = GetRating(airbnbData)
scoreDict['rating'] = rating
print('rating: {}'.format(len(rating)))
totalScore = rating * (busScore + roomScore + priceScore - crimeScore) / 100
totalScore[totalScore < 0] = 0
scoreDict['total'] = totalScore
print('total score: {}'.format(len(totalScore)))
mergeDf = MergeScore(airbnbData, scoreDict)
print(mergeDf.head())
mergeDf.to_csv('totalscore.csv')