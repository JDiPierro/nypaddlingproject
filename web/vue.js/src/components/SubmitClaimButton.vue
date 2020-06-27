<template>
  <v-dialog
    v-model="dialog"
    width="500"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn :color="triggerBtnColor"
             class="ml-2"
             v-bind="attrs"
             v-on="on"
      >
        {{triggerBtnText}}
      </v-btn>
    </template>

    <v-card>
      <v-card-title
        class="headline primary lighten-3"
        primary-title
      >
        Submit a Change
      </v-card-title>

      <v-card-text>
        <h2 class="text-center my-4 blue--text darken-3">Thank you for your contribution!</h2>
        <p>
          <a href="https://paddling.com">Paddling.com</a> reviews change submissions manually so they may take a bit to show up. We typically see changes go through within 2 weeks.
        </p>
        <p>Our admin team will periodically check the page for <b class="black--text">{{claim.location.name}}</b> and reward you with points when the update has processed.</p>
        <p>
          Let us know what changes you submitted so we know what to look out for when reviewing your submission.
        </p>
        <v-list class="transparent">
          <v-list-item>
            <v-checkbox v-model="addedPhotos" label="Added Photos"></v-checkbox>
          </v-list-item>
          <v-list-item>
            <v-checkbox v-model="updatedDescription" label="Updated Description"></v-checkbox>
          </v-list-item>
          <v-list-item>
            <v-checkbox v-model="updatedAmenities" label="Updated Amenities"></v-checkbox>
          </v-list-item>
          <v-list-item class="mt-4">
            <v-textarea v-model="notes" label="Notes" outlined placeholder="Optional"></v-textarea>
          </v-list-item>
        </v-list>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-actions>
        <v-btn
          color="error"
          @click="dialog = !dialog"
        >
          Cancel
        </v-btn>
        <v-spacer></v-spacer>
        <v-btn
          color="success"
          @click="submitChange()"
        >
          Submit
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
  import { mapActions } from 'vuex';

  export default {
    name: "SubmitClaimButton",
    props: ['claim'],
    data () {
      return {
        dialog: false,

        // Form fields
        addedPhotos: false,
        updatedDescription: false,
        updatedAmenities: false,
        notes: ""
      }
    },
    computed: {
      triggerBtnColor() {
        if(this.claim['status'] === 'initial') {
          return "success"
        }
        return "amber"
      },
      triggerBtnText() {
        if(this.claim['status'] === 'initial') {
          return "Submit Change"
        }
        return "Update Change"
      }
    },
    beforeMount() {
      if(this.claim['status'] === 'submitted') {
        this.addedPhotos = this.claim['update_info']['addedPhotos']
        this.updatedDescription = this.claim['update_info']['updatedDescription']
        this.updatedAmenities = this.claim['update_info']['updatedAmenities']
        this.notes = this.claim['update_info']['notes']
      }
    },
    methods: {
      ...mapActions('locations', ['submit']),
      submitChange() {
        this.submit({
          location_id: this.claim.location_id,
          update_info: {
            addedPhotos: this.addedPhotos,
            updatedDescription: this.updatedDescription,
            updatedAmenities: this.updatedAmenities,
            notes: this.notes
          }
        })
        // TODO: Check success
        this.dialog = false
      }
    }
  }
</script>

<style scoped>

</style>
