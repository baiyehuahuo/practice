// using the delete operator
#include <iostream>
#include <cstring>
using std::cin;
using std::cout;
using std::endl;

char *get_name(void);

int main()
{
    char *name;
    name = get_name();
    cout << name << " at " << (int *)name << endl;
    delete[] name;

    name = get_name();
    cout << name << " at " << (int *)name << endl;
    delete[] name;
    return 0;
}

char *get_name()
{
    char tmp[80];
    cout << "Enter last name: ";
    cin >> tmp;
    char *pn = new char[strlen(tmp) + 1];
    strcpy(pn, tmp);
    return pn;
}