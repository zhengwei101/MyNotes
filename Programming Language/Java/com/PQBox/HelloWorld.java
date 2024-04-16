package com.PQBox;

public class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello World");
    System.out.println(new Person("HanMeiMei", 20)); // HanMeiMei's age is 20

    Employee employee = new Employee("Junior", 30, 4452, "Software Manager");
    System.out.println(employee.GetName() + " " + employee.getAge() + " (" + employee.getEmployeeId() + ")");

    Playable football = new Football();
    Playable basketball = new Basketball();

    // Call the "play" method on each Playable object to play different sports
    football.play();
    football.score();
    basketball.play();
    basketball.score();
  }
}
