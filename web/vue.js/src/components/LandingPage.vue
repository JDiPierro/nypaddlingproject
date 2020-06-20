<template>
  <div>
    <v-parallax :src="require('../assets/vbanner.jpg')">
      <v-layout class="center-text" align-center column justify-center>
        <h1 class="display-2 font-weight-thin mb-3">Upstate New York Paddling Project</h1>
        <LoginWithFB/>
      </v-layout>
    </v-parallax>

    <v-layout class="padding-top-bottom" justify-center>
      <v-flex class="center-text" xs12 sm10>
        <h1>A group effort to document boat launch locations in and around Upstate New York.</h1>
      </v-flex>
    </v-layout>

    <v-layout justify-center>
      <v-flex xs12 sm10>
        <v-divider></v-divider>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
import LoginWithFB from '../components/LoginWithFB.vue';

export default {
  components: {
    LoginWithFB,
  },
}
</script>

<style scoped>
.center-text {
  text-align: center;
}

.center-image {
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.padding-top-bottom {
  padding-top: 50px;
  padding-bottom: 50px;
}
</style>
