<template>
  <div class="list">
    <v-row class="mb-8" align="center">
      <v-col cols="12" sm="">
        <h1 class="text-h2">Manage</h1>
      </v-col>
    </v-row>
    <v-fade-transition>
      <div v-show="pageLoadStatus === STATUS.COMPLETE">
        <v-row>
          <v-col cols="12">
            <v-card
              v-for="application in applications"
              :key="application.id"
              elevation="8"
              class="mb-4 mx-auto"
            >
              <v-card-title>{{ application.name }}</v-card-title>
              <v-card-subtitle>{{ application.description }}</v-card-subtitle>
              <v-divider inset />

              <v-card-text class="py-8">
                <v-expansion-panels multiple hover>
                  <v-expansion-panel>
                    <v-expansion-panel-header disable-icon-rotate>
                      <b>service-be</b>
                      <template v-slot:actions>
                        <v-icon color="success"> mdi-server </v-icon>
                      </template>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content>
                      Lorem ipsum dolor sit amet, consectetur adipiscing elit,
                      sed do eiusmod tempor incididunt ut labore et dolore magna
                      aliqua. Ut enim ad minim veniam, quis nostrud exercitation
                      ullamco laboris nisi ut aliquip ex ea commodo consequat.
                    </v-expansion-panel-content>
                  </v-expansion-panel>
                  <v-expansion-panel>
                    <v-expansion-panel-header disable-icon-rotate>
                      <b>service-fe</b>
                      <template v-slot:actions>
                        <v-icon color="error"> mdi-server </v-icon>
                      </template>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content>
                      Lorem ipsum dolor sit amet, consectetur adipiscing elit,
                      sed do eiusmod tempor incididunt ut labore et dolore magna
                      aliqua. Ut enim ad minim veniam, quis nostrud exercitation
                      ullamco laboris nisi ut aliquip ex ea commodo consequat.
                    </v-expansion-panel-content>
                  </v-expansion-panel>
                </v-expansion-panels>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </div>
    </v-fade-transition>
    <v-fade-transition>
      <v-overlay
        v-show="
          pageLoadStatus === STATUS.PRE_LOADING ||
          pageLoadStatus === STATUS.LOADING
        "
        opacity="0"
        absolute
      >
        <v-progress-circular indeterminate size="64" />
      </v-overlay>
    </v-fade-transition>
    <v-expand-transition>
      <div v-show="pageLoadStatus === STATUS.ERROR">
        <v-alert
          type="error"
          text
          transition="scroll-y-transition"
          class="mt-0"
        >
          Failed loading services!
        </v-alert>
      </div>
    </v-expand-transition>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import api from "@/apis";
import { STATUS } from "@/constants";
import { ApplicationInfo } from "@/apis/datatransfers";
export default Vue.extend({
  data() {
    return {
      STATUS,
      pageLoadStatus: STATUS.PRE_LOADING,
      applications: new Array<ApplicationInfo>(),
    };
  },
  created() {
    api.application.list().then((response) => {
      const data = response.data.data;
      for (let i = 0; i < data.length; i++) {
        this.applications.push(new ApplicationInfo(data[i]));
        console.log(new ApplicationInfo(data[i]));
      }
      this.pageLoadStatus = STATUS.COMPLETE;
    });
  },
});
</script>
