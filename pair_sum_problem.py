def getPairsCount(arr, n, k):
    # code here
    lookup = {}
    count = 0
    for i in range(n):       
        if k-arr[i] in lookup.keys():
            count += 1
        if arr[i] not in lookup.keys():
            lookup[arr[i]] = True
    return count
k = [1, 1, 1, 1]
print(getPairsCount(k,len(k),2))