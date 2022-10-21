<script lang="ts">
  import FileList from "./FileList.svelte"

  export let address: string = "C:/Users/simon/"

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

  let SubmitAddress;

  //GetThumbnailAsBase64(name).then(res => document.querySelector<HTMLImageElement>("#logo").src = "data:image/png;base64," + res)*/
  
</script>

<main>
  <!--<img alt="Wails logo" id="logo" src="">-->
  <div class="input-box" id="input">
    <input type=checkbox bind:checked={showhiddenFiles} on:change={updateHiddenFiled} id="cb-hiddenfiles"> <label for="cb-hiddenfiles">Show Hidden Files</label>
    <button class="btn" on:click={() => SubmitAddress(true)}>^</button>
    <input autocomplete="off" bind:value={address} class="input" id="name" type="text"/>
    <button class="btn" on:click={() => SubmitAddress(false)}>-></button>
  </div>
  <div class="flex-container">
    <FileList bind:fileViewAddress={address} bind:submitPath={SubmitAddress}/>
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
