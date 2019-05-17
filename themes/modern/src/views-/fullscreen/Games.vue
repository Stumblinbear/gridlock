<template>
  <div>
    <div class="uk-padding">
      <vk-grid gutter="small" matched class="uk-flex-center">
        <div v-for="(game, id) in games" :key="id" class="uk-width-medium">
          <vk-card type="secondary"
            class="game uk-background-cover uk-background-blend-overlay"
            v-bind:style="{ backgroundImage: 'url(' + game.url + ')', paddingTop: '200px' }"
            v-focus @click="console.log('wew')">
            <vk-label slot="badge">New!</vk-label>

            <vk-card-title>{{ game.name }}</vk-card-title>
            <p>{{ game.description }}</p>
          </vk-card>
        </div>
      </vk-grid>
    </div>
  </div>
</template>

<style>
  .game {
    transition: 1s;
  }
  .game:focus, .game:hover {
    outline: 2px solid #FFF;
    transform: scale(1.1);
    z-index: 10;
  }
</style>

<script>
  export default {
    data() {
      return {
        games: {
          "link-to-the-past": {
            url: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            name: 'Link to the Past',
            description: 'I like descriptions uwu'
          }, "breath-of-the-wild": {
            url: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            name: 'Breath of the Wild',
            description: 'Breathhhhh'
          }, "for-honor": {
            url: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            name: 'For Honor',
            description: 'for victory uwuwuwuu'
          }, "for-honfor": {
            url: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            name: 'For Honor',
            description: 'for victory uwuwuwuu'
          }, "for-honour": {
            url: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            name: 'For Honor',
            description: 'for victory uwuwuwuu'
          }, "for-honori": {
            url: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            name: 'For Honor',
            description: 'for victory uwuwuwuu'
          }
        }
      }
    }
  }
</script>