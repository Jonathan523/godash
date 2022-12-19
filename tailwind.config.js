/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/colors/themes")["[data-theme=garden]"],
          primary: "#f28c18",
          secondary: "rgba(70,70,70,0.7)",
        },
        dark: {
          ...require("daisyui/src/colors/themes")["[data-theme=halloween]"],
          secondary: "#b9b9b9",
        },
      },
    ],
    darkTheme: "dark",
  },
};
