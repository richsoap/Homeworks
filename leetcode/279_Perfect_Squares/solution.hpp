#include <vector>
#include <queue>
using namespace std;
class Solution {
  public:
    int numSquares(int n) {
      if(n <= 0)
        return 0;

      vector<int> perfectSquares;
      vector<int> cntPerfectSquares(n);
      for(int i = 1;i * i <= n;i ++) {
        perfectSquares.push_back(i*i);
        cntPerfectSquares[i*i-1] = 1;
      }
      if(perfectSquares.back() == n)
        return 1;
      queue<int> searchQ;
      for(auto& i:perfectSquares)
        searchQ.push(i);
      int currCntPerfectSquares = 1;
      while(!searchQ.empty()) {
        currCntPerfectSquares ++;
        int searchQSize = searchQ.size();
        for(int i = 0;i < searchQSize;i ++) {
          int tmp = searchQ.front();
          for(auto& j:perfectSquares) {
            if(tmp + j == n)
              return currCntPerfectSquares;
            else if((tmp + j < n) && (cntPerfectSquares[tmp + j - 1] == 0)) {
              cntPerfectSquares[tmp + j - 1] = currCntPerfectSquares;
              searchQ.push(tmp + j);
            } else if(tmp + j > n)
              break;
          }
          searchQ.pop();
        }
      }
      return 0;
    }
};
