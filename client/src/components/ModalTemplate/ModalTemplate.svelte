<script lang="ts">
    export let show:         Boolean
    export let onHide:       () => {}
    let modalRef: any = null

    const closeModal = (e: Event) => {
        if(e.target == modalRef){
            onHide()
        }
    }
</script>

<div class={`modal ${show ? "slidein" : "slideout"}`} bind:this={modalRef} on:click={closeModal} tabindex="0" role="button" on:keyup={null}>
    
</div>

<style lang="scss">
    .slidein{
        left: 0;
    }

    .slideout{
        left: 100vw;
    }

    .modal{
        display: grid;
        position: fixed;
        top: 0;
        z-index: 2;
        width: 100%;
        height: 100%;

        background-color: rgba($color: #000000, $alpha: 0.4);
        transition: 0.3s;
    }
</style>