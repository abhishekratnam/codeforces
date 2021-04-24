def isGoodarray(arr,ln):
    sum = 0
    for i in range(ln):
        sum += a[i]
    if sum % 2 != 0:
        return 0
    part = [0] * ((sum // 2)+ 1)
    for i in range((sum // 2)+ 1):
        part[i] = 0
    for i in range(ln) :
        for j in range(sum // 2, arr[i] - 1, -1) :
            if (part[j - arr[i]] == 1 or j == arr[i]) :
                part[j] = 1

    return part[sum // 2]
ele = int(input())
a = list(map(int,input().strip().split()))[:ele]
length = len(a)
if isGoodarray(a, length) == 1:
    print(0)
