import math
def intersection(arr_1,arr_2):
    intersection = []
    for i in arr_1:
        for j in arr_2:
            if i==j:
                intersection.append(i)

    print(intersection)
def union(arr_1,arr_2):
    l = len(arr_1)
    m = len(arr_2)
    i,j = 0,0
    result = []
    while i < l and j < m:
        if arr_1[i] < arr_2[j]:
            result.append(arr_1[i])
            i += 1
        elif  arr_2[j] < arr_1[i]:
            result.append(arr_2[j])
            j += 1
        else:
            result.append(arr_2[j])
            j += 1
            i += 1
    while i < l:
        result.append(arr_1[i])
        i += 1
    while j < m:
        result.append(arr_2[j])
        j += 1
    print(result,"\n")

arr_1 = [1, 2, 4, 5, 6 ]
arr_2 = [ 2, 3, 5, 7 ]
intersection(arr_1,arr_2)
union(arr_1,arr_2)
