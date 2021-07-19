/* eslint-disable @typescript-eslint/no-explicit-any */
export class ApplicationInfo {
  id?: string;
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
    this.name = source["name"];
    this.description = source["description"];
    this.createdAt = source["createdAt"];
    this.updatedAt = source["updatedAt"];
  }
}
export class ServiceInfo {
  id?: string;
  name: string;
  description: string;
  endpoint: string;
  type: string;
  config: string;
  showcase: boolean;
  createdAt: number;
  updatedAt: number;

  static createFrom(source: any = {}): ServiceInfo {
    return new ServiceInfo(source);
  }

  constructor(source: any = {}) {
    if ("string" === typeof source) source = JSON.parse(source);
    this.id = source["id"];
    this.name = source["name"];
    this.description = source["description"];
    this.endpoint = source["endpoint"];
    this.type = source["type"];
    this.config = source["config"];
    this.showcase = source["showcase"];
    this.createdAt = source["createdAt"];
    this.updatedAt = source["updatedAt"];
  }
}
