const plugin = require('tailwindcss/plugin');

module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '[data-pc-layout="color-header"]': {
      'body': {
        '@apply relative': {},
        '&:not([class*="bg-"])': {
          '@apply bg-primary-500': {},
        },
      },
      '.pc-header': {
        '@apply absolute bg-transparent': {},
        '.pc-head-link': {
          '@apply text-white after:bg-[rgba(255,255,255,0.15)]': {},
          'svg, i': {
            '@apply text-white': {},
          }
        },
      },
      '.pc-sidebar': {
        '@apply absolute border-0 bg-transparent h-full': {},
        '.navbar-wrapper': {
          '@apply h-full': {},
        },
        '.navbar-content': {
          '@apply relative p-[18px_0_10px] rounded-[0_12px_0_0] shadow-[inset_0_0_1px_1px_#fff] shadow-theme-border dark:shadow-themedark-border bg-white dark:bg-themedark-cardbg': {},
          minHeight: `calc(100% - theme(spacing.header-height))`
        },
      },
      '.pc-footer': {
        '@apply bg-theme-bodybg dark:bg-themedark-bodybg': {},
      },
      '.pc-container': {
        '@apply pt-[140px]': {},
        '.pc-content': {
          '@apply bg-theme-bodybg dark:bg-themedark-bodybg min-h-[calc(100vh_-_273px)]': {},
        },
        '.page-header': {
          '@apply mt-[-140px] p-0': {},
          'h2': {
            '@apply text-white': {},
          },
          '.breadcrumb': {
            '@apply mb-[5px]': {},
            
            '.breadcrumb-item, a': {
              '@apply text-white': {}
            },
          },
        },
      },
    }
  });
});
