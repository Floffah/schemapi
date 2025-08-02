import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Schemapi",
  description: "Documentation for Schemapi",
  base: "/schemapi",
  themeConfig: {
    editLink: {
      pattern: 'https://github.com/floffah/schemapi/edit/main/docs/:path'
    },

    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Get Started', link: '/getting-started' }
    ],

    sidebar: [
      {
        text: 'Writing Schemapi',
        items: [
          { text: 'Syntax Overview', link: '/syntax' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
