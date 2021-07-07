n = list(map(int,input().split()))
k = []
l = []
for i in n:
    if i>=0:
        k.append(i)
    else:
        l.append(i)
print(k+l)
print(l+k)
    