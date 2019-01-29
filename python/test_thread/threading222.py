#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Author: Fuvism Team <contact@fuvism.com>
# Copyright (c) 2016 - FUVISM
import logging
import threading
import time


# coding here
logging.basicConfig(
    level=logging.DEBUG,
    format='[%(levelname)s] (%(threadName)-10s) %(message)s',
    filename='/tmp/testlog.log',
    filemode='w'
    )   

 
def worker():
    logging.debug("Fuvism Threading")
    time.sleep(1)
    return


if __name__ == "__main__":
    t = threading.Thread(target=worker)
    t.start()
