//
// navbar.js
// Theme module
//

import { Collapse } from "bootstrap";

const collapses = document.querySelectorAll(".navbar-nav .collapse");

collapses.forEach((collapse) => {
  // Init collapses
  const collapseInstance = new Collapse(collapse, {
    toggle: false,
  });

  // Hide sibling collapses on `show.bs.collapse`
  collapse.addEventListener("show.bs.collapse", (e) => {
    e.stopPropagation();

    const closestCollapse = collapse.parentElement.closest(".collapse");
    const siblingCollapses = closestCollapse.querySelectorAll(".collapse");

    siblingCollapses.forEach((siblingCollapse) => {
      const siblingCollapseInstance = Collapse.getInstance(siblingCollapse);

      if (siblingCollapseInstance === collapseInstance) {
        return;
      }

      siblingCollapseInstance.hide();
    });
  });

  // Hide nested collapses on `hide.bs.collapse`
  collapse.addEventListener("hide.bs.collapse", (e) => {
    e.stopPropagation();
    const childCollapses = collapse.querySelectorAll(".collapse");
    childCollapses.forEach((childCollapse) => {
      const childCollapseInstance = Collapse.getInstance(childCollapse);
      childCollapseInstance.hide();
    });
  });

  if (collapse.querySelector(".nav-open")) {
    setTimeout(() => {
      // 把最顶层的先打开
      collapseInstance.show();
      // 查找子
      const childCollapses = collapse.querySelectorAll(".collapse");
      childCollapses &&
        childCollapses.length > 0 &&
        childCollapses.forEach((childCollapse) => {
          const childCollapseInstance = Collapse.getInstance(childCollapse);
          childCollapseInstance.show();
        });
    }, 200);
  }
});
