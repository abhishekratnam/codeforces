def getDuplicates(nums):
    k = {}
    result = []
    for i in nums:
        if i in k:
            k[i] += 1      
        else:
            k[i] = 1
    
    for j in k.keys():
        if k[j]>1:
            result.append(j)
    return result
arr= [3, 3,  9,9, 12, 16, 20]
n = getDuplicates(arr)
print(n)
        