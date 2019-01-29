#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: SimonLuo <simonluo@thstack.com>

import time
from celery import Celery

celery = Celery('myapp', broker='redis://localhost:6379/0') # , backend='redis://localhost:6379/0')
#celery.conf.update(CELERY_RESULT_BACKEND = 'redis://localhost:6379/0')
celery.conf.update(
        {'CELERY_RESULT_BACKEND': 'redis://localhost:6379/0'})
# celery.conf.CELERY_RESULT_BACKEND = 'redis://localhost:6379/0'

@celery.task(bind=True)
def add(self, x, y):
    time.sleep(60)
    return x + y


if __name__ == '__main__':
    add.apply_async(args=[1, 2])
