package thymeleaf

var Q_TH_ATTRIBUTE = `(
(attribute_name) @tag
(#match? @tag "th:*")
)`
var Q_TH_REPLACE = `(
(attribute_name) @tag
(#eq? @tag "th:replace")
)`
var Q_TH_REPLACE_INSERT = `(
(attribute_name) @tag
(#any-of @tag
    "th:insert"
    "th:replace"
)
)`
var Q_TH_INSERT = `(
(attribute_name) @tag
(#eq? @tag "th:insert")
)`

var Q_ATTRIBUTE_VALUE = `
(quoted_attribute_value) @tag
`
