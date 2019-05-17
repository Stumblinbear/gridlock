<template>
    <div class="scroll-panel-content">
        <slot />
    </div>
</template>

<style>
    .scroll-panel-content {
        width: 100%;
        padding: 1em;

        background-color: rgba(0, 0, 0, .2);

        white-space: nowrap;
        overflow-x: scroll;
        scrollbar-width: none; /* Firefox */
        -ms-overflow-style: none;  /* IE 10+ */
    }

    .scroll-panel-content.gradient {
        background: linear-gradient(to right,
                        #191919 0%,
                        #19191955 5%,
                        #19191955 95%,
                        #191919 100%);
        border-left: 5px solid #191919;
        border-right: 5px solid #191919;
    }

    .scroll-panel-content::-webkit-scrollbar { /* WebKit */
        width: 0;
        height: 0;
    }
</style>