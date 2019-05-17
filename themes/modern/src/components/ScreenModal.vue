<template>
    <span>
        <slot name="activator" :on="showModal"></slot>

        <div class="screen-overlay" ref="modal"
            v-focus-section="this.id" @click="hideModal"
            v-if="this.shown">
            <b-container fluid style="padding:0;height:100%">
                <b-row style="height:100%">
                    <b-col
                        @click.stop=""
                        :xs="xs ? xs : 12" :offset-xs="12 - (xs ? xs : 12)"
                        :sm="sm ? sm : 12" :offset-sm="12 - (sm ? sm : 12)"
                        :md="md ? md : 8" :offset-md="12 - (md ? md : 8)"
                        :lg="lg ? lg : 6" :offset-lg="12 - (lg ? lg : 6)"
                        :xl="xl ? xl : 5" :offset-xl="12 - (xl ? xl : 5)">
                        <div class="screen-modal" @click.stop="">
                            <template v-if="bg">
                                <img
                                    v-if="bg.type.indexOf('image') !== -1"
                                    class="screen-modal-bg"
                                    :src="bg.src" />
                                <video v-focus
                                    v-else-if="bg.type.indexOf('video') !== -1"
                                    class="screen-modal-bg"
                                    muted autoplay>
                                    <source :type="bg.type" :src="bg.src" />
                                </video>
                                <div v-else>{{ bg }}</div>
                            </template>

                            <div class="screen-modal-container">
                                <div class="screen-modal-close" v-focus="this.id"
                                    @click.stop="hideModal">
                                    <fa-icon icon="times" />
                                </div>

                                <template v-if="cover">
                                    <img
                                        v-if="cover.type.indexOf('image') !== -1"
                                        class="screen-modal-cover"
                                        :src="cover.src" />
                                    <video v-focus
                                        v-else-if="cover.type.indexOf('video') !== -1"
                                        class="screen-modal-cover"
                                        autoplay>
                                        <source :type="cover.type" :src="cover.src" />
                                    </video>
                                    <div v-else>{{ cover }}</div>
                                </template>

                                <div class="screen-modal-content">
                                    <slot></slot>
                                </div>
                            </div>
                        </div>
                    </b-col>
                </b-row>
            </b-container>
        </div>
    </span>
</template>

<script>
    function uuidv4() {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
            var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }

    export default {
        props: ['cover', 'bg',
            'xs', 'sm', 'md', 'lg', 'xl'],
        data() {
            return {
                id: 'modal-' + uuidv4(),
                shown: false
            }
        },
        methods: {
            showModal() {
                this.shown = true;

                SpatialNavigation.disable('main');

                SpatialNavigation.enable(this.id);
                setTimeout(() => {
                    SpatialNavigation.focus(this.id, true);
                }, 800);
            },
            hideModal() {
                this.shown = false;

                SpatialNavigation.enable('main');
                SpatialNavigation.remove(this.id);
            }
        }
    }
</script>

<style>
    .screen-overlay {
        position: fixed;
        left: 0;
        top: 0;

        width: 100%;
        height: 100%;
        overflow: hidden;

        background-color: rgba(0, 0, 0, 0.4);
        animation-name: fadeInBg;
        animation-duration: 0.4s;

        scrollbar-width: thin; /* Firefox */
        -ms-overflow-style: thin;  /* IE 10+ */
        z-index: 999999;
    }

    .screen-modal {
        display: flex;
        justify-content: center;
        align-items: center;

        width: 100%;
        height: 100%;
        
        color: #FFF;
        background-color: #222;

        border-top-left-radius: 1em;
        border-top-right-radius: 1em;
        box-shadow: 0 0 10px #111;
        box-sizing: border-box;

        overflow: hidden;
        
        animation-name: slideIn;
        animation-duration: .5s;
    }

    .screen-overlay .row > div { padding: 0; }
    .screen-overlay .row { margin: 0; }

    .screen-modal .screen-modal-bg {
        /*min-width: 100%;*/
        min-height: 100%;

        border-top-left-radius: 1em;
        border-top-right-radius: 1em;
        box-sizing: border-box;

        filter: brightness(.15);
        z-index: 0;
    }

    .screen-modal .screen-modal-container {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;

        width: 100%;
        height: 100%;
    }

    .screen-modal-close {
        position: absolute;
        top: 0;
        right: 0;
        margin: .25em;
        padding: .25em .75em;

        font-size: 2em;
        color: rgba(255, 255, 255, .5);
        
        animation-name: slideIn;
        animation-duration: .8s;

        z-index: 1000;
    }

    .screen-modal-close:focus,
    .screen-modal-close:hover {
        color: rgba(255, 255, 255, 1);
        box-shadow: none;
    }

    .screen-modal .screen-modal-cover {
        width: 100%;
        max-height: 30em;

        border-top-left-radius: 1em;
        border-top-right-radius: 1em;
        box-sizing: border-box;
        overflow: hidden;
        
        object-fit: cover;
    }

    .screen-modal video.screen-modal-cover {
        padding: 2px;
    }

    .screen-modal .screen-modal-content {
        padding: 1em;

        z-index: 1000000;
    }

    .screen-modal .screen-modal-content * {
        max-width: 100%;

        white-space: normal;
        word-wrap: break-word;
    }

    /* Add Animation */
    @keyframes slideIn {
        from { transform: translateY(100%) }
        to { transform: translateY(0%) }
    }

    @keyframes fadeInBg {
        from { background-color: rgba(0, 0, 0, 0); }
        to { background-color: rgba(0, 0, 0, 0.4); }
    }

    @media (min-width: 768px) {
        .screen-modal {
            border-top-right-radius: 0;
            border-bottom-left-radius: 1em;
        }

        .screen-modal .screen-modal-bg {
            border-top-right-radius: 0;
            border-bottom-left-radius: 1em;
        }

        .screen-modal .screen-modal-cover {
            border-top-right-radius: 0;
            border-bottom-left-radius: 1em;
        }
        
        @keyframes slideIn {
            from { transform: translateX(100%) }
            to { transform: translateX(0%) }
        }

        @keyframes fadeInBg {
            from { background-color: rgba(0, 0, 0, 0); }
            to { background-color: rgba(0, 0, 0, 0.4); }
        }
    }
</style>