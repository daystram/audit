export enum SERVICE_STATUS {
  OK = "OK",
  WARNING = "WARNING",
  ERROR = "ERROR",
}

export enum SERVICE_TYPE {
  HTTP = "HTTP",
  TCP = "TCP",
  UDP = "UDP",
}

export const SERVICE_TYPE_LIST = [
  { text: SERVICE_TYPE.HTTP, value: "http" },
  { text: SERVICE_TYPE.TCP, value: "tcp" },
  { text: SERVICE_TYPE.UDP, value: "udp" },
];
