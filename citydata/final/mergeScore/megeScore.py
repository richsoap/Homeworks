import pandas as pd
import numpy as np
import math

schoolPosition = {'Longitude':-2.233771, 'Latitude':53.46679}

rc = 6378.137
rj = 6356.725

def ToRad(deg):
  return deg * math.pi / 180

def GetR(lat):
  return rj + (rc - rj) * (90-lat) / 90

def One2OneDistanceSquare(a, b):
  degree = a['Latitude']
  r = GetR(degree)
  sr = r * math.cos(ToRad(degree))

  deltaLatitude = ToRad(a['Latitude'] - b['Latitude']) 
  deltaLongitude = ToRad(a['Longitude'] - b['Longitude']) 
  return (deltaLatitude * r)**2 + (deltaLongitude * sr)**2

def One2OneDistance(a, b):
  return math.sqrt(One2OneDistanceSquare(a, b))

def OneDistanceSqure(point, targetsdf):
  result = np.zeros(len(targetsdf))
  for index, row in targetsdf.iterrows():
    result[index] = One2OneDistanceSquare(point, row)
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
  areas = np.zeros(len(airbnb))
  minindexs = [0]*len(airbnb)
  score = LoadScore(crime, 'score_0')
  area = LoadScore(crime, 'area')
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
    areas[index] = area[index]
    for i in range(len(minindex)):
      if minindex[i]:
        minindexs[index] = i
        break
  return result, areas, minindexs

def CommuteScoreFunc(distance):
  if distance < 0.1:
    return 100
  elif distance < 0.4:
    return 80
  elif distance < 1:
    return 70
  elif distance < 3:
    return 50
  else:
    return 50 * 9 / (distance**2)

def GetCommuteScore(airbnb, school):
  result = np.zeros(len(airbnb))
  distance = [0]*len(airbnb)
  for index, row in airbnb.iterrows():
    distance[index] = One2OneDistance(row, school)
    result[index] = CommuteScoreFunc(distance[index])
  return result, distance

def GetBusDistance(airbnb, bus):
  result = np.zeros(len(airbnb))
  tempdf = []
  for index, row in airbnb.iterrows():
    if index % 10 == 0:
      print("{}/{}".format(index, len(airbnb)))
    distance = OneDistance(row, bus)
    result[index] = np.min(distance)
    row['bus distance'] = result[index]
    tempdf.append(row)
  pd.DataFrame(tempdf).to_csv("distance.csv")
  return result

def GetBusScore(airbnb, bus):
  distance = GetBusDistance(airbnb, bus)
  #print(distance)
  allScore = 50*(0.1)/distance
  allScore[allScore>=50] = 50
  return allScore, distance

def GetRoomScore(airbnb):
  result = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    result[index] = row['bed_score'] * 100
  return result

def GetPrice(airbnb):
  result = [0]*len(airbnb)
  for index, row in airbnb.iterrows():
    result[index] = row['price']
  return result

def GetPriceScore(airbnb):
  price = GetPrice(airbnb)
  #print('price {}'.format(price))
  delta = np.max(price) - np.min(price)
  #print('delta {}'.format(delta))
  score = price - np.min(price)
  #print('score {}'.format(score))
  score = score / delta
  #print('score {}'.format(score))
  score = 1 - score
  return score * 100

def GetPriceBonus(airbnb, minindexs):
  pricesCount = {}
  prices = GetPrice(airbnb)
  for index in range(len(prices)):
    minIndex = minindexs[index]
    if minIndex not in pricesCount:
      pricesCount [minIndex] = [0,0]
    pricesCount[minIndex][0] += prices[index]
    pricesCount[minIndex][1] += 1
  
  for key in pricesCount:
    pricesCount[key][0] /= pricesCount[key][1]

  result = [0]*len(prices)
  average = [0]*len(prices)

  for index,row in airbnb.iterrows():
    minIndex = minindexs[index]
    average[index] = pricesCount[minIndex][0]
    price = prices[index]
    if price > average[index]:
      result[index] = max(30 - (price - average[index]), 0)
    else:
      result[index] = min(30 + (average[index] - price), 50)
  return result, average

def GetRating(airbnb):
  rating = np.zeros(len(airbnb))
  for index, row in airbnb.iterrows():
    rating[index] = row['rating']
  return rating

def GetSelfScore(airbnb, bus):
  busScore = GetBusScore(airbnb, bus)
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
#airbnbData = pd.read_csv("airbnb_short.csv")

crimeData = pd.read_csv("crime.csv")
busData = pd.read_csv("bus.csv")

scoreDict = {}
print('load data')

crimeScoreData = pd.read_csv('crimescore.csv')
crimeScore = crimeScoreData['crime']
areaScore = crimeScoreData['area']
minindexs = crimeScoreData['minindexs']

#crimeScore, areaScore, minindexs = GetCrimeScore(airbnbData, crimeData)
scoreDict['crime'] = crimeScore
scoreDict['area'] = areaScore
scoreDict['minindexs'] = minindexs

print("crime score: {}".format(len(crimeScore)))
df = pd.DataFrame(scoreDict)
df.to_csv("crimescore.csv")


commuteScore, distance = GetCommuteScore(airbnbData, schoolPosition)
scoreDict['commute'] = commuteScore
scoreDict['commute distance'] = distance
print('commute score: {}'.format(len(commuteScore)))

busScore, distance = GetBusScore(airbnbData, busData)
scoreDict['bus'] = busScore
scoreDict['bus distance'] = distance
print("bus bonus: {}".format(len(busScore)))

roomScore = GetRoomScore(airbnbData)
scoreDict['room'] = roomScore
print('room score: {}'.format(len(roomScore)))

priceScore = GetPriceScore(airbnbData)
scoreDict['price score'] = priceScore
print('price score: {}'.format(len(priceScore)))

priceBonus, average = GetPriceBonus(airbnbData, minindexs)
scoreDict['price bonus'] = priceBonus
scoreDict['average'] = average
print('price bonus: {}'.format(len(priceBonus)))

rating = GetRating(airbnbData)
scoreDict['rating'] = rating
print('rating: {}'.format(len(rating)))

totalScore = rating * (commuteScore + priceBonus + busScore + roomScore + priceScore - crimeScore) / 100
#totalScore = rating * (commuteScore + priceBonus + roomScore + priceScore - crimeScore) / 100

totalScore[totalScore < 0] = 0
scoreDict['total'] = totalScore
print('total score: {}'.format(len(totalScore)))
mergeDf = MergeScore(airbnbData, scoreDict)
print(mergeDf.head())
mergeDf.to_csv('totalscore.csv')
