import requests

to_post = {"name": "w", "pwd": "123"}

r = requests.post("http://127.0.0.1:2333/post", to_post)

print(r.text)