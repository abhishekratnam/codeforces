import re
import urllib
import requests
from bs4 import BeautifulSoup
url = ["http://gadyakosh.org/gk/%E0%A4%B2%E0%A5%80%E0%A4%B2%E0%A4%BE%E0%A4%A7%E0%A4%B0_%E0%A4%AE%E0%A4%82%E0%A4%A1%E0%A4%B2%E0%A5%8B%E0%A4%88_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%AA%E0%A4%82%E0%A4%95%E0%A4%9C_%E0%A4%AC%E0%A4%BF%E0%A4%B7%E0%A5%8D%E0%A4%9F_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%A1%E0%A5%89._%E0%A4%A4%E0%A5%81%E0%A4%B2%E0%A4%B8%E0%A5%80%E0%A4%B0%E0%A4%BE%E0%A4%AE_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%B5%E0%A4%BF%E0%A4%B7%E0%A5%8D%E0%A4%A3%E0%A5%81_%E0%A4%A8%E0%A4%BE%E0%A4%97%E0%A4%B0_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_-_2_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%B5%E0%A4%BF%E0%A4%B7%E0%A5%8D%E0%A4%A3%E0%A5%81_%E0%A4%96%E0%A4%B0%E0%A5%87_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%B8%E0%A4%B5%E0%A4%BF%E0%A4%A4%E0%A4%BE_%E0%A4%B8%E0%A4%BF%E0%A4%82%E0%A4%B9_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2","http://gadyakosh.org/gk/%E0%A4%B5%E0%A4%BF%E0%A4%B7%E0%A5%8D%E0%A4%A3%E0%A5%81_%E0%A4%A8%E0%A4%BE%E0%A4%97%E0%A4%B0_%E0%A4%B8%E0%A5%87_%E0%A4%AC%E0%A4%BE%E0%A4%A4%E0%A4%9A%E0%A5%80%E0%A4%A4_/_%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2"]
# url = []
results = []
# status, response = requests.get('http://gadyakosh.org/gk/%E0%A4%95%E0%A5%81%E0%A4%AE%E0%A4%BE%E0%A4%B0_%E0%A4%AE%E0%A5%81%E0%A4%95%E0%A5%81%E0%A4%B2')
# for link in BeautifulSoup(response, parse_only=SoupStrainer('a')):
#     if link.has_attr('href'):
#         url.append(link['href'])
for i in url:
  resp = requests.get(i)
  if resp.status_code == 200:
      soup = BeautifulSoup(resp.content, "html.parser")
      title =  soup.find_all('span',class_="gk-article-cust-heading")
      subtitle = soup.find_all('span' ,class_="gk-article-cust-subheading") 
      results.append(title)
      results.append(subtitle)
      for para in soup.find_all('p'):            
          results.append(para.get_text())
for i in results:
  print(i," ")
    

