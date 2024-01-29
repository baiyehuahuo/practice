#include<iostream>
#include<fstream>
#include<cstdlib>
int main()
{
    using namespace std;
    const int SIZE = 60;
    char data_file[SIZE];

    cout << "Enter name of date file: ";
    cin >> data_file;

    ifstream inFile;
    inFile.open(data_file);
    if (!inFile.is_open())
    {
        exit(EXIT_FAILURE);
    }

    double sum = 0, tmp;
    int count = 0;
    while(inFile.good())
    {
        inFile >> tmp;
        sum += tmp;
        count++;   
    }
    if (inFile.eof())
        cout << "End of file reached. " << endl;
    else if (inFile.fail())
        cout << "Input terminated by data mismatch." << endl;
    else 
        cout << "Input terminated for unkonwn reason." << endl;
    if (count == 0) 
        cout << "No data processed." << endl;
    else
    {
        cout << "Items read: " << count << endl;
        cout << "Sum: " << sum << endl;
        cout << "Average: " << sum / count << endl;
    }
    inFile.close();
    return 0;
}
