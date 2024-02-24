<script lang="ts">
    export let show:         Boolean
    export let onHide:       () => {}
    export let imageURL:     string
    let modalRef: any = null

    const closeModal = (e: Event) => {
        if(e.target == modalRef){
            onHide()
        }
    }
</script>

<div class={`modal ${show ? "slidein" : "slideout"}`} bind:this={modalRef} on:click={closeModal} tabindex="0" role="button" on:keyup={null}>
    <div class="image-wrapper">
        <span class="close" on:click={onHide} role="button" tabindex="0" on:keyup={null}>&times;</span>
        <img src={imageURL} alt="">
    </div>
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

        .image-wrapper{
            margin: auto;

            span{
                position: absolute;
                font-size: 25px;
                padding: 0px 10px;
                transition: 0.3s;
                color: rgba(0, 0, 0, 0.4);
                // border: 2px solid red;

                @media screen and (min-width: 992px) {
                    font-size: 50px;  
                    padding: 0px 20px; 
                }

                &:hover{
                    cursor: pointer;
                    color: white;
                }
            }

            img{
                width: 90vw;
                height: auto;

                @media screen and (min-width: 992px) {
                    height: 80vh;
                    width: auto;
                }
            }
        }
    }

</style>