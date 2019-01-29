#coding:utf-8                                                                   

import requests                                                                 

url = "http://sendcloud.sohu.com/webapi/mail.send.json"                         

API_USER = 'Fuvism_no_reply'
API_KEY = 'EW0IkG194DHhb6c1'

params = {                                                                      
    "api_user": API_USER, # 使用api_user和api_key进行验证                       
    "api_key" : API_KEY,                                             
    "to" : "simonluo@fuvism.com;xcluo@thstack.com", # 收件人地址, 用正确邮件地址替代, 多个地址用';'分隔  
    "from" : "no-reply@fuvism.com", # 发信人, 用正确邮件地址替代     
    "fromname" : "no-reply",                                                    
    "subject" : "Hi ",
    "html": "你 %content% 太棒了！你已成功的从SendCloud发送了一封测试邮件，接下来快登录前台去完善账户信息吧！",
    "resp_email_id": "true",
}                                                                               

r = requests.post(url, data=params)

print r.text

