#include<iostream>
const int ArSize = 16;
int main()
{
    using namespace std;
    long long facts[ArSize];
    facts[0] = facts[1] = 1LL;
    for(int i = 2; i < ArSize; i++)
    {
        facts[i] = facts[i-1] * i;
    }
    for(int i = 0; i < ArSize; i++ )
    {
        cout << i << "! = " << facts[i] << endl; 
    }
    return 0;
}