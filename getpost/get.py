import requests

req = requests.get("http://127.0.0.1:2333")

print(req.text)