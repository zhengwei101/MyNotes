-- https://wiki.luatos.com/
-- http://www.lua.org/manual/

print("Hello World!")

a = 1       -- global variable
local b = 2 -- local variable
c, d, e, f, g = 3, 4, 5, 6 -- structure binding

print(g, h) -- nil

-- type: number
a = 0x1A
b = 3.14e10

a = 7
b = 5
print("+", a + b)
print("-", a - b)
print("*", a * b)
print("/", a / b)   --浮点除法
print("%", a % b)
print("^", a ^ b)   --乘方
--print("//", a // b) --向下取整除法
--print("<<", 1 << 3) --左移
--print(">>", 8 >> 3) --右移

-- type: string
print("--------------------------------- string")
local name1 = "abc \n abc"
local name2 = 'def'

local name3 = [[
line2\n
line3]]
print("[[]]", name3) -- 转义字符在这里不起作用

local name4 = name1..name2
print(name4)

local str = tostring(456)
local num = tonumber("789")
local failed = tonumber("123abc") -- nil，转换失败，返回nil

local length = #str --返回str的长度

-- function
print("--------------------------------- function")
function fun(x, y)
    print("x:", x, "y:", y)
end
fun(314, 3.14)
fun(255) -- y:nil
local ret = fun()
print(ret) -- nil

fun2 = function(a, b)
    return a, b
end

a, b = fun2(314, 3.14)

-- 数组
print("--------------------------------- table number_index")
tab = {1, 2, 3, 4, 5, 6, 7, 8, 9, "ac", {}, function() end}
print(tab[1]) -- starts from 1

print(tab[100]) -- nil
tab[100] = 123 -- add new


local length_of_tab = #tab -- length

table.insert(tab, "insert to tail")
table.insert(tab, 2, "insert to second")

local removed = table.remove(tab, 1)


print("--------------------------------- table string_index")
person = {
    name = "Cherno",
    age = 24,
    gender = "male",
    [",;:$"] = "bad variable name" -- bad name
}
print(person.name)

print(person["gender"])
print(person[",;:$"])

person["mom"] = "Cherno's mom" -- add new
print(person["pap"]) -- nil

-- 全局表
print("--------------------------------- _G")
print(_G) -- lua中所有全局变量都放在全局表_G中

global_variable = "global_variable"
local local_variable = "local_variable"
print(_G["global_variable"])
print(_G["local_variable"]) -- nil

-- table.insert
print(_G["table"])
print(_G["table"]["insert"])

-- 真假
print("--------------------------------- true & false")
local T = true  -- 0 and anything but nil, false
local F = false -- nil, false

print(">",  1>2)
print(">=", 1>=2)
print("<",  1<2)
print("<=", 1<=2)
print("==", 1==2)
print("~=", 1~=2)


T = 0
F = nil
print(T and F) -- nil , 注意这里是返回某一个操作数的值，而不是常规意义的布尔值
print(T or F) -- 0
print(not T) -- false

print(T > 0 and "YES" or "NO") -- NO

print("--------------------------------- if")
if F then
    print(false)
elseif T then
    print(true)
else
    print("Error")
end

print("--------------------------------- for")
for i = 1, 10 do
    print(i)
end

for i = 1, 10, 2 do
    print(i)
end

for i = 10, 1, -1 do
    print(i)
end

for i = 1, 10 do
    local i = 4766
    print(i)
end

for i = 1, 10 do
    print(i)
    if i == 5 then break end
end

print("--------------------------------- while")
local it = 10
while it > 1 do
    print(it)
    it = it - 1
    if(it == 5) then break end
end

print("--------------------------------- Appendix")
local string_char = string.char(0x30, 0x31, 0x32, 0x33, 0x00) -- 0x00 ok
local askii_code = string.byte(string_char, 1)