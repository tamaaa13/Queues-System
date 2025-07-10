const plugin = require('tailwindcss/plugin');

module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '.badge': {
      '@apply px-[0.8em] py-[0.45em] inline-block text-[.75em] leading-[0.9] font-medium rounded-md': {}
    }
  });
});
