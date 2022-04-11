#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        int max_length = 0;
        int length = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] == 1) {
                ++length;
            } else {
                max_length = max(max_length, length);
                length = 0;
            }
        }
        max_length = max(max_length, length);
        return max_length;
    }
};

int main() {

};
