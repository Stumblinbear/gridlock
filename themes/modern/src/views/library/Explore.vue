<template>
  <b-row align-h="center">
    <b-col>
      <div class="mb-4 px-2 py-2">
        <h4>Recently Played</h4>
        <scroll-panel>
          <game-card
            v-for="(game, id) in allGames" :key="id"
            :game="game"
            :tag="'played ' + Math.round((id + 1) * 1.7) + ' days ago'"
            title="slide-up"
            shape="horizontal"
            size="lg"
            brighten
            unbordered />
        </scroll-panel>
      </div>
      <div class="mb-4 px-2 py-2">
        <h4>Recently Played</h4>
        <scroll-panel>
          <game-card
            v-for="(game, id) in allGames" :key="id"
            :game="game"
            title="overlay"
            shape="vertical"
            size="sm"
            brighten />
        </scroll-panel>
      </div>
      <div class="mb-4 px-2 py-2">
        <h4>Recently Played</h4>
        <scroll-panel>
          <game-card
            v-for="(game, id) in allGames" :key="id"
            :game="game"
            title="overlay"
            shape="square"
            size="sm"
            brighten />
        </scroll-panel>
      </div>
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
        allGames: [
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
    },
    computed: {
      games() {
        if(!this.filter.search) return this.allGames;

        let list = [];

        for(let game of this.allGames) {
          if(game.name.toLowerCase().indexOf(this.filter.search.toLowerCase()) !== -1)
            list.push(game);
        }

        return list;
      }
    }
  }
</script>