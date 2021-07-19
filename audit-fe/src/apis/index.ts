import axios, { AxiosInstance, AxiosResponse } from "axios";
import { ACCESS_TOKEN } from "@daystram/ratify-client";
import { authManager, refreshAuth } from "@/auth";
import router from "@/router";
import { ApplicationInfo, ServiceInfo } from "./datatransfers";

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
    list: function (): Promise<AxiosResponse> {
      return apiClient.get(`application/`);
    },
    get: function (id: string): Promise<AxiosResponse> {
      return apiClient.get(`application/${id}`);
    },
    create: function (
      applicationInfo: ApplicationInfo
    ): Promise<AxiosResponse> {
      return apiClient.post(`application/`, applicationInfo);
    },
    update: function (
      applicationInfo: ApplicationInfo
    ): Promise<AxiosResponse> {
      return apiClient.put(
        `application/${applicationInfo.id}`,
        applicationInfo
      );
    },
    delete: function (applicationId: string): Promise<AxiosResponse> {
      return apiClient.delete(`application/${applicationId}`);
    },
    service: {
      list: function (applicationId: string): Promise<AxiosResponse> {
        return apiClient.get(`application/${applicationId}/service/`);
      },
      get: function (
        applicationId: string,
        serviceId: string
      ): Promise<AxiosResponse> {
        return apiClient.get(
          `application/${applicationId}/service/${serviceId}`
        );
      },
      create: function (
        applicationId: string,
        serviceInfo: ServiceInfo
      ): Promise<AxiosResponse> {
        return apiClient.post(
          `application/${applicationId}/service/`,
          serviceInfo
        );
      },
      update: function (
        applicationId: string,
        serviceInfo: ServiceInfo
      ): Promise<AxiosResponse> {
        return apiClient.put(
          `application/${applicationId}/service/${serviceInfo.id}`,
          serviceInfo
        );
      },
      delete: function (
        applicationId: string,
        serviceId: string
      ): Promise<AxiosResponse> {
        return apiClient.delete(
          `application/${applicationId}/service/${serviceId}`
        );
      },
    },
  },
};
