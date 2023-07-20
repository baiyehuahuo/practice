// using the new operator
#include <iostream>
int main()
{
    using std::cout;
    using std::endl;

    int nights = 1001;
    int *pt = new int;
    *pt = 1001;

    cout << "nights value = " << nights << ": locations = " << &nights << endl;

    cout << "int value = " << *pt << ": location = " << pt << endl;

    double *pd = new double;
    *pd = 10000001.0;

    cout << "double value = " << *pd << ": locations = " << pd << endl;
    cout << "location of pointer pd : " << &pd << endl;
    cout << "size of pt = " << sizeof(pt) << ": size of *pt = " << sizeof(*pt) << endl;
    cout << "size of pd = " << sizeof(pd) << ": size of *pd = " << sizeof(*pd) << endl;
    return 0;
}