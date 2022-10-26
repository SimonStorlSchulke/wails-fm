<script lang="ts">
  import { GetTreeHTML } from "../wailsjs/go/main/App.js";
  import { writable } from 'svelte/store';
  import { selectedSidebarFolder } from "./AppState";


  function AddHtmlToTree(html: string, element: HTMLElement) {
    let dom = document.createElement("div");
    dom.innerHTML = html;
    let ul: HTMLElement = dom.childNodes[0] as HTMLElement;
    let liv: HTMLElement = dom.querySelector("li") as HTMLElement
    element.parentElement.append(ul);

    ul["expanded"] = false

    let selector: HTMLElement = ul.querySelector(".tree-expander")

    if (!selector) {
      return
    }

    liv.addEventListener("click", (e) => {
      selectedSidebarFolder.set((e.target as HTMLElement).getAttribute("data-path"))
    })

    selector.addEventListener("click", (e) => {
      e.stopPropagation()
      if (ul["expanded"] == true) {
        selector.classList.remove("expanded")
        ul["expanded"] = false;        
        ul.querySelectorAll("ul").forEach((ul)=>ul.remove())
      } 
      
      else {
        selector.classList.add("expanded")
        ul["expanded"] = true;
        let fullPath: string = ul.getAttribute("data-path");

        GetTreeHTML(fullPath).then((items) => {
          items.forEach((li) => {
            AddHtmlToTree(li, e.target as HTMLElement);
          });
        });
      }

    });
  }

  GetTreeHTML("").then((mp) => {
    const ft: HTMLElement = document.querySelector(".filetree-spawner");
    mp.forEach((li) => {
      AddHtmlToTree(li, ft);
    });
  });
</script>

<div class="wrapper" style="float: left;">
  <ul style="margin-left: -20px;">
    <div class="filetree-spawner"/>
  </ul>
</div>

<style>
  .wrapper {
    overflow: scroll;
    max-height: 100%;
    width: calc(100% - 52px);
    max-width: calc(100% - 52px);
  }
</style>
