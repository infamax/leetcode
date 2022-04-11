#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    const int START = -1;
    const int FINISH = 1;
    vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
        vector<pair<int, int>> events;
        events.reserve((intervals.size() + newInterval.size()) * 2);

        for (const auto& interval : intervals) {
            events.emplace_back(interval[0], START);
            events.emplace_back(interval[1], FINISH);
        }

        events.emplace_back(newInterval[0], START);
        events.emplace_back(newInterval[1], FINISH);

        sort(begin(events), end(events));

        vector<vector<int>> res;
        vector<int> tmp;
        int k = 0;
        for (int i = 0; i < events.size(); ++i) {
            if (k == 0 && events[i].second == START) {
                tmp.push_back(events[i].first);
            }

            if (events[i].second == START) {
                ++k;
            }

            if (events[i].second == FINISH) {
                --k;
            }

            if (k == 0 && events[i].second == FINISH) {
                tmp.push_back(events[i].first);
                res.push_back(tmp);
                tmp.clear();
            }
        }
        return res;
    }
};

int main() {
    return 0;
}