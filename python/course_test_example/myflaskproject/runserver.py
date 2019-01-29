#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Fuvism <contact@fuvism.com>
# Copyright (c) 2016 - FUVISM
from myflaskproject import app


app.run(host='0.0.0.0', port=5000, threaded=True, debug=True)
