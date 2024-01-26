import numpy as np
value = 1, 2, 3, 4, 5, 6, 7, 8, 9
arr = np.array(list(map(int,input().split())))
# change_array = np.array([1, 2, 3, 4, 5, 6, 7, 8, 9
# ])
reshape = np.reshape(arr,(3,3))
print(reshape)