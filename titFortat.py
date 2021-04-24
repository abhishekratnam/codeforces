def titFortat(lst,n,k):
    Yi = 0
    Xi = 0
    Yj = 0
    Xj = 0
    sort = sorted(lst)
    y,x = sort[-1],sort[-2]
    for index in range(len(lst)):
        if lst[index]==y:
            Yi = index 
            Yj = index-1
        elif lst[index]==x:
            Xi = index 
            Xj = index-1
    return sort

cases = int(input())
lst = []
j = 0
times = []
for i in range(cases):
    n, k = map(int, input().split())
    lst = list(map(int,input().strip().split()))[:n]
    times.append(titFortat(lst,n,k))
print(times)
