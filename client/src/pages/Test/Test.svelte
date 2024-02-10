<script lang="ts">
    import { onMount } from "svelte";
    import { Image } from "svelte-bootstrap-icons";

    let imageURL:    any 

    const onFileSelected = (e: any)=>{
        let image = e.target.files[0];
        let reader = new FileReader();
        reader.readAsDataURL(image);
        reader.onload = e => {
            if (e.target) {
                imageURL = e.target.result
            }
        };
    }


    onMount(() => {
        imageURL = null
    })


</script>

<main>
    <div class="icon-wrapper image-select">
        <label for="file-input">
            <Image width={25} height={25} />
        </label>
        <input id="file-input" type="file" accept="image/x-png,image/gif,image/jpeg"  on:change={(e)=>onFileSelected(e)} bind:this={imageURL} />
    </div>        
    <div class="image-wrapper">
        <img src={imageURL} alt="" width="200px" height="auto">
    </div>
    {#if imageURL}
        <h1>image</h1>
    {:else}
        <h1>no image</h1>
    {/if}
    <span class="close" on:click={() => imageURL = null} role="button" tabindex="0" on:keyup={null}>
        &times;
    </span>
</main>

<style lang="scss">
    input{
        display: none;
    }
</style>