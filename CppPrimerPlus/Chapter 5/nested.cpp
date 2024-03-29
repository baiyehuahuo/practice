#include<iostream>
const int Cities = 5;
const int Years = 4;
int main()
{
    using namespace std;
    const char * cities[Cities] =
    {
        "Gribble City",
        "Gribbletown",
        "New Gribble",
        "San Gribble",
        "Gribble Vista",
    };
    int max_temps[Years][Cities] = 
    {
        {96,100,87,101,105},
        {96,98,91,107,104},
        {97,101,93,108,107},
        {98,103,95,109,108},
    };
    cout <<"Maximum temperatures for 2008 - 2011: " << endl;
    for (int city = 0; city < Cities; city++) 
    {
        cout << cities[city] << ":";
        for (int year = 0; year < Years; year++)
        {
            cout << "\t" << max_temps[year][city];
        }
        cout << endl;
    }
    return 0;
}