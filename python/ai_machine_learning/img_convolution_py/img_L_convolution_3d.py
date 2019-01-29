import sys
from PIL import Image
import numpy as np

print("Convolution(Image) with 3*3 Kernel......")
im = Image.open('./test.png').convert('L')
print 'Image Size: ', im.size
# im.show()
# im.save('./test_L.png')

img = im.load()
M_kernel = np.mat([[1, 1, 1], [0, 0, 0], [-1, -1, -1]])
print 'Kernel Shape: ', M_kernel.shape
print M_kernel

# Convolution
total = 500 - 3 + 1     # n - m + 1
print('Total computing(n-m+1): %d'% total)
L_rest = [[0 for j in range(total)] for i in range(total)]

for i in range(total-2):
	print('\rConvolution Run %d/%d'%(i, total)),
	sys.stdout.flush()

	for j in range(total):
		point = M_kernel[0, 0] * img[i, j] + \
			M_kernel[0, 1] * img[i, j+1] + \
			M_kernel[0, 2] * img[i, j+2] + \
			M_kernel[1, 0] * img[i+1, j] + \
			M_kernel[1, 1] * img[i+1, j+1] + \
			M_kernel[1, 2] * img[i+1, j+2] + \
			M_kernel[2, 0] * img[i+3, j] + \
			M_kernel[2, 1] * img[i+3, j+1] + \
			M_kernel[2, 2] * img[i+3, j+2] 
		
        L_rest[i][j] = point

print '\nConvolution Results Shape: ', np.matrix(L_rest).shape

for i in range(len(L_rest)):
    line = L_rest[i]
    for j in range(len(line)):
        img[i, j] = L_rest[i][j]
# im.show('abcd')
