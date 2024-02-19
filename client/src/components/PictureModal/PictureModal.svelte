<script lang="ts">
    export let show:         Boolean
    export let onHide:       () => {}
    export let modalHeader:  string
    export let modalContent: any
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

</style>