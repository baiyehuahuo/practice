// small arrays of integer
#include <iostream>

int main()
{
    using std::cout;
    using std::endl;

    int yams[3];
    yams[0] = 7;
    yams[1] = 8;
    yams[2] = 6;

    int yamcosts[3] = {20, 30, 5};
    cout << "Total yams = " << yams[0] + yams[1] + yams[2] << endl;
    cout << "The package with " << yams[1] << " yams costs " << yamcosts[1] << " cents per yam." << endl;

    int total = yamcosts[0] * yams[0] + yamcosts[1] * yams[1];
    total = total + yamcosts[2] * yams[2];
    cout << "The total yam expense is " << total << " cents." << endl;

    cout << "\nSize of yams array is " << sizeof(yams) << " bytes." << endl;
    cout << "Size of one element is " << sizeof(yams[0]) << " bytes." << endl;
    return 0;
}