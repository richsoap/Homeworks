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

int goDeep(map<int, int>& record, map<int, int>::iterator maxit, int K) {
  if(K == 0 || maxit->first == 1)
    return maxit->first;
  //cerr<<"goDeep in "<<maxit->first<<": "<<maxit->second<<endl;
  int result = maxit->first;
  int i;
  for(i = 1;i * maxit->second <= K && i < maxit->first;i ++) {
    auto splitres = split(maxit->first, i + 1);
    for(auto& line : splitres) {
      //cerr<<"split "<<line[0]<<" "<<line[1]<<" record: "<<record[line[0]]<<endl;
      record[line[0]] += (line[1] * maxit->second);
    }
    int temp = goDeep(record, getMaxIt(maxit), K - i * maxit->second);
    for(auto& line : splitres)
      record[line[0]] -= (line[1] * maxit->second);
    if(temp > result)
      break;
    result = temp;
  }
  //cerr<<"res: "<<result<<endl;
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
  for(auto& it:record){
    //cerr<<"rec "<<it.first<<" "<<it.second<<endl;
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
