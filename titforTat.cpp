#include <bits/stdc++.h>

using namespace std;
#define ll long long
int main()
{   
   ll t;
   cin>>t;
   while(t--)
   {
     ll n,i,k;
     cin>>n>>k;ll a[n];
     for(i=0;i<n;i++)
     {cin>>a[i];}
    
     if((a[0]-k)>=0)
     {cout<<a[0]-k<<" ";
     if(n==2)
     {cout<<a[1]+k;}
     else{
     a[2]+=k;
     for(i=1;i<n;i++)
     cout<<a[i]<<" ";
     }}
     else
     {
       cout<<0<<" ";
       if(n==2)
       {cout<<a[0]+a[1];}
       else{
       a[2]+=a[0];
     for(i=1;i<n;i++)
     cout<<a[i]<<" ";
     }}
     cout<<endl;
     
     }
}