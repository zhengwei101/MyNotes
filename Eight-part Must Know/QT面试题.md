# QT面试题

## 一，信号与槽（Signal and Slot)

> 信号和槽是用于对象之间的通信，它是Qt的核心机制，在Qt编程中有着广泛的应用。

举例1，在一个十字路口，信号灯变成了绿色，对面的汽车看到后就启动了。信号灯就是发送信号的对象，绿灯亮是它发送的信号 (signal)，汽车是接收对象，汽车行驶是汽车对信号的响应，也叫槽 (slot)。

举例2，比如在一个主窗口内有一个关闭按钮，如果点击这个按钮窗口就会关闭，那么关闭按钮是发送信号的对象，它发送的信号(signal)是点击，接收信号的对象是窗口，响应信号的槽(slot)是关闭窗口。

在Qt中，发送对象、发送的信号、接收对象、槽可以通过很多种方式连接。

我们下面通过一些例子逐一做演示。

### Qt 4 使用宏

> 在Qt 4的版本中，主要通过connect + 宏的方式进行通信连接。

connect(发送对象，信号，接收对象，槽函数)， 其中发送信号和槽函数需要用SIGNAL() 和 SLOT()来进行声明。

connect 函数声明如下：

```cpp
[static] QMetaObject::Connection QObject::connect(const QObject *sender, const char *signal, const QObject *receiver, const char *method, Qt::ConnectionType type = Qt::AutoConnection)
```

比如点击按钮关闭窗口的例子，代码可以这样写：

```cpp
connect(ui->pushButton, SIGNAL(clicked()), this, SLOT(close()));
```

如果想自定义槽函数，需要先将槽函数的声明添加到类的slots中。

比如我们对一个`QLineEdit`控件添加一个接收`textEdited`信号的槽函数`onTextEdited`

```cpp
class MainWindow : public QMainWindow
{
    Q_OBJECT
public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
private slots:
    void onTextEdited(QString);
private:
    Ui::MainWindow *ui;
};
```

然后实现函数，并用connect与信号连接

```cpp
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    connect(ui->pushButton, SIGNAL(clicked()), this, SLOT(close()));    
    connect(ui->lineEdit, SIGNAL(textEdited(QString)), this, SLOT(onTextEdited(QString)));
}

void MainWindow::onTextEdited(QString s)
{
    qDebug() << s; 
}
```

这样写的好处是信号和槽参数很直观，但缺点是因为使用宏，编译时不做类型检查，如果有问题的话，在运行的时候才会发现。

### 使用Qt Creator 界面添加信号的槽函数

 另外一种方式不需要使用 ``connect` 函数，可以通过Qt Creator 界面来完成发送信号和槽函数的连接。

 1. 比如我们右键点击一个按钮，然后选择“转到槽”
 2. 选择信号，我们点击QAbstractButton的clicked()信号，表示按钮被点击
 3. 接下来，Qt Creator会自动为我们生成如下代码，首先是槽函数的声明：

 ```cpp
class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

private slots:
    void on_pushButton_clicked();
    
private:
    Ui::MainWindow *ui;
};
 ```

然后是槽函数的实现

 ```cpp
void MainWindow::on_pushButton_clicked()
{

}
 ```

使用这种方法我们不需要使用connect函数将信号与槽函数做连接。

这里槽函数的命名有一定的规则，一般是 `on_objectname_signal` 这样来命名的。

这种方法优点是减少了自己手动敲代码的工作量，缺点是究竟有哪些信号与槽函数做了连接不易被发现，没有``connect`函数看起来直观。

### 使用Qt 5 新 connect 函数

> Qt5 推出了新的`connect`函数，不需要使用`SIGNAL()`和`SLOT()`宏，可以在编译时做类型检查。

connect函数的声明如下：

```cpp
[static] QMetaObject::Connection QObject::connect(const QObject *sender, PointerToMemberFunction signal, const QObject *context, Functor functor, Qt::ConnectionType type = Qt::AutoConnection)
```

同样用代码实现点击按钮关闭窗口，并且添加一个`QLineEdit`控件，发送`textEdited`信号，由`onTextChanged()`函数作为槽函数响应。

使用这种方法，槽函数的声明不需要放到`slots`中，只要像普通的函数一样声明就可以了，类型需要与`textEdit`信号保持一致

```cpp
class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
    void textChanged(QString);

private:
    Ui::MainWindow *ui;
};
```

使用 connect 将信号与槽函数连接，不需要再使用`SIGNAL()` 和 `SLOT()` 宏

```cpp
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    connect(ui->pushButton, &QPushButton::clicked, this, &MainWindow::close);
    connect(ui->lineEdit, &QLineEdit::textEdited, this, &MainWindow::textChanged);
}

void MainWindow::textChanged(QString s)
{
    qDebug() << s;
}
```

### 使用函数指针

在 Qt5 版本的`connect` 函数里，信号与槽函数的参数其实都是函数指针，当信号或槽函数有重载时，使用函数指针可以明确告诉编译器使用哪一个重载函数，避免歧义。

下面的例子虽然没有使用重载，不过我们改成通过使用函数指针来向connect传递槽函数参数。

首先还是声明两个槽函数，分别响应点击按钮信号，和textEdited信号：

```cpp
class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow();
    void onButtonPushed();
    void onTextEdited(QString);    

private:
    Ui::MainWindow *ui;
};
```

然后对函数做简单的实现：

```cpp
void MainWindow::onButtonPushed()
{
    this->close();
}

void MainWindow::onTextEdited(QString s)
{
    qDebug() << s;
}
```

最后声明函数指针，并且将它们与信号连接

```cpp
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    void(MainWindow:: *btnClicked)() = &MainWindow::onButtonPushed;
    void(MainWindow:: *textEdited)(QString) = &MainWindow::onTextEdited;
    connect(ui->pushButton, &QPushButton::clicked, this, &btnClicked);
    connect(ui->lineEdit, &QLineEdit::textEdited, this, textEdited);
}
```

### 使用Lambda表达式

使用 Lambda表达式的好处是代码的书写更加方便快捷。在connect 函数中，槽函数参数我们可以改用Lambda表达式的方式来进行传参。

```cpp
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    connect(ui->pushButton, &QPushButton::clicked, this, [=](){
        this->close();
    });
    connect(ui->lineEdit, &QLineEdit::textEdited, this, [=](QString s){
        qDebug() << s;
    });
}
```

使用Lambda表达式，我们就不需要在类中对槽函数做任何的声明了。

Lambda表达式是`C++11`的内容，在比较低的 Qt版本中，要注意在Pro项目文件中加入 `CONFIG += C++11`。

`Qt`当中组件之间通过信号与槽的方式进行通信非常地高效，对于开发者来说也很简单。使用 `Qt5`版本的开发者建议使用上面后三种新的方式进行连接。

补充一点，信号和槽之间不是一一对应的关系。

一个信号可以对应多个槽，比如点击一个按钮可以触发多个不同的响应；

一个槽也可以响应多个不同的信号，比如点击按钮可以关闭窗口，点击左上角的小叉也可以关闭窗口。

信号和槽之间只要通过connect 函数连接就建立了耦合关系，如果想解除连接可以使用disconnect 函数。
