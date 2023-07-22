// some type combinations
#include <iostream>
struct antarctiaca_years_end
{
    /* data */
    int year;
};

int main()
{
    using std::cout;
    using std::endl;

    antarctiaca_years_end s01, s02, s03;
    s01.year = 1998;
    antarctiaca_years_end *pa = &s02;
    pa->year = 1999;
    antarctiaca_years_end trio[3];
    trio[0].year = 2003;
    cout << trio->year << endl;
    const antarctiaca_years_end *arp[3] = {&s01, &s02, &s03};
    cout << arp[1]->year << endl;
    const antarctiaca_years_end **ppa = arp;
    auto ppb = arp;

    cout << (*ppa)->year << endl;
    cout << (*(ppb + 1))->year << endl;
}
