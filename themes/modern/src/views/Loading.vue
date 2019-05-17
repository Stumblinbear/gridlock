<template>
  <div>
    <b-navbar class="titlebar">
      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
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

    <b-container>
      <b-row>
        <b-col class="mt-5" md="4" offset-md="4">
          <b-card>
            <b-card-text class="text-center">
              <b-spinner variant="primary" type="grow" label="Loading"></b-spinner>
            </b-card-text>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
  export default {
    mounted() {
      setTimeout(() => {
        this.$router.push('windowed');
      }, 5000);
    }
  }
</script>