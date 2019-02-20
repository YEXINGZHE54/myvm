package myvm.example;

import java.lang.System;
import java.lang.Object;
import java.lang.Integer;

public class HelloWorld {
    public static final int Val = 1;
    public static void main(String[] args) {
        int[] arr = {1,2,3};
        System.out.println(arr[0]);
        System.out.println(args[0] + "def");
        System.out.println(int.class.getName());
        try {
            foo();
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }

    public static void foo() throws Exception {
        throw new Exception("abc");
    }
}