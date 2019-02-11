package myvm.example;

import java.lang.System;
import java.lang.Object;

public class HelloWorld {
    public static final int Val = 1;
    public static void main(String[] args) {
        Object o = new Object();
        System.out.println("Hello, world!");
        System.out.println(o.toString());
    }
}