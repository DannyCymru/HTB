#!/bin/python

import requests
import hashlib
from bs4 import BeautifulSoup

url="http://94.237.59.174:34311"

#Proxy requests through burp for debugging
#proxies = {"http": "http://127.0.0.1:8080", "https": "http://127.0.0.1:8080"}

get = requests.get(url,
					#proxies=proxies
					)

cookies = get.cookies

bsoup = BeautifulSoup(get.content, "html.parser")

s2e = bsoup.h3.text

hashed = hashlib.md5(s2e.encode()).hexdigest()

data = {"hash":hashed}

post = requests.post(url,
					#proxies=proxies, 
					cookies=cookies, data=data
					)

flag = BeautifulSoup(post.content, "html.parser")

print(flag.p.text)
