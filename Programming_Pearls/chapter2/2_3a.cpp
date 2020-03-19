#include <vector>
#include <iostream>
using namespace std;

int main() {
  int n,k;
  cin>>n>>k;
  vector<int> result(n, 0);
  int index = n - k;
  for(int i = 0;i < n;i ++)
    cin>>result[(index++) % n];
  for(auto val:result)
    cout<<val<<" ";
  cout<<endl;
  return 0;
}