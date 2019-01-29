#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Fuvism <contact@fuvism.com>
# Copyright (c) 2016 - FUVISM
from myflaskproject import db


class User(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    username = db.Column(db.String(64), index = True, unique = True)
    email = db.Column(db.String(120), index = True, unique = True)
    password = db.Column(db.String(64))
    nickname = db.Column(db.String(64))

    def __repr__(self):
        return '<User %r>' % (self.username)
