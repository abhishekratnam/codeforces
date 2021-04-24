a = []
result = []
for i in range(3):
    a.append(int(input()))
if a[0]<a[1] and a[0]<a[2]:
    result.append(a[0])
elif a[1]>a[0] and a[1]<a[2]:
    result.append(a[1])
elif a[2]>a[1] and a[2]>a[0]:
    result.append(a[2])
print(result)
print(a)
