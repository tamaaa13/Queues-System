const plugin = require('tailwindcss/plugin');

module.exports = plugin(function ({ addComponents }) {
  addComponents({
    '.card': {
      '@apply rounded-xl border relative border-theme-border dark:border-themedark-border mb-6 bg-theme-cardbg dark:bg-themedark-cardbg':
        {},

      '.card-header': {
        '@apply p-[25px] border-b border-theme-border dark:border-themedark-border': {}
      },
      '.card-body': {
        '@apply p-[25px]': {}
      },
      '.card-footer': {
        '@apply p-[25px] border-t border-theme-border dark:border-themedark-border': {}
      },
      '.card-link': {
        '@apply text-primary-500 mr-3 inline-block hover:text-theme-bodycolor dark:hover:text-themedark-bodycolor': {}
      }
    }
  });
});
