<template>
  <div>
    <div class="uk-padding">
      <vk-card type="secondary" padding="small">
        <vk-spinner /> &nbsp; Downloading 5 games...
      </vk-card>
    </div>

    <div class="uk-padding">
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
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        recentlyPlayed: {
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
          }
        }
      }
    }
  }
</script>