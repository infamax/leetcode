#include <string>
#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    void addChars(string& str, int k) {
        string tmp;
        for (int i = 0; i < k; ++i) {
            tmp.push_back('0');
        }
        for (char i : str) {
            tmp.push_back(i);
        }
        str = tmp;
    }

    string addBinary(string a, string b) {
        if (a.size() < b.size()) {
            addChars(a, b.size() - a.size());
        } else if (a.size() > b.size()) {
            addChars(b, a.size() - b.size());
        }
        string res;
        int tens = 0;
        for (int i = b.size() - 1; i >= 0; --i) {
            int sum = (a[i] - '0') + (b[i] - '0');
            res.push_back('0' + ((tens + sum) % 2));
            tens = (sum + tens) / 2;
        }
        if (tens != 0) {
            res.push_back('1');
        }
        reverse(begin(res), end(res));
        return res;
    }
};

int main() {
    Solution sol;
    cout << sol.addBinary("11", "1");
    return 0;
}