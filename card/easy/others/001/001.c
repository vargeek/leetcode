#include </usr/include/stdint.h>
#include <stdio.h>
#include <assert.h>
// Number of 1 Bits
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/565/
// https://leetcode.com/submissions/detail/171385435/

int hammingWeight(uint32_t n) {
    int count = 0;
    while(n != 0){
        
        if (n & 0x01) {
            count++;
        }
        n = n >> 1;
        
    }
    return count;
}

int main(int argc, char const *argv[])
{
    uint32_t inputs[] = {11, 128};
    int outputs[] = {3, 1};
    
    for(size_t i = 0; i < 2; i++)
    {
        
        assert(hammingWeight(inputs[i]) == outputs[i]);

    }
    
    return 0;
}
