<template>
  <div class="list">
    <v-row class="mb-8" align="center">
      <v-col cols="12" sm="">
        <h1 class="text-h2">Manage</h1>
      </v-col>
      <v-col cols="auto">
        <v-btn color="primary darken-2" rounded @click="resetCreateApplication">
          New Application
        </v-btn>
      </v-col>
    </v-row>

    <!-- APPLICATION ALERTS -->
    <v-expand-transition>
      <div v-show="createApplication.successAlert">
        <v-alert type="success" text dense transition="scroll-y-transition">
          Application successfully created!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="updateApplication.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Application successfully updated!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="delApplication.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Application successfully deleted!
        </v-alert>
      </div>
    </v-expand-transition>
    <!-- SERVICE ALERTS -->
    <v-expand-transition>
      <div v-show="createService.successAlert">
        <v-alert type="success" text dense transition="scroll-y-transition">
          Service successfully created!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="updateService.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Service successfully updated!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="delService.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Service successfully deleted!
        </v-alert>
      </div>
    </v-expand-transition>
    <!-- SPINNER -->
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
    <!-- PAGE ERROR -->
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

    <!-- APPLICATION LIST -->
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
                {{ application.name }}
                <v-spacer />
                <v-btn
                  text
                  rounded
                  plain
                  color="error"
                  @click="resetDeleteApplication(application)"
                >
                  Delete
                </v-btn>
                <v-btn
                  text
                  rounded
                  plain
                  color="warning"
                  class="ml-4"
                  @click="resetUpdateApplication(application)"
                >
                  Edit
                </v-btn>
              </v-card-title>
              <v-card-subtitle>{{ application.description }}</v-card-subtitle>
              <v-divider inset />

              <v-card-text class="pt-6">
                <v-col cols="12">
                  <v-row v-if="application.services.length" class="mb-6">
                    <v-expansion-panels multiple hover>
                      <v-expansion-panel
                        v-for="service in application.services"
                        :key="service.id"
                      >
                        <v-expansion-panel-header>
                          <template v-slot:default>
                            <v-row no-gutters align="center">
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
                              <v-col cols="12" sm="2">
                                <v-icon
                                  :color="service.enabled ? 'primary' : 'grey'"
                                >
                                  mdi-chart-box
                                </v-icon>
                                <v-icon
                                  :color="service.showcase ? 'primary' : 'grey'"
                                >
                                  mdi-radar
                                </v-icon>
                              </v-col>
                            </v-row>
                          </template>
                        </v-expansion-panel-header>
                        <v-expansion-panel-content>
                          <v-row>
                            <v-col cols="12">
                              <v-text-field
                                v-model.trim="service.endpoint"
                                label="Endpoint"
                                disabled
                                :prepend-icon="'mdi-link'"
                              />
                              <v-select
                                v-model.trim="service.type"
                                label="Type"
                                disabled
                                :items="SERVICE_TYPE_LIST"
                                :prepend-icon="'mdi-cogs'"
                              />
                              <v-textarea
                                v-model.trim="service.config"
                                class="configuration-editor"
                                label="Configuration"
                                hint="Service tracking JSON configuration"
                                disabled
                                :prepend-icon="'mdi-code-tags'"
                              />
                            </v-col>
                          </v-row>
                          <v-row>
                            <v-col rows="auto">
                              <v-btn
                                text
                                rounded
                                plain
                                block
                                color="error"
                                @click="
                                  resetDeleteService(application, service)
                                "
                              >
                                Delete
                              </v-btn>
                            </v-col>
                            <v-col rows="auto">
                              <v-btn
                                text
                                rounded
                                plain
                                block
                                color="warning"
                                @click="
                                  resetUpdateService(application, service)
                                "
                              >
                                Edit
                              </v-btn>
                            </v-col>
                          </v-row>
                        </v-expansion-panel-content>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-row>
                  <v-row>
                    <v-btn
                      rounded
                      plain
                      block
                      color="primary"
                      @click="resetCreateService(application)"
                    >
                      New Service
                    </v-btn>
                  </v-row>
                </v-col>
              </v-card-text>
            </v-card>
            <div class="text--disabled" v-if="!applications.length">
              No registered applications
            </div>
          </v-col>
        </v-row>
      </div>
    </v-fade-transition>

    <!-- APPLICATION/CREATE DIALOG -->
    <v-dialog
      v-model="createApplication.creating"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="createApplication.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">New Application</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  createApplication.formLoadStatus === STATUS.IDLE ||
                  createApplication.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelCreateApplication();
                  }
                "
              >
                Cancel
              </v-btn>
              <v-btn
                text
                rounded
                class="ml-4"
                color="success"
                :disabled="createApplication.formLoadStatus === STATUS.LOADING"
                @click="confirmCreateApplication"
              >
                <div v-if="createApplication.formLoadStatus !== STATUS.LOADING">
                  Create
                </div>
                <div v-else>Creating</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="createApplication.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed creating application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="createApplication.application.name"
                :error-messages="createApplicationNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="createApplication.formLoadStatus === STATUS.LOADING"
                @input="$v.createApplication.application.name.$touch()"
                @blur="$v.createApplication.application.name.$touch()"
                :prepend-icon="'mdi-application'"
              />
              <v-text-field
                v-model.trim="createApplication.application.description"
                :error-messages="createApplicationDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="createApplication.formLoadStatus === STATUS.LOADING"
                @input="$v.createApplication.application.description.$touch()"
                @blur="$v.createApplication.application.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- APPLICATION/UPDATE DIALOG -->
    <v-dialog
      v-model="updateApplication.updating"
      class="ml-4"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="updateApplication.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">Update Application</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  updateApplication.formLoadStatus === STATUS.IDLE ||
                  updateApplication.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelUpdateApplication();
                  }
                "
              >
                Cancel
              </v-btn>
              <v-btn
                text
                rounded
                class="ml-4"
                color="success"
                :disabled="
                  updateApplication.formLoadStatus === STATUS.LOADING ||
                  !updateApplicationChanged
                "
                @click="confirmUpdateApplication"
              >
                <div v-if="updateApplication.formLoadStatus !== STATUS.LOADING">
                  Update
                </div>
                <div v-else>Updating</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="updateApplication.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed updating application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="updateApplication.application.name"
                :error-messages="updateApplicationNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="updateApplication.formLoadStatus === STATUS.LOADING"
                @input="$v.updateApplication.application.name.$touch()"
                @blur="$v.updateApplication.application.name.$touch()"
                :prepend-icon="'mdi-application'"
              />
              <v-text-field
                v-model.trim="updateApplication.application.description"
                :error-messages="updateApplicationDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="updateApplication.formLoadStatus === STATUS.LOADING"
                @input="$v.updateApplication.application.description.$touch()"
                @blur="$v.updateApplication.application.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- APPLICATION/DELETE DIALOG -->
    <v-dialog v-model="delApplication.deleting" width="546" overlay-opacity="0">
      <v-card :loading="delApplication.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">Delete Application</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  delApplication.formLoadStatus === STATUS.IDLE ||
                  delApplication.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelDeleteApplication();
                  }
                "
              >
                Cancel
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="delApplication.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed deleting application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              Are you sure you want to delete
              <b>{{ delApplication.application.name }}</b> application? All of
              this application's services, incidents, and logs will be deleted.
              This action is <b>irreversible</b>!
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn
                outlined
                rounded
                color="error"
                block
                :disabled="delApplication.formLoadStatus === STATUS.LOADING"
                @click="confirmDeleteApplication"
              >
                <div v-if="delApplication.formLoadStatus !== STATUS.LOADING">
                  Delete Application
                </div>
                <div v-else>Deleting</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>

    <!-- SERVICE/CREATE DIALOG -->
    <v-dialog
      v-model="createService.creating"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="createService.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">New Service</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  createService.formLoadStatus === STATUS.IDLE ||
                  createService.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelCreateService();
                  }
                "
              >
                Cancel
              </v-btn>
              <v-btn
                text
                rounded
                class="ml-4"
                color="success"
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                @click="confirmCreateService"
              >
                <div v-if="createService.formLoadStatus !== STATUS.LOADING">
                  Create
                </div>
                <div v-else>Creating</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="createService.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed creating service!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="createService.service.name"
                :error-messages="createServiceNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                @input="$v.createService.service.name.$touch()"
                @blur="$v.createService.service.name.$touch()"
                :prepend-icon="'mdi-server'"
              />
              <v-text-field
                v-model.trim="createService.service.description"
                :error-messages="createServiceDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                @input="$v.createService.service.description.$touch()"
                @blur="$v.createService.service.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
              <v-text-field
                v-model.trim="createService.service.endpoint"
                :error-messages="createServiceEndpointErrors"
                :counter="50"
                label="Endpoint"
                hint="URL with protocol or IP address"
                required
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                @input="$v.createService.service.endpoint.$touch()"
                @blur="$v.createService.service.endpoint.$touch()"
                :prepend-icon="'mdi-link'"
              />
              <v-select
                v-model.trim="createService.service.type"
                label="Type"
                required
                :items="SERVICE_TYPE_LIST"
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-cogs'"
              />
              <v-textarea
                v-model.trim="createService.service.config"
                :error-messages="createServiceConfigErrors"
                class="configuration-editor"
                label="Configuration"
                hint="Service tracking JSON configuration"
                required
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                @input="$v.createService.service.config.$touch()"
                @blur="$v.createService.service.config.$touch()"
                :prepend-icon="'mdi-code-tags'"
              />
              <v-checkbox
                v-model.trim="createService.service.enabled"
                label="Enable service tracking"
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-chart-box'"
              />
              <v-checkbox
                v-model.trim="createService.service.showcase"
                label="Show service in Monitor page"
                required
                :disabled="createService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-radar'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- SERVICE/UPDATE DIALOG -->
    <v-dialog
      v-model="updateService.updating"
      class="ml-4"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="updateService.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">Update Service</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  updateService.formLoadStatus === STATUS.IDLE ||
                  updateService.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelUpdateService();
                  }
                "
              >
                Cancel
              </v-btn>
              <v-btn
                text
                rounded
                class="ml-4"
                color="success"
                :disabled="
                  updateService.formLoadStatus === STATUS.LOADING ||
                  !updateServiceChanged
                "
                @click="confirmUpdateService"
              >
                <div v-if="updateService.formLoadStatus !== STATUS.LOADING">
                  Update
                </div>
                <div v-else>Updating</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="updateService.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed updating service!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="updateService.service.name"
                :error-messages="updateServiceNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                @input="$v.updateService.service.name.$touch()"
                @blur="$v.updateService.service.name.$touch()"
                :prepend-icon="'mdi-server'"
              />
              <v-text-field
                v-model.trim="updateService.service.description"
                :error-messages="updateServiceDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                @input="$v.updateService.service.description.$touch()"
                @blur="$v.updateService.service.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
              <v-text-field
                v-model.trim="updateService.service.endpoint"
                :error-messages="updateServiceEndpointErrors"
                :counter="50"
                label="Endpoint"
                hint="URL with protocol or IP address"
                required
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                @input="$v.updateService.service.endpoint.$touch()"
                @blur="$v.updateService.service.endpoint.$touch()"
                :prepend-icon="'mdi-link'"
              />
              <v-select
                v-model.trim="updateService.service.type"
                label="Type"
                required
                :items="SERVICE_TYPE_LIST"
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-cogs'"
              />
              <v-textarea
                v-model.trim="updateService.service.config"
                :error-messages="updateServiceConfigErrors"
                class="configuration-editor"
                label="Configuration"
                hint="Service tracking JSON configuration"
                required
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                @input="$v.updateService.service.config.$touch()"
                @blur="$v.updateService.service.config.$touch()"
                :prepend-icon="'mdi-code-tags'"
              />
              <v-checkbox
                v-model.trim="updateService.service.enabled"
                label="Enable service tracking"
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-chart-box'"
              />
              <v-checkbox
                v-model.trim="updateService.service.showcase"
                label="Show service in Monitor page"
                required
                :disabled="updateService.formLoadStatus === STATUS.LOADING"
                :prepend-icon="'mdi-radar'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- SERVICE/DELETE DIALOG -->
    <v-dialog v-model="delService.deleting" width="546" overlay-opacity="0">
      <v-card :loading="delService.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">Delete Service</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  delService.formLoadStatus === STATUS.IDLE ||
                  delService.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelDeleteService();
                  }
                "
              >
                Cancel
              </v-btn>
            </v-col>
          </v-row>
        </v-card-title>
        <v-divider inset />
        <v-card-text class="pt-4">
          <v-expand-transition>
            <div v-show="delService.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed deleting service!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              Are you sure you want to delete
              <b>{{ delService.service.name }}</b> service? All of this
              services's incidents, and logs will be deleted. This action is
              <b>irreversible</b>!
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn
                outlined
                rounded
                color="error"
                block
                :disabled="delService.formLoadStatus === STATUS.LOADING"
                @click="confirmDeleteService"
              >
                <div v-if="delService.formLoadStatus !== STATUS.LOADING">
                  Delete Service
                </div>
                <div v-else>Deleting</div>
              </v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import api from "@/apis";
import {
  SERVICE_STATUS,
  SERVICE_TYPE,
  SERVICE_TYPE_LIST,
  STATUS,
} from "@/constants";
import { ApplicationInfo, ServiceInfo } from "@/apis/datatransfers";
import { ipAddressWithPort, isJson } from "@/utils/validators";
import { maxLength, required, or, url } from "vuelidate/lib/validators";
export default Vue.extend({
  data() {
    return {
      STATUS,
      SERVICE_STATUS,
      SERVICE_TYPE,
      SERVICE_TYPE_LIST,
      pageLoadStatus: STATUS.PRE_LOADING,
      applications: new Array<ApplicationInfo>(),
      // TODO: rename
      createApplication: {
        creating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
      },
      updateApplication: {
        updating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
        old: new ApplicationInfo(),
      },
      delApplication: {
        deleting: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
      },
      createService: {
        creating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
        service: new ServiceInfo(),
      },
      updateService: {
        updating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
        service: new ServiceInfo(),
        old: new ServiceInfo(),
      },
      delService: {
        deleting: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
        service: new ServiceInfo(),
      },
    };
  },

  computed: {
    createApplicationNameErrors() {
      const errors: string[] = [];
      if (!this.$v.createApplication.application?.name.$dirty) return errors;
      !this.$v.createApplication.application.name.required &&
        errors.push("Name required");
      !this.$v.createApplication.application.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    createApplicationDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.createApplication.application?.description.$dirty)
        return errors;
      !this.$v.createApplication.application.description.required &&
        errors.push("Description required");
      !this.$v.createApplication.application.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    updateApplicationNameErrors() {
      const errors: string[] = [];
      if (!this.$v.updateApplication.application?.name.$dirty) return errors;
      !this.$v.updateApplication.application.name.required &&
        errors.push("Name required");
      !this.$v.updateApplication.application.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    updateApplicationDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.updateApplication.application?.description.$dirty)
        return errors;
      !this.$v.updateApplication.application.description.required &&
        errors.push("Description required");
      !this.$v.updateApplication.application.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    updateApplicationChanged() {
      return (
        this.$data.updateApplication.application.name !==
          this.$data.updateApplication.old.name ||
        this.$data.updateApplication.application.description !==
          this.$data.updateApplication.old.description
      );
    },

    createServiceNameErrors() {
      const errors: string[] = [];
      if (!this.$v.createService.service?.name.$dirty) return errors;
      !this.$v.createService.service.name.required &&
        errors.push("Name required");
      !this.$v.createService.service.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    createServiceDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.createService.service?.description.$dirty) return errors;
      !this.$v.createService.service.description.required &&
        errors.push("Description required");
      !this.$v.createService.service.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    createServiceEndpointErrors() {
      const errors: string[] = [];
      if (!this.$v.createService.service?.endpoint.$dirty) return errors;
      !this.$v.createService.service.endpoint.required &&
        errors.push("Endpoint required");
      !this.$v.createService.service.endpoint.maxLength &&
        errors.push("Endpoint too long");
      !this.$v.createService.service.endpoint.endpoint &&
        errors.push("Endpoint invalid");
      return errors;
    },
    createServiceConfigErrors() {
      const errors: string[] = [];
      if (!this.$v.createService.service?.config.$dirty) return errors;
      !this.$v.createService.service.config.required &&
        errors.push("Configuration required");
      !this.$v.createService.service.config.json && errors.push("Invalid JSON");
      return errors;
    },
    updateServiceNameErrors() {
      const errors: string[] = [];
      if (!this.$v.updateService.service?.name.$dirty) return errors;
      !this.$v.updateService.service.name.required &&
        errors.push("Name required");
      !this.$v.updateService.service.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    updateServiceDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.updateService.service?.description.$dirty) return errors;
      !this.$v.updateService.service.description.required &&
        errors.push("Description required");
      !this.$v.updateService.service.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    updateServiceEndpointErrors() {
      const errors: string[] = [];
      if (!this.$v.updateService.service?.endpoint.$dirty) return errors;
      !this.$v.updateService.service.endpoint.required &&
        errors.push("Endpoint required");
      !this.$v.updateService.service.endpoint.maxLength &&
        errors.push("Endpoint too long");
      !this.$v.updateService.service.endpoint.endpoint &&
        errors.push("Endpoint invalid");
      return errors;
    },
    updateServiceConfigErrors() {
      const errors: string[] = [];
      if (!this.$v.updateService.service?.config.$dirty) return errors;
      !this.$v.updateService.service.config.required &&
        errors.push("Configuration required");
      !this.$v.updateService.service.config.json && errors.push("Invalid JSON");
      return errors;
    },
    updateServiceChanged() {
      return (
        this.$data.updateService.service.name !==
          this.$data.updateService.old.name ||
        this.$data.updateService.service.description !==
          this.$data.updateService.old.description ||
        this.$data.updateService.service.endpoint !==
          this.$data.updateService.old.endpoint ||
        this.$data.updateService.service.type !==
          this.$data.updateService.old.type ||
        this.$data.updateService.service.config !==
          this.$data.updateService.old.config ||
        this.$data.updateService.service.enabled !==
          this.$data.updateService.old.enabled ||
        this.$data.updateService.service.showcase !==
          this.$data.updateService.old.showcase
      );
    },
  },

  validations: {
    createApplication: {
      application: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
      },
    },
    updateApplication: {
      application: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
      },
    },
    createService: {
      service: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
        endpoint: {
          required,
          maxLength: maxLength(50),
          endpoint: or(ipAddressWithPort, url),
        },
        config: {
          required,
          json: or(isJson),
        },
      },
    },
    updateService: {
      service: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
        endpoint: {
          required,
          maxLength: maxLength(50),
          endpoint: or(ipAddressWithPort, url),
        },
        config: {
          required,
          json: or(isJson),
        },
      },
    },
  },

  created() {
    this.loadApplications();
  },

  methods: {
    loadApplications() {
      this.applications = new Array<ApplicationInfo>();
      api.application
        .list()
        .then((response) => {
          response.data.data.forEach((application: ApplicationInfo) => {
            this.applications.push(application);
            this.loadServices(application);
          });
          this.pageLoadStatus = STATUS.COMPLETE;
        })
        .catch(() => {
          this.pageLoadStatus = STATUS.ERROR;
        });
    },
    loadServices(application: ApplicationInfo) {
      this.$set(application, "services", new Array<ServiceInfo>()); // solves Vue's nested object reactivity detection
      api.application.service
        .list(application.id as string)
        .then((response) => {
          response.data.data.forEach((service: ServiceInfo) => {
            application.services.push(service);
          });
        })
        .catch(() => {
          this.pageLoadStatus = STATUS.ERROR;
        });
    },
    // APPLICATION/CREATE
    resetCreateApplication() {
      this.createApplication.creating = true;
      this.createApplication.formLoadStatus = STATUS.IDLE;
      this.createApplication.application = new ApplicationInfo();
      this.$v.createApplication.$reset();
    },
    confirmCreateApplication() {
      this.$v.createApplication.$touch();
      if (!this.$v.createApplication.$invalid) {
        this.createApplication.formLoadStatus = STATUS.LOADING;
        api.application
          .create(this.createApplication.application)
          .then(() => {
            this.createApplication.creating = false;
            this.createApplication.successAlert = true;
            this.loadApplications();
            setTimeout(() => {
              this.createApplication.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.createApplication.apiResponseCode = error.response.data.code;
            this.createApplication.formLoadStatus = !this.createApplication
              .apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelCreateApplication() {
      this.createApplication.creating = false;
    },
    // APPLICATION/UPDATE
    resetUpdateApplication(application: ApplicationInfo) {
      this.updateApplication.updating = true;
      this.updateApplication.formLoadStatus = STATUS.IDLE;
      this.updateApplication.application = new ApplicationInfo(application);
      this.updateApplication.old = application;
      this.$v.updateApplication.$reset();
    },
    confirmUpdateApplication() {
      this.$v.updateApplication.$touch();
      if (!this.$v.updateApplication.$invalid) {
        this.updateApplication.formLoadStatus = STATUS.LOADING;
        api.application
          .update(this.updateApplication.application)
          .then(() => {
            this.updateApplication.updating = false;
            this.updateApplication.successAlert = true;
            this.loadApplications();
            setTimeout(() => {
              this.updateApplication.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.updateApplication.apiResponseCode = error.response.data.code;
            this.updateApplication.formLoadStatus = !this.updateApplication
              .apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelUpdateApplication() {
      this.updateApplication.updating = false;
    },
    // APPLICATION/DELETE
    resetDeleteApplication(application: ApplicationInfo) {
      this.delApplication.deleting = true;
      this.delApplication.formLoadStatus = STATUS.IDLE;
      this.delApplication.application = application;
    },
    confirmDeleteApplication() {
      this.delApplication.formLoadStatus = STATUS.LOADING;
      api.application
        .delete(this.delApplication.application.id as string)
        .then(() => {
          this.delApplication.deleting = false;
          this.delApplication.successAlert = true;
          this.loadApplications();
          setTimeout(() => {
            this.delApplication.successAlert = false;
          }, 10000);
        })
        .catch((error) => {
          this.delApplication.apiResponseCode = error.response.data.code;
          this.delApplication.formLoadStatus = !this.delApplication
            .apiResponseCode
            ? STATUS.ERROR
            : STATUS.IDLE;
        });
    },
    cancelDeleteApplication() {
      this.delApplication.deleting = false;
    },
    // SERVICE/CREATE
    resetCreateService(application: ApplicationInfo) {
      this.createService.creating = true;
      this.createService.formLoadStatus = STATUS.IDLE;
      this.createService.application = application;
      this.createService.service = new ServiceInfo();
      this.createService.service.type = SERVICE_TYPE_LIST.filter(
        (type) => type.text === SERVICE_TYPE.HTTP
      )[0].value;
      this.createService.service.config = JSON.stringify({});
      this.$v.createService.$reset();
    },
    confirmCreateService() {
      this.$v.createService.$touch();
      if (!this.$v.createService.$invalid) {
        this.createService.formLoadStatus = STATUS.LOADING;
        api.application.service
          .create(
            this.createService.application.id as string,
            this.createService.service
          )
          .then(() => {
            this.createService.creating = false;
            this.createService.successAlert = true;
            this.loadServices(this.createService.application);
            setTimeout(() => {
              this.createService.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.createService.apiResponseCode = error.response.data.code;
            this.createService.formLoadStatus = !this.createService
              .apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelCreateService() {
      this.createService.creating = false;
    },
    // SERVICE/UPDATE
    resetUpdateService(application: ApplicationInfo, service: ServiceInfo) {
      this.updateService.updating = true;
      this.updateService.formLoadStatus = STATUS.IDLE;
      this.updateService.application = application;
      this.updateService.service = new ServiceInfo(service);
      this.updateService.old = service;
      this.$v.updateService.$reset();
    },
    confirmUpdateService() {
      this.$v.updateService.$touch();
      if (!this.$v.updateService.$invalid) {
        this.updateService.formLoadStatus = STATUS.LOADING;
        api.application.service
          .update(
            this.updateService.application.id as string,
            this.updateService.service
          )
          .then(() => {
            this.updateService.updating = false;
            this.updateService.successAlert = true;
            this.loadServices(this.updateService.application);
            setTimeout(() => {
              this.updateService.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.updateService.apiResponseCode = error.response.data.code;
            this.updateService.formLoadStatus = !this.updateService
              .apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelUpdateService() {
      this.updateService.updating = false;
    },
    // SERVICE/DELETE
    resetDeleteService(application: ApplicationInfo, service: ServiceInfo) {
      this.delService.deleting = true;
      this.delService.formLoadStatus = STATUS.IDLE;
      this.delService.application = application;
      this.delService.service = service;
    },
    confirmDeleteService() {
      this.delService.formLoadStatus = STATUS.LOADING;
      api.application.service
        .delete(
          this.delService.application.id as string,
          this.delService.service.id as string
        )
        .then(() => {
          this.delService.deleting = false;
          this.delService.successAlert = true;
          this.loadServices(this.delService.application);
          setTimeout(() => {
            this.delService.successAlert = false;
          }, 10000);
        })
        .catch((error) => {
          this.delService.apiResponseCode = error.response.data.code;
          this.delService.formLoadStatus = !this.delService.apiResponseCode
            ? STATUS.ERROR
            : STATUS.IDLE;
        });
    },
    cancelDeleteService() {
      this.delService.deleting = false;
    },
  },
});
</script>
