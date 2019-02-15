#include <stdio.h>

bool isPalindrome(int x) {
    if (x == 0)
    {
        return true;
    }
    if (x < 0)
    {
        return false;
    }
    int y = 0;
    int tempX = x;
    while (tempX > 0) {
        y = y*10 + tempX%10;
        tempX /= 10;
    }
    return x == y;
}

