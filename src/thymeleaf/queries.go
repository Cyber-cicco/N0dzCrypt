package thymeleaf

const Q_TH_ATTRIBUTE = `(
(attribute_name) @tag
(#match? @tag "th:*")
)`
const Q_TH_REPLACE = `(
(attribute_name) @tag
(#eq? @tag "th:replace")
)`
const Q_TH_REPLACE_INSERT = `(
(attribute_name) @var
(#any-of @var
    "th:insert"
    "th:replace"
)
)`
const Q_TH_INSERT = `(
(attribute_name) @tag
(#eq? @tag "th:insert")
)`

const Q_ATTRIBUTE_VALUE = `
(quoted_attribute_value) @tag
`
const Q_ATTRIBUTE_HX = `(
(attribute_name) @tag
)`

var pageTags = []string{
    "hx-get",
    "hx-post",
    "hx-put",
    "hx-patch",
    "hx-delete",
    "hx-push-url",
    "href",
    "th:hx-get",
    "th:hx-post",
    "th:hx-put",
    "th:hx-patch",
    "th:hx-delete",
    "th:href",
    "th:hx-push-url",
}
