#include </usr/include/stdint.h>
#include <stdio.h>
#include <assert.h>

// Reverse Bits
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/648/


uint32_t reverseBits(uint32_t n) {
    uint32_t ans = 0;
    for(size_t i = 0; i < 32; i++)
    {
        ans = (ans << 1) | (n & 0x01);
        n >>= 1;        
    }
    
    return ans;
}

int main(int argc, char const *argv[])
{
    uint32_t input = 43261596;
    uint32_t want = 964176192;
    uint32_t got = reverseBits(input);
    assert(got == want);
    
    // while(input != 0){
    //     printf("%d", input & 0x01);
    //     input >>= 1;
    // }
    // printf("\n");

    // while(want != 0){
    //     printf("%d", want & 0x01);
    //     want >>= 1;
    // }
    // printf("\n");
    
    // while(got!=0){
    //     printf("%d", got & 0x01);
    //     got >>= 1;
    // }
    // printf("\n");

    return 0;
}
