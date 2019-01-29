import datetime
from test_celery.test_celery import mytask, mytask2

for i in range(4):
    task_id = mytask.apply_async((i,))
    print 'task_id = ', task_id

# print 'start_at = ', datetime.datetime.now()
# task_id = mytask2.apply_async(countdown=5)
# print 'task_id = ', task_id
