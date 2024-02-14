<script lang="ts">
    import { onMount } from "svelte";
    import { Image } from "svelte-bootstrap-icons";
    
    let imageURL:     any 
    let showImageUrl: boolean = false

    const onFileSelected = (e: any)=>{
        const file = e.target.files[0];
        const reader = new FileReader();

        reader.onload = () => {
            imageURL = reader.result;
        };

        reader.readAsDataURL(file);
    }

    onMount(() => {
        imageURL = null
    })
</script>

<div>
    <div class="icon-wrapper image-select">
        <label for="file-input">
            <Image width={25} height={25} />
            <input id="file-input" type="file" accept="image/x-png,image/gif,image/jpeg"  on:change={(e)=>onFileSelected(e)}  />
        </label>
    </div>        
    <div class="image-wrapper">
        <img src={imageURL} alt="" width="200px" height="auto">
    </div>
    <button on:click={() => showImageUrl = !showImageUrl}>Send the Picture</button>
    {#if imageURL}
        <h1>image</h1>
    {:else}
        <h1>no image</h1>
    {/if}

    {#if showImageUrl}
        <h2>{imageURL}</h2>
    {/if}
    <span class="close" on:click={() => imageURL = null} role="button" tabindex="0" on:keyup={null}>
        &times;
    </span>
</div>

<style lang="scss">
    .icon-wrapper{
        input{
            display: none;
        }
    }
</style>