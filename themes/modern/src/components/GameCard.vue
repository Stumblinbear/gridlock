<template>
    <screen-modal
        :bg="{ type: 'image/png', src: game.art.cover }"
        coverType="full">
        <template v-slot:activator="{ on }">
            <div v-focus
                @click="on"
                :class="'card text-white game-card'
                        + (title ? ' game-card-with-' + title + '-title' : '')
                        + (shape ? ' game-card-shape-' + shape : '')
                        + (size ? ' game-card-size-' + size : '')
                        + (brighten !== undefined ? ' game-card-brighten' : '')
                        + (unbordered !== undefined ? ' game-card-unbordered' : '')">
                <img :src="game.art.cover" class="card-img">
                <div class="game-card-tag">
                    <b-badge>{{ game.tag }}</b-badge>
                </div>
                <div
                    v-if="title"
                    :class="(title ? 'game-card-title-' + title : '')">
                    <h5 class="card-title">{{ game.name }}</h5>
                    <p class="card-text">{{ game.description }}</p>
                </div>
            </div>
        </template>

        <h3>{{ game.name }}</h3>
        <div v-html="game.art.cover"></div>
    </screen-modal>
</template>

<script>
    import ScreenModal from '@/components/ScreenModal.vue';

    export default {
        components: {
            ScreenModal
        },
        props: ['game',
            'title', 'shape', 'size',
            'brighten', 'unbordered']
    }
</script>

<style>
    .game-card {
        display: inline-block;
        margin: .5em;

        overflow: hidden;

        border: 0;
        background-color: rgba(0, 0, 0, 0);
        box-shadow: 0 0 5px #111;
        border-radius: 3px;

        transition: transform .2s;

        vertical-align: middle;
    }

    .game-card:focus, .game-card:hover {
        box-shadow: 0 0 0 4px #FFF;
        transform: scale(1.1);
        z-index: 100;
    }

    .game-card .card-title {
        margin-bottom: 0;

        white-space: normal;
        word-break: break-word;
    }

    .game-card .card-img {
        width: 100%;
        height: 100%;

        border-radius: 0;

        object-fit: cover;
        overflow: hidden;
    }

    /* Sizes */
    .game-card-size-xs { font-size: .5em; }
    .game-card-size-xs .card-title { font-size: 1.6em; }

    .game-card-size-sm { font-size: .75em; }
    .game-card-size-sm .card-title { font-size: 1.3em; }

    .game-card-size-md { font-size: 1em; }
    .game-card-size-md .card-title { font-size: 1.3em; }

    .game-card-size-lg { font-size: 1.25em; }
    .game-card-size-lg .card-title { font-size: 1.3em; }


    /* Brightened selections */
    .game-card.game-card-brighten .card-img {
        filter: brightness(.5);
    }

    .game-card.game-card-brighten:hover .card-img,
    .game-card.game-card-brighten:focus .card-img {
        filter: brightness(1);
    }


    /* Bordered selections */
    .game-card.game-card-unbordered:focus,
    .game-card.game-card-unbordered:hover {
        box-shadow: inherit;
    }


    /* Tag */
    .game-card .game-card-tag {
        position: absolute;
        top: 0;
        right: 0;
        margin: 0 .3em 0 0;
    }

    .game-card .game-card-tag .badge {
        font-weight: 300;
        text-transform: uppercase;
        background-color: rgba(0, 0, 0, .3);
    }

    /* Overlay */
    .game-card .game-card-title-overlay {
        position: absolute;
        bottom: 0;
        right: 0;
        left: 0;

        padding: 1em;
        line-height: 1em;

        background-color: rgba(0, 0, 0, .4);
    }

    .game-card.game-card-brighten:hover .game-card-title-overlay,
    .game-card.game-card-brighten:focus .game-card-title-overlay {
        background-color: rgba(0, 0, 0, .7);
    }

    /* Bottom title */
    .game-card .game-card-title-bottom {
        padding: .5em;

        line-height: 1em;
        background-color: rgba(0, 0, 0, .4);
    }

    .game-card .game-card-title-bottom .card-title {
        margin-bottom: .3em;
    }

    .game-card:hover .game-card-title-bottom,
    .game-card:focus .game-card-title-bottom {
        color: #000;
    }

    .game-card:hover .game-card-title-bottom,
    .game-card:focus .game-card-title-bottom {
        background-color: #FFF;
    }

    /* Floating title */
    .game-card.game-card-with-float-title:focus,
    .game-card.game-card-with-float-title:hover {
        transform: none;
    }

    .game-card .game-card-title-float {
        padding: .5em;
        line-height: 1em;
        font-size: .7em;
    }

    .game-card .game-card-title-float {
        padding: .5em;
        line-height: 1em;
        font-size: .7em;
    }

    /* Slide up title */
    .game-card.game-card-with-slide-up-title,
    .game-card.game-card-with-slide-up-title img {
        border-radius: .2em;
        transform: none;

        background-color: #FFF;

        transform: none;
    }

    .game-card.game-card-with-slide-up-title,
    .game-card.game-card-with-slide-up-title img,
    .game-card.game-card-with-slide-up-title .game-card-title-slide-up {
        transition: border-radius .35s, margin-top .35s, margin-bottom .35s;
    }

    .game-card.game-card-with-slide-up-title .card-title {
        font-size: .9em;
        text-align: center;
    }

    .game-card.game-card-with-slide-up-title .card-text {
        display: none;
    }

    .game-card.game-card-with-slide-up-title .game-card-title-slide-up {
        height: 0;
        padding: 0 .5em;

        color: #000;
    }

    .game-card.game-card-with-slide-up-title:focus img,
    .game-card.game-card-with-slide-up-title:hover img {
        margin-top: calc(-8% - .5em);
        border-radius: 1em;
    }

    .game-card.game-card-with-slide-up-title:focus .game-card-title-slide-up,
    .game-card.game-card-with-slide-up-title:hover .game-card-title-slide-up {
        margin-top: .4em;
        margin-bottom: 8%;

        background-color: #FFF;
    }

    /* Shapes */
    .game-card-shape-fluid .card-img {
        width: 15em;
        height: auto;
    }

    .game-card-shape-square .card-img {
        width: 15em;
        height: 15em;
    }

    .game-card-shape-vertical .card-img {
        width: 15em;
        height: 20em;
    }

    .game-card-shape-horizontal .card-img {
        width: 20em;
        height: 15em;
    }

    .game-card-shape-longboi .card-img {
        width: 20em;
        height: 10em;
    }

    .game-card-shape-circle {
        width: 15em;
        height: 15em;

        border-radius: 20em;

        text-align: center;
    }
    .game-card.game-card-shape-circle .game-card-overlay {
        left: 50%;
        bottom: 50%;
        width: 100%;

        transform: translate(-50%, 50%);
    }
</style>