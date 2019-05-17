<template>
  <div>
    <div class="uk-section uk-section-xsmall uk-section-muted">
      <vk-card type="blank" padding="small">
        <form class="uk-search uk-search-navbar">
          <input class="uk-search-input" type="search" placeholder="Search game library...">
        </form>
      </vk-card>
    </div>

    <hr />
    
    <vk-grid gutter="small" matched class="uk-flex-center">
      <div v-for="(game, id) in games" :key="id" class="uk-width-medium">
        <vk-card type="secondary" hover
          class="uk-background-cover uk-background-blend-overlay"
          v-bind:style="{ backgroundImage: 'url(' + game.url + ')', paddingTop: '150px' }">
          <vk-label slot="badge">New!</vk-label>

          <vk-card-title>{{ game.name }}</vk-card-title>
          <p>{{ game.description }}</p>
          
          <vk-button type="primary">Launch</vk-button>
        </vk-card>
      </div>
    </vk-grid>
  </div>
</template>

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