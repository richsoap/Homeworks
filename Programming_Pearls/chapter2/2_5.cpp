#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;

int main() {
  string buffer;
  int offset;
  cin>>buffer>>offset;
  for(int i = 0;i < buffer.length() - offset;i ++)
    swap(buffer[i], buffer[i+offset]);
  cout<<buffer<<endl;
  return 0;
}