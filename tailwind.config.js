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
          primary: "#0047af",
        },
      },
      "halloween",
    ],
    darkTheme: "halloween",
  },
};
