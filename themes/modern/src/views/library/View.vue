<template>
  <b-container fluid gutter="none" class="my-3">
    <b-row>
      <b-col md="4" lg="3" class="hidden-sm-down">
        <b-form-input v-focus
          class="search"
          size="lg"
          placeholder="Search game library..."
          v-model="filter.search" />

        <hr />

        <b-nav vertical pills class="nav-lib">
          <template v-for="(link, i) in nav">
            <b-nav-item
              :key="i" v-if="link != null"
              :to="(link.route || { name: link.name })"
              :active="$route.name == (link.route || { name: link.name })"
              :class="link.color ? 'nav-color-' + link.color : ''">
              <fa-icon :icon="link.icon" /> {{ link.name }}
            </b-nav-item>
            <hr :key="i" v-else />
          </template>
        </b-nav>
      </b-col>
      <b-col md="8" lg="9" class="my-3">
        <b-container fluid>
          <router-view />
        </b-container>
      </b-col>
    </b-row>
  </b-container>
</template>

<style>
  .search {
    height: calc(1.5em + 1.5rem + 2px);
    padding: 0rem 1rem;
    font-size: 1.8rem;
    line-height: 1.5;
    border-radius: 0;
    font-weight: 300;
    font-family: Roboto;
  }
  
  .search, .search:focus {
    color: #FFF;
    background-color: rgba(0, 0, 0, 0);
    outline: 0;
    box-shadow: none;
    border: 0;
  }

  .nav-lib .nav-link {
    padding: 1em;
    background-color: rgba(0, 0, 0, .2);

    transition: transform .15s;
  }

  .nav-lib .nav-link.active {
    background-color: rgba(150, 150, 150, .5);
  }

  .nav-lib .nav-link:hover,
  .nav-lib .nav-link:focus {
    transform: scale(1.1);
  }

  .nav-lib .svg-inline--fa {
    width: 2em;
  }
</style>

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
        filter: {
          search: '',
          installed: true
        },
        selected: 'explore',

        nav: [
          { icon: 'home', name: 'Explore' },
          { icon: 'cubes', name: 'Collections' },
          null,
          { icon: 'clock', name: 'Recently Played' },
          { icon: 'play', name: 'Available' },
          null,
          { icon: 'book', name: 'All Games' },
          { icon: 'hdd', name: 'Local Games' },
          { icon: 'satellite', name: 'Stream Games' }
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