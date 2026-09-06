// Карманный доктор — интерактив лендинга.
// Заменяет jquery.min.js + index.js живого сайта: аккордеоны блока оплаты
// и отправка целей Яндекс.Метрики (счётчик 26736234) с инлайновых onclick.

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

function onTargetReadPaymentGoogle() {
    reachGoal("READ_PAYMENT");
    reachGoal("READ_PAYMENT_GOOGLE_PLAY");
}

function onTargetReadPaymentApple() {
    reachGoal("READ_PAYMENT");
    reachGoal("READ_PAYMENT_APPSTORE");
}

// Аккордеоны: клик открывает пункт, повторный — закрывает, соседние закрываются.
document.addEventListener("DOMContentLoaded", function () {
    var accordions = document.querySelectorAll(".accord");
    Array.prototype.forEach.call(accordions, function (accord) {
        var items = accord.querySelectorAll(".accord__item");
        Array.prototype.forEach.call(items, function (item) {
            item.addEventListener("click", function () {
                var wasOpen = item.classList.contains("accord__item--show");
                Array.prototype.forEach.call(items, function (i) {
                    i.classList.remove("accord__item--show");
                });
                if (!wasOpen) {
                    item.classList.add("accord__item--show");
                }
            });
        });
    });
});
