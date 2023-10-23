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
    <div class="modal_body">
        <div class="modal_header">
            <div class="header">{ modalHeader }</div>
            <span class="close" on:click={onHide} role="button" tabindex="0" on:keyup={null}>&times;</span>
        </div>
        <div class="modal_content">
            <svelte:component this={modalContent} {onHide}/>
        </div>
        <div class="modal_footer" >
            <button on:click={onHide}>Close</button>
        </div>
    </div>
</div>

<style lang="scss">
    $borderColor: rgb(200, 200, 200);
    $bright_blue: rgb(0, 150, 255);
    $border_radius: 5px;

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

        .modal_body{
            position: relative;
            margin: auto;
            width: 95vw;
            border-radius: 10px;

            .modal_header{
                display: flex;
                // grid-template-columns: auto auto;
                justify-content: space-between;
                align-items: center;
                // border: 2px red solid;
                border-bottom: 1px solid $borderColor;
                background-color: white;
                padding: 10px 15px;
                border-top-left-radius: $border_radius;
                border-top-right-radius: $border_radius;

                .header{
                    // border: 2px solid red;
                    justify-self: right;
                    font-size: 18px;

                    @media screen and (min-width: 768px) {
                        font-size: 25px;
                    }
                }

                .close{
                    // border: 2px solid saddlebrown;
                    justify-self: right;
                    transition: 0.5s;
                    font-size: 40px;
                }

                .close:hover{
                    cursor: pointer;
                    color: gray;
                }
            }

            .modal_content{
                background-color: white;
                padding: 10px;
                border-bottom: 1px solid $borderColor;
                // border: 2px solid yellow;
            }

            .modal_footer{
                display: flex;
                justify-content: center;
                padding: 15px 10px;
                background-color: white;
                border-bottom-left-radius: $border_radius;
                border-bottom-right-radius: $border_radius;

                button{
                    background-color: $bright_blue;
                    border: none;
                    border-radius: 5px;
                    padding: 10px 25px;
                    color: white;
                    font-size: 15px;
                    transition: 0.5s;
                }

                button:hover{
                    cursor: pointer;
                    background-color: grayscale($color: $bright_blue);
                }
            }
        }

        @media screen and (min-width: 768px) {
            .modal_body{
                width: 80vw;
            }
        }

        @media screen and (min-width: 992px) {
            .modal_body{
                width: 50vw;
            }
        }
    }
</style>