package java

var Q_JAVA_STRING = `(
    (variable_declarator value : (string_literal)) @declaration
    (#match? @declaration ".*")
)`

var Q_JAVA_ROUTE = `(
    (field_access object : (identifier) field : (identifier) @test) 
    (#match? @test ".*" )
)`


