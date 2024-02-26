package thymeleaf

import "testing"

func TestMvReplacesAndInsert(t *testing.T) {

    input1 := `
    <div class="bg-white" th:insert="~{page/snail/snail :: snail('3003')}"></div>
    `
    input2 := `
    <div class="bg-white" th:insert=""></div>
    `
    input3 := `
    <div class="bg-white" th:href="page/rat/rat"></div>
    `
    expect1 := `
    <div class="bg-white" th:insert="~{page/snail/rat :: snail('3003')}"></div>
    `
    expect2 := `
    <div class="bg-white" th:insert=""></div>
    `
    expect3 := `
    <div class="bg-white" th:href="page/rat/rat"></div>
    `
    oldName := "page/snail/snail"
    newName := "page/snail/rat"
    output1 := mvReplacesAndInserts(input1, oldName, newName)
    output2 := mvReplacesAndInserts(input2, oldName, newName)
    output3 := mvReplacesAndInserts(input3, oldName, newName)

    if output1 != expect1 {
        t.Fatalf("erreur pour le test 1 : test. Wanted %q, got %q", expect1, output1)
    }
    if output2 != expect2 {
        t.Fatalf("erreur pour le test 2 : test. Wanted %q, got %q", expect2, output2)
    }
    if output3 != expect3 {
        t.Fatalf("erreur pour le test 3 : test. Wanted %q, got %q", expect3, output3)
    }

}

func TestMvReferencesOldName(t *testing.T) {
    input1 := `
<!--/*@thymesVar id="cal" type="fr.diginamic.digilearning.dto.CalendarInfos"*/-->
<!--/*@thymesVar id="week" type="java.util.List<fr.diginamic.digilearning.dto.DayInfos>"*/-->
<!--/*@thymesVar id="dateValue" type="java.lang.String"*/-->
<!--/*@thymesVar id="prev" type="java.lang.String"*/-->
<!--/*@thymesVar id="next" type="java.lang.String"*/-->
<!--/*@thymesVar id="hours" type="java.util.List<fr.diginamic.digilearning.dto.HourInfos>"*/-->
<!--/*@thymesVar id="hourMap" type="java.util.Map<java.time.LocalDateTime, fr.diginamic.digilearning.dto.CoursDto>"*/-->
<!--/*@thymesVar id="adminHourMap" type="java.util.Map<java.time.LocalDateTime, fr.diginamic.digilearning.dto.CoursAdminDto>"*/-->
<!--/*@thymesVar id="dateUtil" type="fr.diginamic.digilearning.utils.DateUtil"*/-->
<!--/*@thymesVar id="coursPrevus" type="java.util.List<fr.diginamic.digilearning.dto.CoursDto>"*/-->
<div id="calendar" class="flex flex-col overflow-x-auto bg-white gap-2 w-full h-full max-w-full">
    <div class="flex flex-col w-full h-full justify-center">
        <div class="flex flex-col-reverse md:flex-row items-center w-full">
            <div class="block text-primary w-full md:w-1/4 text-center bg-grey m-1 p-2 rounded-md lg:hidden">Veuillez vous connecter sur la version PC pour organiser votre emploi du temps</div>
            <div class="flex flex-col md:flex-row gap-1">
                <form class="flex flex-col md:flex-row p-2"  onsubmit="return false;">
                    <div class="flex flex-col md:flex-row gap-2 w-full items-center">
                        <label for="semaine">Semaine :</label>
                        <input
                                name="semaine"
                                id="semaine"
                                th:value="${dateValue}"
                                class="p-2 text-black w-[200px] border-primary border-2 rounded-md"
                                type="date"
                                hx-on::after-request="reInit()"
                                hx-get="/agenda/date"
                                hx-push-url="true"
                                hx-select="#calendar"
                                hx-target="#calendar"
                                hx-vals="*"
                        >
                    </div>
                </form>
                <div class="flex w-full md:w-fit items-center justify-center py-2">
                    <img
                            src="/img/icons/simple-arrow-primary.svg"
                            class="rotate-90 hover:cursor-pointer px-3 py-1 rounded-md hover:bg-lightAccent"
                            th:hx-get="'/agenda/date?semaine=' + ${prev}"
                            hx-push-url="true"
                            hx-on::before-request="reInit()"
                            hx-select="#calendar"
                            hx-target="#calendar"
                    >
                    <img
                            src="/img/icons/simple-arrow-primary.svg"
                            class="-rotate-90 hover:cursor-pointer px-3 py-1 rounded-md hover:bg-lightAccent"
                            th:hx-get="'/agenda/date?semaine=' + ${next}"
                            hx-on::before-request="reInit()"
                            hx-push-url="true"
                            hx-select="#calendar"
                            hx-target="#calendar"
                    >
                </div>
            </div>
            <div class="grow"></div>
            <div th:text="${cal.month()}"></div>
        </div>
    </div>
    <table class="flex-col w-full min-w-[1000px]">
        <thead class="flex w-full gap-1">
        <tr class="flex w-full">
            <th class="px-5 py-2  bg-lightAccent"></th>
            <th
                    class="px-5 py-2 w-1/5 min-w-[200px] bg-lightAccent text-center"
                    th:each="d, iterStat : ${week}"
                    th:text="${d.weekDay()}  + ' ' + ${d.dateNumber()}"
            >
            </th>
        </tr>
        </thead>
        <tbody class="flex flex-col py-2 h-[85vh] w-full">
        <tr class="flex w-full h-[75px]" th:each="h, hiterStat : ${hours}">
            <td class="px-5 h-[50px] py-2" th:text="${h.hour()} + 'h'"></td>
            <td
                    th:id="${dateUtil.getId(w.dateJour(), h.time()).toString()}"
                    th:class="(${iterStat.index} % 2 == 1) ? 'px-5 border-t-grey border-t-4 h-[50px] min-w-[200px] max-w-1/5 w-1/5 bg-white' : 'px-5 min-w-[200px] max-w-1/5 border-t-white border-t-4 h-[50px] w-1/5 bg-grey'"
                    th:data-draggable2="${dateUtil.getLdt(w.dateJour(), h.time()).toString()}"
                    th:each="w, iterStat: ${week}" >
                    <th:block th:if="${hourMap.containsKey(dateUtil.getLdt(w.dateJour(), h.time())) && !_user.isAdministrateur()}">
                        <th:block
                            th:replace="~{pages/agenda/fragments/agenda.calendar.cours :: cours(
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time())).getId()},
                                ${dateUtil.getLdt(w.dateJour(), h.time()).toString()},
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time())).getDureeEstimee()},
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time()))}
                            )}"></th:block>
                    </th:block>
                    <th:block th:if="${hourMap.containsKey(dateUtil.getLdt(w.dateJour(), h.time())) && _user.isAdministrateur()}">
                        <th:block
                            th:replace="~{pages/agenda/fragments/cours.admin :: cours(
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time())).getId()},
                                ${dateUtil.getLdt(w.dateJour(), h.time()).toString()},
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time())).getDureeEstimee()},
                                ${hourMap.get(dateUtil.getLdt(w.dateJour(), h.time()))}
                            )}"></th:block>
                    </th:block>
                </td>
            </tr>
        </tbody>
    </table>
    <script th:insert="scripts/calendar.js">
    </script>
    <script>
        reInit()
    </script>
</div>
    `
    oldname := "pages/agenda"
    newname := "pages/calendar"
    mvReferencesOfOldName(oldname, newname, input1)
}
