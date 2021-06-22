//
// quill.js
// Theme module
//

import Quill from "quill";

import Data from "bootstrap/js/src/dom/data";

window.Data = Data;

const toggles = document.querySelectorAll("[data-quill]");

toggles.forEach((toggle) => {
  const elementOptions = toggle.dataset.quill
    ? JSON.parse(toggle.dataset.quill)
    : {};

  const defaultOptions = {
    modules: {
      toolbar: [
        ["bold", "italic"],
        ["link", "blockquote", "code", "image"],
        [
          {
            list: "ordered",
          },
          {
            list: "bullet",
          },
        ],
      ],
    },
    theme: "snow",
  };

  const options = {
    ...elementOptions,
    ...defaultOptions,
  };
  const q = new Quill(toggle, options);
  if (toggle.dataset.for) {
    const input = toggle.parentElement.querySelector("[name="+toggle.dataset.for+"]");
    q.on("text-change", () => {
      input.value = q.root.innerHTML;
    });
  }

  Data.set(toggle, "quill", q);
});

// Make available globally
window.Quill = Quill;
