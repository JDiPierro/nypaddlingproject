
<template>
  <div>
    <v-btn color="primary" @click="logInWithFacebook">Login with Facebook</v-btn>
  </div>
</template>
<script>
  import { mapActions } from 'vuex';

  export default {
    name:"LoginWithFB",
    methods: {
      ...mapActions('user', ['login']),
      async logInWithFacebook() {
        let self = this;
        await window.FB.login(function(response) {
          if (response.authResponse) {
            self.login(response.authResponse)
          } else {
            alert("Facebook login failed...");
          }
        });
      },
      async initFacebook() {
        window.fbAsyncInit = function() {
          window.FB.init({
            appId: "2577263505936090",
            cookie: true, // This is important, it's not enabled by default
            version: "v13.0"
          });
        };
      },
      async loadFacebookSDK(d, s, id) {
        var js,
            fjs = d.getElementsByTagName(s)[0];
        if (d.getElementById(id)) {
          return;
        }
        js = d.createElement(s);
        js.id = id;
        js.src = "https://connect.facebook.net/en_US/sdk.js";
        fjs.parentNode.insertBefore(js, fjs);
      }
    },
    async mounted() {
      await this.loadFacebookSDK(document, "script", "facebook-jssdk");
      await this.initFacebook();
    }
  };
</script>
<style>
</style>
