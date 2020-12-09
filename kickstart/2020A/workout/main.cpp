#include <map>
#include <iostream>
#include <vector>
using namespace std;

void print(int casenum, int val) {
  cout<<"Case #"<<casenum<<": "<<val<<endl;
}

map<int, int>::iterator getMaxIt(map<int, int>::iterator it) {
  do {
    it --;
  }while(it->second == 0);
  return it;
}

vector<vector<int>> split(int v, int n) {
  vector<vector<int>> result;
  int delta = v/n;
  int res = 0;
  result.push_back(vector<int>({delta, n-1}));
  if(delta * n == v)
    result[0][1] ++;
  else 
    result.push_back(vector<int>({v - delta * (n - 1), 1}));
  return result;
}

vector<int> getPrinumbers(int n) {
  static vector<int> result;
  if(result.empty())
    result.push_back(2);
  else if(n <= result.back())
    return result;
  for(int i = result.back() + 1;i <= n;i ++) {
    bool isPri = true;
    for(int j = 0;j < result.size() && isPri;j ++)
      isPri = i % result[j] != 0;
    if(isPri)
      result.push_back(i);
  }
  return result;
}

int goDeep(map<int, int>& record, map<int, int>::iterator maxit, int K) {
  if(K == 0 || maxit->first == 1)
    return maxit->first;
  int result = maxit->first;
  vector<int> prinums = getPrinumbers(K);
  for(int i = 0;i < prinums.size() && (prinums[i] - 1) * maxit->second <= K && prinums[i] <= maxit->first;i ++) {
    int dur = prinums[i];
    auto splitres = split(maxit->first, dur);
    for(auto& line : splitres) {
      record[line[0]] += (line[1] * maxit->second);
    }
    int temp = goDeep(record, getMaxIt(maxit), K- (dur - 1) * maxit->second);
    for(auto& line : splitres) {
      record[line[0]] -= (line[1] * maxit->second);
    }
    if(temp > result)
      break;
    result = temp;
  }
  return result;
}

int solve() {
  int N, K;
  cin>>N>>K;
  map<int, int> record;
  int prenum, nownum;
  for(int i = 0;i < N;i ++) {
    cin>>nownum;
    if(i != 0)
      record[nownum - prenum] ++;
    prenum = nownum;
  }
  auto maxit = record.end();
  return goDeep(record, --maxit, K);
}

int main() {
  int T;
  cin>>T;
  for(int i = 1;i <= T;i ++)
    print(i, solve());
  return 0;
}
