<template>
    <screen-modal
        :cover="article.cover"
        :bg="article.cover">
        <template v-slot:activator="{ on }">
            <div v-focus
                :class="'card text-white news-card' + (article.important ? ' news-important' : '')"
                @click="on">
                <img :src="article.thumbnail" class="card-img">
                <div class="news-card-content">
                    <h5 class="card-title"
                        :style="!article.preview ? 'text-align:center' : ''">
                        {{ article.title }}
                    </h5>
                    <p class="card-text" v-if="preview">
                        {{ article.preview.replace(/<\/?[^>]+(>|$)/g, '') }}
                    </p>
                </div>
            </div>
        </template>

        <h3>{{ article.title }}</h3>
        <div v-html="article.preview"></div>
    </screen-modal>
</template>

<script>
    import ScreenModal from '@/components/ScreenModal.vue';

    export default {
        components: {
            ScreenModal
        },
        props: ['article']
    }
</script>

<style>
    .news-card {
        position: relative;
        height: 20em;

        border: 0;
        background-color: rgba(0, 0, 0, .4);
        box-shadow: 0 0 5px #111;
        border-radius: 2px;

        transition: transform .2s;
    }

    .news-card:focus,
    .news-card:hover {
        box-shadow: 0 0 0 4px #FFF;
        transform: scale(1.1);
        z-index: 100;
    }

    .news-card img {
        height: 50%;

        border-radius: 2px;
        box-sizing: border-box;
        overflow: hidden;
        
        object-fit: cover;
    }

    .news-card.news-important {
        border-radius: 5px;
    }

    .news-card.news-important img {
        height: 100%;
        border-radius: 5px;
    }

    .news-card .news-card-content {
        max-height: 50%;
        
        overflow-x: scroll;
        scrollbar-width: none; /* Firefox */
        -ms-overflow-style: none;  /* IE 10+ */
    }

    .news-card .news-card-content::-webkit-scrollbar { /* WebKit */
        width: 0;
        height: 0;
    }

    .news-card .card-title {
        margin: .5em;
    }

    .news-card .card-text {
        margin: 1em;
    }

    .news-card.news-important .news-card-content {
        position: absolute;
        top: 50%;
        left: 0;
        right: 0;

        background-color: rgba(0, 0, 0, .6);
        
        transform: translate(0, -50%);
    }
</style>