import datetime
import time
from asyncgo import app

# import ipdb; ipdb.set_trace()
# app.conf.update(CELERY_IMPORTS=('test_celery.test_celery',))
@app.task(name="test_celery.test_celery.mytask", bind=True, time_limit=180)
def mytask(self, num):
    time.sleep(3)
    return 'aaaaaaaaaaaaaaaaaaaa num=%d'% num

@app.task(bind=True, time_limit=180)
def mytask2(self):
    return 'end_at = ', str(datetime.datetime.now())


# task.apply_async(args=[])
if __name__ == '__main__':
    print app.conf.CELERY_IMPORTS
    # mytask.delay()
