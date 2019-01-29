#!/usr/bin/python
#!coding:utf-8
from myflaskapp import app


app.debug = True                                             
app.run(host='0.0.0.0', port=5000)
