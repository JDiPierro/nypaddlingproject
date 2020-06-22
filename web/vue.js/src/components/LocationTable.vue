<template>
<div>
  <v-data-table
    :headers="headers"
    :items="locations"
  >
    <template v-slot:item.link="{ item }">
      <v-btn color="primary" :href="item.link" target="_blank">Paddling.com</v-btn>
      &nbsp;
      <v-btn color="success">Claim</v-btn>
    </template>
  </v-data-table>
</div>
</template>

<script>
  export default {
    name: "LocationTable",
    props: ["locations"],
    data: () => {
      return {
        headers: [
          { text: 'Name', value: 'name'},
          { text: 'County', value: 'county' },
          { text: 'Created At', value: 'created_at' },
          { text: 'Updated At', value: 'updated_at' },
          { text: 'Description Length', value: 'desc_len' },
          { text: 'Num Photos', value: 'num_photos' },
          { text: 'Links', value: 'link', sortable: false },
        ],
      }
    }
  }

</script>

<style scoped>

</style>
