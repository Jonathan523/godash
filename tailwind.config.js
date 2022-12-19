/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/colors/themes")["[data-theme=light]"],
          primary: "#d07915",
        },
        halloween: {
          ...require("daisyui/src/colors/themes")["[data-theme=halloween]"],
          primary: "#d07915",
        },
      },
    ],
    darkTheme: "halloween",
  },
};
