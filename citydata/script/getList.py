import requests
import time
headers = {"Accept": "application/json, text/plain, */*",
"Accept-Encoding": "gzip, deflate, br",
"Accept-Language": "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,zh-TW;q=0.6",
"Connection": "keep-alive",
"Cookie": "buvid3=234EC8A7-E103-4C59-9EEE-98AE42DA219F110266infoc; LIVE_BUVID=AUTO4015555617493776; sid=c9bgaekm; stardustvideo=1; CURRENT_FNVAL=16; rpdid=|(J~|~Ju|)|k0J'ullY|~m~mk; im_notify_type_18810060=0; fts=1558066919; im_seqno_18810060=24; im_local_unread_18810060=0; pgv_pvi=3242741760; _uuid=CB725600-BB55-8D76-8C7E-C70FEE0187E088633infoc; UM_distinctid=16e34d14f7c1f7-07e2b253ad069c-7711439-16e360-16e34d14f7d34d; laboratory=1-1; CURRENT_QUALITY=64; DedeUserID=18810060; DedeUserID__ckMd5=e02d7297369f3808; SESSDATA=b6fccab7%2C1585761652%2Ce269d331; bili_jct=99bab3000dd045ebeeb6040dffa275cb; bp_t_offset_18810060=362845391290238108; INTVER=1; stardustpgcv=0606",
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
