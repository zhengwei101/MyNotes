# Lua语法快速入门

## Hello World

```lua
print("hello world") --hello world
```

## 注释

单行注释

```lua
-- 两个横线开始的单行注释
```

多行注释

```lua
--[[
    第一行注释的内容
    第二行注释的内容
--]]
```

## 表达式

### 关系操作符

跟其他语言类似，Lua支持 `<`、`>`、`<=`、`>=`、`==`、`~=`。

### 逻辑操作符

逻辑操作有 `and`、`or`和 `not`。

'or'和 'and'都是可短路的（如果已足够进行条件判断则不计算后面的条件表达式）。

### 字符串连接

使用 `..`操作

```lua
print("hello".."world") --helloworld
```

## 数据类型

Lua有如下数据类型：

| 数据类型 | 描述 |
| :------- | :----------------------------------------------------------- |
| nil      | 这个最简单，只有值nil属于该类，表示一个无效值（在条件表达式中相当于false）。 |
| boolean  | 包含两个值：false和true。 |
| number   | 用于表示实数，Lua没有整数类型，只有double型。 |
| string   | 字符串由一对双引号或单引号来表示 |
| function | Lua中，函数是作为第一类值，可以存储在变量里，可以通过参数传递给其他函数，还可以作为其他函数的返回值。 |
| table    | Lua 中的表(table)其实是一个"关联数组"(associative arrays)，数组的索引可以是数字、字符串或表类型。在 Lua 里，table 的创建是通过"构造表达式"来完成，最简单构造表达式是{}，用来创建一个空表。 |

## 变量

在Lua脚本中，如果不把变量声明为 `local`，无论在函数内还是函数外，均为全局变量。将该变量设置为 `nil`时，会回收该变量

例：

```lua
a = 10 -- 全局变量

function foo1()
    print(a) -- 10
    b = 11 -- 全局变量
    local c = 12 -- 局部变量
    print(c) -- 12
end

function foo2()
    print(b) -- 11
    print(c) -- nil
    b = nil
end

function foo3()
    print(b) -- nil
end

foo1()
foo2()
foo3()
```

## 循环

### for

基本语法：

var 从 exp1 变化到 exp2，每次变化以 exp3 为步长递增 var，并执行一次 **"执行体"**。

exp3 是可选的，如果不指定，默认为1。

```lua
for var=exp1,exp2,exp3 do  
    <执行体>  
end  
```

示例：

```lua
-- 依次输出 0，1，2，3
for i = 0, 3 do
    print(i) 
end

-- 依次输出 3，2，1，0
for i = 3, 0, -1 do
    print(i)
end
```

### for循环遍历数组

```lua
--遍历数组，`#arr`表示arr的长度
arr = {"a", "b", "c"}
for i = 1, #arr do
    print(arr[i])
end
```

使用pairs的方式

```lua
-- 遍历数组
arr = {"a", "b", "c"}
for i, v in pairs(arr) do
    print(i, v)
end

-- 输出
-- 1       a
-- 2       b
-- 3       c
```

### while

```lua
i = 0
a = 0
while i <= 10 do
    a = a + i
    i = i + 1
end
print(a)        -- 55
```

## 流程控制

### if

```lua
a = 1
if a > 0 then
    print("a>0")   -- a>0
end
```

### if ... else

```lua
a = 0
if a > 0 then
    print("a>0")
else
    print("a<=0")   -- a<=0
end
```

### if ... elseif ... else

```lua
a = 0
if a > 0 then
    print("a>0")
elseif a == 0 then
    print("a=0")     -- a=0
else
    print("a<0")
end
```

## 函数

在Lua中，函数可以作为一个类型的值，储存在变量中，也可以作为其他函数的返回值

```lua
-- 函数
function add(a, b)
    return a + b
end
function add2()
    return add
end
f1 = add
f2 = add2()
print(add(1, 1)) -- 2
print(f1(1, 2))  -- 3
print(f2(1, 3))  -- 4
```
