#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>

from flask import Flask
from flask import url_for
from flask import render_template

app = Flask(__name__)


@app.route("/")
def index():
    return render_template('hello.html', name=name)


@app.route("/home")
def home():
    return render_template('home.html')


@app.route("/api/<str:api_moduler>/<str:api_func>/", methods=['GET', 'POST'])
def api():
    return "API"
