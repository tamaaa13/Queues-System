const plugin = require('tailwindcss/plugin');
module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '.pc-header': {
      '@apply fixed ltr:right-0 rtl:left-0 ltr:max-lg:left-0 ltr:lg:left-sidebar-width  rtl:max-lg:right-0 rtl:lg:right-sidebar-width flex h-header-height z-[1025] backdrop-blur-[7px] bg-theme-headerbg dark:bg-themedark-headerbg text-theme-headercolor dark:text-themedark-headercolor transition-all duration-200 ease-in-out':
        {},
      '.pc-head-link': {
        '@apply w-11 h-11 rounded-lg relative flex items-center justify-center mx-1 text-theme-headercolor after:!content-[""] after:rounded-full after:z-[1] after:absolute after:transition after:inset-0 after:bg-theme-activebg dark:after:bg-themedark-activebg after:scale-0 hover:after:rounded hover:after:scale-100':
          {},
        'i, svg, img ': {
          '@apply relative z-[5] transition': {}
        },
        i: {
          '@apply text-2xl leading-none': {}
        },
        svg: {
          '@apply w-6 h-6': {}
        },
        '&:hover': {
          'i, svg': {
            '@apply scale-[1.08]': {}
          }
        }
      },
      '.pc-h-item.dropdown': {
        '@apply max-sm:static': {},
        '.dropdown-menu': {
          '@apply max-sm:min-w-[calc(100vw_-_30px)] max-sm:!left-[15px] max-sm:!right-[15px] max-sm:!top-full max-w-full': {},
          '&:not(.dropdown-menu-end)': {
            '@apply rtl:!right-0 rtl:!left-auto': {}
          },
        },
        '&.drp-show': {
          '.dropdown-menu': {
            '@apply max-sm:!transform-none': {}
          }
        }
      },
      '.dropdown-menu.dropdown-notification': {
        '@apply sm:min-w-[450px]': {},
        '.card': {
          '@apply cursor-pointer hover:bg-theme-activebg dark:hover:bg-themedark-activebg': {}
        }
      },
      '.dropdown-menu.dropdown-user-profile': {
        '@apply sm:min-w-[352px]': {},
        '.dropdown-item': {
          '@apply flex items-center rounded-lg justify-between': {}
        }
      }
    },
    '.offcanvas.offcanvas-end.pc-announcement-offcanvas': {
      '@apply w-[474px] right-[-474px] z-[1028]': {},
      '.offcanvas-body': {
        '@apply h-[calc(100vh_-_68px)]': {}
      }
    }
  });
});
