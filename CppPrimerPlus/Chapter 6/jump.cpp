#include<iostream>
int main()
{
    using namespace std;
    const int ArSize = 80;
    char line[ArSize];
    int spaces  = 0;
    cout << "Enter a line of text: " << endl;
    cin.get(line, ArSize);
    cout << "Complete line: " << line << endl;
    cout << "Line through first period: " << endl;
    for (int i = 0; line[i] != '\0'; i++)
    {
        cout << line[i];
        if (line[i] == '.')
            break;
        if (line[i] != ' ')
            continue;
        spaces++;
    }
    cout << "\n" << spaces << " spaces" << endl;
    cout << "Done" << endl;
    return 0;
}