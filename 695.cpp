#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    bool correctIndexes(int i, int j, int n, int m) {
        if (i < 0 || j < 0 || i >= n || j >= m) {
            return false;
        }
        return true;
    }

    int calculateAreaOfIsland(vector<vector<int>>& grid, int i, int j) {
        if (!correctIndexes(i, j, grid.size(), grid[0].size()) || grid[i][j] == 0) {
            return 0;
        }
        grid[i][j] = 0;
        return 1 + calculateAreaOfIsland(grid, i, j + 1) + calculateAreaOfIsland(grid, i, j - 1)
        + calculateAreaOfIsland(grid, i - 1, j) + calculateAreaOfIsland(grid, i + 1, j);
    }

    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int max_area = 0;

        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[i].size(); ++j) {
                if (grid[i][j] == 1) {
                    int area = calculateAreaOfIsland(grid, i, j);
                    max_area = max(max_area, area);
                }
            }
        }

        return max_area;
    }
};

int main() {
    return 0;
}