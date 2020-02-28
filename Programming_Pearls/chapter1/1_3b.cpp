#include <vector>
#include <stdio.h>
using namespace std;

#define UNITSIZE (sizeof(int)*8)
struct bitmap {
  vector<int> record;
  bitmap() {
    int vsize = 10000000 / UNITSIZE;
    record = vector<int>(vsize, 0);
  }
  int get(int index) {
    int loc = index / UNITSIZE;
    int off = index % UNITSIZE;
    return (record[loc] >> off) & 1;
  }
  void set(int index) {
    int loc = index / UNITSIZE;
    int off = index % UNITSIZE;
    record[loc] = record[loc] | (1 << off);
  }
};

int main() {
  struct bitmap record;
  int val;
  for(int i = 0;i < 1000000 && scanf("%d", &val) > 0;i ++)
    record.set(val);
  for(int i = 0;i < 10000000;i ++)
    if(record.get(i) == 1)
      printf("%d\n", i);
  return 0; 
}