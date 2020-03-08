# B站爬取计划
试图爬取B站所持有的正版番的相关播放信息，例如ID，播放量，追番数等。
## 爬取内容与数据格式
### 播放页
- [x] 番剧ID(id): int
- [x] 番剧名称(title): string
- [x] 播放量(views): int
- [x] 弹幕数目(danmakus): int
- [x] 追番量(follow): int
- [x] 硬币数(coins): int
- [ ] 分享量(share): int
- [x] 评分(score): double
- [x] 评分人数(score_count): int
- [x] 时间戳(timestamp): unix_timestamp
- [x] Tag(tag): string(使用空格分隔)
- [x] (长短)点评数量(long_review_count/short_review_count): int
- [x] 番剧状态(finish): 0为尚未开播，1为连载，2为完结
- [x] 集数(ep): int
## 数据获取
### 播放页链接获取
1. 遍历40w番剧链接������
2. 利用索引页的API找到番剧的ID������
#### 播放页可获取的信息
- mediaInfo: 播放量(views), 弹幕数(danmakus), 追番量(favorites), 转发量(share), 番剧名(title)
- rating: 评分(score), 评分人数(count)
- 详情页的链接
#### 详情页
#### 可获取的信息
- media-tag: 有很多个，都是该番的Tag
- clearfix: 长评和短评
- index_show: 全xxx话/更新至第xxx话
## 操作流程
1. 使用API获得所有的番剧播放页链接
2. 从播放页中获取详情页链接
3. 并从播放页中信息调用相关API或强行正文正则表达以获取相关信息
4. 从详情页的正文中正则表达获得相关信息
## 注意事项
1. 没有抓取到分享量获取所用的HTTP请求
2. 代码鲁棒性不太行，偶尔会因为获取不到内容而崩溃，再运行一次即可
3. 未开播的番的评分等信息使用-1表示
4. 部分港澳台专有番会被404
