import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Alneai's Bookmark",
  description: "A VitePress Site",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '导航', link: '/' },
      { text: '笔记', link: '/note' },
      { text: '存档', link: '/archive' }
    ],

    sidebar: [
      {
        text: '目录',
        items: [
          { text: '导航', link: '/index' },
          { text: '笔记', link: '/note' },
          { text: '存档', link: '/archive' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/Alneai' }
    ]
  },
  vite: {
    server: {
      port: 5200,
      proxy: {
        '/link/': {
          target: 'http://127.0.0.1:5201',
          changeOrigin: true,
        }
      }
    }
  }
})
