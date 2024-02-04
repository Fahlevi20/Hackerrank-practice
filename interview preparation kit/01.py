#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#

def plusMinus(arr):
    # Write your code here
    positive = []
    negative = []
    zero = []
    i = 0
    length= len(arr)
    
    for i in range(length):
        if arr[i] > 0:
            positive.append(1)
        if arr[i] < 0:
            negative.append(1)
        if arr[i] == 0:
            zero.append(1)
        
    print(len(positive)/length)
    print(len(negative)/length)        
    print(len(zero)/length)
 

if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
