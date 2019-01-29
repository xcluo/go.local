#!/usr/bin/python
#!coding:utf-8
from myflaskproject import app, db
from myflaskproject import models
from flask import render_template

@app.route('/')
def home():
    return "Hello, World! <a href='/index/'>Index</a>"

@app.route('/index/')
def index():
    u = models.User(username='john', password='123456', nickname='John', email='john@email.com')
    db.session.add(u)
    db.session.commit()
    return 'Success!'

@app.route('/lists/')
def lists():
    # u = models.User(nickname='john', email='john@email.com')
    users = db.session.query(models.User).all()
    return render_template('lists.html', users=users)
