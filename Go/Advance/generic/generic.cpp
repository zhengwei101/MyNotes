#include <iostream>
using namespace std;

template <typename T>
class Stack
{
private:
    int _size;
    int _top;
    T *_array;

public:
    Stack(int n)
    {
        _size = n;
        _top = 0;
        _array = new T[_size];
    }

    void push(T t)
    {
        if (_top < _size)
        {
            _array[_top++] = t;
        }
    }

    T pop()
    {
        if (_top > 0)
        {
            return _array[--_top];
        }
        return T();
    }
};

int main()
{
    Stack<int> stack(3);
    stack.push(1);
    stack.push(2);
    cout << stack.pop() + stack.pop() << endl;
    Stack<string> sstack(3);
    sstack.push("hello");
    sstack.push("world");
    cout << sstack.pop() + sstack.pop() << endl;
    return 0;
}