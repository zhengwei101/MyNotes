# JavaScript

## ECMAScript

ECMAScript5.1 也叫ECMA-262, 2011年6月发布

ECMAScript6 2015年6月正式发布

## DOM 文档对象模型

1. DOM并不是JavaScript专有的，其他语言也能操作DOM，比如PHP、JAVA、Python等。

## BOM 浏览器对象模型

1. 没有标准。规定了JavaScript操作浏览器的方法和接口。

## 应用领域

- 验证表单
- 特效、交互
- 数据交互
- 游戏（俄罗斯方块、贪吃蛇）
- 桌面应用(Electron)
- 移动应用程序(App)
- 后台(node.js)
- VR

## 块级作用域和常量

let和const的出现让 JS 有了块级作用域，还可以像强类型语言一样定义常量。

由于之前没有块级作用域以及 var 关键字所带来的变量提升，经常给我们的开发带来一些莫名其妙的问题。

下面看几个简单的demo理解理解

```js
//demo 1
function f1() {
    let n = 5;
    if (true) {
        let n = 10;
    }
    console.log(n); // 5
}

// 嵌套循环不会相互影响
for (let i = 0; i < 3; i++) {
    console.log("out", i);
    for (let i = 0; i < 2; i++) {
        console.log("in", i);
    }
}

// 引用离开作用域的b会报错
if (true) {
    let b = 'zfpx';
}
//console.log(b);// ReferenceError: b is not defined

//demo2
const PI = 3.1415; // PI为常量
console.log(PI);   // 3.1415
//PI = 3;          // TypeError: "PI" is read-only
```

重复定义会报错

```js
//demo3
if (true) {
    let a = 1;
    //let a = 2; //Identifier 'a' has already been declared
}
```

不存在变量的提升

```js
//demo4
//console.log(i) //ReferenceError: Cannot access 'i' before initialization
let i = 10;
```

## 函数闭包

一般闭包要解决的问题就是要想办法间接的获取函数内部数据的使用权。

写一个函数，函数内定义一个新函数，返回新函数，用新函数获取函数内的数据

```js
function foo(){
    var num = Math.random();
    function func() {
        return num;
    }
    return func;
}

var f = foo();

// f可以直接访问num, 而且多次访问，访问的也是同一个，并不会返回新的num
var res1 = f();
var res2 = f();
console.log(res1);
console.log(res2);

```

写一个函数，函数内定义一个对象，对象中绑定多个函数（方法），返回对象，

利用对象的方法访问函数内的数据

```js
function foo() {
    var num = Math.random();
    //分别定义get和set函数，使用对象进行返回
    return {
        //get_num负责获取数据
        get_num: function () {
            return num;
        },
        //set_num负责设置数据
        set_num: function (value) {
            num = value;
        }
    }
}

var obj = foo();
console.log(obj.get_num());
obj.set_num(111);
console.log(obj.get_num()); // 111
```


## 解构

解构意思就是分解一个东西的结构，可以用一种类似数组的方式定义N个变量，可以将一个数组中的值按照规则赋值过去。

```js
var [name, age] = ['hyan', 8];
console.log(name, age); // hyan 8

var [x, y] = getVal(),                         //函数返回值的解构
    [name, , age] = ['ddan', 'male', 'secrect']; //数组解构

function getVal() {
    return [1, 2];
}

console.log('x:' + x + ', y:' + y);//输出：x:1, y:2
console.log('name:' + name + ', age:' + age);//输出：name:ddan, age:secrect

```

数组、对象和字符串的解构赋值

```js
'use strict';
// 数组的解构赋值
let [x, [[y], z]] = [1, [[2], 3]];
console.log(x); // 1
console.log(y); // 2
console.log(z); // 3

// 对象的解构赋值
var { foo, bar } = { foo: "aaa", bar: "bbb" };
console.log(foo);   // "aaa"
console.log(bar );  // "bbb"

// 字符串的解构赋值
const [a, b, c, d, e] = 'hello';
console.log(a + b + c + e); // 'helo'
```

## Arrows 箭头函数

箭头函数简化了函数的的定义方式，一般以 "=>" 操作符左边为输入的参数，而右边则是进行的操作以及返回的值(Inputs=>outputs)。

箭头函数根本没有自己的this，导致内部的this就是外层代码块的this。

正是因为它没有this，所以也就不能用作构造函数，从而避免了this指向的问题。

请看下面的例子。

```js
let [a, b] = [1, 2];

//以前
function add(a, b) {
    console.log(a + b);
}
add(a, b);

//现在
let new_add = (a, b) => console.log(a + b);
new_add(a, b);

//forEach
var numbers = [1, 2, 3, 4];
numbers.forEach(function (item, index, array) {
    console.log(item + "\t" + index + "\t" + array);
});

var array = [7, 8, 9];

//传统写法
array.forEach(function (v, i, a) {
    console.log(v);
});

//ES6, 输入参数如果多于一个要用()包起来，函数体如果有多条语句需要用{ }包起来
array.forEach(v => console.log(v));
```

## Template Strings 字符串模板

ES6中允许使用反引号 ` 来创建字符串，此种方法创建的字符串里面可以包含由美元符号加花括号包裹的变量${vraible}。

```js
//产生一个随机数
var num = Math.random();

//将这个数字输出到console
console.log(`your num is ${num}`);

let name = 'guorong';
let age = 18;

console.log(`${name} was ${age}`)
```
