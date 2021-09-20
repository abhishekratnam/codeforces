def isSubset( a1, a2, n, m):
    k ={}
    count = 0
    for i in a2:
        if i not in k.keys():
            k[i]=True
    for i in a1:
        if i in k:
            count += 1
        
            
            
    if count == len(a2):
        return "Yes"
    return "No"
    
    