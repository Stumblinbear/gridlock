<template>
  <div>
    <vk-navbar :transparent="true">
      <vk-navbar-nav>
        <vk-navbar-logo class="uk-logo" id="title">Gridlock</vk-navbar-logo>
        <vk-nav-item v-for="(link, i) in nav.left" :key="i"
          :href="'#' + link.path"
          :icon="link.icon" :title="link.name"
          :active="$route.path == link.path"
          v-focus />
      </vk-navbar-nav>
      <vk-navbar-nav slot="right">
        <vk-nav-item v-for="(link, i) in nav.right" :key="i"
          :href="'#' + link.path"
          :icon="link.icon" :title="link.name"
          :active="$route.path == link.path"
          v-focus />
        <vk-nav-item icon="close" @click="exit"
          v-focus />
      </vk-navbar-nav>
    </vk-navbar>

    <router-view />
  </div>
</template>

<style>
  :focus {
    outline: 2px solid #FFF;
  }
</style>

<script>
  export default {
    data() {
      return {
        show: false,

        nav: {
          left: [
            { icon: 'home', name: 'Home', path: '/windowed/home' },
            { icon: 'bookmark', name: 'Library', path: '/windowed/games' }
          ],
          right: [
            { icon: 'cog', name: 'Settings', path: '/windowed/settings' }
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