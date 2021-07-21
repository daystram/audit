<template>
  <div class="list">
    <v-row class="mb-8" align="center">
      <v-col cols="12" sm="">
        <h1 class="text-h2">Monitor</h1>
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
              <v-card-title>
                {{ application.name }}<v-spacer />
                <v-icon
                  v-if="getApplicationStatus(application) === SERVICE_STATUS.OK"
                  color="success"
                >
                  mdi-check-circle-outline
                </v-icon>
                <v-icon
                  v-else-if="
                    getApplicationStatus(application) === SERVICE_STATUS.WARNING
                  "
                  color="warning"
                >
                  mdi-alert-outline
                </v-icon>
                <v-icon
                  v-else-if="
                    getApplicationStatus(application) === SERVICE_STATUS.ERROR
                  "
                  color="error"
                >
                  mdi-alert-octagon-outline
                </v-icon>
                <v-icon v-else>mdi-alert-octagon-outline</v-icon>
              </v-card-title>
              <v-card-subtitle>{{ application.description }}</v-card-subtitle>
              <v-divider inset />

              <v-card-text class="pt-6">
                <v-expansion-panels multiple hover>
                  <v-expansion-panel
                    v-for="service in application.services"
                    :key="service.id"
                  >
                    <v-expansion-panel-header disable-icon-rotate>
                      <template v-slot:default>
                        <v-row no-gutters>
                          <v-col cols="12" sm="4" class="text-truncate">
                            <b>{{ service.name }}</b>
                          </v-col>
                          <v-col
                            cols="12"
                            sm="6"
                            class="text--secondary text-truncate"
                          >
                            {{ service.description }}
                          </v-col>
                          <v-col cols="12" sm="2"></v-col>
                        </v-row>
                      </template>
                      <template v-slot:actions>
                        <v-icon
                          v-if="getServiceStatus(service) === SERVICE_STATUS.OK"
                          color="success"
                        >
                          mdi-server
                        </v-icon>
                        <v-icon
                          v-else-if="
                            getServiceStatus(service) === SERVICE_STATUS.WARNING
                          "
                          color="warning"
                        >
                          mdi-server
                        </v-icon>
                        <v-icon
                          v-else-if="
                            getServiceStatus(service) === SERVICE_STATUS.ERROR
                          "
                          color="error"
                        >
                          mdi-server
                        </v-icon>
                        <v-icon
                          v-else-if="
                            getServiceStatus(service) ===
                            SERVICE_STATUS.DISABLED
                          "
                          color="grey"
                        >
                          mdi-server
                        </v-icon>
                        <v-icon v-else>mdi-server</v-icon>
                      </template>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content>
                      <v-row>
                        <v-col cols="12">
                          <div>WIP</div>
                        </v-col>
                      </v-row>
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
import { SERVICE_STATUS, STATUS } from "@/constants";
import { ApplicationInfo, ServiceInfo } from "@/apis/datatransfers";
export default Vue.extend({
  data() {
    return {
      STATUS,
      SERVICE_STATUS,
      pageLoadStatus: STATUS.PRE_LOADING,
      applications: new Array<ApplicationInfo>(),
    };
  },
  created() {
    this.loadMonitor();
  },
  methods: {
    loadMonitor() {
      this.applications = new Array<ApplicationInfo>();
      api.monitor
        .get()
        .then((response) => {
          response.data.data.forEach((application: ApplicationInfo) => {
            this.applications.push(application);
          });
          this.pageLoadStatus = STATUS.COMPLETE;
        })
        .catch(() => {
          this.pageLoadStatus = STATUS.ERROR;
        });
    },
    getApplicationStatus(application: ApplicationInfo) {
      if (application.services.length === 0) return SERVICE_STATUS.WARNING;
      // TODO
      return SERVICE_STATUS.OK;
    },
    getServiceStatus(service: ServiceInfo) {
      if (!service.enabled) return SERVICE_STATUS.DISABLED;
      // TODO
      return SERVICE_STATUS.OK;
    },
  },
});
</script>
