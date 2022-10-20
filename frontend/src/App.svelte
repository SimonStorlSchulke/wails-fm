<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {GetFolderFilepaths, GreetJohn, GetLocalFile, GetThumbnailAsBase64} from '../wailsjs/go/main/App.js'
  import {WindowFullscreen} from '../wailsjs/runtime/runtime.js'

  let resultText: string = "Please enter your name below 👇"
  let name: string = "C:/Users/simon/"

  var stringToHTML = function (str) {
    var parser = new DOMParser();
    var doc = parser.parseFromString(str, 'text/html');
    return doc.body;
  };

  function sanitizePath(path: string): string {
    path = path.replace(/([^:]\/)\/+/g, "$1");
    path = path.replaceAll("\\", "/");
    if (!path.endsWith("/")) {
      path += "/"
    }
    return path;
  }

  function intoFolder() {
    console.log("up")
  }

  function submitPath(moveUp: boolean): void {
    document.querySelector(".file-area").innerHTML = "";
    name = sanitizePath(name)
    if (moveUp) {
      let els: string[] = name.split("/")
      let num: number = els[els.length - 2].length + 1
      console.log(num)
      name = name.slice(0, name.length - num)
    }
    GetFolderFilepaths(name).then(fi => {
      document.querySelector(".file-area").append(stringToHTML(fi))
      document.querySelectorAll(".dir").forEach(dir => {
        dir.addEventListener("click", function() {
          let o: string = dir.querySelector("p").innerText
          name = name + "/" + o
          submitPath(false)
        })
      })
    })

    

    //GetThumbnailAsBase64(name).then(res => document.querySelector<HTMLImageElement>("#logo").src = "data:image/png;base64," + res)
  }


  resultText = "Testing"

</script>

<main>
  <!--<img alt="Wails logo" id="logo" src="">-->
  <div class="input-box" id="input">
    <button class="btn" on:click={() => submitPath(true)}>^</button>
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={() => submitPath(false)}>-></button>
  </div>
  <div class="flex-container">
    <div class="file-area"></div>
    <div class="inspector">Inspector</div>
  </div>
</main>

<style>

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
    width: 600px;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
