#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> res(n, vector<int>(n,0));
        int i = 0, j = 0;
        int board_left = -1;
        int board_right = n;
        int board_up = 0;
        int board_down = n;
        bool fl_right = true;
        bool fl_left = false;
        bool fl_up = false;
        bool fl_down = false;
        for (int k = 1; k <= n * n; ++k) {

            if (fl_right) {
                res[i][j] = k;
                ++j;
                if (j == board_right) {
                    fl_right = false;
                    fl_down = true;
                    --j;
                    ++i;
                    --board_right;
                }
            } else if (fl_down) {
                res[i][j] = k;
                ++i;
                if (i == board_down) {
                    fl_down = false;
                    fl_left = true;
                    --i;
                    --j;
                    --board_down;
                }
            } else if (fl_left) {
                res[i][j] = k;
                --j;
                if (j == board_left) {
                    fl_left = false;
                    fl_up = true;
                    ++j;
                    --i;
                    ++board_left;
                }
            } else {
                res[i][j] = k;
                --i;
                if (i == board_up) {
                    fl_up = false;
                    fl_right = true;
                    ++i;
                    ++j;
                    ++board_up;
                }
            }

        }

        return res;
    }
};

int main() {
    return 0;
}