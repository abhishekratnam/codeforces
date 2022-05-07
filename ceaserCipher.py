#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'caesarCipher' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING s
#  2. INTEGER k
#

def caesarCipher(s, k):
    # Write your code here
    original_alphabet = ["a","b","c","d","e","f","g","h","i",
    "j","k","l","m","n","o","p","q","r","s","t","u","v","w",
    "x","y","z"]
    if k>26:
        k = k % 26 
    rotated_alphabet = original_alphabet[k:]+original_alphabet[:k]
    final = {}
    for a,b in zip(original_alphabet, rotated_alphabet):
        final[a] = b
        final[a.upper()] = b.upper()
    result = []
    for char in s:
        if not char.isalnum():
            result.append(char)
            continue
        elif char.isnumeric():
            result.append(char)
            continue
        else:
            if char in final:
                result.append(final[char])
            
    return ''.join(result)            
            
    

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    s = input()

    k = int(input().strip())

    result = caesarCipher(s, k)

    fptr.write(result + '\n')

    fptr.close()
