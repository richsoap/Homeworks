import csv
import sys
import requests
import json
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from bs4 import BeautifulSoup
import time
import datetime
import re

# some useful regex
numMatcher = re.compile("[0-9]+")
def getNum(url):
  if len(url) == 0:
    return '0'
  result = numMatcher.findall(url)
  if len(result) == 0:
    return '0'
  else:
    return result[0]

# csv format
mediaInfoHead = ['coins', 'danmakus', 'follow', 'views']
headers = ['id', 'title', 'coins', 'danmakus', 'follow', 'views', 'score', 'score_count', 'tags', 'long_review_count', 'short_review_count', 'state', 'ep', 'timestamp']

# start location
start = 1
if len(sys.argv) != 1:
  start = int(sys.argv[1])

# chrome config
options = Options()
options.headless = True
driver = webdriver.Chrome(executable_path="/home/richsoap/tools/chromedriver/chromedriver", chrome_options=options)

# some API
mediaInfoAPI = "https://api.bilibili.com/pgc/web/season/stat?season_id={}" #Get
ratingInfoAPI = "https://api.bilibili.com/pgc/review/user?media_id={}&ts={}" #Get

# rating is related to timestamp
timestamp = int(time.mktime(datetime.datetime.now().timetuple()))

count = 0
with open('bilibili_bangumi_{}.csv'.format(timestamp), 'w+') as cf:
  with open('extractedlist.txt', 'r') as f:
    csvWriter = csv.DictWriter(cf, headers)
    for url in f:
      count += 1
      if count < start:
        continue
      result = {}
# Get MediaInfo
      mediaInfo = json.loads(requests.get(mediaInfoAPI.format(getNum(url))).content)
      for name in mediaInfoHead:
        result[name] = mediaInfo['result'][name]
# Get Title
      driver.get(url)
      soup = BeautifulSoup(driver.page_source, 'lxml')
      if soup.find(attrs={'class': 'media-title'}) == None:
        continue
      coverInfo = soup.find(attrs = {'class':'media-title'}).attrs
      coverId = getNum(coverInfo['href'])
      result['title'] = coverInfo['title']
      result['id'] = coverId
# Get Rating Info
      reviewInfo = json.loads(requests.get(ratingInfoAPI.format(coverId, timestamp)).content)['result']['media']
      if 'rating' in reviewInfo:
        result['score'] = reviewInfo['rating']['score']
        result['score_count'] = reviewInfo['rating']['count']
      else:
        result['score'] = -1
        result['score_count'] = -1
      result['timestamp'] = timestamp
# Get State Info
      result['ep'] = getNum(reviewInfo['new_ep']['index_show'])
      if result['ep'] == '0':
        result['state'] = 0
      elif reviewInfo['new_ep']['index_show'].startswith('全'):
        result['state'] = 2
      else:
        result['state'] = 1
# Get Review Info
      driver.get("https:" + coverInfo['href'])
      soup = BeautifulSoup(driver.page_source, 'lxml')
      
      lis = soup.find_all("li")
      result['long_review_count'] = 0
      result['short_review_count'] = 0
      for li in lis:
        if li.text.startswith("长评"):
          result['long_review_count'] = getNum(li.text)
        if li.text.startswith("短评"):
          result['short_review_count'] = getNum(li.text)
      
      tagstr = ""
      tags = soup.find_all(attrs={'class':'media-tag'})
      for tag in tags:
        tagstr += (tag.text + " ")
      result['tags'] = tagstr

# Write to CSV
      if count == 1:
        csvWriter.writeheader()
      csvWriter.writerow(result)

      print(count)
      time.sleep(1)
    
driver.quit()
