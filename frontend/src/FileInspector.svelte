<script lang="ts">
    import { onMount } from "svelte";
    import { selectedFilePaths } from "./AppState";
    import { GetFileDetailsSingle, SetFileMeta, GetFileMeta } from "../wailsjs/go/main/App.js";

    let fileDetails;
    let selFilesNum = 1;

    selectedFilePaths.subscribe((selFiles) => {
        if (selFiles) {
            selFilesNum = selFiles.length;
        }
        if (!selFiles) {return}
        if (selFilesNum == 1) {
            GetFileDetailsSingle(selFiles[0]).then((details) => {
                fileDetails = details
            })
        } else if (selFilesNum > 1) {
            fileDetails = "multiple Selected";
        } else {
            fileDetails = "Nothing Selected";
        }
    });
    let metaComment;
    let metaColor;


    function UpdateMeta(dom) {
        GetFileMeta(fileDetails.QID).then( meta => {
                metaComment = meta.Comment;
                metaColor = meta.Color;
            })
    }

</script>

<div class="inspector">
    {#key fileDetails}
        {#if fileDetails}
            {#if fileDetails == "multiple Selected"}
                <p> {selFilesNum} Files selected</p>
            {:else if  fileDetails == "Nothing Selected"}
                <p>Nothing selected</p>
            {:else}
                <h3>{fileDetails.Name}</h3>
                <p>{fileDetails.MimeType}</p>
                <p>{fileDetails.CreationTime}</p>
                <p>Modified: {fileDetails.ModifiedTime}</p>
                <table use:UpdateMeta>
                    <tr>
                      <td>
                          <label for="filecolor">File Color</label>
                        </td>
                        <td>
                            <input bind:value={metaColor} type="color" id="filecolor" name="filecolor"><br>
                        </td>
                    </tr>
                    <tr>
                        <td>
                          <label for="filecomment">Comment</label>
                      </td>
                      <td>
                          <input bind:value={metaComment}  type="text" id="filecomment" name="filecomment" placeholder="enter file comment"><br>
                      </td>
                    </tr>
                  </table>
                
                <button on:click={() => {
                    SetFileMeta(fileDetails.QID, {
                        Tags: [],
                        Comment: metaComment,
                        Color: metaColor
                    })
                    }}>Save</button>
            {/if}
        {/if}
    {/key}
</div>

<style>
    .inspector {
        padding: 10px;
        min-width: 300px;
        max-width: 300px;
    }


    h3 {
        word-wrap: break-word;
    }
</style>
