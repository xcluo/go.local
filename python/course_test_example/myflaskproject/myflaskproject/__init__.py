#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Fuvism <contact@fuvism.com>
# Copyright (c) 2016 - FUVISM
from flask import Flask


app = Flask(__name__)

from flask.ext.sqlalchemy import SQLAlchemy
app.config.from_object('myflaskproject.config')
db = SQLAlchemy(app)

from myflaskproject import views
from myflaskproject import models
