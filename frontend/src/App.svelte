<script lang="ts">
  import FileVIew from "./FileVIew.svelte"
  import logo from './assets/images/logo-universal.png'
  import {GetFolderAPI, GetLocalFile, GetThumbnailAsBase64, OpenWithDefaultApp} from '../wailsjs/go/main/App.js'
  import {WindowFullscreen} from '../wailsjs/runtime/runtime.js'


  let resultText: string = "Please enter your name below 👇"
  let name: string = "C:/Users/simon/"

  var stringToHTML = function (str) {
    var parser = new DOMParser();
    var doc = parser.parseFromString(str, 'text/html');
    return doc.body;
  };

  let showhiddenFiles: boolean = false;
  function updateHiddenFiled() {
    let hiddenFiles = document.querySelectorAll(".file-hidden")
    if (!showhiddenFiles) {
      hiddenFiles.forEach(filecard => {
        (filecard as HTMLElement).classList.add("hidden");
      });
    } else {
      hiddenFiles.forEach(filecard => {
        (filecard as HTMLElement).classList.remove("hidden");
      });
    }
  }


  function sanitizePath(path: string): string {
    path = path.replace(/([^:]\/)\/+/g, "$1");
    path = path.replaceAll("\\", "/");
    if (!path.endsWith("/")) {
      path += "/"
    }
    return path;
  }

  function submitPathOld(moveUp: boolean): void {
    //document.querySelector(".file-area").innerHTML = "";
    name = sanitizePath(name)
    if (moveUp) {
      let els: string[] = name.split("/")
      let num: number = els[els.length - 2].length + 1
      name = name.slice(0, name.length - num)
    }
    /*GetFolderFilepaths(name).then(fi => {
      //document.querySelector(".file-area").append(stringToHTML(fi))
      document.querySelectorAll(".dir").forEach(dir => {
        dir.addEventListener("click", function() {
          let o: string = dir.querySelector("p").innerText
          name += o
          submitPathOld(false)
        })
      })
      updateHiddenFiled()
    })*/

    //GetThumbnailAsBase64(name).then(res => document.querySelector<HTMLImageElement>("#logo").src = "data:image/png;base64," + res)
  }
  let folderData;
  function submitPath(moveUp: boolean): void {
    name = sanitizePath(name)
    if (moveUp) {
      let els: string[] = name.split("/")
      let num: number = els[els.length - 2].length + 1
      name = name.slice(0, name.length - num)
    }
    GetFolderAPI(name).then(fi => {
      folderData = fi;
      console.log(fi.Files)
    })
  }

  function folderDoubleclicked(foldername) {
    name = sanitizePath(name)
    name += foldername.detail.foldername
    submitPath(false)
    }

  window.onload = function() {submitPath(false)}
</script>

<main>
  <!--<img alt="Wails logo" id="logo" src="">-->
  <div class="input-box" id="input">
    <input type=checkbox bind:checked={showhiddenFiles} on:change={updateHiddenFiled} id="cb-hiddenfiles"> <label for="cb-hiddenfiles">Show Hidden Files</label>
    <button class="btn" on:click={() => submitPath(true)}>^</button>
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
    <button class="btn" on:click={() => submitPath(false)}>-></button>
  </div>
  <div class="flex-container">
    <div class="file-area">
      {#key folderData}
      <section>
        {#if folderData}
          {#if folderData.Files}
            {#each folderData.Files as fileinfo }
            <FileVIew on:message={folderDoubleclicked} fileinfo={fileinfo} isDir={true}/>
          {/each}
          {/if}
        {/if}
      </section>

      <section>
        {#if folderData}
        {#if folderData.Folders}
          {#each folderData.Folders as folderinfo }
          <FileVIew fileinfo={folderinfo} isDir={false}/>
        {/each}
        {/if}
        {/if}
      </section>
      {/key}
    </div>
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
