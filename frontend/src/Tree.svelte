<script lang="ts">
  import { GetTreeHTML } from "../wailsjs/go/main/App.js";

  let selectedDir;

  function AddHtmlToTree(html: string, element: HTMLElement) {
    let dom = document.createElement("div");
    dom.innerHTML = html;
    let ul: HTMLElement = dom.childNodes[0] as HTMLElement;
    element.parentElement.append(ul);

    ul["expanded"] = false

    let selector: HTMLElement = ul.querySelector("span")

    if (!selector) {
      return
    }

    ul.addEventListener("click", (e) => {
      e.stopPropagation()
      selectedDir = ul.getAttribute("data-path")
    })

    selector.addEventListener("click", (e) => {
      
      e.stopPropagation()
      let el = e.target as HTMLElement

      if (ul["expanded"] == true) {
        ul["expanded"] = false;

        ul.setAttribute("expanded", "false");
        
        ul.querySelectorAll("ul").forEach((ul)=>ul.remove())
        
      } 
      
      else {
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
  }
</style>
