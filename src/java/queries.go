package java

var Q_JAVA_FINAL = `(
    (variable_declarator value : (string_literal)) @declaration
    (#match? @declaration ".*")
)`

