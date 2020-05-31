#include <iostream>

using namespace std;

class father {
    public:
    virtual void Hello() {
        cout<<"Hello world"<<endl;
    }
};

class son: public father {
    public:
    virtual void Hello() {
        cout<<"Fuck you"<<endl;
    }
};

class daughter: public father {
    public:
    virtual void Hello() {
        cout<<"What the Hell?"<<endl;
    }
};

int main() {
    father fa;
    fa.Hello();
    son so;
    so.Hello();
    daughter dau;
    dau.Hello();
    father& fakeFa = so;
    fakeFa.Hello();
    fakeFa = dau;
    fakeFa.Hello();
    return 0;
}