import { toast } from "@zerodevx/svelte-toast";

export const success = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "white",
      "--toastColor": "black",
      "--toastBarBackground": "green",
    },
  });

export const warning = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "white",
      "--toastColor": "black",
      "--toastBarBackground": "yellow",
    },
  });

export const failure = (m) =>
  toast.push(m, {
    theme: {
      "--toastBackground": "white",
      "--toastColor": "black",
      "--toastBarBackground": "red",
    },
  });
