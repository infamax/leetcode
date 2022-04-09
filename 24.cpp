#include <iostream>
#include <vector>

using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
 };
class Solution {
public:
    void swapTwoNodes(ListNode* node1, ListNode* node2) {
        if (node2 == nullptr) {
            return;
        }
        node1->next = node2->next;
        node2->next = node1;
    }
    ListNode* swapPairs(ListNode* head) {
        ListNode* node1 = head;
        int iter = 0;
        while (node1 != nullptr) {
            ListNode* node2 = node1->next;
            swapTwoNodes(node1, node2);
            if (iter == 0) {
                head = node1;
            }
            node1 = node1->next;
            if (node1 == nullptr) {
                break;
            }
            node1 = node1->next;
            ++iter;
        }
        return head;
    }
};