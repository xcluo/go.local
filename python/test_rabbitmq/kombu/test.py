#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Longgeek <longgeek@gmail.com>

import kombu_rpc

kombu_rpc.cast("{'a': 1}")
#result = kombu_rpc.call('{"b": 2}')
#print result
