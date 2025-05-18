/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{js,jsx,ts,tsx}", "index.html", "*.{js,ts,jsx,tsx,mdx}"],
  theme: {
    extend: {},
  },
    plugins: [require("tailwindcss-animate")],
};