package com.PQBox;

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
