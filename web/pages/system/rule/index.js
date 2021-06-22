const Data = window.Data;
const Modal = window.Modal;
import { save, init } from "@shoppehub/fastapi";
import axios from "axios";

init("https://api.chemball.com");

export default (function () {
  var modal = document.getElementById("system_rule_new");

  document.querySelector(".btn-new").addEventListener("click", function (e) {
    const data = Modal.getInstance(modal) || new Modal(modal);
    data.toggle(this);
  });

  // var formModal = Modal.getInstance(modal);
  // modal.addEventListener("show.bs.modal", function (event) {
  //   console.log(222,event);
  //   // Button that triggered the modal
  //   var button = event.relatedTarget;
  //   // Extract info from data-bs-* attributes
  //   var recipient = button.getAttribute("data-bs-whatever");
  //   // If necessary, you could initiate an AJAX request here
  //   // and then do the updating in a callback.
  //   //
  //   // Update the modal's content.
  //   var modalTitle = modal.querySelector(".modal-title");
  //   var modalBodyInput = modal.querySelector(".modal-body input");

  //   modalTitle.textContent = "New message to " + recipient;
  //   modalBodyInput.value = recipient;
  // });
  modal.addEventListener("hide.bs.modal", function (event) {
    console.log(111, event);

    const form = modal.querySelector("form");

    console.log(form);
  });

  // Fetch all the forms we want to apply custom Bootstrap validation styles to

  const submitBtn = modal.querySelector(".form-submit");
  const loadingBtn = modal.querySelector(".form-loading");
  submitBtn.addEventListener("click", async function (event) {
    submitBtn.classList.add("d-none");
    loadingBtn.classList.remove("d-none");

    var forms = modal.querySelector(".needs-validation");

    if (!forms.checkValidity()) {
      event.preventDefault();
      event.stopPropagation();
    }
    forms.classList.add("was-validated");

    // const qe = modal.querySelector("[data-quill]");
    // qe.classList.add("is-invalid");

    const data = {};

    modal.querySelectorAll(".form-control").forEach((item) => {
      data[item.name] = item.value;
    });

    if (!data.desc) {
      modal.querySelector('[name="desc"]').classList.add("is-invalid");
    } else {
      modal.querySelector('[name="desc"]').classList.remove("is-invalid");
    }

    var result;
    try {
      result = await save({
        collection: "system/rule",
        body: data,
      });
    } catch (error) {
      window.alert.warn("网络错误，请待会再试");
    }

    submitBtn.classList.remove("d-none");
    loadingBtn.classList.add("d-none");
    if (result.data.success) {
      window.alert.success("操作成功啦");

      Modal.getInstance(modal).hide();

      const content = await fetch("/system/rule/list").then((res) => res.text());
      document.getElementById("main-list").innerHTML = content
    } else {
      window.alert.warn("服务器出错啦");
      console.log(result.data.errMessage);
    }
  });
})();
