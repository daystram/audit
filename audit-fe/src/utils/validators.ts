export function ipAddressWithPort(urlString: string): boolean {
  const ipRegex = /^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(:\d{1,5})?$/;
  const addressRegex = /^([-a-zA-Z0-9@:%._+~#=]{2,256}\.[a-z]{2,})(:\d{1,5})?$/;
  return ipRegex.test(urlString) || addressRegex.test(urlString);
}

export function isJson(jsonString: string): boolean {
  try {
    JSON.parse(jsonString);
    return true;
  } catch {
    return false;
  }
}
