# -*- coding: utf-8 -*-

import requests
import json

secret = {'corpid': '==', 'corpsecret': '=='}

tk = requests.get('https://qyapi.weixin.qq.com/cgi-bin/gettoken', params=secret)

print(tk.text)

tokenj = json.loads(tk.text)
token = tokenj['access_token']

# token = '=='

# tag = {"tagname": "部门老总"}

# r = requests.post('https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token={}'.format(token), data=json.dumps(tag))

# print(r.text)

req = {'access_token': token,
       'agentid': '=='
}

r = requests.get('https://qyapi.weixin.qq.com/cgi-bin/agent/get', params=req)
print(r.text)
