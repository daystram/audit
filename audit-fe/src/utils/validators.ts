export function ipAddressWithPort(urlString: string): boolean {
  const regex = /^(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(:\d{1,5})?$/;
  return regex.test(urlString);
}
