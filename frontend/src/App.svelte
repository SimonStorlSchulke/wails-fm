<script lang="ts">
  import FileList from "./FileList.svelte"
  import FileInspector from "./FileInspector.svelte"
  import Sidebar from "./Sidebar.svelte"
  import { GetHomeDir } from "../wailsjs/go/main/App.js";
  import { zoomLevel } from "./AppState";


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

  async function GoHome() {
    await GetHomeDir().then((adr) => {
      address = adr
      console.log(address)
    })
    SubmitAddress(false)
  }

  const tree = {
		label: "USA", children: [
			{label: "Florida", children: [
				{label: "Jacksonville"},
				{label: "Orlando", children: [
					{label: "Disney World"},
					{label: "Universal Studio"},
					{label: "Sea World"},
				]},
				{label: "Miami"},
			]},
			{label: "California", children: [
				{label: "San Francisco"},
				{label: "Los Angeles"},
				{label: "Sacramento"},
			]},
		],
	}

  let selectedDir;

</script>

<main>

  <div class="input-box" id="input">
    <button class="btn" on:click={GoHome}>Home</button>
    <input type=checkbox bind:checked={showhiddenFiles} on:change={updateHiddenFiled} id="cb-hiddenfiles"> <label for="cb-hiddenfiles">Show Hidden</label>
    <button class="btn" on:click={() => SubmitAddress(true)}>^</button>
    <input autocomplete="off" bind:value={address} class="input" id="name" type="text"/>
    <button class="btn" on:click={() => SubmitAddress(false)}>></button>

    Zoom <input type="range" min="1" max="20" bind:value={$zoomLevel} class="slider" id="myRange">

  </div>

  <div class="flex-container">
    <Sidebar {tree}/>
    <FileList bind:fileViewAddress={address} bind:submitPath={SubmitAddress} bind:showhiddenFiles={showhiddenFiles}/>
    <FileInspector/>
  </div>

</main>

<style>

  .input-box {
    background: rgba(0, 0, 0, 0.543);
    text-align: left;
    padding: 0px;
    height: 36px;
  }

  .input-box .btn {
    min-width: 30px;
    height: 100%;
    margin: 0px;
    cursor: pointer;
    font-weight: bold;
    background-image: linear-gradient(rgb(52, 69, 105),rgb(33, 45, 69));
    color: white;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(rgb(70, 95, 137),rgb(36, 52, 85));
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgb(41, 54, 82);
    -webkit-font-smoothing: antialiased;
    width: 700px;
    color: white;
  }

  .input-box .input:hover {
    border: none;
  }

  .input-box .input:focus {
    border: none;
    background-color: rgb(44, 60, 95);
  }

</style>
