let modal = document.getElementById("settingsModal");
let modalOverlay = document.getElementById("modal-overlay");
let openButton = document.getElementById("settings-button");
let closeButton = document.getElementById("close-button");

const initCheck = async () => {
  try {
    let isSet = await eel.is_user_set()();
    if (isSet) {
      setScreen();
    } else {
      modal.classList.toggle("closed");
      modalOverlay.classList.toggle("closed");
    }
    console.log(isSet);
    return isSet;
  } catch (e) {
    console.error(e);
  }
};

(async () => {
  try {
    let details = await eel.read_settings()();
    document.getElementById("name").value = details.Name;
    document.getElementById("textField").innerText = `Hello, ${details.Name}`;
    document.getElementById("dob").value = details.DOB;
    //document.getElementById("colors").value = details.Colors;
  } catch (e) {
    console.error(e);
  }
})();

initCheck();
let dob;

openButton.addEventListener("click", function () {
  modal.classList.toggle("closed");
  modalOverlay.classList.toggle("closed");
});

closeButton.addEventListener("click", function () {
  console.log("pressed");
  modal.classList.toggle("closed");
  modalOverlay.classList.toggle("closed");
});

const form = document.getElementById("myForm");
form.addEventListener("submit", async (event) => {
  event.preventDefault();
  modal.classList.toggle("closed");
  modalOverlay.classList.toggle("closed");
  await eel.update_dob(
    form["name"].value,
    form["dob"].value,
    ""
  )();

  setScreen();
});
const setScreen = async () => {
  try {
    let weeks = await eel.get_weeks_done()();
    const squares = document.querySelector(".squares");
    squares.innerHTML = "";
    for (var i = 0; i < weeks; i++) {
      squares.insertAdjacentHTML("beforeend", `<li data-level="done"></li>`);
    }
    for (var i = 0; i < 4680 - weeks; i++) {
      squares.insertAdjacentHTML(
        "beforeend",
        `<li data-level="incomplete"></li>`
      );
    }
  } catch (e) {
    console.error(e);
  }
};
const weeks = document.querySelector(".week");
for (var i = 1; i <= 52; i++) {
  weeks.insertAdjacentHTML("beforeend", `<li>${i}</li>`);
}
const ages = document.querySelector(".age");
for (var i = 0; i < 90; i++) {
  ages.insertAdjacentHTML("beforeend", `<li>${i}</li>`);
}

(function (d, s, id) {
  var js,
    sjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) return;
  js = d.createElement(s);
  js.id = id;
  js.src = "https://sdk.snapkit.com/js/v1/create.js";
  sjs.parentNode.insertBefore(js, sjs);
})(document, "script", "snapkit-creative-kit-sdk");
