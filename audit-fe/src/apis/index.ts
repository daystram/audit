import axios, { AxiosInstance, AxiosResponse } from "axios";
import { ACCESS_TOKEN } from "@daystram/ratify-client";
import { authManager, refreshAuth } from "@/auth";
import router from "@/router";
import { ApplicationInfo } from "./datatransfers";

const apiClient: AxiosInstance = axios.create({
  baseURL: `${
    import.meta.env.DEV ? import.meta.env.VUE_APP_DEV_BASE_API : ""
  }/api/`,
});

apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      refreshAuth(router.currentRoute.fullPath);
    }
    return Promise.reject(error);
  }
);

const withAuth = () => ({
  headers: {
    Authorization: `Bearer ${authManager.getToken(ACCESS_TOKEN)}`,
  },
});

export default {
  application: {
    get: function (id: string): Promise<Response> {
      return apiClient.get(`application/${id}`);
    },
    list: function (): Promise<AxiosResponse> {
      return apiClient.get(`application/list`);
    },
    create: function (application: ApplicationInfo): Promise<AxiosResponse> {
      return apiClient.post(`application`, application);
    },
    update: function (application: ApplicationInfo): Promise<AxiosResponse> {
      return apiClient.put(`application/${application.id}`, application);
    },
    delete: function (hash: string): Promise<AxiosResponse> {
      return apiClient.delete(`application/${hash}`);
    },
  },
};
