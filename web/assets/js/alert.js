//
// autosize.js
// Dashkit module
//

function alertFn(key, msg) {
  const root = document.querySelector(key);
  root.classList.remove("d-none");
  document.querySelector(key + " .alert span").innerHTML = msg;
  setTimeout(() => {
    root.classList.add("d-none");
  }, 3000);
}

const alert = {
  success: (msg) => {
    alertFn("#alert-success", msg);
  },
  warn: () => {
    alertFn("#alert-warn", msg);
  },
};

// Make available globally
window.alert = alert;
