package com.PQBox;

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

    @Override //注解，可以不加，加了表示下面的函数是被重载的
    public String GetName() {
        return super.GetName() + ", " + jobTitle;
    }
}