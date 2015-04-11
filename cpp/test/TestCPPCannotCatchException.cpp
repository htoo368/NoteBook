#include <iostream>

void foo()
{
    using namespace std;
    int bValue;
    if (bValue)
    {
        // do A
        cout << "if: " << bValue;
    }
    else
    {
        cout << "else: " << bValue;
        //      // do B
    }
}

int main(int argc, char* argv[])
{
    foo();
    return 0;
}
