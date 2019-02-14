#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int reverse(int x) {
    int result = 0;
    while (x != 0) {
        int temp = x % 10;
        if (result > INT_MAX / 10 || result < INT_MIN / 10) {
            return 0;
        }
        
        result = 10*result+temp;
        x /= 10;//如果写成x = x / 10 速度慢1倍，内存减少不明显，gcc编译器不会主动优化，因为/=这个是c里的运算符
    }
    return result;
}