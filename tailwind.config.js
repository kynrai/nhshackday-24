/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Poppins", "sans-serif"],
      },
      colors: {
        "dhcw-blue": "#688cb4",
        "nhs-blue-1": "#0072CE",
        "nhs-grey-0": "#FFFFFF",
        "nhs-grey-1": "#E8EDEE",
        "nhs-grey-2": "#768692",
        "nhs-grey-3": "#425563",
        "nhs-green-2": "#009639",
        "nhs-purple-2": "#330072",
      },
    },
  },
  plugins: [],
};
