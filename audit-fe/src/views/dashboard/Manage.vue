<template>
  <div class="list">
    <v-row class="mb-8" align="center">
      <v-col cols="12" sm="">
        <h1 class="text-h2">Manage</h1>
      </v-col>
      <v-col cols="auto">
        <v-btn color="primary darken-2" rounded @click="resetCreate">
          New Application
        </v-btn>
      </v-col>
    </v-row>
    <v-expand-transition>
      <div v-show="create.successAlert">
        <v-alert type="success" text dense transition="scroll-y-transition">
          Application successfully created!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="update.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Application successfully updated!
        </v-alert>
      </div>
    </v-expand-transition>
    <v-expand-transition>
      <div v-show="del.successAlert">
        <v-alert type="info" text dense transition="scroll-y-transition">
          Application successfully deleted!
        </v-alert>
      </div>
    </v-expand-transition>
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
                  @click="resetDelete(application)"
                >
                  Delete
                </v-btn>
                <v-btn
                  text
                  rounded
                  plain
                  color="warning"
                  class="ml-4"
                  @click="resetUpdate(application)"
                >
                  Update
                </v-btn>
              </v-card-title>
              <v-card-subtitle>{{ application.description }}</v-card-subtitle>
              <v-divider inset />

              <v-card-text class="pt-6">
                <v-expansion-panels multiple hover>
                  <v-expansion-panel
                    v-for="service in application.services"
                    :key="service.id"
                  >
                    <v-expansion-panel-header>
                      <b>{{ service.name }}</b>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content>
                      {{ service.description }}
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

    <!-- APPLICATION/CREATE DIALOG -->
    <v-dialog
      v-model="create.creating"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="create.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">New Application</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  create.formLoadStatus === STATUS.IDLE ||
                  create.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelCreate();
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
                :disabled="create.formLoadStatus === STATUS.LOADING"
                @click="confirmCreate"
              >
                <div v-if="create.formLoadStatus !== STATUS.LOADING">
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
            <div v-show="create.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed creating application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="create.application.name"
                :error-messages="createNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="create.formLoadStatus === STATUS.LOADING"
                @input="$v.create.application.name.$touch()"
                @blur="$v.create.application.name.$touch()"
                :prepend-icon="'mdi-application'"
              />
              <v-text-field
                v-model.trim="create.application.description"
                :error-messages="createDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="create.formLoadStatus === STATUS.LOADING"
                @input="$v.create.application.description.$touch()"
                @blur="$v.create.application.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- APPLICATION/UPDATE DIALOG -->
    <v-dialog
      v-model="update.updating"
      class="ml-4"
      width="546"
      persistent
      overlay-opacity="0"
    >
      <v-card :loading="update.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto">Update Application</v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  update.formLoadStatus === STATUS.IDLE ||
                  update.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelUpdate();
                  }
                "
              >
                Cancel
              </v-btn>
              <v-btn
                text
                rounded
                class="ml-4"
                color="warning"
                :disabled="
                  update.formLoadStatus === STATUS.LOADING || !updateChanged
                "
                @click="confirmUpdate"
              >
                <div v-if="update.formLoadStatus !== STATUS.LOADING">
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
            <div v-show="update.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed updating application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              <v-text-field
                v-model.trim="update.application.name"
                :error-messages="updateNameErrors"
                :counter="20"
                label="Name"
                required
                :disabled="update.formLoadStatus === STATUS.LOADING"
                @input="$v.update.application.name.$touch()"
                @blur="$v.update.application.name.$touch()"
                :prepend-icon="'mdi-application'"
              />
              <v-text-field
                v-model.trim="update.application.description"
                :error-messages="updateDescriptionErrors"
                :counter="50"
                label="Description"
                required
                :disabled="update.formLoadStatus === STATUS.LOADING"
                @input="$v.update.application.description.$touch()"
                @blur="$v.update.application.description.$touch()"
                :prepend-icon="'mdi-text'"
              />
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>
    <!-- APPLICATION/DELTE DIALOG -->
    <v-dialog v-model="del.deleting" width="546" overlay-opacity="0">
      <v-card :loading="del.formLoadStatus === STATUS.LOADING">
        <v-card-title>
          <v-row no-gutters align="center">
            <v-col cols="auto"> Delete Application </v-col>
            <v-spacer />
            <v-col cols="auto">
              <v-btn
                v-if="
                  del.formLoadStatus === STATUS.IDLE ||
                  del.formLoadStatus === STATUS.ERROR
                "
                text
                rounded
                color="error"
                @click="
                  () => {
                    cancelDelete();
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
            <div v-show="del.formLoadStatus === STATUS.ERROR">
              <v-alert type="error" text dense transition="scroll-y-transition">
                Failed deleting application!
              </v-alert>
            </div>
          </v-expand-transition>
          <v-row>
            <v-col cols="12">
              Are you sure you want to delete <b>{{ del.application.name }}</b
              >? All of this application's services, incidents, and logs will be
              deleted. This action is <b>irreversible</b>!
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-btn
                outlined
                rounded
                color="error"
                block
                :disabled="del.formLoadStatus === STATUS.LOADING"
                @click="confirmDelete"
              >
                <div v-if="del.formLoadStatus !== STATUS.LOADING">Delete</div>
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
import { SERVICE_STATUS, STATUS } from "@/constants";
import { ApplicationInfo, ServiceInfo } from "@/apis/datatransfers";
import { maxLength, required } from "vuelidate/lib/validators";
export default Vue.extend({
  data() {
    return {
      STATUS,
      SERVICE_STATUS,
      pageLoadStatus: STATUS.PRE_LOADING,
      applications: new Array<ApplicationInfo>(),
      create: {
        creating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
      },
      update: {
        updating: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
        old: new ApplicationInfo(),
      },
      del: {
        deleting: false,
        formLoadStatus: STATUS.IDLE,
        successAlert: false,
        apiResponseCode: "",
        application: new ApplicationInfo(),
      },
    };
  },

  computed: {
    createNameErrors() {
      const errors: string[] = [];
      if (!this.$v.create.application?.name.$dirty) return errors;
      !this.$v.create.application.name.required && errors.push("Name required");
      !this.$v.create.application.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    createDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.create.application?.description.$dirty) return errors;
      !this.$v.create.application.description.required &&
        errors.push("Description required");
      !this.$v.create.application.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    updateNameErrors() {
      const errors: string[] = [];
      if (!this.$v.update.application?.name.$dirty) return errors;
      !this.$v.update.application.name.required && errors.push("Name required");
      !this.$v.update.application.name.maxLength &&
        errors.push("Name too long");
      return errors;
    },
    updateDescriptionErrors() {
      const errors: string[] = [];
      if (!this.$v.update.application?.description.$dirty) return errors;
      !this.$v.update.application.description.required &&
        errors.push("Description required");
      !this.$v.update.application.description.maxLength &&
        errors.push("Description too long");
      return errors;
    },
    updateChanged() {
      return (
        this.$data.update.application.name !== this.$data.update.old.name ||
        this.$data.update.application.description !==
          this.$data.update.old.description
      );
    },
  },

  validations: {
    create: {
      application: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
      },
    },
    update: {
      application: {
        name: { required, maxLength: maxLength(20) },
        description: { required, maxLength: maxLength(50) },
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
          });
          this.pageLoadStatus = STATUS.COMPLETE;
        })
        .catch(() => {
          this.pageLoadStatus = STATUS.ERROR;
        });
    },

    resetCreate() {
      this.create.creating = true;
      this.create.formLoadStatus = STATUS.IDLE;
      this.create.application = new ApplicationInfo();
      this.$v.create.$reset();
    },
    confirmCreate() {
      this.$v.create.$touch();
      if (!this.$v.create.$invalid) {
        this.create.formLoadStatus = STATUS.LOADING;
        api.application
          .create(this.create.application)
          .then(() => {
            this.create.creating = false;
            this.create.successAlert = true;
            this.loadApplications();
            setTimeout(() => {
              this.create.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.create.apiResponseCode = error.response.data.code;
            this.create.formLoadStatus = !this.create.apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelCreate() {
      this.create.creating = false;
    },

    resetUpdate(application: ApplicationInfo) {
      this.update.updating = true;
      this.update.formLoadStatus = STATUS.IDLE;
      this.update.application = new ApplicationInfo(application);
      this.update.old = application;
      this.$v.update.$reset();
    },
    confirmUpdate() {
      this.$v.update.$touch();
      if (!this.$v.update.$invalid) {
        this.update.formLoadStatus = STATUS.LOADING;
        api.application
          .update(this.update.application)
          .then(() => {
            this.update.updating = false;
            this.update.successAlert = true;
            this.loadApplications();
            setTimeout(() => {
              this.update.successAlert = false;
            }, 10000);
          })
          .catch((error) => {
            this.update.apiResponseCode = error.response.data.code;
            this.update.formLoadStatus = !this.update.apiResponseCode
              ? STATUS.ERROR
              : STATUS.IDLE;
          });
      }
    },
    cancelUpdate() {
      this.update.updating = false;
    },

    resetDelete(application: ApplicationInfo) {
      this.del.deleting = true;
      this.del.formLoadStatus = STATUS.IDLE;
      this.del.application = application;
    },
    confirmDelete() {
      this.del.formLoadStatus = STATUS.LOADING;
      api.application
        .delete(this.del.application.id as string)
        .then(() => {
          this.del.deleting = false;
          this.del.successAlert = true;
          this.loadApplications();
          setTimeout(() => {
            this.del.successAlert = false;
          }, 10000);
        })
        .catch((error) => {
          this.create.apiResponseCode = error.response.data.code;
          this.create.formLoadStatus = !this.create.apiResponseCode
            ? STATUS.ERROR
            : STATUS.IDLE;
        });
    },
    cancelDelete() {
      this.del.deleting = false;
    },
  },
});
</script>
