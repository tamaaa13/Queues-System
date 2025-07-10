import colors from 'tailwindcss/colors';

const generateColorPalette = (color) => ({
  DEFAULT: color[500],
  50: color[50],
  100: color[100],
  200: color[200],
  300: color[300],
  400: color[400],
  500: color[500],
  600: color[600],
  700: color[700],
  800: color[800],
  900: color[900],
  950: color[950]
});

const presets = {
  defaultTheme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#4680FF',
          50: '#e3ecff',
          100: '#c8d9ff',
          200: '#a3c0ff',
          300: '#7ea6ff',
          400: '#6293ff',
          500: '#4680ff',
          600: '#3f78ff',
          700: '#376dff',
          800: '#2f63ff',
          900: '#2050ff',
          950: '#1b43d5'
        }
      }
    }
  },
  themes: [
    {
      name: 'preset-2',
      extend: {
        colors: {
          primary: generateColorPalette(colors.red)
        }
      }
    },
    {
      name: 'preset-3',
      extend: {
        colors: {
          primary: generateColorPalette(colors.orange)
        }
      }
    },
    {
      name: 'preset-4',
      extend: {
        colors: {
          primary: generateColorPalette(colors.amber)
        }
      }
    },
    {
      name: 'preset-5',
      extend: {
        colors: {
          primary: generateColorPalette(colors.yellow)
        }
      }
    },
    {
      name: 'preset-6',
      extend: {
        colors: {
          primary: generateColorPalette(colors.lime)
        }
      }
    },
    {
      name: 'preset-7',
      extend: {
        colors: {
          primary: generateColorPalette(colors.green)
        }
      }
    },
    {
      name: 'preset-8',
      extend: {
        colors: {
          primary: generateColorPalette(colors.emerald)
        }
      }
    },
    {
      name: 'preset-9',
      extend: {
        colors: {
          primary: generateColorPalette(colors.teal)
        }
      }
    },
    {
      name: 'preset-10',
      extend: {
        colors: {
          primary: generateColorPalette(colors.cyan)
        }
      }
    },
    {
      name: 'preset-11',
      extend: {
        colors: {
          primary: generateColorPalette(colors.sky)
        }
      }
    },
    {
      name: 'preset-12',
      extend: {
        colors: {
          primary: generateColorPalette(colors.blue)
        }
      }
    },
    {
      name: 'preset-13',
      extend: {
        colors: {
          primary: generateColorPalette(colors.indigo)
        }
      }
    },
    {
      name: 'preset-14',
      extend: {
        colors: {
          primary: generateColorPalette(colors.violet)
        }
      }
    },
    {
      name: 'preset-15',
      extend: {
        colors: {
          primary: generateColorPalette(colors.purple)
        }
      }
    },
    {
      name: 'preset-16',
      extend: {
        colors: {
          primary: generateColorPalette(colors.fuchsia)
        }
      }
    },
    {
      name: 'preset-17',
      extend: {
        colors: {
          primary: generateColorPalette(colors.pink)
        }
      }
    },
    {
      name: 'preset-18',
      extend: {
        colors: {
          primary: generateColorPalette(colors.rose)
        }
      }
    }
  ]
};

export default presets;
