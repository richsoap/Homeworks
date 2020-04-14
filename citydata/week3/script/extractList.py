import re
urlmatcher = re.compile("https://www.bilibili.com/bangumi/play/[a-zA-Z0-9]+")
with open("rawlist.txt") as f:
  for line in f:
    urlresult = urlmatcher.findall(line)
    for url in urlresult:
      print(url)
