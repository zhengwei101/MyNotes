# 从C++到Java:基本语法

## Java相较于C++的区别

- Java不用为内存分配考虑了，自带垃圾回收的Java让你再也不用delete内存。
- Java取消了多继承，但是也可以通过接口实现多继承。
- Java没有重载运算符。
- Jave有**反射机制**，允许程序在运行时自我检查，同时也允许对其内部的成员进行操作，C++没有提供这样的特性。

## 基础语法

## Hello World

```java
//HelloWorld.java
package com.PQBox;
//import java.io.*;

public class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello World");
  }
}
```

这里来看一下Java和C++语法上的区别：

- 第一行`package`语句：

  比起C++，Java可以更加灵活地控制自己的工程代码层级。
  
  `package`是包命令，这一行表示这个`HelloWorld.java`文件在`com/PQBox`文件夹下。
  
  也就是说，如果你的工程根目录在`~/Documents`下，那么你的这个文件所在路径就是`~/Documents/com/PQBox/HelloWorld.java`。

- 第二行被注释掉了，但同时展现了两个点：

  - 注释：和C++一模一样。以`//`表示行注释，以`/** **/`表示块注释
  - `import`语句：由于你可以用`package`来打包你的java文件，那么相应的你也可以通过`import`来引入文件，相当于C++的`#include`。这里引入的是Java标准IO，载入了`Java_installation/java/io/`文件夹下所有Java文件。

- 第五行的类声明：

  和C++一样类由`class`声明，不同的是类前面可以加上访问修饰符，这里`public`表示这个java文件是对所有工程，包括其他工程可见的（包外可见）。
  
  一个源文件只能有一个`public`类。类最后不需要加**分号**（结构体，枚举后面都不需要）。

  **需要注意的是：Java的`main`函数所在的文件名称必须和类名一致。**

- 第六行`main`函数的声明：

  函数的声明也和C++一样（Java里习惯叫函数为方法），只不过每个方法前面都需要加上访问修饰符表示方法的访问权限（不加表示包内访问）。
  
  这里的`public`表示其他包可以访问。

  这里的`String[] args`其实和`int argc, char** argv`一样，只不过Java里面用`String`表示字符串，数组有`length`变量可以得到数组大小，所以没必要加上`int argc`，最后就简化成这样了。

- 第七行`System.out.println()`是Java的控制台输出函数，`println`是`print line`的简写，也有我们熟悉的`printf`函数：`System.out.printf()`。

### **为什么main函数要写在类中？**

因为Java遵循**万事万物皆对象**的说法，**不允许将方法或者变量（Java习惯称字段）的声明暴露在类外**。

当我们运行Java的时候，其实是Java虚拟机JVM调用`HelloWorld.main()`函数。

在Java中，一切都是在类中运作的。

## 变量的声明

### 普通变量

普通变量和C++一样，但是要注意一下基本数据类型的名称变化，8种基本数据类型如下：

`byte`,`short`,`int`,`long`,`float`,`double`,`char`,`boolean`

其中`byte`是表示8位字节的数据类型（也是整数类型中的一份子），取值`-128~127`。

### Null还是nullptr?

答案是都不是，Java为了表示空，定义了自己的类型`null`，千万别搞混了。

### 静态变量

使用`static`

### 常量

和C++不一样，使用`final`：

```java
final int a = 10;    //直接赋值
final int a;
a = 20;    //延迟赋值
```

### 数组

数组就比较特殊了，需要使用`new`关键字创建：

数组是一个容器，当它被创建后，不仅元素的类型是确定的，元素的个数也是确定的。

换句话说，数组的长度是确定的，不可能再变长或者变短。

因此，数组可以使用一个字段（length）来表示长度。

```java
int a[] = new int[4];
int[] a = new int[4];
//方括号放在前面后面都可以，但是里面不能有东西

int[] a = {1, 2, 3, 4} //初值列初始化，不需要new

int[] a;
a = new int[4]; //延迟初始化，在需要的时候使用new，在初始化之前a=null

System.out.println(a.length);// 获取数组的长度

String str = "hello";
System.out.println(str.length()); // 获取字符串的长度

```

这代码简直就像C++中利用指针创建数组：

```c++
int* a = new int[4];
```

二维数组也差不多：

```java
int[][] a = new int[2][3];

//也可以像指针一样，先new一维，再new剩下的
int[][] a = new int[2][];    
for(int i=0;i<a.length;i++)
  a = new int[3];

/*这里可以想象这段代码是这样的:
int** a = new int[2];
for(int i=0;i<2;i++)
    a[i]=new int[3];
*/
```

#### delete呢？

Java中没有delete这种东西，因为根本不需要。

在你使用`new`在内存中分配内存之后，Java虚拟机JVM会使用**垃圾回收机制**，在对象不再需要的时候自动回收内存。

妈妈再也不用担心我没有delete了。

#### Arrays类

Java的`java.util.Arrays`类是一个静态类（所有的类成员都是静态的），专门用于对数组操作，比如`Arrays.sort()`用于排序，`Arrays.fill()`用于向数组里填充元素等。

### 类的实例化

类的实例化也是使用new：

```java
String str = new String("haha");    //在new之后调用构造函数
```

和C++不一样的是，如果你调用的是默认构造函数，也得加上括号：

```java
String str = new String;   //Error!
String str = new String(); //Right!
```

同样JVM会帮你回收内存，不需要delete。

## 函数（方法）声明

和C++一模一样，只是要注意方法前面的访问修饰符。

### 传值还是传引用？

由于Java将指针概念隐藏了起来，导致我们没有办法使用指针，

那么就带来一个问题：函数参数什么时候传值，什么时候传引用呢？

其实和Python差不多：

- **八种基本数据类型都是传值的**
- **数组，类，接口等其他数据类型是传引用的**

那么这个时候就有人要问了：那我要将基本数据类型传引用岂不是非法？

那也不是。

Java给每个基本数据类型一个对应的类，称为**闭包**，只要你使用其相应的类对象传值就可以了。

### 闭包

每一个基本数据类型对应的类称为闭包，比如`int`对应`Integer`，`char`对应`Character`。

所有的闭包名称都是基本数据类名称的全称，并且首字母大写：

```java
int     - Integer
byte    - Byte
short   - Short
long    - Long
char    - Character
double  - Double
float   - Float
boolean - Boolean
```

每个闭包都可以通过传递基本数据类型来构造，并且也都可以像原本类型一样进行操作:

```java
Integer i = new Integer(i);
i=i+2;
```

但是不能将不兼容类型放入来构造:

```java
Integer i = new Integer(3.3f);    //不能将float传给int的闭包
```

### 装箱与拆箱

基本数据类型与包装类的转换被称为**装箱**和**拆箱**

装箱(boxing) 是将值类型转换为引用类型，如 int 转 Integer

装箱过程是通过调用包装类的valueOf方法实现的。

拆箱(unboxing) 是将引用类型转换为值类型，如 Integer 转 int

拆箱过程是通过调用包装类的xxxValue方法实现的，（xxx代表对应的基本数据类型）。

```java
  Integer i  10; //装箱(boxing)
  int n = i;    //拆箱(unboxing)
```

### 强转的方法

Java的开发者很痛恨C/C++中使用`(double)value;`这种强制转换方式，所以在Java中一律去掉了这种转换法则。

而且还规定，不同类型的数据之间不能相互转换。

比如`boolean`就不能和数类(`int`, `float`等)转换。

而作为最基础的整型提升还保留了下来（就是说`int`可以隐式转换为`float`这种）。

如果想要转换的话，必须构建对应类型的闭包，使用闭包的方法转换:

```java
int i = new Double(3.3).intValue();
```

一般转换的方法都是`<type>Value()`格式。

## 条件判断

和C++一模一样，都是短路的，不说了。

## 循环结构

和C++一模一样，就是for循环多了一种**foreach**循环（Java8新增）：

```java
int[] a = {1, 2, 3, 4};
for(int i : a){
  System.out.println(i);
}
```

其实和C++11的foreach也一样，只不过Java中没有auto关键字，所以你必须显式地写上变量类型。

## 枚举类型

和C++一样，使用`enum`关键字。

但是由于Java的关系，没有办法和数型之间相互转换。

而且枚举类型本身也属于类，所以在同一个文件中不能同时有public enum和public class。

## 包

所谓包，其实如果你学过Python的话就非常好理解了，就像是Python里面的包一样，通过文件夹来将源代码分层。

包的路径不使用`/`或者`\`作为分割，而是`.`，

也就是说，如果你的工程根目录在`~/Documents`，

那么你的`com.PQBox`包就在`~/Documents/com/PQBox`。

使用`package`可以告诉Java当前文件在哪个包中，**`package`语句必须是在代码的第一行** 。

再举个例子，如果你使用的是C++，那么你可能要这样管理你的工程：

```java
-MapEditor_Project
  -include
      - MapEditor.hpp
      -common
          - header.hpp
          - structs.hpp
  -src
      - MapEditor.cpp
      -common
          - structs.cpp
      - main.cpp
```

然后你得在`MapEditor.cpp`中写上`#include "include/MapEditor.hpp"`（假设你没有改变头文件搜索路径），

在`structs.cpp`中写`#include "include/common/structs.hpp"`

最后在`main.cpp`中以同样的方式包含你要的头文件。

但是在Java中，你可以这样:

```java
-MapEditor_Project    //Java文件可不分什么头文件源文件哦
  -src
    - MapEditor.java
    - Main.java
    -common
      - header.java
      - structs.java
```

然后你得在`MapEditor.java, Main.java`中写`package src;`，

在`header.java, structs.java`中写`package src.common;`，

然后在`Main.java`中使用`import src.common.header;`来引用`header.java`文件(**注意不能直接使用`import common.header;`必须指定包的全路径**)

实际上，包和C++的namespace一样开辟了新的命名空间，不同的包在不同的命名空间内，所以不同包内的类可以重名。

## 访问修饰符

由于所有字段和方法都必须在类里面，所以先了解一下访问修饰符。

和C++一样，存在`public、protected 、private`三种访问修饰符关键字。

但是访问权限却有四个，因为**不写访问修饰符也是一种访问权限——包内访问**

**注意不能在类前加`protected`**

`public`让其他包的代码可以访问这个包内的类和方法。

`protected`不能让其他包访问这个包内的信息，同一个包内也不能访问，类内可访问，并且继承下去之后仍然是`protcected`。

`private`包外包内都不能访问，继承之后仍为`private`。

不写（默认）访问权限是包内访问，其他的包不能访问，同一个包可以相互访问。

## 面向对象

### 类

使用`class`声明，类内的方法和字段全部必须加上访问修饰符，类末尾不需要分号：

```java
public class Person { // 包内可见
    public Person(String name, int age) {
      this.name = name; // 使用this来指代类中的字段和方法
      this.age = age;
    }
  
    public String GetName() {
      return name;
    }
  
    public int GetAge() {
      return age;
    }
  
    public String toString() { // 特殊方法
      return name + "'s age is " + age;
    }
  
    public boolean IsMale() {
      return isMale;
    }
  
    private String name;
    private int age;
    private boolean isMale = false;
  }
```

和Python一样，Java没有重载运算符，为了弥补这个缺陷，Java保留了一些特殊函数，

这些函数有一定的作用，你可以重写来实现自己的功能，最常用的就是`toString()`方法，

它会在类被输出的时候自动调用这个方法：

```java
System.out.println(new Person("HanMeiMei", 20)); // HanMeiMei's age is 20
```

在Java中，所有的变量都可以在声明时直接赋值，如果没有赋值，对象会是null，基本数据类型会是0，boolean变量会是false。

#### 构造函数

构造函数和C++一样的意义，没给构造函数的时候会默认给个空的。

#### 析构函数？

有析构函数函数吗？

没有，不过如果你硬要有一个的话，可以参考`finalize()`函数。

### 继承

继承的话使用`extends`关键字：

```java
class Child extends Person { //等价于C++的 class Child : public Person
  //...
}
```

没错，你没办法像C++一样控制继承的权限，也就是说你每次继承都是public继承。

Java取消了多继承，这意味着`extends`后面只能跟一个类名。

#### 所有类的基类

和Python一样，Java有着*万类之根*的类`Object`，这个类主要有一些RTTI的处理和定义`toString()`方法。

任何不继承的类默认继承`Object`。

#### 无法继承的类

在类前面加上`final`可以防止类被继承，如果继承了会报编译时错误。

### 重载

和C++一样，注意重载是不能改变访问权限的，比如你的方法在父类是`public`，到了子类必须仍然是`public`，不然会报编译时错误。

在重载中如果想要调用父类的方法，需要使用`super`关键字，super关键字代指当前类的父类:

```java
// Child class Employee
public class Employee extends Person {
    private int employeeId;
    private String jobTitle;

    public Employee(String name, int age, int employeeId, String jobTitle) {
        super(name, age); // 调用父类的构造函数
        this.employeeId = employeeId;
        this.jobTitle = jobTitle;
    }

    public int getEmployeeId() {
        return employeeId;
    }

    public int getAge() {
        return super.GetAge(); // 调用父类的对应方法
    }

    @Override
    public String GetName() {
        return super.GetName() + ", " + jobTitle;
    }
}
```

### 虚类

使用`abstract`来声明虚类和虚方法，如果一个类中有一个虚方法，那么这个类必须冠以`abstract`:

```java
abstract class Action{
  public Action(){
    System.out.println("Action!");
  }
  public abstract void ShakeHands();
  public abstract void Bow();
}
```

### 接口

当虚类只有虚函数，没有字段的时候就变成了接口：

```java
public interface Action {
    void ShakeHands();
    void Bow();
} 
```

接口使用`interface`关键字声明。

和类的区别在于：**接口的方法不能被实现**。

其实接口就是C++中没有成员变量的纯虚类。

但是和C++纯虚类一样，嘴上说着不能实现函数，其实还是可以实现的，但是你要冠以`default`关键字:

```java
public default void ShakeHands(){
  System.out.println("shake");
}
```

这样如果你不想实现这个函数，就是使用这个默认函数。

#### 实现接口

接口可以被实现：

```java
class Employee extends Person implements Action{
  public void ShakeHands(){
    System.out.println("child is shake your hand");
  }
  @Override    //注解，可以不加，加了表示下面的函数是被重载的
  public void Bow(){
    System.out.println("Bow");
  }
}
```

使用`implements`来实现接口，后面可以跟**多个接口**。

所以我们可以通过**接口**来实现**多继承**。

实现接口必须将接口所有的方法全部实现。

这里推荐使用`@Override`注解，他可以帮你辨别方法是否是被重写的，以方便在编译时找出重写错误。

### 静态方法和字段

当然是使用`static`来实现了。

但是别的类在使用的时候，不是用`::`而是使用`.`:

```java
class Person{
  static int PersonNum = 0;
}

class CalculatePerson{
  public int calcu(){
    return Person.PersonNum;
  }
}
```

## 注解

注解是个新东西，详见[廖雪峰老师的博客](https://www.liaoxuefeng.com/wiki/1252599548343744/1255945389098144)

## 异常

### 处理异常

格式和C++一模一样：

```java
try{
  //捕获异常
}catch(Exception e){
  //处理异常
}finally{
  //可选，不管异常出不出现都执行的代码块
}
```

但是Java中有一个至今我都感觉很烦的事情：

**如果一个方法可能会抛出异常（其方法体内有throw语句），那么你必须在使用这个方法的时候捕捉异常**，不然会有编译时错误（除了`RuntimeException`,`Error`及它们的子类）。

所以你常常会看见Java中存在大量的`try..catch`块，严重影响代码美观。

一般而言`Error`是不需要捕获的严重异常（你就算写了`try_catch`块也不用对其进行处理），如果碰到了整个程序适合直接挂。

`Exception`应当是可处理的异常。

### 抛出异常

#### 创建异常类

要抛出异常，你可以使用系统的异常类，或者自己创建。

方法是继承`Throwable`类并且改写其中的方法(一般可以改写`toString()`方法以便找到错误所在)。

而且你的函数内可能抛出什么异常，你就需要使用`throws`关键字指定函数要抛出什么异常：

```java
public void throwException(int type) throws IOException, NullPointerException{
  if(type==1)
    throw new IOException();
  else
    throw new NullPointerException;
}
```

而且你不能写出这样的代码:

```java
public void throwException(int type) throws IOException, NullPointerException{
  throw new IOException();
  throw new NullPointerException;
}
```

这样Java会认为一个函数不能同时抛出两个异常而给出编译时错误。

#### 使用throw抛出异常

这个很显然了，和C++一样：

```c++
throw new Exception();
```

### 断言

使用`assert`断言，会抛出`AssertException`。

```java
assert 1==false;
```

## 日志打印

Java自带了`java.util.logging`模块可以打印日志，

或者你可以下载使用广受好评的第三方模块`log4j`。

## 泛型

### 泛型的声明

和C++一样Java也拥有泛型：

```java
//泛型类
class Pair<T1, T2>{
  public Pair(T1 value1, T2 value2){
    first = value1;
    second = value2;
  }
  public T1 first;
  public T2 second;
}

//泛型接口
public interface Iterator<T>{
  T next();
  boolean hasNext();
}

//泛型方法（在非泛型类中的方法或者静态方法）
public static <T> void Function(T value){
  //...
}
```

Java中免去了`template`这种关键字。

泛型类和泛型接口直接在名称后面加上`<T1,T2,...>`即可，泛型方法则是在**返回值前**加上范型。

而且要注意的是：**静态方法是不受类的泛型类型影响的**，也就是说你不能这样写：

```java
class A<T>{
  public static void Say(T msg){
    System.out.println(msg);
  }
}
```

必须要将静态方法变为泛型方法才行:

```java
public static <T> void Say(T msg){}
```

当然这里的T和泛型类的T已经不是一个东西了。

### 泛型的使用

```java
Pair<Integer, String> pair = new Pair<>(21, "ads"); //new后面的泛型可以不加，编译器会推导
Function<String>("abs"); //泛型方法的使用
```

### 泛型的继承

和C++一样：泛型类可以继承泛型类，非泛型类只能继承泛型类的特化类。

也就是说可以这样写：

```java
class A<T> extends B<T>{
}

class C extends B<Integer>{
}
```

需要注意的是泛型类型必须是类类型，导致这种情况是因为`泛型擦除`。

### 泛型擦除及其他

泛型擦除的内容很多，参考[廖雪峰老师的博客](https://www.liaoxuefeng.com/wiki/1252599548343744/1265104600263968)。

extends和super通配符也一并参考。
