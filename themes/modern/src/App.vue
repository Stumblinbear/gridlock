<template>
  <div id="window">
    <b-navbar id="titlebar">
      <b-navbar-nav>
        <b-nav-item v-for="(link, i) in nav.left" :key="i"
          :to="(link.route || { name: link.name })"
          :active="$route.name == (link.route || { name: link.name })"
          :class="link.color ? 'nav-color-' + link.color : ''">
          <fa-icon :icon="link.icon" /> {{ link.name }}
        </b-nav-item>
      </b-navbar-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item v-for="(link, i) in nav.right" :key="i"
          :to="(link.route || { name: link.name })"
          :active="$route.name == (link.route || { name: link.name })"
          :class="link.color ? 'nav-color-' + link.color : ''">
          <fa-icon :icon="link.icon" /> {{ link.name }}
        </b-nav-item>
        <b-nav-item>
          <fa-icon icon="expand" href="'#/fullscreen'" />
        </b-nav-item>

        <div class="titlebar-controls">
          <b-nav-item class="titlebar-minimize">
            <svg x="0px" y="0px" viewBox="0 0 10 1">
              <rect fill="#FFF" width="10" height="1"></rect>
            </svg>
          </b-nav-item>
          <b-nav-item class="titlebar-resize">
            <svg class="fullscreen-svg" x="0px" y="0px" viewBox="0 0 10 10">
              <path fill="#FFF" d="M 0 0 L 0 10 L 10 10 L 10 0 L 0 0 z M 1 1 L 9 1 L 9 9 L 1 9 L 1 1 z "/>
            </svg>
            <svg class="maximize-svg" x="0px" y="0px" viewBox="0 0 10 10">
              <mask id="Mask">
                  <rect fill="#000" width="10" height="10"></rect>
                  <path fill="#FFF" d="M 3 1 L 9 1 L 9 7 L 8 7 L 8 2 L 3 2 L 3 1 z"/>
                  <path fill="#FFF" d="M 1 3 L 7 3 L 7 9 L 1 9 L 1 3 z"/>
              </mask>
              <path fill="#FFF" d="M 2 0 L 10 0 L 10 8 L 8 8 L 8 10 L 0 10 L 0 2 L 2 2 L 2 0 z" mask="url(#Mask)"/>
            </svg>
          </b-nav-item>
          <b-nav-item class="titlebar-close" @click="exit">
            <svg x="0px" y="0px" viewBox="0 0 10 10">
              <polygon fill="#FFF" points="10,1 9,0 5,4 1,0 0,1 4,5 0,9 1,10 5,6 9,10 10,9 6,5"></polygon>
            </svg>
          </b-nav-item>
        </div>
      </b-navbar-nav>
    </b-navbar>

    <div id="window-content">
      <router-view />
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        show: false,

        nav: {
          left: [
            { icon: 'bookmark', name: 'Library', route: { name: 'Explore' } },
            { icon: 'newspaper', name: 'News' },
            { icon: 'store', name: 'Store' },
            { icon: 'cog', name: 'Settings' }
          ],
          right: [
            { icon: 'user', color: 'blue', route: { name: 'User' } }
          ]
        }
      }
    },
    methods: {
      exit() {
        window.close();
      }
    }
  }
</script>