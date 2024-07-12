# 现代JSON库使用指南

> I don’t know who is the king of kings,
> But I do know nlohmann::json is the data-structure of data-structures.

此数据结构一出，简直是惊天地泣鬼神，有如下的好处：

- 现代风格，代码简洁；
- 万能型数据结构（能装载任何类型的数据，包括自定义数据）；
- 蕴含`OO`思想，可以作为`OO`数据库使用；

到底有没有这么好，读完本专题你就知道了。

关于JSON(JavaSrcipt Object Notation)的详细格式可以参考[官网](https://www.json.org/)，十分简短明了。
这里就举几个例子帮助理解。

```json
{
    "pi": 3.141,
    "happy": true,
    "name": "Niels",
    "nothing": null,
    "answer": {
        "everything": 42
    },
    "list": [1, 2, 3],
    "object": {
        "currency": "USD",
        "value": 42.99
    }
}
```

最重要的概念是object，用`{}`括起来表示一个object。所有JSON一头一尾都是`{`和`}`，说明整个JSON都是一个object。

在object里面可以存放多种数据格式，统称 member。

最最常见的是`key/value`对，其中key是string类型，需要用`" "`括起来。

value可以是任何常见的类型，比如int，bool，string，null指针。

每个member之间用 `,`分隔。

value里还可以**嵌套**放置object。

value还可以是数组，用`[]`括起来，每一个值用`,`分隔。

上面代码中最后一个value是一个object，object内部有两个member。

可以看出来JSON整体呈现一种层次结构, 各个member之间的顺序并不重要。

了解了JSON之后遇到的问题就是如何解析、输出JSON，这个不需要自己写了，有许多优秀的库来做这件事。

下面介绍这个库的基础操作，看了这些你就会明白为什么叫做现代C++风格。

## 插入普通member

```cpp
#include <iostream>
#include <iomanip>

#include "nlohmann/json.hpp"
using json = nlohmann::json; //命名空间， nlohmann是作者的简称

using namespace std;

int main()
{
 json myJson; //可以把json类型理解为一种数据结构
 
 //加入普通的member, key/value对
 myJson["pi"] = 3.14;                //数字
 myJson["pass"] = true;              //bool
 myJson["region"] = string("Asia");  //字符串，需要显示使用string
 myJson["nope"] = nullptr;           //空指针
 
 //复杂一点的情况
 myJson["vector"] = {1,2,3};         //数组
 myJson["1st"]["2nd"] = string("object inside"); //嵌套object: 写法1
 myJson["1st"]["3rd"] = 121;
 myJson["moreobj"] = {{"obj1", "hello"}, {"obj2", "world"}}; //嵌套object: 写法2
 
 cout << myJson << endl; //非格式化输出json
 cout << setw(4) << myJson << endl; //格式化输出json，重载了setw
 
 return 0;
}
```

很清晰明了，直接使用`[]`插入key，后面跟value，这个用法与map很相似。

对于数组，可以用初始化列表

对于嵌套object，可以有两种写法

- 第一种是连续n个`[]`
- 另一种还是借助初始化列表

比较起来还是前者更加灵活，只要这样理解即可：n个`[]`代表深入了n层

另外跟cout配合确实很人性化，输出格式化的结果如下：

```json
{
    "1st": {
        "2nd": "object inside",
        "3rd": 121
    },
    "moreobj": {
        "obj1": "hello",
        "obj2": "world"
    },
    "nope": null,
    "pass": true,
    "pi": 3.14,
    "region": "Asia",
    "vector": [1, 2, 3]
}
```

应该注意到在插入member的时候我们没有指定value的类型，这个会自动推断出来。

## 获取value

看了下面的代码就知道为什么叫现代C++了，因为设计跟STL一致

```cpp
//获取数据
//方式一, find迭代器
auto objJson = myJson.find("region");
if (objJson != myJson.end()) {
    string s = objJson.value(); //必须显示指明接收value值的类型
}
try {
    //方式二, at
    bool pass = myJson.at("pass"); //如果不存在会抛 out_of_range 异常
    //方式三, []
    vector<int> v = myJson["vector"]; //如果不存在会抛 type_error 异常
}
catch (nlohmann::detail::exception& e) {
    string error = std::format("json throw an error:{}, Try to Fix it!", e.what());
    cout << error << endl;
}
//获取嵌套object的方式类似
int n = myJson["1st"]["3rd"];
double d = myJson["1st"]["3rd"].get<double>(); //可以显示指定value的类型
```

三种方式，就是个map。

- 最安全的是find，但代码书写最麻烦
- at简便许多，也能够告知是否存在key
- `[]`最简单，不过也最容易出现意外

当然对于嵌套object还是用`[]`最直观

如果使用auto则得不到value的值，而是一个类似object的东西，
这其实也说明了`[]`返回的到底是什么东西，以及为什么可以不断`[]`逐层深入。

从获取value也可以知道一件事：**如果不知道JSON的结构则根本无法解析该JSON**

## 格式化JSON

处理JSON最常见的一种情况是网络上传过来一个JSON串或者有一个JSON文件，需要从中提取出相应信息，然后继续计算。
对于这种情况`json for modern c++`提供了非常简便的方法。

只要使用命名空间里面的`parse`函数即可

如果传入的字符串不符合JSON规范呢？
会抛出`parse_error`异常，记得使用时配合`try-catch`!

还有一点要说明，网络传输JSON通常都是UTF8编码，这个情况`json for modern c++`是可以处理的，得益于UTF8可以放到string里面，所以解析起来其实是一样的。

另外如果value是二进制也没有关系，一样放到string里解析即可。

去格式化JSON就更简单了，直接用dump即可。

```cpp
string strNoFormat = json.dump();
```

以上就是nlohmann::json的基础用法，其实任何json库都支持这些功能，没什么大不了的。

接下来介绍nlohmann::json的神奇扩展用法，**nlohmann的json绝对不是json那么简单。**

## 一、完美支持STL

```cpp
json myJson;  //万能数据结构
//1.可以装载n为vector
vector<vector<int>>form;
form.push_back(vector<int>({ 1,2,3,4,5 }));
form.push_back(vector<int>({ 6,7,8,9,10 }));
myJson["2D"] = form;

auto formR = myJson["2D"].get<vector<vector<int>>>();
```

再来点更复杂的

```c++
//2.可以装载key为string的关联容器
map<string, double> wageTable;
wageTable["Bob"] = 12.5;
wageTable["Tom"] = 8.9;
wageTable["Mary"] = 15.8;
myJson["WageTable"] = wageTable;

map<string, double> Rtable = myJson["WageTable"];

cout << myJson.dump();
```

输出结果：

```json
{"WageTable":{"Bob":12.5,"Mary":15.8,"Tom":8.9}}
```

## 二、支持自定义结构体

更神奇的是还可以支持自定义类型，只要我们给出该类型与json的转换关系即可。

```c++
namespace magic {
struct Person
{
    string Name;
    int Age = 0;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(Person, Name, Age)
};

struct DownParam
{
    string Url;
    string SaveDir;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(DownParam, Url, SaveDir)
};
} // namespace magic

//3.支持自定义结构体
magic::Person p = {"Peter", 18};
magic::DownParam d = {"http://a/b.txt", "/user/download"};
json j = p;
cout << j << endl;

j.clear();

j = d;
cout << j << endl;

json myJson;
myJson["Person"] = p;
myJson["DownParam"] = d;

auto re = myJson["Person"].get<magic::Person>();
cout << re.Name << ", " << re.Age << endl;

auto dw = myJson["DownParam"].get<magic::DownParam>();
cout << dw.Url << ", " << dw.SaveDir << endl;
```

输出

```sh
{"Age":18,"Name":"Peter"}
{"SaveDir":"/user/download","Url":"http://a/b.txt"}
Peter, 18
http://a/b.txt, /user/download
```

## 三、支持自定义类

```c++
class ComplexOne
{
private:
    vector<string> vStr;
    std::unordered_map<string, magic::Person> hashNameIndex;

public:
    vector<string> GetVStr() { return vStr; }

    void AddStr(const string& str) { vStr.emplace_back(str); }

    void AddPair(string key, magic::Person p) { hashNameIndex[key] = p; }

    friend void to_json(json& j, const ComplexOne& cp);
    friend void from_json(const json& j, ComplexOne& cp);
};

void to_json(json& j, const ComplexOne& cp)
{
    j = json{{"vstr", cp.vStr}, {"hash_name_index", cp.hashNameIndex}};
}

void from_json(const json& j, ComplexOne& cp)
{
    j.at("vstr").get_to(cp.vStr);
    j.at("hash_name_index").get_to(cp.hashNameIndex);
}

int main()
{
    //4. 支持自定义类
    magic::Person p = {"Peter", 18};

    ComplexOne one;
    one.AddStr("apple");
    one.AddStr("banana");
    one.AddPair("001", p);

    json j = one;
    cout << j << endl; //输出：{"hash_name_index":{"001":{"Age":18,"Name":"Peter"}},"vstr":["apple","banana"]}

    ComplexOne anotherOne = j;

    for (auto& v : anotherOne.GetVStr()) {
        cout << v << endl;
    }

 return 0;
}
```

这个类看起来比较复杂了，主要是成员变量复杂。
然后**声明了两个友元**，提供类与json相互转化的方法，之所以声明友元是因为转化需要直接操作成员变量。
这里还涉及到了自定义结构体Person的转化，因为我们已提供了转化方法，所以这里也可以直接写。

如果打印出节点，可以看到json内部是如何储存这个ComplexOne类的。
本质上还是字符串，也就是说对于任何一个自定义类，当我们将其插入json时，json自动调用`to_json`把这个类翻译成字符串。
当我们get时，json调用`from_json`实例化一个新的类并返回。
拿到这个新的类，我们可以使用类成员函数操作，就跟正常的一样。

当然使用自定义类有几个注意事项，可以去官网查看。

还有json也不是完全万能的，比如不能放指针，这个跟json格式不符。

其他的一些用法，比如获得指针或引用，以及二进制压缩等，可以去[官网](https://github.com/nlohmann/json/)查看。

这种万能型数据结构，特别适合做信使，在不同的模块间进行传递，这也是**OO编程**的利器。

## 四、json蕴含的OO思想

我们使用的数据库，都是所谓的关系型数据库，使用SQL语句来操作。
更具体一点，数据库里面都是各种表，表与表进行某种链接，表内存储数据实体。
这种关系型数据库与面向对象思想是不兼容的。

但是json就很有意思了，你可以把它视为一种面向对象的数据库。怎么讲？
面向对象通常会出现继承，这是一个树状图，而json的存储方式刚好是这样的。

比如可以写如下代码：

```cpp
json OOD;
OOD["Animal"]["Bird"]["Sparrow"] = vector<Sparrow>()
```

这背后蕴含的是sparrow这个数据实体是`animal->bird->sparrow`这个路径下自然延展的结果。

```cpp
OOD["Animal"]["Bird"]["Swallow"] = map<string, Swallow>()
```

如果我们的每只燕子都有名字，那么可以通过map查询。

得益于json库支持自定义类型，我们可以把class存储到json里面，并根据需要获取拷贝或引用。

这跟OO又有什么关系呢？

我们可以轻易获取`OOD["Animal"]["Bird"]`这个节点，这意味着什么？

这意味着Bird节点下的Sparrow和Swallow都一并获得了。换言之我们可以按照OO的思路存储数据，并且可以轻易地在继承体系中获取任意一个节点之下的数据。
不需要关系型数据库的知识和操作，我们用json存储数据自然而然就是OO了。

## `json::array`和`json::object` 函数用法

显性指定类型或表达一些特定的意图，函数：`json::array`和`json::object`可满足您的需求：

```c++
// a way to express the empty array []
json empty_array_explicit = json::array();

// ways to express the empty object{}
json empty_object_implicit = json({});
json empty_object_explicit = json::object();

// a way to express an _array_ of key/value pairs [["currency", "USD"], ["value", 42.99]]
json array_not_object = json::array({ {"currency", "USD"}, {"value", 42.99} });
```

## 序列化/反序列化

```cpp
To/from strings
```

您可以通过附加_json到字符串来创建JSON值（反序列化）：

```cpp
// create object from string literal
json j = "{ \"happy\": true, \"pi\": 3.141 }"_json;

// or even nicer with a raw string literal
auto j2 = R"(
  {
    "happy": true,
    "pi": 3.141
  }
)"_json;

if (j.at("happy")) {
    cout << "happy is true" << endl;
}
```

或者使用`json::parse()`明确表达

```cpp
// parse explicitly
auto j3 = json::parse("{ \"happy\": true, \"pi\": 3.141 }");
```

您还可以获取JSON值的字符串表示形式（序列化）：

```cpp
// explicit conversion to string
std::string s = j.dump();    // {\"happy\":true,\"pi\":3.141}

// serialization with pretty printing
// pass in the amount of spaces to indent
std::cout << j.dump(4) << std::endl;
// {
//     "happy": true,
//     "pi": 3.141
// }
```

注意序列化和赋值之间的区别：

```cpp
// store a string in a JSON value
json j_string = "this is a string";

// retrieve the string value (implicit JSON to std::string conversion)
std::string cpp_string = j_string;

// retrieve the string value (explicit JSON to std::string conversion)
auto cpp_string2 = j_string.get<std::string>();

// retrieve the serialized value (explicit JSON serialization)
std::string serialized_string = j_string.dump();

// output of original string
std::cout << cpp_string << " == " << cpp_string2 << endl;

// output of serialized value
std::cout << j_string << " == " << serialized_string << std::endl;
```

不同之处是：`.dump()`始终返回序列化值，而`.get<std::string>()`返回最初存储的字符串值。
**请注意，该库仅支持UTF-8，当您在库中存储具有不同编码的字符串时，调用dump()可能会抛出异常。**

## 二进制格式(CBOR, MessagePack, and UBJSON)

虽然 JSON 是一种无处不在的数据格式，但它并不是一种适合数据交换（例如通过网络）的那种非常紧凑的格式。

因此，该库支持：

- [BJData](https://json.nlohmann.me/features/binary_formats/bjdata/) (Binary JData),
- [BSON](https://json.nlohmann.me/features/binary_formats/bson/) (Binary JSON),
- [CBOR](https://json.nlohmann.me/features/binary_formats/cbor/) (简明二进制对象表示，Concise Binary Object Representation),
- [MessagePack](https://json.nlohmann.me/features/binary_formats/messagepack/), and
- [UBJSON](https://json.nlohmann.me/features/binary_formats/ubjson/) (通用二进制规范，Universal Binary JSON)

以有效地将JSON值编码为字节向量（byte vectors ）和解码此类向量（vectors）。

```cpp
// create a JSON value
json j = R"({"compact": true, "schema": 0})"_json;

// serialize to CBOR
std::vector<std::uint8_t> v_cbor = json::to_cbor(j);

// 0xA2, 0x67, 0x63, 0x6F, 0x6D, 0x70, 0x61, 0x63, 0x74, 0xF5, 0x66, 0x73, 0x63, 0x68, 0x65, 0x6D, 0x61, 0x00

// roundtrip
json j_from_cbor = json::from_cbor(v_cbor);

// serialize to MessagePack
std::vector<std::uint8_t> v_msgpack = json::to_msgpack(j);

// 0x82, 0xA7, 0x63, 0x6F, 0x6D, 0x70, 0x61, 0x63, 0x74, 0xC3, 0xA6, 0x73, 0x63, 0x68, 0x65, 0x6D, 0x61, 0x00

// roundtrip
json j_from_msgpack = json::from_msgpack(v_msgpack);

// serialize to UBJSON
std::vector<std::uint8_t> v_ubjson = json::to_ubjson(j);

// 0x7B, 0x69, 0x07, 0x63, 0x6F, 0x6D, 0x70, 0x61, 0x63, 0x74, 0x54, 0x69, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6D, 0x61, 0x69, 0x00, 0x7D

// roundtrip
json j_from_ubjson = json::from_ubjson(v_ubjson);

//other 
```

## 使用注意事项

### 1. `nlohmann::json`支持隐式转换，但是不要使用，因为行为不确定 (??)

看下面的例子及输出结果

```cpp
long long lnum = 327221437;

nlohmann::json a;
a["num1"] = 327220625;

nlohmann::json b;
b["num2"] = 900000;

if (lnum - a["num1"] > b["num2"])
{
    cout << lnum << endl;
    cout << a["num1"] << endl;
    cout << b["num2"] << endl;
    cout << lnum - a["num1"] << endl;
    cout << lnum - a["num1"].get<long long>() << endl;
}
```

`lnum - a["num1"] > b["num2"]`并不成立，但返回为true，

如果写成`lnum - a["num1"].get<long long>() > b["num2"].get<long long>()`就没问题

> **修正： 在VS2022中，(lnum - a["num1"] > b["num2"])返回为false**

#### 2.调用`json::parse(iter.begin(), iter.end())`的时候需要注意

作者也说了，是`[iter.begin(), iter.end())`，左闭右开的，所以如果是一个数组保存了json字符串，那么就是a, `a+len`，而不是`a, a+len-1`

#### 3.`json::parse` 如果是非法的json会直接丢一个异常

可以通过`json::accept`判断是否合法

#### 4. `multiple json from file`

一种不错的方法是将整个文件解析为JSON，然后遍历该JSON并将值添加到向量中。

```cpp
#include <nlohmann/json.hpp>
#include <vector>
#include <fstream>

using nlohmann::json;

int main()
{
    std::vector <json> allJson;
    std::ifstream i("test.json");
    json j = json::parse(i);
    for (auto it = j.begin(); it != j.end(); ++it)
    {
        if (j.is_array())
        {
            allJson.push_back(it.value());
        }
        else
        {
            // If it's not an array, it's an object, and from what I understood you want the key too.
            allJson.push_back({it.key(), it.value()});
        }
    }
    return 0;
}
```

## 使用NlohmannJson写JSON保留插入顺序

在使用过程中，遇到了一个问题是没办法保持插入的顺序，每个插入的键值对会按照字符串的顺序排列的，因为其内部用到了std:map。
查看了github的主页说明，是这么说的：

> By default, the library does not preserve the insertion order of object elements. This is standards-compliant, as the JSON standard defines objects as "an unordered collection of zero or more name/value pairs". If you do want to preserve the insertion order, you can specialize the object type with containers like tsl::ordered_map (integration) or nlohmann::fifo_map (integration).

这段话的意思是JSON标准的定义是零个或多个键值对对的无序集合，如果要保证插入顺序，可以使用`tsl::ordered_map(integration)`或`nlohmann::fifo_map(integration)`等容器专门化对象类型。
`nlohmann::fifo_map`同样在github上找到，“专门化对象类型”的意思是`nlohmann/json`组件内部用到了很多`std`容器，只需要将其替换成可以保存插入顺序的容器就可以了，也就是`nlohmann::fifo_map`。

重新找了一些英文资料，最终找到的解决方案如下：

```cpp
#include "json.hpp"
#include "fifo_map.hpp"
#include <iostream>

using namespace nlohmann;

// A workaround to give to use fifo_map as map, we are just ignoring the 'less' compare
template<class K, class V, class dummy_compare, class A>
using my_workaround_fifo_map = fifo_map<K, V, fifo_map_compare<K>, A>;
using my_json = basic_json<my_workaround_fifo_map>;

int main()
{
    my_json j;
    j["f"] = 5;
    j["a"] = 2;
    my_json j2 = {
      {"pi", 3.141},
      {"happy", true},
      {"name", "Niels"},
      {"nothing", nullptr},
      {"answer", {
        {"everything", 42}
      }},
      {"list", {1, 0, 2}},
      {"object", {
        {"currency", "USD"},
        {"value", 42.99}
      }}
    };

    std::cout << j.dump(4) << std::endl;
    std::cout << j2.dump(4) << std::endl;

    return 0;
}
```

运行后，可以看到输出的JSON不再是字符串顺序而是插入顺序.

### 使用`nlohmann::json`的开源库

1. JsonRpcCXX

### 参考资料

<https://dins.site/coding-lib-json-intro-chs/>
