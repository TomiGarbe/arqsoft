export const wait = () =>
  new Promise((res) => setTimeout(res, Math.random() * 2500));
