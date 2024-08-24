"""
-+---+---+---+---+---+---+
 | P | y | t | h | o | n |
-+---+---+---+---+---+---+
   0   1   2   3   4   5   6
  -6  -5  -4  -3  -2  -1
"""

s = "Python"
print(s[0])
print(s[-6])

print("##########旧式模板##########")
name = "Mars"
print("%s %s" % ("Hello", name))
print("PI: %.2f" % 3.14)

print("##########新式模板##########")
welcome = {"action": "Hello", "name": "Mars"}
print("%(action)s %(name)s" % welcome)
print("{0} {1}!".format("Again", "Mars"))
print("{action} {name}!".format(**welcome))  # 通过**welcome这种形式把Dictionary中的值读出来，传递给format方法。

print("---------String operations---------")
stringInDoubleQuotes = "Hello 'Python!'"
stringInSingleQuotes = 'Hello "Python!"'
stringInTripleQuotes = '''Hello "Python!"
This might be a 'long string'
going through multiple lines.'''

print(stringInTripleQuotes)

aNumber = 123
aString = str(aNumber)
# aString[0] = '5' # 字符串在Python中是只读的。一旦创建完成，就不能像C语言一样用位置去修改了

print(int('123'))

action = "Hello "
print(action[0:3])  # 字符串截取

names = "Mars!"
# print(dir(action)) # 查看字符串类型支持的所有方法，可以使用dir方法
# print(help(action.count)) #查看某个方法的具体帮助，可以使用help方法

welcome = action + names
print(welcome)
print(welcome.upper())
print(welcome.lower())

trim = " Hello "
print(trim)
print(trim.strip())  # 去掉字符串的首尾空格

print("----------list operations----------")

list1 = []
list2 = list()

num_list = [8, 6, 5, 7, 2]
mixList = [1, "one", 1.0]

nestedList = [num_list, mixList]

# print(mixList + num_list)

# print(mixList.extend(num_list))  # extend的返回值是None
# print(mixList)

# num_list.append(9)
# print(num_list)

# num_list.append([11,12])
# print(num_list)

# list作为一个类对象，是个引用类型，因此，要想保留sort，extend或append之前的值，
# 单纯的把它赋值给另一个变量是不行的, 
# 必须使用deep copy的做法，明确调用list的copy方法：
num_list1 = num_list.copy()

print("sort before ", num_list1)
num_list.sort()
print("sort after  ", num_list)

# 这样，对numList的操作就不会影响到numList1了。

# 访问list中的元素
print("num_list[0]    ", num_list[0])
print("num_list[0:3]  ", num_list[0:3])
print("num_list[0:-2] ", num_list[0:-2])

num_list.insert(1, 999)  # insert的第一个参数是要插入的位置，第二个参数是要插入的值。
print("after insert ", num_list)

num_list.insert(100, 2)  # 如果你使用的位置超过了list的最大长度，insert就会把元素插在list末尾
print("after insert ", num_list)

# 第一种，是使用pop方法, 从list中删除元素
print(num_list.pop(0))
print(num_list)

# 第二种，是使用remove方法,它直接删除list中第一个和参数值相等的元素，并且，没有返回值。
print(num_list.remove(5))
print(num_list)

# 第三种，是使用全局的del函数，我们直接来看代码：
del (num_list[0:5:2])  # 表示从位置0开始，每两个元素删掉一个，一直到位置5。因此，执行之后，结果就是[6, 8]了。
print("after del ", num_list)
