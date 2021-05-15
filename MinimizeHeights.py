
def getMinDiff(self, arr, n, k):
    # code here
    inc = []
    dec = []
    for i in arr:
        if i+k >k:
            inc.append(i+k)
    for j in arr:
        if j-k >=0:
            dec.append(j-k)
    
    inc = sorted(inc)
    dec = sorted(dec)
    if inc:
        return inc[-1] - inc[0]
    if dec:
        return dec[-1] - dec[0]
