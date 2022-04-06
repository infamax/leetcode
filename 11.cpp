#include <algorithm>
#include <iostream>
#include <vector>

using namespace  std;

/* Native solution
class Solution {
public:
    int maxArea(vector<int>& height) {
        int res = -1;
        int last_height = height[height.size()- 1];
        for (int i = 0; i < height.size(); ++i) {
            for (int j = i + 1; j < height.size(); ++j) {
                int current_res = min(height[i], height[j]) * (j - i);
                res = max(current_res, res);
            }
        }
        return res;
    }
};
 */

class Solution {
public:
    int maxArea(vector<int>& height) {
        int i = 0;
        int j = height.size() - 1;
        int res = 0;
        while (i < j) {
            res = max(res, min(height[i], height[j]) * (j - i));
            if (height[i] < height[j]) {
                ++i;
            } else {
                --j;
            }
        }
        return res;
    }
};


int main() {
    return 0;
}
