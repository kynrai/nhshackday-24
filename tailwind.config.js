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
      },
    },
  },
  plugins: [],
}

