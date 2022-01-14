// importScripts("//s.thsi.cn/js/vis-client/billboard/testpwa/js/precache-manifest.js");

importScripts('//s.thsi.cn/js/common/workbox/4.3.1/workbox-sw.js');
workbox.setConfig({ modulePathPrefix: '//s.thsi.cn/js/common/workbox/4.3.1' });
workbox.core.skipWaiting();

workbox.core.clientsClaim();

/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */
self.__precacheManifest = [].concat(self.__precacheManifest || []);
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});

workbox.routing.registerRoute(/.*s\.thsi\.cn.*\.(js|css)/, new workbox.strategies.CacheFirst(), 'GET');