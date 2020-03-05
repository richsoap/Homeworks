# B站爬取计划
## 爬取内容
### 播放页
[x] 番剧ID(id)
[x] 番剧名称(title)
[x] 播放量(views)
[x] 弹幕数目(danmakus)
[x] 追番量(follow)
[x] 硬币数(coins)
[ ] 分享量(share)
[x] 评分(score)
[x] 评分人数(score_count)
[x] 时间戳(timestamp)
### 详情页
[x] Tag(tag)
[x] (长短)点评数量(long_review_count/short_review_count)
[x] 番剧状态[完结/连载](finish)
[x] 集数(ep)
## 数据获取
### 播放页
1. 遍历(40w), 不好
2. 从索引页找, 棒
#### 关键字段
- mediaInfo: 播放量(views), 弹幕数(danmakus), 追番量(favorites), 转发量(share), 番剧名(title)
- rating: 评分(score), 评分人数(count)
#### 详情页
从class: media-title的超链接中跳转
#### 关键字段
- media-tag: 有很多个，都是该番的Tag
- clearfix: 长评和短评
- index_show: 全xxx话/更新至第xxx话
## 操作流程
### 先爬取所有番剧的url
### 再爬取具体内容
## 问题
没有抓取到分享量获取所用的HTTP请求，还需要更多研究
