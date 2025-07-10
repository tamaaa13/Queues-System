const plugin = require('tailwindcss/plugin');
module.exports = plugin(function ({ addComponents, theme }) {
  addComponents({
    '.preset-btn.btn': {
      '@apply flex items-center justify-center h-[60px] p-[5px] w-full border-2 border-theme-border dark:border-themedark-border':{},
      "&.active":{
        '@apply border-primary-500':{},
      },
      '&.btn-img':{
        '@apply h-auto rounded-lg':{},
        'img':{
          '@apply w-full':{},
        }
      }
    },
    '.theme-color': {
      '@apply relative flex gap-1':{},
      '>a': {
        '@apply relative w-10 h-10 rounded-xl inline-flex items-center justify-center overflow-hidden transition-all':{},
        '&:after': {
          '@apply content-[""] absolute inset-[4px] rounded-lg bg-white/30 z-10 scale-0 transition-all':{},
        },
        'i': {
          '@apply text-[20px] text-white opacity-70 transition-all':{},
        },
        '&:hover': {
          'i': {
            '@apply opacity-100':{},
          },
        },
        '&.active': {
          '&:after': {
            '@apply scale-100':{},
          },
          'i': {
            '@apply opacity-100':{},
          },
        }
      }
    }
  });
});
