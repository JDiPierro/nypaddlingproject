<template>
<div>
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
        clearable
        outlined
      ></v-text-field>
    </v-col>
  </v-row>
  <v-data-table
    id="location-table"
    :headers="headers"
    :items="locations"
    :items-per-page="itemsPerPage"
    :search="searchFilter"
  >
    <template v-slot:item.name="{ item }">
      <h2>{{item.name}}</h2>
    </template>

    <template v-slot:item.desc_len="{ item }">
      <v-chip
        :color="descBadgeColor(item.desc_len)"
      >
        {{item.desc_len}}
      </v-chip>
    </template>

    <template v-slot:item.num_photos="{ item }">
      <v-chip
        :color="photoBadgeColor(item.num_photos)"
      >
        {{item.num_photos}}
      </v-chip>
    </template>

    <template v-slot:item.link="{ item }">
      <v-btn color="primary" :href="item.link" target="_blank">Paddling.com</v-btn>
      &nbsp;
      <ClaimButton
        :location_id="item._id"
        :claims="item.claims"
      ></ClaimButton>
    </template>
  </v-data-table>
</div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex';
  import { default as ClaimButton } from './ClaimButton';

  export default {
    name: "LocationTable",
    props: ["locations"],
    components: {
      ClaimButton
    },
    data: () => {
      return {
        selectedCounty: null,
        searchFilter: null,
        itemsPerPage: 12
      }
    },
    computed: {
      ...mapGetters('locations', ['counties']),
      headers () {
        return [
          { text: 'Name', value: 'name'},
          {
            text: 'County',
            value: 'county',
            width: 100,
            filter: (value, search, item) => {
              if (!this.selectedCounty) {
                return true
              }
              return item.county === this.selectedCounty
            }
          },
          { text: 'Created At', value: 'created_at', filterable: false, width: 150 },
          { text: 'Updated At', value: 'updated_at', filterable: false, width: 150 },
          { text: 'Description Length', value: 'desc_len', filterable: false, width: 50 },
          { text: 'Num Photos', value: 'num_photos', filterable: false, width: 50 },
          { text: 'Links', value: 'link', sortable: false, filterable: false, width: 350 },
        ]
      }
    },
    methods: {
      ...mapActions('locations',['claim']),
      descBadgeColor(len_s) {
        const len = parseInt(len_s)
        if(len === 0) {
          return "error"
        }
        if(len < 100) {
          return "warning"
        }
        return "success"
      },
      photoBadgeColor(num_s) {
        const num = parseInt(num_s)
        if(num === 0) {
          return "error"
        }
        if(num < 2) {
          return "warning"
        }
        return "success"
      }
    },
  }
</script>

<style scoped>
  #location-table >>> tr {
    height: 60px;
  }
</style>
