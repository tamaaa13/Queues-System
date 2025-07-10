const plugin = require('tailwindcss/plugin');

module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '[data-pc-layout="compact"]': {
      '.pc-sidebar': {
        '.pc-user-card': {
          '@apply hidden': {}
        },
        '.pc-micon': {
          '@apply h-[27px] w-[46px]': {},
          'svg': {
            '@apply h-[26px] w-[26px]': {}
          }
        },
        '.m-header': {
          '.logo.logo-sm': {
            '@apply w-[50px]': {}
          },
          '.badge':{
            '@apply hidden': {}
          }
        },
        '.pc-navbar': {
          '> .pc-item': {
            '@apply m-[4px_14px]': {},
            '&.pc-hasmenu': {
              '> .pc-submenu': {
                '@apply absolute left-[94px] top-0 h-screen min-w-[186px] bg-theme-bodybg dark:bg-themedark-bodybg': {}
              },
            },
            '> .pc-hasmenu': {
              '.pc-link': {
                '@apply p-[12px_30px_12px_45px] after:left-[30px]': {}
              },
            },
            '.pc-hasmenu': {
              '.pc-hasmenu': {
                '.pc-link': {
                  '@apply p-[12px_30px_12px_52px] after:left-[40px]': {}
                },
                '.pc-hasmenu': {
                  '.pc-link': {
                    '@apply p-[12px_30px_12px_70px] after:left-[52px]': {}
                  },
                }
              }
            },
            '> .pc-link': {
              '@apply p-[14px_12px] inline-block': {},
              '.pc-micon': {
                '@apply mr-0': {},
                'i': {
                  '@apply text-[22px] align-middle': {},
                }
              }
            }
          }
        },
        '&:not(.pc-compact-submenu-active)': {
          '@apply w-sidebar-collapsed-width': {},
          '.m-header': {
            '@apply p-[16px_12px] w-sidebar-collapsed-width': {},
            '> a': {
              '@apply my-0 mx-auto': {},
            }, 
            '.logo': {
              '&.logo-lg': {
                '@apply hidden': {},
              },
              '&.logo-sm': {
                '@apply inline-block': {},
              }
            }  
          },
          '@media (min-width: 1025px)': {
            '~ .pc-footer,~ .pc-container': {
              '@apply ml-sidebar-collapsed-width': {},
            }
          },
          '@media (max-width: 1024.98px)': {
            '.m-header': {
              '@apply w-sidebar-collapsed-width': {},
              '.b-brand': {
                '@apply w-[50px] overflow-hidden my-0 mx-auto': {},
              },
            },
          }
        },
        '@media (max-width: 1024.98px)':{
          '&:not(.mob-sidebar-active)': {
            '@apply -left-sidebar-collapsed-active-width': {},
          },
        },
        '.navbar-content': {
          '@apply w-sidebar-collapsed-width': {},
          height: `calc(100vh - theme(spacing.header-height))`
        },
        '.pc-compact-submenu': {
          '@apply relative': {},
          '&::after': {
            '@apply content-[""] absolute left-0 top-0 h-[calc(100%_-_30px)] w-px bg-theme-border dark:bg-themedark-border': {},
          },
          '.pc-compact-title': {
            '@apply mb-[14px] p-[20px_18px] border-b border-b-theme-border dark:border-b-themedark-border': {},
            '.avtar i': {
              '@apply text-[18px]': {},
            },
            'h5': {
              '@apply font-semibold': {},
            },
          },
          '.pc-compact-list': {
            height: `calc(100vh - theme(spacing.header-height) - 80px)`,
            '.simplebar-content': {
              '> .pc-submenu': {
                '> .pc-item': {
                  '@apply before:left-[15px]': {},
                  '> .pc-link': {
                    '@apply p-[10px_16px]': {},
                  },
                  '> .pc-submenu': {
                    '> .pc-item': {
                      '@apply before:left-[30px]': {},
                      '> .pc-link': {
                        '@apply p-[10px_16px_10px_30px]': {},
                      },  
                      '> .pc-submenu': {
                        '> .pc-item': {
                          '@apply before:left-[45px]': {},
                          '> .pc-link': {
                            '@apply p-[10px_16px_10px_45px]': {},
                          },  
                        }
                      }
                    }
                  }
                },
              },
            },
          },
        },
        '&.pc-compact-submenu-active': {
          '.navbar-content': {
            '@apply w-sidebar-collapsed-active-width': {},
            '&::before': {
              '@apply absolute content-[""] top-0 left-[92px] w-px h-full bg-theme-border dark:bg-themedark-border': {},
            },
          },
          '.navbar-wrapper': {
            '@apply !flex flex-wrap w-sidebar-collapsed-active-width': {},
            '.m-header': {
              '@apply w-full border-b border-b-theme-border dark:border-b-themedark-border': {},
            },
            '.pc-compact-submenu': {
              '@apply flex-1': {},
              width: `calc(100% - theme(spacing.sidebar-collapsed-width))`,
              '.pc-compact-list .simplebar-content > .pc-submenu': {
                '@apply !block': {},
              },
            },
          },
          '@media (min-width: 1025px)': {
            '~ .pc-footer,~ .pc-container': {
              '@apply ml-sidebar-collapsed-active-width': {},
            },
          }
        },
        '&.pc-sidebar-hide': {
          '@apply w-0': {},
          '~ .pc-footer, ~ .pc-container ': {
            '@apply ml-0': {},
          },
          '.pc-navbar': {
            '> .pc-item.pc-hasmenu.pc-trigger > > .pc-submenu': {
              '@apply !hidden': {},
            }
          }
        },
        '.pc-badge,.pc-caption,.pc-mtext,.pc-navbar > li > a > .pc-arrow': {
          '@apply hidden': {},
        }
      },
      '@media (min-width: 1025px)': {
        '.pc-header': {
          '@apply left-sidebar-collapsed-width': {},
        }
      },
      '.pc-sidebar-hide .pc-header': {
        '@apply left-0': {},
      },
      '.pc-compact-submenu-active': {
        '@media (min-width: 1025px)': {
          '&.pc-sidebar-hide .pc-header': {
            '@apply left-0': {},
          },
          '.pc-header': {
            '@apply left-sidebar-collapsed-active-width': {},
          }
        }
      }
    }
  });
});
