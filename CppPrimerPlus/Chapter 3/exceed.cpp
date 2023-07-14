// exceeding some integer limits
#include <iostream>
#define ZERO 0
#include <climits>

int main()
{
    using std::cout;
    using std::endl;
    short sam = SHRT_MAX;
    unsigned short sue = sam;

    cout << "Sam has " << sam << " dollars and Sue has " << sue << " dollars deposited." << endl;
    cout << "Add $1 to each account." << endl;
    sam = sam + 1;
    sue = sue + 1;
    cout << "Now Sam has " << sam << " dollars and Sue has " << sue << " dollars deposited." << endl;
    sam = ZERO;
    sue = ZERO;
    cout << "Sam has " << sam << " dollars and Sue has " << sue << " dollars deposited." << endl;
    cout << "Task $1 from each account." << endl;
    sam = sam - 1;
    sue = sue - 1;
    cout << "Sam has " << sam << " dollars and Sue has " << sue << " dollars deposited." << endl;
    return 0;
}