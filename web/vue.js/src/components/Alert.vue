<template>
  <v-alert v-model="isVisible" :type="type" transition="scale-transition" dismissible outline>{{value}}</v-alert>
</template>

<script>
export default {
  props: ["value", "type"],
  methods: {
    isVisible: function() {
      if (this.value !== '') {
        return true;
      } else {
        return false;
      }
    }
  }
}
</script>
