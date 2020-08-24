Sample keyword query with many terms that takes 45+ seconds in the Golang implementation.
This same query returns instantly in the Java implementation.

Golang code generated with:
```
java -Xmx500M -cp antlr-4.8-complete.jar org.antlr.v4.Tool -o v4parser -package v4parser -visitor -listener -Dlanguage=Go VirgoQueryLexer.g4 VirgoQuery.g4
```

CPU/memory profiles generated with:
```
go test -run ^$ -bench . -cpuprofile cpu.prof -memprofile mem.prof ./v4parser/minimal_test.go

```
