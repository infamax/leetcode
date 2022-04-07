#include <algorithm>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    int lastStoneWeight(vector<int>& stones) {
        if (stones.empty()) {
            return 0;
        }

        if (stones.size() == 1) {
            return stones.front();
        }

       sort(begin(stones), end(stones));

       while (stones.size() > 1) {
           int last1 = stones.back();
           stones.pop_back();
           int last2 = stones.back();
           stones.pop_back();
           if (last1 != last2) {
               stones.push_back(last1 - last2);
           }
           sort(begin(stones), end(stones));
       }

        if (stones.empty()) {
            return 0;
        }
        return stones.front();
    }
};

int main() {
    return 0;
}