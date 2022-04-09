#include <algorithm>
#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>

using namespace std;

class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> frequents;
        for (int num : nums) {
            ++frequents[num];
        }

        vector<int> items_by_frequents;
        items_by_frequents.reserve(frequents.size());

        for (const auto& p : frequents) {
            items_by_frequents.push_back(p.second);
        }

        sort(rbegin(items_by_frequents), rend(items_by_frequents));
        items_by_frequents.resize(k);
        unordered_set<int> s(begin(items_by_frequents), end(items_by_frequents));
        vector<int> res;
        for (const auto& p : frequents) {
            if (s.find(p.first) != end(s)) {
                res.push_back(p.first);
            }
        }
        return res;
    }
};

int main() {
    return 0;
}