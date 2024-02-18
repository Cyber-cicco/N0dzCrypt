package thymeleaf

import "testing"

func TestMvReplacesAndInsert(t *testing.T) {

    input1 := `
    <div class="bg-white" th:insert="~{page/snail/snail :: snail('3003')}"></div>
    `
    input2 := `
    <div class="bg-white" th:insert=""></div>
    `
    expect1 := `
    <div class="bg-white" th:insert="~{page/snail/rat :: snail('3003')}"></div>
    `
    expect2 := `
    <div class="bg-white" th:insert=""></div>
    `
    oldName := "page/snail/snail"
    newName := "page/snail/rat"
    output1 := mvReplacesAndInserts(input1, oldName, newName)
    output2 := mvReplacesAndInserts(input2, oldName, newName)

    if output1 != expect1 {
        t.Fatalf("erreur pour le test 1 : test. Wanted %q, got %q", expect1, output1)
    }
    if output2 != expect2 {
        t.Fatalf("erreur pour le test 1 : test. Wanted %q, got %q", expect2, output2)
    }

}
