#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

class Solution {
public:
    bool CorrectIndex(int i, int j, int n, int m) {
        if (i < 0 || j < 0 || i >= n || j >= m) {
            return false;
        }
        return true;
    }

    void floodFill(vector<vector<int>>& field, int i, int j, int oldColor, int newColor,
                   vector<vector<bool>>& visited) {

        visited[i][j] = true;
        if (field[i][j] != oldColor) {
            return;
        }
        field[i][j] = newColor;


        if (CorrectIndex(i, j - 1, field.size(), field[0].size())
            && !visited[i][j - 1]) {
            floodFill(field, i, j - 1, oldColor, newColor, visited);
        }

        if (CorrectIndex(i, j + 1, field.size(), field[0].size())
            && !visited[i][j + 1]) {
            floodFill(field, i, j + 1, oldColor, newColor, visited);
        }

        if (CorrectIndex(i - 1, j, field.size(), field[0].size()) &&
            !visited[i - 1][j]) {
            floodFill(field, i - 1, j, oldColor, newColor, visited);
        }

        if (CorrectIndex(i + 1, j, field.size(), field[0].size()) &&
                !visited[i + 1][j]) {
            floodFill(field, i + 1, j, oldColor, newColor, visited);
        }
    }

    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int newColor) {
        vector<vector<int>> res = image;
        int oldColor = image[sr][sc];
        vector<vector<bool>> visited(image.size(), vector<bool>(image[0].size(), false));
        floodFill(res, sr, sc, oldColor, newColor, visited);
        return res;
    }
};

int main() {
    return 0;
}