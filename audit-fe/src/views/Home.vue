<template>
  <v-container fill-height fluid class="gradient-bg">
    <v-col>
      <v-row>
        <v-col>
          <h1 class="text-sm-h1 text-h3 text-center">
            <Logo />
          </h1>
          <p
            class="mt-12 mb-12 text-subtitle-1 ma-auto text-center app-subtitle"
          >
            Service monitor and incident manager.
          </p>
        </v-col>
      </v-row>
      <v-row align="center" justify="center" class="mt-0">
        <v-btn
          :to="{ name: 'dashboard:monitor' }"
          elevation="6"
          rounded
          x-large
          class="ma-2"
          v-text="'View Dashboard'"
          color="primary darken-2"
        />
      </v-row>
      <v-row align="center" justify="center" class="mt-0">
        <v-btn
          href="https://github.com/daystram/audit"
          x-large
          class="mt-12"
          text
          plain
          rounded
        >
          View on GitHub <v-icon class="ml-1" v-text="'mdi-github'" />
        </v-btn>
      </v-row>
    </v-col>
    <div class="app-version text-overline text--disabled">
      {{ appVersion || "" }}
    </div>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Logo from "@/components/Logo.vue";
export default Vue.extend({
  components: {
    Logo,
  },
  data() {
    return {
      appVersion: process.env.VUE_APP_VERSION,
    };
  },
  computed: {
    isAuthenticated() {
      return false;
    },
  },
});
</script>
