#include <bits/stdc++.h>
using namespace std;
int maxSubarraySum(int arr[], int n){
        int curSum =0;
        int maxSum =0;
        // Your code here
        for(int i=0;i<n;i++){
            curSum = curSum + arr[i];
            if(curSum > maxSum){
                maxSum = curSum;
            }
            if(curSum<0){
                curSum = 0;
            }
            
        }
        return maxSum;
}
int main(){
    int t,n;
    cin>>t;
    while(t--){
        cin>>n;
        int a[n];
        for(int i=0;i<n;i++){
            cin>>a[i];
        }
        cout<<maxSubarraySum(a,n)<<endl;
        
    }
}