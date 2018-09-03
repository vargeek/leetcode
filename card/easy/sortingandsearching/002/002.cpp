#include <stdio.h>

// First Bad Version
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/774/
// https://leetcode.com/problems/first-bad-version/description/

// Forward declaration of isBadVersion API.

bool isBadVersion(int version) {
    // return version >= 4;
    return version >= 1702766719;
}

int firstBadVersion(int n) {
    int left = 1;
    int right = n;
    
    int mid = 0;
    int first = 0;
    while(left <= right){
        mid = ((long)left + right) / 2;
        if (isBadVersion(mid)) {
            first = mid;
            right = mid - 1;
        } else {
            left = mid + 1;
        }   
    }
    return first;
}

int main(int argc, char const *argv[])
{

    printf("%d\n", firstBadVersion(2126753390));
    return 0;
}
