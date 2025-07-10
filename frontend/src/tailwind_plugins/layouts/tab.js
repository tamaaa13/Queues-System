const plugin = require('tailwindcss/plugin');

module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '[data-pc-layout="tab"]': {
      '.pc-sidebar': {
        '.tab-container': {
          '@apply flex items-start': {},
          '.tab-sidemenu, .tab-link': {
            height: `calc(100vh - theme(spacing.header-height))`,
          },
          '.tab-sidemenu': {
            '@apply w-sidebar-tab-width border-r border-r-theme-border dark:border-r-themedark-border': {},
            '.nav-link': {
              '@apply w-[50px] h-[50px] flex items-center justify-center transition-all text-theme-sidebarcolor dark:text-themedark-sidebarcolor hover:bg-theme-sidebarcolor/20 dark:hover:bg-themedark-sidebarcolor/20 cursor-pointer font-semibold text-[18px] p-0 my-1 mx-auto rounded-lg': {},
              'svg': {
                '@apply w-[22px] h-[22px]': {},
              },
              '&:focus,&.active': {
                '@apply text-primary-500 bg-primary-500/20': {},
              },
            }
          },
          '.tab-link': {
            '@apply flex-auto w-[1%] min-w-0': {},
          }
        },
        '.pc-caption, .pc-user-card': {
          '@apply hidden': {}
        },
        '.pc-mtext': {
          '@apply ml-2.5': {}
        },
        '.pc-micon': {
          '@apply mr-0': {}
        },
        '&:not(.pc-sidebar-hide)': {
          '@apply w-sidebar-tab-navbar-width': {},
          '.navbar-wrapper': {
            '@apply w-sidebar-tab-navbar-width': {}
          }
        },
        '.pc-navbar': {
          '> .pc-item': {
            '> .pc-submenu': {
              '.pc-link': {
                '@apply p-[12px_30px_12px_45px] after:left-[30px]': {}
              }
            },
            '.pc-submenu': {
              '.pc-submenu': {
                '.pc-link': {
                  '@apply p-[12px_30px_12px_52px] after:left-[40px]': {}
                },
                '.pc-submenu': {
                  '.pc-link': {
                    '@apply p-[12px_30px_12px_70px] after:left-[52px]': {}
                  }
                }
              }
            }
          }
        },
        '@media (min-width: 1025px)': {
          '&:not(.pc-sidebar-hide)': {
            '~ .pc-header': {
              '@apply left-sidebar-tab-navbar-width': {}
            },
            '~ .pc-container, ~ .pc-footer': {
              '@apply ml-sidebar-tab-navbar-width': {}
            }
          }
        },
        '@media (max-width: 1024px)': {
          '&:not(.mob-sidebar-active)': {
            '@apply -left-sidebar-tab-navbar-width': {}
          }
        }
      }
    }
                  
    // ===============================================
    // [data-pc-layout='tab'] {
    //   .pc-sidebar {
    //     @media (min-width: 1025px) {
    //     }
    //     @media (max-width: 1024px) {
    //     }
    //   }
    //   &[data-pc-direction='rtl'] {
    //     .pc-sidebar {
    //       .pc-navbar {
    //         > .pc-item {
    //           > .pc-submenu {
    //             .pc-link {
    //               padding: 12px 45px 12px 30px;
    //               &::after {
    //                 right: 30px;
    //               }
    //             }
    //           }
    //           .pc-submenu {
    //             .pc-submenu {
    //               .pc-link {
    //                 padding: 12px 52px 12px 30px;
    //                 &::after {
    //                   right: 40px;
    //                 }
    //               }
    //               .pc-submenu {
    //                 .pc-link {
    //                   padding: 12px 70px 12px 30px;
    //                   &::after {
    //                     right: 52px;
    //                   }
    //                 }
    //               }
    //             }
    //           }
    //         }
    //       }
    //       @media (min-width: 1025px) {
    //         &:not(.pc-sidebar-hide) {
    //           ~ .pc-header {
    //             left: 0;
    //             right: $sidebar-tab-navbar-width;
    //           }
    //           ~ .pc-container,
    //           ~ .pc-footer {
    //             margin-left: 0;
    //             margin-right: $sidebar-tab-navbar-width;
    //           }
    //         }
    //       }
    //       @media (max-width: 1024px) {
    //         &:not(.mob-sidebar-active) {
    //           left: 0;
    //           right: -#{$sidebar-tab-navbar-width};
    //         }
    //       }
    //       .tab-container .tab-sidemenu {
    //         border-left: 1px solid var(--bs-border-color);
    //         border-right: none;
    //       }
    //     }
    //   }
    // }
    // ===============================================
  });
});
