import requests
import time
headers = {"Accept": "application/json, text/plain, */*",
"Accept-Encoding": "gzip, deflate, br",
"Accept-Language": "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,zh-TW;q=0.6",
"Connection": "keep-alive",
"Cookie": "",
"Host": "api.bilibili.com",
"Origin": "https://www.bilibili.com",
"Referer": "https://www.bilibili.com/anime/index",
"Sec-Fetch-Dest": "empty",
"Sec-Fetch-Mode": "cors",
"Sec-Fetch-Site": "same-site",
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36"}

params = {
"season_version": "-1",
"area": "-1",
"is_finish": "-1",
"copyright": "-1",
"season_status": "-1",
"season_month": "-1",
"year": "-1",
"style_id": "-1",
"order": "3",
"st": "1",
"sort": "0",
"page": "1",
"season_type": "1",
"pagesize": "20",
"type": "1"
}

for i in range(152):
  params["page"] = str(i + 1)
  response = requests.get("https://api.bilibili.com/pgc/season/index/result", params = params, headers = headers)
  print(response.content)
  time.sleep(1)
