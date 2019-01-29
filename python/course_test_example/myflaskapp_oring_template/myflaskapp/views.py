#!/usr/bin/python
#!coding:utf-8
from myflaskapp import app

@app.route('/')
# @app.route('/index')
def index():
    return "Hello, World!"
