#include<bits/stdc++.h> 
using namespace std 
int main(){
    int n; cin>>n; 
    int arr[n] ; 
    for(int i=0;i<n;i++){
        cin>>arr[i] ; 
    } 
    sort(arr,arr+n); 
    int ans = 0 ;
    for(int i=0;i<n;i++){
        if(arr[i] != i+1){
            ans = i+1   ; break;
        }
    }  
    cout<<ans<<endl;
    return 0;
}