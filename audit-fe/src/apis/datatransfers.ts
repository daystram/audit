/* eslint-disable @typescript-eslint/no-explicit-any */
export class ApplicationInfo {
  id?: string;
  services: Array<ServiceInfo>;
  name: string;
  description: string;
  createdAt?: number;
  updatedAt?: number;

  static createFrom(source: any = {}): ApplicationInfo {
    return new ApplicationInfo(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.id = source["id"];
    this.services = source["services"];
    this.name = source["name"];
    this.description = source["description"];
    this.createdAt = source["createdAt"];
    this.updatedAt = source["updatedAt"];
  }
}

export class ServiceInfo {
  id?: string;
  reports: Array<ReportInfo>;
  name: string;
  description: string;
  endpoint: string;
  type: string;
  config: string;
  enabled: boolean;
  showcase: boolean;
  createdAt: number;
  updatedAt: number;

  static createFrom(source: any = {}): ServiceInfo {
    return new ServiceInfo(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.id = source["id"];
    this.reports = source["reports"];
    this.name = source["name"];
    this.description = source["description"];
    this.endpoint = source["endpoint"];
    this.type = source["type"];
    this.config = source["config"];
    this.enabled = source["enabled"];
    this.showcase = source["showcase"];
    this.createdAt = source["createdAt"];
    this.updatedAt = source["updatedAt"];
  }
}

export class ReportInfo {
  latency: number;
  timestamp: number;

  static createFrom(source: any = {}): ReportInfo {
    return new ReportInfo(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.latency = source["latency"];
    this.timestamp = source["timestamp"];
  }
}
