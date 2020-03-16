#include <iostream>
#include <sys/time.h>
#include <unistd.h>
#include "solution.hpp"
using namespace std;
int main(int argc, char* args[]) {
  if(argc != 3) {
    cout<<"wrong argv"<<endl;
    return -1;
  }
  struct timeval start,end;
  int maxval = atoi(args[1]);
  int delta = atoi(args[2]);
  int val = 0;
  long timestamp;
  Solution solution;
  while(val < maxval) {
    gettimeofday(&start, 0);
    solution.numSquares(val);
    gettimeofday(&end, 0);
    val += rand() % delta;
    timestamp = end.tv_usec - start.tv_usec + 1000000 * (end.tv_usec - start.tv_usec);
    if(timestamp < 0)
      continue;
    printf("%d %ld\n",val, timestamp);
  }
  return 0;
}
