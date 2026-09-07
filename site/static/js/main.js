// Карманный доктор — интерактив лендинга.
// Заменяет jquery.min.js + index.js живого сайта:
// отправка целей Яндекс.Метрики (счётчик 26736234) с инлайновых onclick.

var METRIKA_ID = 26736234;

function reachGoal(name) {
    if (window.ym) {
        window.ym(METRIKA_ID, "reachGoal", name);
    } else if (window["yaCounter" + METRIKA_ID]) {
        window["yaCounter" + METRIKA_ID].reachGoal(name);
    }
}

function onTargetGoToGooglePlay() {
    reachGoal("GO_TO_STORE");
    reachGoal("GO_TO_GOOGLE_PLAY");
}

function onTargetGoToAppstore() {
    reachGoal("GO_TO_STORE");
    reachGoal("GO_TO_APPSTORE");
}
