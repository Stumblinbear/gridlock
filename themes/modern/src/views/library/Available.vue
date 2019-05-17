<template>
  <b-row align-h="center">
    <b-col cols="12" v-focus
      v-for="(game, id) in games" :key="id"
      class="d-flex justify-content-between align-items-center py-3 px-3"
      :to="{ name: 'Explore' }">
      {{ game.name }}
      <b-badge variant="primary" pill>Installed</b-badge>
    </b-col>

    <b-col
      v-for="(game, id) in games" :key="id">
      <game-card
        :game="game"
        title="bottom"
        shape="longboi"
        size="sm"
        brighten="true"
        tag="Installed" />
    </b-col>
  </b-row>
</template>

<script>
  import ScrollPanel from '@/components/ScrollPanel.vue';
  import GameCard from '@/components/GameCard.vue';

  export default {
    components: {
      ScrollPanel,
      GameCard
    },
    data() {
      return {
        games: [
          {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }, {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }, {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }, {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }, {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }, {
            name: 'Link to the Past',
            description: 'I like descriptions uwu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/254024-the-legend-of-zelda-a-link-to-the-past-snes-front-cover.jpg',
            }
          }, {
            name: 'Breath of the Wild',
            description: 'Breathhhhh',
            art: {
              cover: 'https://cdn.vox-cdn.com/uploads/chorus_asset/file/7802443/NintendoSwitch_TLOZBreathoftheWild_boxart.jpg',
            }
          }, {
            name: 'For Honor',
            description: 'for victory uwuwuwuu',
            art: {
              cover: 'https://www.mobygames.com/images/covers/l/380874-for-honor-xbox-one-front-cover.jpg',
            }
          }
        ]
      }
    }
  }
</script>