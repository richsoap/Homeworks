#include <stdio.h>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
  vector<int> record(1000000, 0);
  for(int i = 0;i < 1000000 && scanf("%d", &record[i]) > 0;i ++);
  sort(record.begin(), record.end());
  for(auto val:record)
    printf("%d\n", val);
  return 0;
}