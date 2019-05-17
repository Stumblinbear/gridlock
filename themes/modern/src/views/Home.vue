<template>
  <div>
    <b-carousel
      class="w-100"
      style="height:30em;overflow:hidden;box-shadow: 0 0 10px #111"
      v-model="slide"
      :interval="120000"
      controls
      indicators
      background="#ababab"
      img-width="1024"
      img-height="480">
      <b-carousel-slide
          v-for="(game, id) in recentlyPlayed" :key="id"
          :caption="game.name"
          :text="game.description"
          class="h-100">
        <div
          slot="img">
          <img
            class="d-block img-fluid w-100"
            width="1024"
            height="480"
            :src="game.art.play"
            style="filter: brightness(.4); object-size:cover">
        </div>
      </b-carousel-slide>
    </b-carousel>

    <b-container class="mt-3">
      <b-row>
        <b-col>
          <h3>Up Next</h3>
        </b-col>
      </b-row>

      <b-row align-h="center">
        <b-col
          v-for="(game, id) in recentlyPlayed" :key="id"
          style="max-width:15em;height:20em;margin-bottom:1em">
          <game-card v-focus
            :bg="game.art.cover"
            :name="game.name"
            :description="game.description"/>
        </b-col>
      </b-row>
    </b-container>

    <!--<div class="uk-padding">
      <h4>Recently Played</h4>

      <vk-grid gutter="small" matched>
        <div v-for="(game, id) in recentlyPlayed" :key="id" class="uk-width-medium">
          <vk-card type="secondary" v-focus
            class="uk-background-cover uk-background-blend-overlay"
            v-bind:style="{ backgroundImage: 'url(' + game.url + ')', paddingTop: '150px' }">
            <vk-label slot="badge">5 days ago</vk-label>

            <vk-card-title>{{ game.name }}</vk-card-title>
            <p>{{ game.description }}</p>
            
            <vk-button type="primary">Launch</vk-button>
          </vk-card>
        </div>
      </vk-grid>
    </div>
    <div class="uk-padding">
      <h4>Suggestions</h4>

      <vk-grid gutter="small" matched>
        <div v-for="(game, id) in recentlyPlayed" :key="id" class="uk-width-medium">
          <vk-card type="secondary" v-focus
            class="uk-background-cover uk-background-blend-overlay"
            v-bind:style="{ backgroundImage: 'url(' + game.url + ')', paddingTop: '150px' }">
            <vk-label slot="badge">New!</vk-label>

            <vk-card-title>{{ game.name }}</vk-card-title>
            <p>{{ game.description }}</p>
            
            <vk-button type="primary">Launch</vk-button>
          </vk-card>
        </div>
      </vk-grid>
    </div>-->
  </div>
</template>

<script>
  import GameCard from '@/components/GameCard.vue';

  export default {
    components: {
      GameCard
    },
    data() {
      return {
        slide: 0,
        sliding: null,

        recentlyPlayed: {
          "link-to-the-past": {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
              play: 'https://i.kinja-img.com/gawker-media/image/upload/s--TEZMLFfb--/c_fill,fl_progressive,g_center,h_900,q_80,w_1600/whc79shzn5hazqr5tfz7.jpg'
            }
          }, "breath-of-the-wild": {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
              play: 'http://cdn1.expertreviews.co.uk/sites/expertreviews/files/2017/12/the_legend_of_zelda_breath_of_the_wild_the_champions_ballard_dlc_dated.png'
            }
          }, "for-honor": {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
              play: 'https://gameranx.com/wp-content/uploads/2017/02/ForHonor8.jpg'
            }
          }, "for-hfonor": {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
              play: 'https://gameranx.com/wp-content/uploads/2017/02/ForHonor8.jpg'
            }
          }
        }
      }
    }
  }
</script>