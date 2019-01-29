#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Fuvism <contact@fuvism.com>
# Copyright (c) 2016 - FUVISM
import os


basedir = os.path.abspath(os.path.dirname(__file__))

SQLALCHEMY_DATABASE_URI = 'sqlite:///' + os.path.join(basedir, 'db.sqlite3')
SQLALCHEMY_MIGRATE_REPO = os.path.join(basedir, 'db_repository')
