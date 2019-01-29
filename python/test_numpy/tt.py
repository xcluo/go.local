import math


def sigmoid(Xi):
    return 1.0 / (1 + math.exp(-Xi))

for i in range(1, 151):
    Xi = i / 10.0
    print "Xi=%f, Wi=%d, sigmoid=%f" % (Xi, i+1, sigmoid(Xi))

# print '\n'
# for i in range(1, 11):
#     print "Xi=%f, Wi=%f, sigmoid=%f" % (i, 5, sigmoid(i, 5))
