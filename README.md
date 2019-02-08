# 关于

这是一个用于学习目的的以GO实现的VM，能加载类并对class文件做解释执行

可以尝试实现不同的内存堆GC算法，观察不同算法的指标与影响

## 编译class文件
```
javac -d ./java/target ./java/src/myvm/example/HelloWorld.java
```

## 编译VM
```
# setup GOPATH first
make build
```

## 运行class
```
# setup JAVA_HOME first, where to search jre jars
./bin/java -cp ./java/target myvm.example.HelloWorld
```