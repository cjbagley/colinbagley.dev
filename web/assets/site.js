const AUTO = "auto";
const DARK = "dark";
const LIGHT = "light";

const DARK_THEME = "dark-theme";
const LIGHT_THEME = "light-theme";

const themeToggle = document.getElementById('theme-toggle');
const lightInput = document.getElementById(LIGHT_THEME);
const darkInput = document.getElementById(DARK_THEME);
const currentTheme = localStorage.getItem('theme') || AUTO;
const prefersDarkScheme = window.matchMedia("(prefers-color-scheme: dark)");

if (currentTheme == DARK || (currentTheme == AUTO && prefersDarkScheme.matches)) {
    setDarkTheme();
} else {
    setLightTheme();
}

themeToggle.addEventListener("change", function (el) {
    let theme = el.target.value;
    switch (theme) {
        case LIGHT:
            setLightTheme();
            break;
        case DARK:
            setDarkTheme();
            break;
        default:
            document.body.classList.toggle(DARK_THEME);
            document.body.classList.toggle(LIGHT_THEME);
    }
    localStorage.setItem("theme", theme);
});

function setLightTheme() {
    document.body.classList.remove(DARK_THEME);
    document.body.classList.add(LIGHT_THEME);
    darkInput.removeAttribute("checked");
    lightInput.setAttribute("checked", "checked");
}

function setDarkTheme() {
    document.body.classList.add(DARK_THEME);
    document.body.classList.remove(LIGHT_THEME);
    lightInput.removeAttribute("checked");
    darkInput.setAttribute("checked", "checked");
}