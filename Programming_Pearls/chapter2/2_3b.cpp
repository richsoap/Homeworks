#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;
void swaptimes(vector<int>& record, int start, int delta, int times) {
  int n = record.size();
  int temp = record[start];
  while(times --) {
    record[start] = record[(start + delta) % n];
    start = (start + delta) % n;
  }
  record[start] = temp;
}

int main() {
  int n,k,newn;
  cin>>n>>k;
  if(n % k == 0)
    newn = n;
  else
    newn = (n/k+1)*k;
  vector<int> record(newn, 0);
  for(int i = 0;i < n;i ++)
    cin>>record[i];
  for(int i = 0;i < k;i ++)
    swaptimes(record, i, k, newn / k - 1);
  for(int i = 0;i < k;i ++)
    swaptimes(record, n - k + i, newn-n, 1);
  for(int i = 0;i < n;i ++)
    cout<<record[i]<<" ";
  cout<<endl;
  return 0;
}