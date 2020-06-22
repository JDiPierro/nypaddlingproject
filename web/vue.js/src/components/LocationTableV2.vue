<template>
<v-card>
  <v-toolbar
    flat
  >
    <v-row
      class="mb-6"
    >
      <v-col md="4">
        <v-select
          v-model="selectedCounty"
          :items="counties"
          clearable
          outlined
          label="Filter by County"
        ></v-select>
      </v-col>
      <v-col
        md="4"
        offset-md="4"
      >
        <v-text-field
          label="Search"
          v-model="searchFilter"
          outlined
        ></v-text-field>
      </v-col>
    </v-row>
  </v-toolbar>
</v-card>
</template>

<script>
  export default {
    name: "LocationTableV2",
    data: () => {
      return {
        selectedCounty: "",
        searchFilter: "",
        counties: ["Saratoga", "Albany", "Ulster"]
      }
    }
  }
</script>

<style scoped>

</style>
