// import type { Config } from "tailwindcss";
// import typography from "@tailwindcss/typography";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js,go}"],
  theme: {
    extend: {
      colors: {
        bgcolor: "rgba(var(--bgcolor))",
        lively: "rgba(var(--lively))",
        lovely: "rgba(var(--lovely))",
        maintext: "rgba(var(--maintext))",
        subtext: "rgba(var(--subtext))",
        hinttext: "rgba(var(--hinttext))",
        err: "rgba(var(--err))",
      },
    },
  },
  plugins: [],
}

