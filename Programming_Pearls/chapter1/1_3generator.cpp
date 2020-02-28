#include <iostream>
#include <vector>
#include <cstdlib>
#define RANGE 10000000
using namespace std;
int main() {
  vector<int> record(RANGE, 0);
  int val;
  int count;
  for(int i = 0;i < 1000000;i ++) {
    int count = 0;
    while(count++ < 100) {
      val = rand() % RANGE;
      if(record[val] == 0) {
        cout<<val<<endl;
        record[val] = 1;
        break;
      }
    }
    if(count >= 100) {
      cerr<<"Error"<<endl;
      return 1;
    }
  }
  return 0;
}
