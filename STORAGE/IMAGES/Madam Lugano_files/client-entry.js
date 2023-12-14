/* eslint-disable */
/**
 * THIS FILE IS GENERATED AUTOMATICALLY.
 * DO NOT EDIT.
 *
 * You are probably looking on adding startup/initialization code.
 * Use "quasar new boot <name>" and add it there.
 * One boot file per concern. Then reference the file(s) in quasar.config.js > boot:
 * boot: ['file', ...] // do not add ".js" extension to it.
 *
 * Boot files are your "main.js"
 **/


import { createApp } from '/node_modules/.q-cache/vite/spa/deps/vue.js?v=6164f9a1'







import '/node_modules/@quasar/extras/fontawesome-v6/fontawesome-v6.css'

import '/node_modules/@quasar/extras/roboto-font/roboto-font.css'

import '/node_modules/@quasar/extras/material-icons/material-icons.css'




// We load Quasar stylesheet file
import '/node_modules/quasar/dist/quasar.sass'




import '/src/css/app.scss'

import '/node_modules/@quasar/quasar-ui-qcalendar/src/index.sass'


import createQuasarApp from '/.quasar/app.js?t=1701277900538'
import quasarUserOptions from '/.quasar/quasar-user-options.js'






console.info('[Quasar] Running SPA.')


const publicPath = `/`

async function start ({
  app,
  router
  , store
}, bootFiles) {
  

  
  let hasRedirected = false
  const getRedirectUrl = url => {
    try { return router.resolve(url).href }
    catch (err) {}

    return Object(url) === url
      ? null
      : url
  }
  const redirect = url => {
    hasRedirected = true

    if (typeof url === 'string' && /^https?:\/\//.test(url)) {
      window.location.href = url
      return
    }

    const href = getRedirectUrl(url)

    // continue if we didn't fail to resolve the url
    if (href !== null) {
      window.location.href = href
      window.location.reload()
    }
  }

  const urlPath = window.location.href.replace(window.location.origin, '')

  for (let i = 0; hasRedirected === false && i < bootFiles.length; i++) {
    try {
      await bootFiles[i]({
        app,
        router,
        store,
        ssrContext: null,
        redirect,
        urlPath,
        publicPath
      })
    }
    catch (err) {
      if (err && err.url) {
        redirect(err.url)
        return
      }

      console.error('[Quasar] boot error:', err)
      return
    }
  }

  if (hasRedirected === true) {
    return
  }
  

  app.use(router)
  

  

    

    
      app.mount('#q-app')
    

    

  

}

createQuasarApp(createApp, quasarUserOptions)

  .then(app => {
    // eventually remove this when Cordova/Capacitor/Electron support becomes old
    const [ method, mapFn ] = Promise.allSettled !== void 0
      ? [
        'allSettled',
        bootFiles => bootFiles.map(result => {
          if (result.status === 'rejected') {
            console.error('[Quasar] boot error:', result.reason)
            return
          }
          return result.value.default
        })
      ]
      : [
        'all',
        bootFiles => bootFiles.map(entry => entry.default)
      ]

    return Promise[ method ]([
      
      import('/src/boot/apollo.ts?t=1701276810619'),
      
      import('/src/boot/i18n.ts?t=1701276815934'),
      
      import('/src/boot/api.ts?t=1701276810619'),
      
      import('/node_modules/.q-cache/vite/spa/deps/@quasar_quasar-app-extension-qcalendar_src_boot_register_js.js?v=6164f9a1')
      
    ]).then(bootFiles => {
      const boot = mapFn(bootFiles).filter(entry => typeof entry === 'function')
      start(app, boot)
    })
  })

