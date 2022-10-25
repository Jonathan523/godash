/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["corporate", "halloween"],
    logs: false,
    darkTheme: "halloween",
  },
};
