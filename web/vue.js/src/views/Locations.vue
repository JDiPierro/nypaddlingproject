<template id="app">
  <v-main>
    <v-container>
      <h1>Locations</h1>
      <LocationTable
        :locations="locations"
      ></LocationTable>
    </v-container>
  </v-main>
</template>

<script lang="ts">
  import { default as LocationTable } from '../components/LocationTable.vue'
  import { mapState, mapActions } from 'vuex'

  export default{
    components: {
      LocationTable,
    },
    computed: {
      ...mapState('locations', ['locations'])
    },
    methods: {
      ...mapActions('locations', ['loadAll']),
    },
    beforeMount () {
      this.loadAll()
    }
  }
</script>

<style scoped lang="scss">
</style>
