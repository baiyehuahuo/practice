// floating-point types
#include <iostream>

int main()
{
    using std::cout;
    using std::endl;

    cout.setf(std::ios_base::fixed, std::ios_base::floatfield);
    float tub = 10.0 / 3.0;
    double mint = 10.0 / 3.0;
    const float million = 1.0e6;

    cout << "tub = " << tub << ", a million tubs = " << million * tub << ", and ten million tubs = " << 10 * tub * million << endl;

    cout << "mint = " << mint << " and a million mint = " << million * mint << endl;

    return 0;
}