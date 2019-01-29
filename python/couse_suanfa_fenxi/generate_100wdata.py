# -*- coding:utf-8 -*-
#!/usr/bin/python

import random

fp = open('data.txt', 'w+')

tt = []
for i in range(10000*100):
    num = random.randint(10000, 10000*100*10000)
    #tt.append(str(num))
    fp.write(str(num)+"\n")

fp.close()


