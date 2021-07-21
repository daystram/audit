import axios, { AxiosInstance, AxiosResponse } from "axios";
import { ACCESS_TOKEN } from "@daystram/ratify-client";
import { authManager, refreshAuth } from "@/auth";
import router from "@/router";
import { ApplicationInfo, ServiceInfo } from "./datatransfers";

const apiClient: AxiosInstance = axios.create({
  baseURL: `${
    process.env.NODE_ENV === "development"
      ? process.env.VUE_APP_DEV_BASE_API
      : ""
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
  monitor: {
    get: function (): Promise<AxiosResponse> {
      return apiClient.get(`monitor/`);
    },
  },
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
      const body = new ApplicationInfo();
      body.name = applicationInfo.name;
      body.description = applicationInfo.description;
      return apiClient.post(`application/`, body);
    },
    update: function (
      applicationInfo: ApplicationInfo
    ): Promise<AxiosResponse> {
      const body = new ApplicationInfo();
      body.name = applicationInfo.name;
      body.description = applicationInfo.description;
      return apiClient.put(`application/${applicationInfo.id}`, body);
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
        const body = new ServiceInfo();
        body.name = serviceInfo.name;
        body.description = serviceInfo.description;
        body.endpoint = serviceInfo.endpoint;
        body.type = serviceInfo.type;
        body.config = JSON.stringify(JSON.parse(serviceInfo.config)); // cleanup JSON
        body.enabled = !!serviceInfo.enabled; // sets to false if unset
        body.showcase = !!serviceInfo.showcase;
        return apiClient.post(`application/${applicationId}/service/`, body);
      },
      update: function (
        applicationId: string,
        serviceInfo: ServiceInfo
      ): Promise<AxiosResponse> {
        const body = new ServiceInfo();
        body.name = serviceInfo.name;
        body.description = serviceInfo.description;
        body.endpoint = serviceInfo.endpoint;
        body.type = serviceInfo.type;
        body.config = JSON.stringify(JSON.parse(serviceInfo.config)); // cleanup JSON
        body.enabled = !!serviceInfo.enabled; // sets to false if unset
        body.showcase = !!serviceInfo.showcase;
        return apiClient.put(
          `application/${applicationId}/service/${serviceInfo.id}`,
          body
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
