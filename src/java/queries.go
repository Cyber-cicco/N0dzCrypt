package java

var Q_JAVA_STRING = `(
    (variable_declarator value : (string_literal)) @declaration
    (#match? @declaration ".*")
)`

var Q_JAVA_HTTP_ANNOTATION = `(
    (annotation name : (identifier) @ann arguments: (annotation_argument_list) ) 
    (#match? @ann ".*Mapping" )
)`

var Q_JAVA_ROUTE = `(
    (field_access object : (identifier) field : (identifier) @test) 
    (#match? @test ".*" )
)`


