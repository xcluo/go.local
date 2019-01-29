import sys
from PIL import Image
import numpy as np


def run_convolution(L_img, L_kernel, title_prex=0):
    M_img = np.mat(L_img)
    print 'Matrix Image Shape: ', M_img.shape

    M_kernel = np.mat(L_kernel)
    print 'Matrix Kernel Shape: ', M_kernel.shape
    print M_kernel

    # Computing
    img_m, img_n = M_img.shape
    ker_m, ker_n = M_kernel.shape
    con_m, con_n = img_m - ker_m + 1, img_n - ker_n + 1
    L_rest = [[0 for n in range(con_n)] for m in range(con_m)]
    for m in range(con_m):
        print('\rConvolution Run %d/%d'%(m, con_m)),
        sys.stdout.flush()
        for n in range(con_n):
            point = 0
            for k_m in range(ker_m):
                for k_n in range(ker_n):
                    point += M_kernel[k_m, k_n] * M_img[m+k_m, n+k_n]
            L_rest[m][n] = point
    
    print '\nConvolution Results Shape: ', np.matrix(L_rest).shape
    
    img = Image.new('L', [con_m, con_n])
    imgL = img.load()
    for i in range(con_m):
        line = L_rest[i]
        for j in range(con_n):
            imgL[i, j] = L_rest[i][j]
    img.show(title='L%s' % title_prex)

    return L_rest

if __name__ == "__main__":
    print("Convolution(Image) with m*n Kernel......")

    image_filename = './test.png'
    image_filename = './photo_airline.jpeg'
    L_kernel = []

    # Start
    img = Image.open(image_filename).convert('L')
    imgL = img.load()
    img_m, img_n = img.size
    L_img = [[imgL[m, n] for n in range(img_n)] for m in range(img_m)]
    
    L_kernel = [[1, 0, -1], [1, 0, -1], [1, 0, -1]]
    L_img = run_convolution(L_img, L_kernel)
#     L_kernel = [[1, 1, 1], [0, 0, 0], [-1, -1, -1]]
#     L_img = run_convolution(L_img, L_kernel)

#     L_kernel = [[1, 1, 1], [0, 0, 0], [-1, -1, -1]]
#     L_img = run_convolution(L_img, L_kernel)
#     L_kernel = [[1, 0, -1], [1, 0, -1], [1, 0, -1]]
#     L_img = run_convolution(L_img, L_kernel)
# 
#     L_kernel = [[1, 1, 1], [0, 0, 0], [-1, -1, -1]]
#     L_img = run_convolution(L_img, L_kernel)
#     L_kernel = [[1, 0, -1], [1, 0, -1], [1, 0, -1]]
#     L_img = run_convolution(L_img, L_kernel)

    # L_kernel = [[1, 1, 1, 1], [0, 0, 0, 0], [0, 0, 0, 0], [-1, -1, -1, -1]]
    # L_img = run_convolution(L_img, L_kernel)
