#include <algorithm>
#include <iostream>
#include <vector>
#include <numeric>

using namespace std;

class Solution {
public:
    int calPoints(vector<string>& ops) {
        vector<int> nums;
        for (const string& op : ops) {
            if (op == "C") {
                nums.pop_back();
            } else if (op == "D") {
                nums.push_back(nums.back() * 2);
            } else if (op == "+") {
                int a1 = nums[nums.size() - 1];
                int a2 = nums[nums.size() - 2];
                nums.push_back(a1 + a2);
            } else {
                nums.push_back(stoi(op));
            }
        }
        return accumulate(begin(nums), end(nums), 0);
    }
};

int main() {
    return 0;
}